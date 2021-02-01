package tree

import (
    "testing"
)

var (
    rbt *RBTree = &RBTree{}
)

// init
func init() {
    values := []int{5, 1, 2, 4, 8, 3}
    for _, value := range values {
        rbt.Insert(value)
    }
}

// TestRBTreeSort
func TestRBTreeSort(t *testing.T) {
    // 前序遍历
    rbt.PreOrder()
    // 中序遍历
    rbt.InOrder()
    // 后序遍历
    rbt.PostOrder()
    // 层级遍历
    rbt.LevelOrder()
}

// TestRBTreeQuery
func TestRBTreeQuery(t *testing.T) {
    rbNode, ok := rbt.Query(2)
    if ok {
        rbNode.printNode()
    }
}

// TestRBTreeInsert
func TestRBTreeInsert(t *testing.T) {
    rbt.Insert(9)
    //
    rbt.InOrder()
    rbt.LevelOrder()
}

// TestRBTreeDelete
func TestRBTreeDelete(t *testing.T) {
    rbt.Delete(2)
    //
    rbt.InOrder()
    rbt.LevelOrder()
}
