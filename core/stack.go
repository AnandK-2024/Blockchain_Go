package core

type Stack struct {
	data []any
}

func NewStack(size int) *Stack {
	return &Stack{
		data: make([]any, size),
	}
}

func (s *Stack) Data() any {
	return s.data
}

func (s *Stack) Push(item any) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() any {
	index := len(s.data) - 1
	value := s.data[index]
	s.data = s.data[:index]
	return value
}

func (s *Stack) len() int {
	return len(s.data)
}

func (s *Stack) peek() any {
	return s.data[s.len()-1]
}
