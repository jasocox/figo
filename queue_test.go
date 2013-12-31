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

func Test_CanIsEmpty(t *testing.T) {
  q := New()
  _ = q.IsEmpty()
}

func Test_CanLen(t *testing.T) {
  q := New()
  _ = q.Len()
}

func Test_CanAddOneItem(t *testing.T) {
  q := New()
  if q.Len() != 0 {
    t.Error("Didn't get the right size")
    return
  }

  q.Push("Nada")
  if q.Len() != 1 {
    t.Error("Should have increased in size")
    return
  }

  val := q.Pop()
  if val != "Nada" {
    t.Error("Did not get the correct value")
    return
  }

  if q.Len() != 0 {
    t.Error("Should have decreased in size")
    return
  }
}
