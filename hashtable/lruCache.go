package hashtable

// LRU 哈希表和双向链表实现

type DulListedNode struct {
	key int
	val int
	prev *DulListedNode
	next *DulListedNode
}

type LRUCache struct {
	hashMap map[int]*DulListedNode
	head *DulListedNode
	tail *DulListedNode
	length int
	capacity int
}

func Constructor(capacity int) LRUCache {
	head := &DulListedNode {
		-1,
		-1,
		nil,
		nil,
	}

	tail := &DulListedNode {
		-1,
		-1,
		head,
		nil,
	}

	head.next = tail

	return LRUCache{
		hashMap: make(map[int]*DulListedNode, capacity),
		head: head,
		tail: tail,
		length: 0,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.hashMap[key]
	if !ok {
		return -1
	}

	this.addToHead(val)

	return val.val
}

func (this *LRUCache) Put(key, value int) {
	val, ok := this.hashMap[key]
	if ok {
		val.val = value
		this.addToHead(val)
		return
	}

	// otherwise
	this.length++
	if this.capacity < this.length {
		this.removeLast()
	}

	new := DulListedNode{
		key: key,
		val: value,
		prev: nil,
		next: nil,
	}
	this.hashMap[key] = &new
	this.addToHead(&new)
}

func (this *LRUCache) addToHead(element *DulListedNode) {
	if element == this.head {
		return
	}

	// element处移除element
	if element.next != nil {
		element.next.prev = element.prev
	}
	if element.prev != nil  {
		element.prev.next = element.next
	}

	// head处

	element.next = this.head.next
	this.head.next.prev = element

	this.head.next = element
	element.prev = this.head

}

func (this *LRUCache) removeLast() {
	elemntToRemove := this.tail.prev
	elemntToRemove.prev.next = this.tail
	this.tail.prev = elemntToRemove.prev

	delete(this.hashMap, elemntToRemove.key)
	this.length--
}


