package queue

import "container/list"

type Queue interface {
  Push(interface{})
  Pop() interface{}
  IsEmpty() bool
  Len() int
}

type queue struct {
  l *list.List
}

func New() Queue {
  return queue{list.New()}
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
