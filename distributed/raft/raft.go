package raft

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func (rn *RaftNode) Run() {
	rn.BaseParameters(rn.cluster)
	electionTimer := time.NewTimer(rn.electionTimeout)
	heartbeatTimer := time.NewTicker(rn.heartbeatInterval)

	for {
		switch rn.state {
		case "Follower", "Candidate":
			select {
			case <-electionTimer.C:
				rn.startElection()
				electionTimer.Reset(rn.electionTimeout)
			case msg := <-rn.inbox:
				rn.handleMessage(msg)
				if msg.Kind == "AppendEntries" {
					electionTimer.Reset(rn.electionTimeout)
				}
			}
		case "Leader":
			select {
			case <-heartbeatTimer.C:
				rn.sendHeartbeats()
			case msg := <-rn.inbox:
				rn.handleMessage(msg)
			}
		}
	}
}

func (rn *RaftNode) startElection() {
	rn.state = "Candidate"
	rn.term++
	rn.votedFor = rn.id
	votes := 1

	for _, nodeID := range rn.cluster {
		if nodeID == rn.id {
			continue
		}
		go func(id int) {
			msg := Message{
				Kind: "RequestVote",
				Term: rn.term,
				From: rn.id,
				To:   id,
			}
			rn.outbox <- msg
		}(nodeID)
	}

	// Ожидание ответов
	for votes <= len(rn.cluster)/2 {
		select {
		case msg := <-rn.inbox:
			if msg.Kind == "RequestVoteResponse" && msg.Term == rn.term && msg.Success {
				votes++
			}
		case <-time.After(rn.electionTimeout):
			return
		}
	}

	rn.state = "Leader"
	rn.leaderID = rn.id
	for _, id := range rn.cluster {
		rn.nextIndex[id] = len(rn.log)
		rn.matchIndex[id] = -1
	}
}

func (rn *RaftNode) handleAppendEntries(args AppendEntriesArgs) AppendEntriesReply {
	if args.Term < rn.term {
		return AppendEntriesReply{Term: rn.term, Success: false}
	}

	rn.term = args.Term
	rn.state = "Follower"
	rn.leaderID = args.LeaderID

	if args.PrevLogIndex >= 0 && (len(rn.log) <= args.PrevLogIndex || rn.log[args.PrevLogIndex].Term != args.PrevLogTerm) {
		return AppendEntriesReply{Term: rn.term, Success: false}
	}

	rn.log = append(rn.log[:args.PrevLogIndex+1], args.Entries...)
	if args.LeaderCommit > rn.commitIndex {
		rn.commitIndex = min(args.LeaderCommit, len(rn.log)-1)
	}

	return AppendEntriesReply{Term: rn.term, Success: true, MatchIndex: len(rn.log) - 1}
}

func (rn *RaftNode) sendHeartbeats() {
	for _, nodeID := range rn.cluster {
		if nodeID == rn.id {
			continue
		}

		prevLogIndex := rn.nextIndex[nodeID] - 1
		prevLogTerm := -1
		if prevLogIndex >= 0 {
			prevLogTerm = rn.log[prevLogIndex].Term
		}

		args := AppendEntriesArgs{
			Term:         rn.term,
			LeaderID:     rn.id,
			PrevLogIndex: prevLogIndex,
			PrevLogTerm:  prevLogTerm,
			Entries:      rn.log[rn.nextIndex[nodeID]:],
			LeaderCommit: rn.commitIndex,
		}

		msg := Message{
			Kind:         "AppendEntries",
			Term:         args.Term,
			From:         rn.id,
			To:           nodeID,
			Entries:      args.Entries,
			LeaderID:     args.LeaderID,
			PrevLogIndex: args.PrevLogIndex,
			PrevLogTerm:  args.PrevLogTerm,
			LeaderCommit: args.LeaderCommit,
		}

		rn.outbox <- msg
	}
}

func (rn *RaftNode) updateCommitIndex() {
	for n := len(rn.log) - 1; n > rn.commitIndex; n-- {
		count := 1 // Лидер учитывает свою копию
		for _, id := range rn.cluster {
			if id != rn.id && rn.matchIndex[id] >= n {
				count++
			}
		}
		if count > len(rn.cluster)/2 {
			rn.commitIndex = n
			break
		}
	}
}

func (rn *RaftNode) handleAppendEntriesReply(reply AppendEntriesReply, nodeID int) {
	if reply.Term > rn.term {
		rn.term = reply.Term
		rn.state = "Follower"
		return
	}

	if reply.Success {
		rn.nextIndex[nodeID] = reply.MatchIndex + 1
		rn.matchIndex[nodeID] = reply.MatchIndex
		rn.updateCommitIndex()
	} else {
		rn.nextIndex[nodeID]--
		if rn.nextIndex[nodeID] < 0 {
			rn.nextIndex[nodeID] = 0
		}
	}
}

