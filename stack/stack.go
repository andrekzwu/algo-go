package stack

import (
    "andrezhz/algo-go/list"
    "sync"
)

type Stacker interface {
    Empty() bool
    Size() int
    Push(node *list.List_head)
    Pop() *list.List_head
    Top() *list.List_head
}

type StackImpl struct {
    size       int
    stack_head *list.List_head
    mutex      sync.Mutex
}

func StackImplInit() *StackImpl {
    return &StackImpl{
        stack_head: list.List_head_init(),
        mutex:      sync.Mutex{},
    }
}

// Empty
func (si *StackImpl) Empty() bool {
    if si == nil || si.stack_head == nil {
        return true
    }
    return list.List_empty(si.stack_head)
}

// Size
func (si *StackImpl) Size() int {
    if si == nil || si.stack_head == nil {
        return 0
    }
    return si.size
}

// Push
func (si *StackImpl) Push(node *list.List_head) {
    if si == nil || si.stack_head == nil {
        return
    }
    si.mutex.Lock()
    defer si.mutex.Unlock()
    //
    si.size++
    list.List_add_tail(node, si.stack_head)
}

// Pop
func (si *StackImpl) Pop() *list.List_head {
    if si == nil || si.stack_head == nil {
        return nil
    }
    si.mutex.Lock()
    defer si.mutex.Unlock()
    //
    if list.List_empty(si.stack_head) {
        return nil
    }
    si.size--
    return list.List_entry_last(si.stack_head)
}

// Pop
func (si *StackImpl) Top() *list.List_head {
    if si == nil || si.stack_head == nil {
        return nil
    }
    if list.List_empty(si.stack_head) {
        return nil
    }
    return list.List_last(si.stack_head)
}
