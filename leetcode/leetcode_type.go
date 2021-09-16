package leetcode

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func (n *Node) GetVal() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("%d", n.Val)
}

// printNode
func printNode(head *Node) {
	fmt.Println("")
	s := head
	for s != nil {
		fmt.Printf("[%s,%s],", s.GetVal(), s.Random.GetVal())
		s = s.Next
	}
}

func (n *ListNode) GetVal() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("%d", n.Val)
}

// printListNode
func printListNode(head *ListNode) {
	fmt.Println("")
	s := head
	for s != nil {
		fmt.Printf("%s,", s.GetVal())
		s = s.Next
	}
	fmt.Println("")
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
