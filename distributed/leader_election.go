package distributed

import (
	"fmt"
	"math/rand"
	"time"
)

// Узел
type Node struct {
	ID         int
	alive      bool
	Inbox      chan Message
	leaderID   int
	startBully bool
	// Для Bully: no big container, just slice of other nodeIDs
	NextID int // Для Ring: nextID int (кольцо)
	// Для Raft: term, state, log ...
	localCount int
}

// Сообщение
type Message struct {
	From    int
	To      int
	Content string
	Data    int
}

// Диспетчер
type Dispatcher struct {
	nodes map[int]chan Message // Карта: ID узла -> Канал узла
}

type LocalData struct {
	localCount int
}

// ============================

// Создание нового диспетчера
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		nodes: make(map[int]chan Message),
	}
}

// Регистрация узла в диспетчере
func (d *Dispatcher) RegisterNode(nodeID int, inbox chan Message) {
	d.nodes[nodeID] = inbox
	//fmt.Println("Node", nodeID, "registered")
}

// Отправка сообщения от одного узла к другому
func (d *Dispatcher) Send(from, to int, content string) {
	d.nodes[to] <- Message{
		From:    from,
		To:      to,
		Content: content,
	}
}

func (d *Dispatcher) SendWithData(from, to int, content string, data int) {
	d.nodes[to] <- Message{
		From:    from,
		To:      to,
		Content: content,
		Data:    data,
	}
}

func (d *Dispatcher) SendToAll(from int, content string) {
	for to, inbox := range d.nodes {
		if to != from {
			inbox <- Message{
				From:    from,
				To:      to,
				Content: content,
			}
		}
	}
}

func (d *Dispatcher) SendHight(from int, content string) {
	for to, inbox := range d.nodes {
		if to > from {
			inbox <- Message{
				From:    from,
				To:      to,
				Content: content,
			}
		}
	}
}

// ============================

// Проверка жив ли узел
func (n *Node) IsAlive() bool {
	return n.alive
}

func (n *Node) SetAlive(alive bool) {
	n.alive = alive
}

func (n *Node) GetLeaderID() int {
	return n.leaderID
}

func (n *Node) SetLeaderID(leaderID int) {
	n.leaderID = leaderID
}

func (n *Node) Bully(d *Dispatcher) {
	n.alive = true
	for msg := range n.Inbox {
		if !n.alive {
			return
		}
		fmt.Printf("Node %d received message from %d: %s\n", n.ID, msg.From, msg.Content)

		switch msg.Content {
		case "ELECTION":
			if n.ID > msg.From {
				d.Send(n.ID, msg.From, "OK")
				time.Sleep(100 * time.Millisecond)
				n.StartBully(d)
			}
		case "COORDINATOR":
			if msg.From > n.leaderID {
				n.SetLeaderID(msg.From)
				n.startBully = false // Остановим дальнейшие выборы
			}
		case "OK":
			// Если получили "OK", продолжаем ждать
			n.startBully = false
		}
	}
}

func (n *Node) StartBully(d *Dispatcher) {
	n.startBully = true
	d.Send(n.ID, n.ID, "startBully")
	d.SendHight(n.ID, "ELECTION")

	time.Sleep(500 * time.Millisecond)

	select {
	case msg := <-n.Inbox:
		if msg.Content == "OK" {
			n.startBully = false
			n.Bully(d)
			return
		}
	default:
		// Продолжаем выполнять проверку, если нет входящего сообщения
	}

	// Если спустя 2 секунды не было ответа, считаем себя координатором
	if n.startBully {
		fmt.Printf("--> Node %d: I'm the boss\n", n.ID)
		d.SendToAll(n.ID, "COORDINATOR")
		n.SetLeaderID(n.ID)
		n.startBully = false
	}

	time.Sleep(1 * time.Second)
}

// RING
func (n *Node) StartRing(d *Dispatcher) {
	d.Send(n.ID, n.NextID, "ELECTION")
}

func (n *Node) Ring(d *Dispatcher) {
	n.alive = true
	for msg := range n.Inbox {
		if !n.alive {
			return
		}
		fmt.Printf("Node %d received message from %d: %s\n", n.ID, msg.From, msg.Content)

		switch msg.Content {
		case "ELECTION":
			if msg.From > n.ID {
				d.Send(msg.From, n.NextID, "ELECTION")
			} else if msg.From < n.ID {
				d.Send(n.ID, n.NextID, "ELECTION")
			} else if msg.From == n.ID {
				fmt.Printf("--> Node %d: I'm the boss\n", n.ID)
				d.Send(n.ID, n.NextID, "COORDINATOR")
				n.SetLeaderID(n.ID)
			}
		case "COORDINATOR":
			if msg.From == n.ID {
				return
			} else {
				n.SetLeaderID(msg.From)
				d.Send(msg.From, n.NextID, "COORDINATOR")
			}
		}
	}
}

func (n *Node) StartGlobalCollection(d *Dispatcher) {
	d.SendToAll(n.ID, "COLLECT")
	received := 0
	sum := 0
	for msg := range n.Inbox {
		if msg.Content == "COLLECT_REPLY" {
			sum += msg.Data
			received++
			if received == len(d.nodes)-1 {
				fmt.Println("{Lider}   sum:", sum)
				return
			}
		}
	}
}

func (n *Node) GlobalCollection(d *Dispatcher) {
	n.alive = false
	d.Send(n.ID, n.ID, "")
	if n.ID == n.leaderID {
		return
	}
	for msg := range n.Inbox {
		if msg.Content == "COLLECT" {
			d.SendWithData(n.ID, n.leaderID, "COLLECT_REPLY", n.localCount)
		}
	}
}

func (n *Node) SetLocalCount() {
	n.localCount = rand.Intn(50) + 50
}

func (n *Node) GetLocalCount() int {
	return n.localCount
}
