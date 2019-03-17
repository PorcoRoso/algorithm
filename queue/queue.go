// 队列：head tail
//
package queue

import (
	"github.com/cheekybits/genny/generic"
	"sync"
)

type Item generic.Type

// 普通队列实现
type ArrayQueue struct {
	items []Item
	lock sync.Mutex
}

// 创建队列
func (q *ArrayQueue) NewArrQueue() *ArrayQueue {
	q.items = []Item{}
	return q
}

// 入队列
func (q *ArrayQueue) EnArrQueue(item Item) {
	q.lock.Lock()
	q.items = append(q.items, item)
	q.lock.Unlock()
}

// 出队列
func (q *ArrayQueue) DeArrQueue() *Item {
	q.lock.Lock()
	item := q.items[0]
	q.items = q.items[1:]
	q.lock.Unlock()

	return &item
}

// 获取队列的第一个元素，不移除
func (q *ArrayQueue) Front() *Item {
	q.lock.Lock()
	item := q.items[0]
	q.lock.Unlock()
	return &item
}

// 判空
func (q *ArrayQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// 获取队列的长度
func (q *ArrayQueue) Size() int {
	return len(q.items)
}

// 循环队列实现
// 无锁并发队列实现


