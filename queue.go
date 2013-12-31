package queue

type Queue interface {}

type q struct {}

func New() Queue {
  return q{}
}
