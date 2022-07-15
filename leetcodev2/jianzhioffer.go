package leetcodev2

import (
	"unsafe"
)

// 09. 用两个栈实现队列
type CQueue struct {
	instack  []int
	outstack []int
}

func Constructor() CQueue {
	return CQueue{instack: make([]int, 0), outstack: make([]int, 0)}

}

func (this *CQueue) AppendTail(value int) {
	this.instack = append(this.instack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.outstack) == 0 {
		if len(this.instack) == 0 {
			return -1
		}
		for len(this.instack) > 0 {
			this.outstack = append(this.outstack, this.instack[len(this.instack)-1])
			this.instack = this.instack[0 : len(this.instack)-1]
		}
	}
	v := this.outstack[len(this.outstack)-1]
	this.outstack = this.outstack[:len(this.outstack)-1]
	return v
}

// 30. 包含min函数的栈
type MinStack struct {
	stack    []int
	minstack []int
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{stack: make([]int, 0), minstack: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minstack) == 0 {
		this.minstack = append(this.minstack, x)
	} else {
		this.minstack = append(this.minstack, min(this.minstack[len(this.minstack)-1], x))
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[0 : len(this.stack)-1]
	this.minstack = this.minstack[0 : len(this.minstack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	if len(this.minstack) == 0 {
		return 0
	}
	min := this.minstack[len(this.minstack)-1]
	return min
}

// [06] 从尾到头打印链表
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	stack := make([]int, 0)
	s := head
	for s.Next != nil {
		stack = append([]int{s.Val}, stack...)
		s = s.Next
	}
	stack = append([]int{s.Val}, stack...)
	return stack
}

// [24] 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var p, r *ListNode
	s := head
	for s != nil {
		p = s
		s = s.Next
		p.Next = r
		r = p
	}
	return r
}

// [35] 复杂链表的复制
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	mapPos := make(map[uintptr]int)
	mapIndex := make(map[int]int)
	mapnIndex := make(map[int]*Node)
	var nh, s *Node
	r := head
	i := 0
	for r != nil {
		m := &Node{Val: r.Val}
		if s == nil {
			s = m
		} else {
			s.Next = m
			s = m
		}
		if nh == nil {
			nh = s
		}
		mapnIndex[i] = s
		mapPos[uintptr(unsafe.Pointer(r))] = i
		//
		i++
		r = r.Next
	}
	r = head
	j := 0
	for r != nil {
		mapIndex[j] = -1
		if index, ok := mapPos[uintptr(unsafe.Pointer(r.Random))]; ok {
			mapIndex[j] = index
		}
		j++
		r = r.Next
	}
	//
	s = nh
	x := 0
	for s != nil {
		rindex := mapIndex[x]
		s.Random = mapnIndex[rindex]
		//
		x++
		s = s.Next
	}
	return nh
}

// [05] 替换空格
func replaceSpace(s string) string {
	if len(s) == 0 {
		return ""
	}
	bytes := make([]byte, 0)
	for _, ch := range []byte(s) {
		if ch == ' ' {
			bytes = append(bytes, []byte("%20")...)
			continue
		}
		bytes = append(bytes, ch)
	}
	return string(bytes)
}

// [58] 左旋字符串
func reverseLeftWords(s string, n int) string {
	bytes := []byte(s)[n:]
	bytes = append(bytes, []byte(s)[:n])
	return string(bytes)
}

// [03] 数组中重复的数字
func findRepeatNumber(nums []int) int {
	mapRepeated := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := mapRepeated[num]; ok {
			return num
		}
		mapRepeated[num] = struct{}{}
	}
	return -1
}

// [08] 在排序数组中查找数字l
func search(nums []int, target int) int {
	isStart := false
	times := 0
	for _, num := range nums {
		if num == target {
			isStart = true
			times++

		} else if isStart {
			return times
		}
	}
	return 0
}

// [53] 0~n-1中缺失的数字
func missingNumber(nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			return nums[i] - 1
		}
	}
	if nums[0] > 0 {
		return 0
	}
	if nums[len(nums)-1] < len(nums) {
		return len(nums)
	}
	return -1
}

// [04] 二维数组中的查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	n := len(matrix[0]) - 1
	m := 0
	for m < len(matrix) && n >= 0 {
		if matrix[m][n] == target {
			return true
		}
		if target < matrix[m][n] {
			n--
		} else {
			m++
		}
	}
	return false
}

// [11] 旋转数组的最小数字
func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	low := 0
	hight := len(numbers) - 1
	for low < hight {
		prev := low + (hight-low)/2
		if numbers[prev] > numbers[hight] {
			low = prev + 1
		} else if numbers[prev] < numbers[hight] {
			hight = prev
		} else {
			hight--
		}
	}
	return numbers[low]
}
