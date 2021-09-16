package queue

import (
    "andrezhz/algo-go/list"
    "sync"
)

type Queuer interface {
    InitQueue()
    EnQueue(node *list.List_head)
    DlQueue() *list.List_head
    GetHead() *list.List_head
    Empty() bool
    Size() int
}

type QueueImpl struct {
    size       int
    queue_head *list.List_head
    mutex      sync.Mutex
}

// NewQueueImpl
func NewQueueImpl() *QueueImpl {
    return &QueueImpl{
        queue_head: list.List_head_init(),
        mutex:      sync.Mutex{},
    }
}

// InitQueue
func (qi *QueueImpl) InitQueue() {
    qi.size = 0
    list.Init_list_head(qi.queue_head)
}

// EnQueue
func (qi *QueueImpl) EnQueue(node *list.List_head) {
    if qi == nil || qi.queue_head == nil {
        return
    }
    qi.mutex.Lock()
    defer qi.mutex.Unlock()
    //
    qi.size++
    list.List_add_tail(node, qi.queue_head)
}

// DlQueue
func (qi *QueueImpl) DlQueue() *list.List_head {
    if qi == nil || qi.queue_head == nil {
        return nil
    }
    qi.mutex.Lock()
    defer qi.mutex.Unlock()
    //
    if list.List_empty(qi.queue_head) {
        return nil
    }
    qi.size--
    return list.List_entry_first(qi.queue_head)
}

// GetHead()
func (qi *QueueImpl) GetHead() *list.List_head {
    if qi == nil || qi.queue_head == nil {
        return nil
    }
    qi.mutex.Lock()
    defer qi.mutex.Unlock()
    //
    if list.List_empty(qi.queue_head) {
        return nil
    }
    return list.List_first(qi.queue_head)
}

//  Empty()
func (qi *QueueImpl) Empty() bool {
    if qi == nil || qi.queue_head == nil {
        return true
    }
    return list.List_empty(qi.queue_head)
}

// Size
func (qi *QueueImpl) Size() int {
    if qi == nil || qi.queue_head == nil {
        return 0
    }
    return qi.size
}
