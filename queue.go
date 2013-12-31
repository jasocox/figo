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
  lock sync.Mutex
}

func New() Queue {
  return queue{l: list.New()}
}

func (q queue) Push(o interface{}) {
  q.lock.Lock()
  defer q.lock.Unlock()

  q.l.PushBack(o)
}

func (q queue) Pop() interface{} {
  q.lock.Lock()
  defer q.lock.Unlock()

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
