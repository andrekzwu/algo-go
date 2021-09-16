package list

import (
    "testing"
)

// TestReversSimpleList
func TestReversSimpleList(t *testing.T) {
    values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    head := initSingleList(values)
    printSimpleList(head)
    newHead := reversSingleList(head)
    printSimpleList(newHead)
}
