package stack

import "errors"

var (
	ErrEmpty = errors.New("stack is empty")
)

type Stack struct {
	data []interface{}
}

func New() *Stack {
	data := make([]interface{}, 0)

	return &Stack{data: data}
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(x interface{}) error {
	s.data = append(s.data, x)
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmpty
	}

	index := len(s.data) - 1

	value := s.data[index]
	s.data = s.data[:index]

	return value, nil
}
