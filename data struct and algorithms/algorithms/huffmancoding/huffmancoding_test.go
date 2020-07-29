package huffmancoding

import (
	"fmt"
	"testing"
)
func TestHuffmanCoding1(t *testing.T) {
	characters := []byte{byte('a'), byte('b'), byte('c'), byte('d'), byte('e'), byte('f')}
	frequencys := []int{5, 9, 12, 13, 16, 45}
	codings := HuffmanCoding(characters, frequencys)
	for key, value := range codings {
		fmt.Printf("%c	%s\n", key, string(value))
	}
}
