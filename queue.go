package queue

import (
	"container/list"
	"sync"
)

type Queue interface {
	Push(interface{})
	Pop() interface{}
	IsEmpty() bool
	Len() int

	Front() *list.Element
	Next(*list.Element) *list.Element
}

type queue struct {
	l *list.List
}

type aqueue struct {
	q    *queue
	lock sync.Mutex
}

func New() Queue {
	return queue{list.New()}
}

func NewSync() Queue {
	return aqueue{q: &queue{list.New()}}
}

func (q queue) Push(o interface{}) {
	q.l.PushBack(o)
}

func (q queue) Pop() interface{} {
	e := q.l.Front()
	if e == nil {
		return nil
	}

	return q.l.Remove(e)
}

func (q queue) IsEmpty() bool {
	return q.l.Len() == 0
}

func (q queue) Len() int {
	return q.l.Len()
}

func (q queue) Front() *list.Element {
	return q.l.Front()
}

func (q queue) Next(e *list.Element) *list.Element {
	if e == nil {
		return e
	}

	return e.Next()
}

func (q aqueue) Push(o interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.q.Push(o)
}

func (q aqueue) Pop() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.q.Pop()
}

func (q aqueue) IsEmpty() bool {
	return q.q.IsEmpty()
}

func (q aqueue) Len() int {
	return q.q.Len()
}

func (q aqueue) Front() *list.Element {
	return q.q.Front()
}

func (q aqueue) Next(e *list.Element) *list.Element {
	return q.q.Next(e)
}
