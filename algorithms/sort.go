package algorithms

import "ws/graph"

func Mergesort(edges []graph.Edge) []graph.Edge {
	if len(edges) <= 1 {
		return edges
	}

	mid := len(edges) / 2
	left := Mergesort(edges[:mid])  // Рекурсивно сортируем левую часть
	right := Mergesort(edges[mid:]) // Рекурсивно сортируем правую часть

	return merge(left, right) // Объединяем отсортированные части
}

func merge(left, right []graph.Edge) []graph.Edge {
	result := make([]graph.Edge, 0, len(left)+len(right))
	i, j := 0, 0

	// Сравниваем элементы и добавляем в результат
	for i < len(left) && j < len(right) {
		if left[i].GetWeight() < right[j].GetWeight() {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Добавляем оставшиеся элементы из левой части
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Добавляем оставшиеся элементы из правой части
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}
