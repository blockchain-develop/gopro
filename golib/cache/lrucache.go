package cache

import "container/list"

type LruCache struct {
	l        *list.List
	m        map[string]*list.Element
	capacity int
}

func NewLruCache(cap int) *LruCache {
	lc := &LruCache{
		l:        list.New(),
		m:        make(map[string]*list.Element),
		capacity: cap,
	}
	return lc
}

func (lc *LruCache) Set(k string, v interface{}) {
	if ele, hit := lc.m[k]; hit {
		lc.l.MoveToBack(ele)
		ele.Value = v
		return
	}
	ele := lc.l.PushBack(v)
	lc.m[k] = ele
	if lc.l.Len() > lc.capacity {
		lc.l.Remove(lc.l.Front())
	}
}

func (lc *LruCache) Get(k string) (interface{}, bool) {
	if ele, hit := lc.m[k]; hit {
		lc.l.MoveToBack(ele)
		return ele.Value, true
	}
	return nil, true
}
