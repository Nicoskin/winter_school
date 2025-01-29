package test

import (
	"fmt"
	"testing"
	"time"
	"ws/distributed"
)

func TestMain4(t *testing.T) {

	// Создаем диспетчер
	d := distributed.NewDispatcher()

	// Создаем узлы
	nodes := []*distributed.Node{
		{ID: 1, NextID: 2, Inbox: make(chan distributed.Message, 10)},
		{ID: 2, NextID: 3, Inbox: make(chan distributed.Message, 10)},
		{ID: 3, NextID: 1, Inbox: make(chan distributed.Message, 10)},
	}

	// Регистрация и запуск
	id_for_leader := -1
	for _, node := range nodes {
		d.RegisterNode(node.ID, node.Inbox)
		node.SetLeaderID(id_for_leader)
		node.SetAlive(true)
		if node.ID == 3 {
			node.SetAlive(false)
		}
		go node.Ring(d)
	}

	time.Sleep(1 * time.Second)

	nodes[0].StartRing(d)

	time.Sleep(2 * time.Second)
	fmt.Println("\nИтоговый результат:")
	for _, node := range nodes {
		fmt.Println("Node", node.ID, "leader is", node.GetLeaderID())
	}

	time.Sleep(1 * time.Second)
	nodes[3].SetAlive(true)
	nodes[3].StartRing(d)

	fmt.Println("\nИтоговый результат:")
	for _, node := range nodes {
		fmt.Println("Node", node.ID, "leader is", node.GetLeaderID())
		close(node.Inbox)
	}
}
