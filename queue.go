package queue

type Queue interface {
  Push(interface{})
  Pop() interface{}
}

type queue struct {}

func New() Queue {
  return queue{}
}

func (q queue) Push(o interface{}) {
}

func (q queue) Pop() interface{} {
  return nil
}
