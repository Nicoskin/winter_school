package algorithms

type PriorityQueue struct {
	items []Item
}

type Item struct {
	vertex, dist int
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{items: make([]Item, 0)}
}

func (pq *PriorityQueue) Push(item Item) {
	pq.items = append(pq.items, item)
	pq.heapifyUp(len(pq.items) - 1)
}

func (pq *PriorityQueue) Pop() (item Item) {
	min_item := pq.items[0]
	pq.items[0] = pq.items[len(pq.items)-1]
	pq.items = pq.items[:len(pq.items)-1]
	pq.heapifyDown(0)
	return min_item
}

func (pq *PriorityQueue) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if pq.items[parentIndex].dist <= pq.items[index].dist {
			break
		}
		pq.items[parentIndex], pq.items[index] = pq.items[index], pq.items[parentIndex]
		index = parentIndex
	}
}

func (pq *PriorityQueue) heapifyDown(index int) {
	lastIndex := len(pq.items) - 1
	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		smallest := index

		if leftChildIndex <= lastIndex && pq.items[leftChildIndex].dist < pq.items[smallest].dist {
			smallest = leftChildIndex
		}
		if rightChildIndex <= lastIndex && pq.items[rightChildIndex].dist < pq.items[smallest].dist {
			smallest = rightChildIndex
		}
		if smallest == index {
			break
		}
		pq.items[index], pq.items[smallest] = pq.items[smallest], pq.items[index]
		index = smallest
	}
}
