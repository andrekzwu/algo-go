package leetcode

import (
	"andrezhz/algo-go/list"
	"sync"
	"unsafe"
)

type LRUCALLBACK func(key, value interface{})
type KEY interface{}

// LRUCache lru cache
type LRUCache struct {
	MaxEntries int
	// callback
	OnEvicted LRUCALLBACK
	//
	cache_head *list.List_head
	cache      map[interface{}]*list.List_head
	//
	mutex sync.Mutex
}

// CacheEntry
type CacheEntry struct {
	node  list.List_head
	key   interface{}
	value interface{}
}

var (
	offSetCacheEntry         = CacheEntry{}
	CacheEntryOffSet uintptr = uintptr(unsafe.Pointer(&offSetCacheEntry.node)) - uintptr(unsafe.Pointer(&offSetCacheEntry))
)

// ListHead2CacheEntry
func ListHead2CacheEntry(pNode *list.List_head) *CacheEntry {
	return (*CacheEntry)(unsafe.Pointer(uintptr(unsafe.Pointer(pNode)) - CacheEntryOffSet))
}

// NewCacheEntry
func NewCacheEntry(key, value interface{}) *CacheEntry {
	return &CacheEntry{
		key:   key,
		value: value,
		node:  list.List_head{},
	}
}

// NewLRUCache
func NewLRUCache(maxEntries int, callback LRUCALLBACK) *LRUCache {
	return &LRUCache{
		MaxEntries: maxEntries,
		OnEvicted:  callback,
		cache_head: list.List_head_init(),
		cache:      make(map[interface{}]*list.List_head, 0),
		mutex:      sync.Mutex{},
	}
}

// Move2Front
func (lruc *LRUCache) Move2Front(pNode *list.List_head) {
	list.List_del(pNode)
	list.List_add_head(pNode, lruc.cache_head)
}

// Push2Front
func (lruc *LRUCache) Push2Front(pNode *list.List_head) {
	list.List_add_head(pNode, lruc.cache_head)
}

// RemoveOldest
func (lruc *LRUCache) RemoveOldest() {
	if lruc.cache == nil || lruc.cache_head == nil {
		return
	}
	pNode := list.List_entry_last(lruc.cache_head)
	if pNode != nil {
		lruc.removeNode(pNode)
	}
}

// removeNode
func (lruc *LRUCache) removeNode(pNode *list.List_head) {
	if lruc.cache == nil || lruc.cache_head == nil || pNode == nil {
		return
	}
	list.List_del(pNode)
	cacheEntry := ListHead2CacheEntry(pNode)
	delete(lruc.cache, cacheEntry.key)
	if lruc.OnEvicted != nil {
		lruc.OnEvicted(cacheEntry.key, cacheEntry.value)
	}
}

// Remove
func (lruc *LRUCache) Remove(key interface{}) {
	if lruc.cache == nil || lruc.cache_head == nil {
		return
	}
	pNode, hit := lruc.cache[key]
	if !hit {
		return
	}
	lruc.removeNode(pNode)
}

// Add
func (lruc *LRUCache) Add(key, value interface{}) {
	if lruc == nil || lruc.cache == nil || lruc.cache_head == nil {
		return
	}
	lruc.mutex.Lock()
	defer lruc.mutex.Unlock()
	// hit
	if pNode, ok := lruc.cache[key]; ok {
		cacheEntry := ListHead2CacheEntry(pNode)
		lruc.Move2Front(pNode)
		cacheEntry.value = value
		return
	}
	// add
	cacheEntry := NewCacheEntry(key, value)
	lruc.cache[key] = &cacheEntry.node
	lruc.Push2Front(&cacheEntry.node)
	if lruc.MaxEntries != 0 && len(lruc.cache) > lruc.MaxEntries {
		lruc.RemoveOldest()
	}
}

// Get
func (lruc *LRUCache) Get(key interface{}) (*CacheEntry, bool) {
	if lruc.cache == nil || lruc.cache_head == nil {
		return nil, false
	}
	pNode, hit := lruc.cache[key]
	if !hit {
		return nil, false
	}
	cacheEntry := ListHead2CacheEntry(pNode)
	lruc.Move2Front(pNode)
	return cacheEntry, true
}

// Len
func (lruc *LRUCache) Len() int {
	if lruc.cache == nil || lruc.cache_head == nil {
		return 0
	}
	return len(lruc.cache)
}

// Clear
func (lruc *LRUCache) Clear() {
	if lruc == nil || lruc.cache == nil || lruc.cache_head == nil {
		return
	}
	if lruc.OnEvicted != nil {
		for pNode := list.List_entry_last(lruc.cache_head); pNode != nil; {
			cacheEntry := ListHead2CacheEntry(pNode)
			lruc.OnEvicted(cacheEntry.key, cacheEntry.value)
			lruc.removeNode(pNode)
		}
	}
	lruc.cache = nil
	lruc.cache_head = nil
}
