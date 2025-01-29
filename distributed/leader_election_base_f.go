package distributed

import (
	"math/rand"
)

// Узел
type Node struct {
	ID         int
	alive      bool
	Inbox      chan Message
	leaderID   int
	startBully bool
	NextID     int // Для Ring: nextID int (кольцо)
	localCount int
	// Для Raft: term, state, log ...
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

func (n *Node) SetLocalCount() {
	n.localCount = rand.Intn(50) + 50
}

func (n *Node) GetLocalCount() int {
	return n.localCount
}
