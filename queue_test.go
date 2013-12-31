package queue

import "testing"

func TestNew(t *testing.T) {
  _ = New()
}

func Test_CanPush(t *testing.T) {
  q := New()
  q.Push("Nada")
}

func Test_CanPop(t *testing.T) {
  q := New()
  _ = q.Pop()
}
