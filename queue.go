package queue

type Queue interface {
  Push(interface{})
  Pop() interface{}
  IsEmpty() bool
  Len() int
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

func (q queue) IsEmpty() bool {
  return true
}

func (q queue) Len() int {
  return 0
}
