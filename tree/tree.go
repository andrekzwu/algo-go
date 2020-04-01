// Copyright 2020, andrezhz. All rights reserved.
// Use of this source code is governed by a BSD-style
// licego nse that can be found in the LICENSE file.go
package tree

import (
    "fmt"
)

type CLUE_TAG byte

var (
    LINK   CLUE_TAG = 0
    THREAD CLUE_TAG = 1
)

type TreeNode struct {
    Value  string
    LChild *TreeNode
    RChild *TreeNode
    LTag   CLUE_TAG
    RTag   CLUE_TAG
}

func (tn *TreeNode) GetValue() string {
    if tn == nil {
        return ""
    }
    return tn.Value
}

func (tn *TreeNode) GetLChild() *TreeNode {
    if tn == nil || tn.LChild == nil {
        return nil
    }
    return tn.LChild
}

func (tn *TreeNode) GetRChild() *TreeNode {
    if tn == nil || tn.RChild == nil {
        return nil
    }
    return tn.RChild
}

func (tn *TreeNode) SetLChild(lTag CLUE_TAG, node *TreeNode) {
    if tn == nil {
        return
    }
    tn.LTag = lTag
    tn.LChild = node
}

func (tn *TreeNode) SetRChild(rTag CLUE_TAG, node *TreeNode) {
    if tn == nil {
        return
    }
    tn.RTag = rTag
    tn.RChild = node
}

func (tn *TreeNode) GetLTag() CLUE_TAG {
    if tn == nil {
        return 0
    }
    return tn.LTag
}

func (tn *TreeNode) GetRTag() CLUE_TAG {
    if tn == nil {
        return 0
    }
    return tn.RTag
}

func (tn *TreeNode) SetLTag(lTag CLUE_TAG) {
    if tn == nil {
        return
    }
    tn.LTag = lTag
}

func (tn *TreeNode) SetRTag(rTag CLUE_TAG) {
    if tn == nil {
        return
    }
    tn.RTag = rTag
}

func NewTreeNode(value string) *TreeNode {
    return &TreeNode{Value: value}
}

func InitTree() *TreeNode {
    // root level 1
    rootNode := NewTreeNode("A")
    // level 2
    nodeB := NewTreeNode("B")
    nodeC := NewTreeNode("C")
    rootNode.LChild, rootNode.RChild = nodeB, nodeC
    // level 3
    nodeD := NewTreeNode("D")
    nodeE := NewTreeNode("E")
    nodeF := NewTreeNode("F")
    nodeG := NewTreeNode("G")
    nodeB.LChild, nodeB.RChild = nodeD, nodeE
    nodeC.LChild, nodeC.RChild = nodeF, nodeG
    // level 4
    nodeH := NewTreeNode("H")
    nodeI := NewTreeNode("I")
    nodeJ := NewTreeNode("J")
    nodeD.LChild, nodeD.RChild = nodeH, nodeI
    nodeE.LChild = nodeJ
    return rootNode
}

// Pre
func PreOrder(node *TreeNode) {
    if node == nil {
        return
    }
    fmt.Printf(node.GetValue())
    // left child tree
    PreOrder(node.GetLChild())
    // right child tree
    PreOrder(node.GetRChild())
}

// InOrder
func InOrder(node *TreeNode) {
    if node == nil {
        return
    }
    // left child tree
    InOrder(node.GetLChild())
    fmt.Printf(node.GetValue())
    // right child tree
    InOrder(node.GetRChild())
}

// PostOrder
func PostOrder(node *TreeNode) {
    if node == nil {
        return
    }
    // left child tree
    PostOrder(node.GetLChild())
    // right child tree
    PostOrder(node.GetRChild())
    fmt.Printf(node.GetValue())
}

// LevelOrder
func LevelOrder(quques []*TreeNode) {
    if len(quques) == 0 {
        return
    }
    subQuques := make([]*TreeNode, 0)
    for _, node := range quques {
        fmt.Printf(node.GetValue())
        if node.GetLChild() != nil {
            subQuques = append(subQuques, node.GetLChild())
        }
        if node.GetRChild() != nil {
            subQuques = append(subQuques, node.GetRChild())
        }
    }
    fmt.Println("")
    LevelOrder(subQuques)
}

var (
    preTreeNode *TreeNode
)

// InThreading
func InThreading(node *TreeNode) {
    if node == nil {
        return
    }
    // left child tree thread
    InThreading(node.GetLChild())

    if node.GetLChild() != nil {
        node.SetLTag(LINK)
    } else {
        node.SetLChild(THREAD, preTreeNode)
    }

    if preTreeNode.GetRChild() != nil {
        preTreeNode.SetRTag(LINK)
    } else {
        preTreeNode.SetRChild(THREAD, node)
    }
    preTreeNode = node
    InThreading(node.GetRChild())
}

// InThreadingPrint
func InThreadingPrint(node *TreeNode) {
    if node == nil {
        return
    }
    if node.GetLTag() == LINK {
        InThreadingPrint(node.GetLChild())
    }

    var pre, post string
    if node.GetLTag() == THREAD {
        pre = node.GetLChild().GetValue()
    }
    if node.GetRTag() == THREAD {
        post = node.GetRChild().GetValue()
    }
    fmt.Printf("%s,{%s,%s}\n", node.GetValue(), pre, post)
    if node.GetRTag() == LINK {
        InThreadingPrint(node.GetRChild())
    }
}