func (rn *RaftNode) handleMessage(msg Message) {
	// Обновляем term при получении сообщения с большим term
	if msg.Term > rn.term {
		rn.term = msg.Term
		rn.state = "Follower"
		rn.votedFor = -1
	}

	switch msg.Kind {
	case "AppendEntries":
		reply := rn.handleAppendEntriesRequest(msg)
		rn.outbox <- Message{
			Kind:       "AppendEntriesResponse",
			Term:       reply.Term,
			From:       rn.id,
			To:         msg.From,
			Success:    reply.Success,
			MatchIndex: reply.MatchIndex,
		}

	case "AppendEntriesResponse":
		rn.handleAppendEntriesResponse(msg)

	case "RequestVote":
		reply := rn.handleRequestVote(msg)
		rn.outbox <- Message{
			Kind:    "RequestVoteResponse",
			Term:    reply.Term,
			From:    rn.id,
			To:      msg.From,
			Success: reply.VoteGranted,
		}

	case "RequestVoteResponse":
		rn.handleRequestVoteResponse(msg)

	case "Heartbeat":
		rn.resetElectionTimer()
	}
}

// Обработчик AppendEntries RPC (запрос)
func (rn *RaftNode) handleAppendEntriesRequest(msg Message) AppendEntriesReply {
	// Конвертируем Message в AppendEntriesArgs
	args := AppendEntriesArgs{
		Term:         msg.Term,
		LeaderID:     msg.LeaderID,
		PrevLogIndex: msg.PrevLogIndex,
		PrevLogTerm:  msg.PrevLogTerm,
		Entries:      msg.Entries,
		LeaderCommit: msg.LeaderCommit,
	}

	// 1. Проверяем term
	if args.Term < rn.term {
		return AppendEntriesReply{Term: rn.term, Success: false}
	}

	// 2. Сбрасываем таймер выборов
	rn.resetElectionTimer()

	// 3. Проверяем логи
	if args.PrevLogIndex >= 0 {
		if len(rn.log) <= args.PrevLogIndex {
			return AppendEntriesReply{Term: rn.term, Success: false}
		}
		if rn.log[args.PrevLogIndex].Term != args.PrevLogTerm {
			return AppendEntriesReply{Term: rn.term, Success: false}
		}
	}

	// 4. Добавляем записи
	if len(args.Entries) > 0 {
		rn.log = append(rn.log[:args.PrevLogIndex+1], args.Entries...)
	}

	// 5. Обновляем commitIndex
	if args.LeaderCommit > rn.commitIndex {
		rn.commitIndex = min(args.LeaderCommit, len(rn.log)-1)
		rn.applyLogs()
	}

	return AppendEntriesReply{
		Term:       rn.term,
		Success:    true,
		MatchIndex: len(rn.log) - 1,
	}
}

// Обработчик ответов на AppendEntries
func (rn *RaftNode) handleAppendEntriesResponse(msg Message) {
	if rn.state != "Leader" {
		return
	}

	if msg.Success {
		rn.matchIndex[msg.From] = msg.MatchIndex
		rn.nextIndex[msg.From] = msg.MatchIndex + 1
		rn.updateCommitIndex()
	} else {
		rn.nextIndex[msg.From] = max(1, rn.nextIndex[msg.From]-1)
		rn.sendAppendEntries(msg.From)
	}
}

// Обработчик RequestVote RPC
func (rn *RaftNode) handleRequestVote(msg Message) RequestVoteReply {
	// Правила предоставления голоса:
	// 1. Candidate's term >= current term
	// 2. Не голосовали в этом term
	// 3. Лог кандидата не менее полный

	lastLogIndex := len(rn.log) - 1
	lastLogTerm := 0
	if lastLogIndex >= 0 {
		lastLogTerm = rn.log[lastLogIndex].Term
	}

	voteGranted := false
	if (msg.Term > rn.term || (msg.Term == rn.term && rn.votedFor == -1)) &&
		msg.LastLogTerm >= lastLogTerm &&
		msg.LastLogIndex >= lastLogIndex {

		rn.votedFor = msg.From
		voteGranted = true
		rn.resetElectionTimer()
	}

	return RequestVoteReply{
		Term:        rn.term,
		VoteGranted: voteGranted,
	}
}

// Сброс таймера выборов (вызывается при получении heartbeat)
func (rn *RaftNode) resetElectionTimer() {
	rn.electionTimeout = time.Duration(rand.Intn(150)+150) * time.Millisecond
}

// Применение коммитнутых логов
func (rn *RaftNode) applyLogs() {
	for rn.lastApplied < rn.commitIndex {
		rn.lastApplied++
		entry := rn.log[rn.lastApplied]
		fmt.Printf("Node %d applying command: %s\n", rn.id, entry.Command)
	}
}
