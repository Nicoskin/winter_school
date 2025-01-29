package raft

import (
	"time"

	"golang.org/x/exp/rand"
)

type RaftNode struct {
	id          int
	term        int
	state       string // "Follower", "Candidate", "Leader"
	log         []LogEntry
	votedFor    int
	leaderID    int
	commitIndex int
	lastApplied int

	nextIndex  map[int]int
	matchIndex map[int]int
	cluster    []int

	inbox             chan Message
	outbox            chan Message
	electionTimeout   time.Duration
	heartbeatInterval time.Duration
}

type LogEntry struct {
	Term    int
	Command string
}

type Message struct {
	Kind         string
	Term         int
	From         int
	To           int
	Entries      []LogEntry
	LeaderID     int
	PrevLogIndex int
	PrevLogTerm  int
	LeaderCommit int
	Success      bool
	MatchIndex   int
}

type AppendEntriesArgs struct {
	Term         int
	LeaderID     int
	PrevLogIndex int
	PrevLogTerm  int
	Entries      []LogEntry
	LeaderCommit int
}

type AppendEntriesReply struct {
	Term       int
	Success    bool
	MatchIndex int
}

func (rn *RaftNode) BaseParameters(cluster []int) {
	rn.state = "Follower"
	rn.term = 0
	rn.votedFor = -1
	rn.leaderID = -1
	rn.commitIndex = -1
	rn.lastApplied = -1
	rn.cluster = cluster
	rn.nextIndex = make(map[int]int)
	rn.matchIndex = make(map[int]int)
	rn.electionTimeout = time.Duration(rand.Intn(150)+150) * time.Millisecond
	rn.heartbeatInterval = 50 * time.Millisecond
}
