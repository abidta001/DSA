package Queue

/* Queue follows the priciple of "First in, First out" FIFO, where the first element added to the queue is the
First one to be removed. example:Queue */

type Queue struct {
	items []int
}

// Enque(Inserting into Queue)
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue(Removing from Queue)
func (q *Queue) Dequeue() int {
	removed := q.items[0]
	q.items = q.items[1:]
	return removed
}
