package list

import (
    "fmt"
)

type ListNode struct {
    value int
    prev  *ListNode
    next  *ListNode
}

// GetValue
func (ln *ListNode) GetValue() int {
    if ln == nil {
        return 0
    }
    return ln.value
}

// Print
func (ln *ListNode) Print() {
    fmt.Printf("%d_%d_%d \n", ln.prev.GetValue(), ln.GetValue(), ln.next.GetValue())
}

// List
type List struct {
    head *ListNode
}

// NewList
func NewList() *List {
    l := &List{}
    l.head = &ListNode{}
    l.head.next = l.head
    l.head.prev = l.head
    return l
}

// Add
func (l *List) Add(value int) {
    if l.head == nil {
        return
    }
    l.add(l.head, l.head.next, value)
}

// add
func (l *List) add(prev, next *ListNode, value int) {
    newNode := &ListNode{value: value}
    //
    next.prev = newNode
    newNode.next = next
    //
    newNode.prev = prev
    prev.next = newNode
}

// AddTail
func (l *List) AddTail(value int) {
    if l.head == nil {
        return
    }
    l.add(l.head.prev, l.head, value)
}

// Delete
func (l *List) Delete(value int) {
    if l.head == nil {
        return
    }
    node := l.Query(value)
    if node == nil {
        return
    }
    l.delete(node.prev, node.next)
    node.next = nil
    node.prev = nil
}

// delete
func (l *List) delete(prev, next *ListNode) {
    prev.next = next
    next.prev = prev
}

// Query
func (l *List) Query(value int) *ListNode {
    s := l.head.next
    for l.head != s {
        if s.value == value {
            return s
        }
        s = s.next
    }
    return nil
}

// PrintList
func (l *List) PrintList() {
    fmt.Println("-------------打印列表")
    if l.head == nil {
        return
    }
    s := l.head.next
    for l.head != s {
        fmt.Printf("%d ", s.value)
        s = s.next
    }
    fmt.Println("")
}
