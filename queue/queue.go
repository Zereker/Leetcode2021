package queue

import "errors"

var (
	ErrEmpty = errors.New("empty queue")
)

type Queue struct {
	data []interface{}
}

func New() *Queue {
	data := make([]interface{}, 0)

	return &Queue{data: data}
}

func (q *Queue) Enqueue(value interface{}) {
	q.data = append(q.data, value)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if !q.IsEmpty() {
		value := q.data[0]
		q.data = q.data[1:]
		return value, nil
	}

	return nil, ErrEmpty
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
