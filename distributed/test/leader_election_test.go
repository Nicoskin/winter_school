package test

import (
	"fmt"
	"testing"
	"time"
	"ws/distributed"
)

func TestMain(t *testing.T) {

	// Создаем диспетчер
	d := distributed.NewDispatcher()

	// Создаем узлы
	nodes := []*distributed.Node{
		{ID: 1, Inbox: make(chan distributed.Message, 10)},
		{ID: 2, Inbox: make(chan distributed.Message, 10)},
		{ID: 3, Inbox: make(chan distributed.Message, 10)},
	}

	// Регистрация и запуск
	id_for_leader := -1
	for _, node := range nodes {
		d.RegisterNode(node.ID, node.Inbox)
		node.SetLeaderID(id_for_leader)
		go node.Start(d)
	}

	// fmt.Println(nodes[0])
	// fmt.Println(d)

	// d.Send(1, 2, "Hello 1->2")
	// d.Send(2, 3, "Hello 2->3")
	// d.Send(3, 1, "Hello 3->1")

	time.Sleep(1 * time.Second)

	go nodes[1].StartBully(d)

	// Завершение
	time.Sleep(2 * time.Second)
	fmt.Println("\nИтоговый результат:")
	for _, node := range nodes {
		fmt.Println("Node", node.ID, "leader is", node.GetLeaderID())
		close(node.Inbox)
	}
}
