package queue

type Queue struct {
	data     []int
}

func NewQueue() *Queue {
	q := &Queue{
		data: make([]int, 0),
	}
	return q
}

func (q *Queue) Add(key int) {
	q.data = append(q.data, key)
}

func (q *Queue) Remove() int {
	if len(q.data) == 0 {
		return -1
	}
	r := q.data[0]
	q.data = q.data[1:]
	return r
}
