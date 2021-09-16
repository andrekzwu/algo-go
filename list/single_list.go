package list

import (
    "fmt"
)

type SingleListNode struct {
    value int
    next  *SingleListNode
}

// initSingleList
func initSingleList(values []int) *SingleListNode {
    var head, p *SingleListNode

    for _, value := range values {
        s := &SingleListNode{
            value: value,
            next:  nil,
        }
        if head == nil {
            head = s
        } else {
            p.next = s
        }
        p = s
    }
    return head
}

// reversSingleList
func reversSingleList(head *SingleListNode) *SingleListNode {
    if head == nil || head.next == nil {
        return head
    }
    var p, n *SingleListNode
    s := head
    for s != nil {
        n = s.next
        s.next = p
        p = s
        s = n
    }
    return p
}

// printSimpleList
func printSimpleList(head *SingleListNode) {
    p := head
    for p != nil {
        fmt.Printf("%d ", p.value)
        p = p.next
    }
    fmt.Println("-------------")
    return
}
