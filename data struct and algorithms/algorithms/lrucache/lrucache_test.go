package lrucache

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLRUCache1(t *testing.T) {
	lru := NewLRUCache(1)
	lru.Add("A", []byte("A"))
	value := lru.Get("A")
	assert.Equal(t, value, []byte("A"))
}

func TestLRUCache2(t *testing.T) {
	lru := NewLRUCache(1)
	lru.Add("A", []byte("A"))
	lru.Add("B", []byte("B"))
	value := lru.Get("B")
	assert.Equal(t, value, []byte("B"))
}

func TestLRUCache3(t *testing.T) {
	lru := NewLRUCache(1)
	lru.Add("A", []byte("A"))
	lru.Add("B", []byte("B"))
	value := lru.Get("A")
	assert.Equal(t, value == nil, true)
}

func TestLRUCache4(t *testing.T) {
	lru := NewLRUCache(2)
	lru.Add("A", []byte("A"))
	lru.Add("B", []byte("B"))
	lru.Add("C", []byte("C"))
	{
		value := lru.Get("A")
		assert.Equal(t, value == nil, true)
	}
	{
		value := lru.Get("B")
		assert.Equal(t, value, []byte("B"))
	}
	{
		value := lru.Get("C")
		assert.Equal(t, value, []byte("C"))
	}
}

func TestLRUCache5(t *testing.T) {
	lru := NewLRUCache(2)
	lru.Add("A", []byte("A"))
	lru.Add("B", []byte("B"))
	lru.Get("A")
	lru.Add("C", []byte("C"))
	{
		value := lru.Get("B")
		assert.Equal(t, value == nil, true)
	}
	{
		value := lru.Get("A")
		assert.Equal(t, value, []byte("A"))
	}
	{
		value := lru.Get("C")
		assert.Equal(t, value, []byte("C"))
	}
}

func TestLRUCache6(t *testing.T) {
	lru := NewLRUCache(2)
	lru.Add("A", []byte("A"))
	lru.Add("B", []byte("B"))
	lru.Add("C", []byte("C"))
	lru.Add("D", []byte("D"))
	{
		value := lru.Get("B")
		assert.Equal(t, value == nil, true)
	}
	{
		value := lru.Get("D")
		assert.Equal(t, value, []byte("D"))
	}
	{
		value := lru.Get("D")
		assert.Equal(t, value, []byte("D"))
	}
}