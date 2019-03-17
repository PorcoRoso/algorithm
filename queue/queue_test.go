package queue

import "testing"

var queue ArrayQueue

func initQueue() *ArrayQueue{
	if queue.items == nil {
		queue = ArrayQueue{}
		queue.NewArrQueue()
		return &queue
	}
	return &queue
}

func TestEnArrQueue(t *testing.T) {
	queue := initQueue()
	queue.EnArrQueue(1)
	queue.EnArrQueue(3)
	queue.EnArrQueue(6)

	if size := queue.Size(); size != 3 {
		
	}
}