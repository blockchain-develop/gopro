package lrucache

type Node struct {
	prev   *Node
	next   *Node
	key    string
	value  []byte
}

type LRUCache struct {
	cache      map[string]*Node
	header     *Node
	tail       *Node
	cap        int
	size       int
}

func NewLRUCache(cap int) *LRUCache {
	lrucache := &LRUCache{
		cache: make(map[string]*Node, 0),
		header: nil,
		tail: nil,
		cap: cap,
		size: 0,
	}
	return lrucache
}

func (self *LRUCache) addNode(key string, value []byte) {
	node := &Node {
		prev: self.tail,
		next: nil,
		key: key,
		value: value,
	}
	self.cache[key] = node
	if self.tail == nil {
		self.header = node
		self.tail = node
	} else {
		self.tail.next = node
		node.prev = self.tail
		self.tail = node
	}
}
func (self *LRUCache) deleteNode() *Node {
	node := self.header
	if node.next == nil {
		self.header = nil
		self.tail = nil
	} else {
		node.next.prev = nil
		self.header = node.next
	}
	node.next = nil
	self.cache[node.key] = nil
	return node
}
func (self *LRUCache) moveNode(node *Node) {
	//
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		self.header = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		self.tail = node.prev
	}
	//
	if self.tail == nil {
		self.header = node
		self.tail = node
	} else {
		self.tail.next = node
		node.prev = self.tail
		self.tail = node
	}
}
func (self *LRUCache) Add(key string, value []byte) {
	if self.size < self.cap {
		self.addNode(key, value)
		self.size ++
	} else {
		self.deleteNode()
		self.addNode(key, value)
	}
}
func (self *LRUCache) Get(key string) []byte {
	node := self.cache[key]
	if node == nil {
		return nil
	} else {
		self.moveNode(node)
		return node.value
	}
}
