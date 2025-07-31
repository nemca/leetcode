package lru_cache

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToFront(node)
		return node.value
	}

	return -1
}

func (this LRUCache) Put(key, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToFront(node)
		return
	}

	if len(this.cache) == this.capacity {
		least := this.tail.prev
		delete(this.cache, least.key)
		this.remove(least)
	}
	node := &Node{key: key, value: value}
	this.cache[key] = node
	this.addToFront(node)
}

func (this LRUCache) moveToFront(node *Node) {
	this.remove(node)
	this.addToFront(node)
}

func (this LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this LRUCache) addToFront(node *Node) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}
