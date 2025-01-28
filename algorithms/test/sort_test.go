package algorithms

import (
	"math/rand"
	"testing"
	"time"
	"ws/algorithms"
	"ws/graph"
)

func TestMain3(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	g := graph.NewGraph()

	for i := 0; i < 1_000_000; i++ {
		g.AddEdge(rand.Intn(100), rand.Intn(100), false, rand.Intn(1000))
	}

	sortedEdges := algorithms.Mergesort(g.GetAllEdges())

	// Проверка результата
	if !isSorted(sortedEdges) {
		t.Error("Результат сортировки некорректен")
	}
}

// Функция для проверки отсортированности массива
func isSorted(edges []graph.Edge) bool {
	for i := 1; i < len(edges); i++ {
		if edges[i-1].GetWeight() > edges[i].GetWeight() {
			return false
		}
	}
	return true
}
