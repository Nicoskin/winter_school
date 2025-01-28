package stack_queue

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{data: make([]int, 0)}
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) > 0 {
		val := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return val, true
	} else {
		return 0, false
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
