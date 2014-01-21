figo - A FIFO queue for Go
==========================

`figo` is a FIFO for Go (Thanks
[@Corey_latislaw](https://twitter.com/corey_latislaw) for the name!). It is
possible to iterate through the queue, if necessary.

There is additionally a thread safe version. Since it adds a bit of overhead,
only use it if you need it.

`figo` has an MIT license and can be viewed in the LICENSE file.

Example Usage
=============

	q := queue.New()
	q.Push(1)
	q.Push(2)
	
	for elem := q.Front(); elem != nil; elem = q.Next(elem) {
		fmt.Printf("%d ", elem.Value)
	}
	
	q.Pop()
	q.Pop()

To make a thread safe version:
	q := queue.NewAsync()
