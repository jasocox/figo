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
}

type queue struct {
  l *list.List
}

type aqueue struct {
  q *queue
  lock sync.Mutex
}

func New() Queue {
  return queue{list.New()}
}

func NewAsync() Queue {
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
