package distributed

import (
	"fmt"
	"time"
)

// Узел
type Node struct {
	ID       int
	alive    bool
	Inbox    chan Message
	leaderID int
	// Для Bully: no big container, just slice of other nodeIDs
	// Для Ring: nextID int (кольцо)
	// Для Raft: term, state, log ...
}

// Сообщение
type Message struct {
	From    int
	To      int
	Content string
}

// Диспетчер
type Dispatcher struct {
	nodes map[int]chan Message // Карта: ID узла -> Канал узла
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

// Метод узла для обработки сообщений
func (n *Node) Start(d *Dispatcher) {
	n.alive = true
	for {
		if !n.alive {
			fmt.Println("--> Node", n.ID, "is dead")
			return
		}
		msg := <-n.Inbox
		fmt.Printf("Node %d received message from %d: %s\n", n.ID, msg.From, msg.Content)
		n.Bully(msg, d)
	}
}
func (n *Node) Bully(msg Message, d *Dispatcher) {
	//fmt.Println("--> Node", n.ID, "received", msg.Content)
	if msg.Content == "ELECTION" {
		if n.ID > msg.From {
			d.Send(n.ID, msg.From, "OK")
		}
	} else if msg.Content == "COORDINATOR" {
		n.SetLeaderID(msg.From)
	}
}

func (n *Node) StartBully(d *Dispatcher) {
	n.SetAlive(false)
	time.Sleep(2000 * time.Millisecond)
	d.SendToAll(n.ID, "ELECTION")

	hasResponse := false
	for {
		select {
		case msg := <-n.Inbox:
			fmt.Printf("1Node %d received message from %d: %s\n", n.ID, msg.From, msg.Content)
			if msg.Content == "OK" {
				hasResponse = true
			}
		case <-time.After(2 * time.Second):
			if !hasResponse {
				fmt.Printf("--> Node %d: I'm the boss\n", n.ID)
				d.SendToAll(n.ID, "COORDINATOR")
				n.SetLeaderID(n.ID)
				time.Sleep(1 * time.Second)
			}

			n.SetAlive(true)
			return
		}
	}
}
