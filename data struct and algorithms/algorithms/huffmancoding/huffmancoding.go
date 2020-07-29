package huffmancoding

type HuffmanNode struct {
	character byte
	frequency  int
	left *HuffmanNode
	right *HuffmanNode
}

func adjustdown(nodes []*HuffmanNode, index int, size int) {
	if index > size/2 -1 {
		return
	}
	child := index * 2 + 1
	if child + 1 <= size - 1 && nodes[child + 1].frequency < nodes[child].frequency {
		child ++
	}
	if nodes[index].frequency <= nodes[child].frequency {
		return
	}
	key := nodes[index]
	nodes[index] = nodes[child]
	nodes[child] = key

	adjustdown(nodes, child, size)
}

func traverse(node *HuffmanNode, code []byte, result map[byte][]byte) {
	if node.left == nil && node.right == nil {
		result[node.character] = make([]byte,len(code))
		copy(result[node.character], code)
		code = code[: len(code) - 1]
		return
	}
	if node.left != nil {
		code = append(code, byte('0'))
		traverse(node.left, code, result)
	}
	code = code[: len(code) - 1]
	if node.right != nil {
		code = append(code, byte('1'))
		traverse(node.right, code, result)
	}
	code = code[: len(code) - 1]
}

func HuffmanCoding(charaters []byte, frequencys []int) (map[byte][]byte) {
	size := len(charaters)
	// build minheap
	huffmanNodes := make([]*HuffmanNode, 0)
	for i := 0;i < size;i ++ {
		huffmanNodes = append(huffmanNodes, &HuffmanNode{
			character: charaters[i],
			frequency: frequencys[i],
			left: nil,
			right: nil,
		})
	}
	for i := size/2 - 1;i >= 0;i -- {
		adjustdown(huffmanNodes, i, size)
	}
	// build huffman tree
	size = len(huffmanNodes)
	for size > 1 {
		first := huffmanNodes[0]
		huffmanNodes[0] = huffmanNodes[size - 1]
		size --
		adjustdown(huffmanNodes, 0, size)
		second := huffmanNodes[0]
		huffmanNodes[0] = &HuffmanNode{
			character: 0,
			frequency: first.frequency + second.frequency,
			left: first,
			right: second,
		}
		adjustdown(huffmanNodes, 0, size)
	}
	root := huffmanNodes[0]
	// coding
	codings := make(map[byte][]byte, 0)
	code := make([]byte, 0)
	traverse(root, code, codings)
	//
	return codings
}
