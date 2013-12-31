package queue

type Queue interface {
  Push(interface{})
}

type queue struct {}

func New() Queue {
  return queue{}
}

func (q queue) Push(o interface{}) {
}
