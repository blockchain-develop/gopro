package baidu

type Node struct {
	next  *Node
	data int
}

func NewNode(data int) *Node {
	node := &Node{
		next: nil,
		data: data,
	}
	return node
}

func Merge(firstList *Node, secondList *Node) *Node {
	firstCurrent := firstList
	secondCurrent := secondList
	merged := firstCurrent
	if firstCurrent.data <= secondCurrent.data {
		merged = firstCurrent
		firstCurrent = firstCurrent.next
	} else {
		merged = secondCurrent
		secondCurrent = secondCurrent.next
	}
	for firstCurrent != nil && secondCurrent != nil {
		if firstCurrent.data <= secondCurrent.data {
			merged.next = firstCurrent
			firstCurrent = firstCurrent.next
		} else {
			merged.next = secondCurrent
			secondCurrent = secondCurrent.next
		}
	}
	if firstCurrent != nil {
		merged.next = firstCurrent
	} else {
		merged.next = secondCurrent
	}
	return merged
}
