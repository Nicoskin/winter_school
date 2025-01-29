package distributed

import (
	"fmt"
	"time"
)

func (n *Node) Start(d *Dispatcher) {
	for {

	}
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
	// n.alive = true
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
