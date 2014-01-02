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
  }

  q.Push("Nada")
  if q.Len() != 1 {
    t.Error("Should have increased in size")
  }

  val := q.Pop()
  if val != "Nada" {
    t.Error("Did not get the correct value")
  }

  if q.Len() != 0 {
    t.Error("Should have decreased in size")
  }
}

func Test_CanAddMultipleItems(t *testing.T) {
  q := New()

  for i:=0; i<10; i++ {
    q.Push(i)
  }
  if q.IsEmpty() {
    t.Error("The queue should not be empty")
  }

  for i:=0; !q.IsEmpty(); i++ {
    v := q.Pop()
    if v != i {
      t.Error("Did not get the correct value")
    }
  }
}

func TestSync(t *testing.T) {
  _ = NewSync()
}

func TestSync_CanPush(t *testing.T) {
  q := NewSync()
  q.Push("Nada")
}

func TestSync_CanPop(t *testing.T) {
  q := NewSync()
  _ = q.Pop()
}

func TestSync_CanIsEmpty(t *testing.T) {
  q := NewSync()
  _ = q.IsEmpty()
}

func TestSync_CanLen(t *testing.T) {
  q := NewSync()
  _ = q.Len()
}

func TestSync_CanAddOneItem(t *testing.T) {
  q := NewSync()
  if q.Len() != 0 {
    t.Error("Didn't get the right size")
  }

  q.Push("Nada")
  if q.Len() != 1 {
    t.Error("Should have increased in size")
  }

  val := q.Pop()
  if val != "Nada" {
    t.Error("Did not get the correct value")
  }

  if q.Len() != 0 {
    t.Error("Should have decreased in size")
  }
}

func TestSync_CanAddMultipleItems(t *testing.T) {
  q := NewSync()

  for i:=0; i<10; i++ {
    q.Push(i)
  }
  if q.IsEmpty() {
    t.Error("The queue should not be empty")
  }

  for i:=0; !q.IsEmpty(); i++ {
    v := q.Pop()
    if v != i {
      t.Error("Did not get the correct value")
    }
  }
}
