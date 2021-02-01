package list

import (
    "testing"
)

// list
var (
    l *List = NewList()
)

// init
func init() {
    values := []int{1, 3, 5, 4, 6, 2, 8, 9, 7}
    for _, value := range values {
        l.Add(value)
    }
    l.PrintList()
}

// TestListAddHead
func TestListAddHead(t *testing.T) {
    l.Add(0)
    l.PrintList()
}

// TestListAddTail
func TestListAddTail(t *testing.T) {
    l.AddTail(11)
    l.PrintList()
}

// TestListQuery
func TestListQuery(t *testing.T) {
    node := l.Query(3)
    node.Print()
}

// TestListDelete
func TestListDelete(t *testing.T) {
    l.Delete(2)
    l.PrintList()
}
