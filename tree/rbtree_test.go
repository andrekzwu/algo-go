package tree

import (
    "fmt"
    "testing"
)

// TestRBTree
func TestRBTree(t *testing.T) {
    values := []int{5, 1, 2}
    rbt := &RBTree{}
    for _, value := range values {
        rbt.Insert(value)
    }
    rbt.InOrder()
    rbt.LevelOrder()
    //
    rbNode, ok := rbt.Query(2)
    if ok {
        fmt.Println("query node value is:", rbNode.value)
        // query successor
        successorNode := rbNode.GetSuccessor()
        if successorNode != nil {
            fmt.Printf("-----------%d successor is %d\n", rbNode.value, successorNode.value)
        }
    }

}

/*
1. 删除节点如果没有子节点，则直接删除
2. 删除节点如果只有一个节点，则用该节点进行替换，然后考虑变色与平衡
3. 删除节点如果有两个节点，则使用期后继节点进行替换，（相当于删除后继节点），然后考虑平衡
4. 所以最终问题，都变成了删除单个节点的情况






*/
