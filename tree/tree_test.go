// Copyright 2020, andrezhz. All rights reserved.
// Use of this source code is governed by a BSD-style
// licego nse that can be found in the LICENSE file.go
package tree

import (
    "fmt"
    "testing"
)

func TestLevelOrderQ(t *testing.T) {
    fmt.Println("-------level orderQ------")
    rootTreeNode := InitTree()
    LevelOrder([]*TreeNode{rootTreeNode})
}

func TestPreOrder(t *testing.T) {
    fmt.Println("-------pre order------")
    rootTreeNode := InitTree()
    PreOrder(rootTreeNode)
    fmt.Println("")
}

func TestInOrder(t *testing.T) {
    fmt.Println("-------in order------")
    rootTreeNode := InitTree()
    InOrder(rootTreeNode)
    fmt.Println("")
}

func TestPostOrder(t *testing.T) {
    fmt.Println("-------post order------")
    rootTreeNode := InitTree()
    PostOrder(rootTreeNode)
    fmt.Println("")
}

func TestInTreading(t *testing.T) {
    fmt.Println("-------in threading order------")
    rootTreeNode := InitTree()
    InThreading(rootTreeNode)
    fmt.Println("-------in threading order print----")
    InThreadingPrint(rootTreeNode)
    fmt.Println("")
}
