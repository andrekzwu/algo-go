package leetcodev2

import (
	"fmt"
	"testing"
)

// TestCQueue offer[09] 单元测试
func TestCQueue(t *testing.T) {
	obj := Constructor()
	obj.AppendTail(1)
	param_2 := obj.DeleteHead()
	fmt.Println("TestCQueue :", param_2)
}

// TestMinStack offer[30] 单元测试
func TestMinStack(t *testing.T) {
	obj := MinStackConstructor()
	obj.Push(1)
	obj.Push(5)
	obj.Push(4)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.Min()
	fmt.Println("TestMinStack :", param_3, param_4)
}

// TestReversePrint offer[06] 单元测试
func TestReversePrint(t *testing.T) {
	head := ListNode{Val: 1}
	l1 := ListNode{Val: 2}
	head.Next = &l1
	l2 := ListNode{Val: 3}
	l1.Next = &l2
	//
	s := reversePrint(&head)
	fmt.Println(s)
}

// TestCopyRandomList [35] 复杂链表的复制
func TestCopyRandomList(t *testing.T) {
	h := &Node{Val: 1}
	l1 := &Node{Val: 2}
	h.Next = l1
	l1.Random = h
	//
	l2 := &Node{Val: 3}
	l1.Next = l2
	h.Random = l2
	l2.Random = h
	//
	nh := copyRandomList(h)
	for nh != nil {
		fmt.Printf("v:%d,n:%d,r:%d\n", nh.val(), nh.Next.val(), nh.Random.val())
		nh = nh.Next
	}
}

// TestReplaceSpace [05] 替换空格
func TestReplaceSpace(t *testing.T) {
	fmt.Println(replaceSpace("a n cd"))
}
