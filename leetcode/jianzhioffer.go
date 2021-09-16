package leetcode

import (
	"fmt"
	"sort"
	"strconv"
)

// -------offer 09
// 用两个栈实现队列
type CQueue struct {
	values []int
}

func Constructor() CQueue {
	return CQueue{
		values: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.values = append(this.values, value)

}

func (this *CQueue) DeleteHead() int {
	if len(this.values) == 0 {
		return -1
	}
	value := this.values[0]
	this.values = this.values[1:]
	return value
}

// ------ooffer 30
type MinStack struct {
	stacks    []int
	minstacks []int
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{stacks: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.stacks = append(this.stacks, x)
	//
	if len(this.minstacks) == 0 || this.minstacks[len(this.minstacks)-1] >= x {
		this.minstacks = append(this.minstacks, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.stacks) == 0 {
		return
	}
	v := this.stacks[len(this.stacks)-1]
	if len(this.minstacks) > 0 && this.minstacks[len(this.minstacks)-1] == v {
		this.minstacks = this.minstacks[:len(this.minstacks)-1]
	}
	this.stacks = this.stacks[:len(this.stacks)-1]

}

func (this *MinStack) Top() int {
	if len(this.stacks) == 0 {
		return 0
	}
	return this.stacks[len(this.stacks)-1]
}

func (this *MinStack) Min() int {
	if len(this.stacks) == 0 || len(this.minstacks) == 0 {
		return 0
	}
	return this.minstacks[len(this.minstacks)-1]
}

// findRepeatNumber ---03
func findRepeatNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return 0
}

// ---offer-04
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	i := len(matrix[0]) - 1
	j := 0
	//
	for i >= 0 && j >= 0 && j < len(matrix) && i < len(matrix[j]) {
		if matrix[j][i] == target {
			return true
		} else if matrix[j][i] > target {
			i--
		} else {
			j++
		}
	}
	return false
}

// ----offer-05
func replaceSpace(s string) string {
	if s == "" {
		return ""
	}
	n := make([]byte, 0)
	space := "%20"
	//
	for _, c := range []byte(s) {
		if c == ' ' {
			n = append(n, []byte(space)...)
		} else {
			n = append(n, c)
		}
	}
	return string(n)
}

// ---offer-07
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	stack := make([]int, 0)
	s := head
	for s != nil {
		stack = append(stack, s.Val)
		s = s.Next
	}
	//
	rts := make([]int, 0)
	for i := len(stack) - 1; i >= 0; i-- {
		rts = append(rts, stack[i])
	}
	return rts
}

// offer-08
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			n++
		}
	}
	return n
}

// offer 10-i
func fib(n int) int {
	f := make(map[int]int, 0)
	f[0] = 0
	f[1] = 1
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
		f[i] = f[i] % 1000000007
	}
	if f[n] == 1000000008 {
		return 1
	}
	return int(f[n] % 1000000007)
}

// offer-10-2
func numWays(n int) int {
	f := make(map[int]int, 0)
	f[0] = 1
	f[1] = 1
	if n <= 0 {
		return 1
	}
	if n <= 1 {
		return 1
	}
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
		f[i] = f[i] % 1000000007
	}
	return f[n]
}

// offer-11
func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	for i := 0; i < len(numbers); i++ {
		if i > 0 && numbers[i] < numbers[i-1] {
			return numbers[i]
		}
	}
	return numbers[0]
}

// offer-24
func reverseListoffer24(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var p *ListNode = nil
	s := head
	for s != nil {
		n := s.Next
		s.Next = p
		//
		p = s
		s = n
	}
	return p
}

// isSubStructure ---offer-26
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	var issub func(a, b *TreeNode) bool
	issub = func(a, b *TreeNode) bool {
		if b == nil {
			return true
		}
		if a == nil {
			return false
		}
		return a.Val == b.Val && issub(a.Left, b.Left) && issub(a.Right, b.Right)
	}

	var dfs func(a, b *TreeNode) bool
	dfs = func(a, b *TreeNode) bool {
		if a == nil {
			return false
		}
		if issub(a, b) {
			return true
		}
		return dfs(a.Left, b) || dfs(a.Right, b)
	}
	if B == nil {
		return false
	}
	return dfs(A, B)
}

// mirrorTree ---offer 27
func mirrorTree(root *TreeNode) *TreeNode {
	var mirror func(r *TreeNode)
	mirror = func(r *TreeNode) {
		if r == nil {
			return
		}
		r.Left, r.Right = r.Right, r.Left
		//
		mirror(r.Left)
		mirror(r.Right)

	}
	mirror(root)
	return root
}

// ---OFFER 28
func isSymmetric(root *TreeNode) bool {
	var symmetric func(left, right *TreeNode) bool
	symmetric = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
	}
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

// offer---32
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderorder32I(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	rlts := make([]int, 0)
	//
	queue = append(queue, root)
	for len(queue) != 0 {
		r := queue[0]
		queue = queue[1:]
		//
		rlts = append(rlts, r.Val)
		if r.Left != nil {
			queue = append(queue, r.Left)
		}
		if r.Right != nil {
			queue = append(queue, r.Right)
		}
	}
	return rlts
}

// offer-32 II
func levelOrderoffer32II(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	rlts := make([][]int, 0)
	var dfs func([]*TreeNode)
	dfs = func(list []*TreeNode) {
		if len(list) == 0 {
			return
		}
		tarr := make([]*TreeNode, 0)
		arr := make([]int, 0)
		for _, v := range list {
			arr = append(arr, v.Val)
			if v.Left != nil {
				tarr = append(tarr, v.Left)
			}
			if v.Right != nil {
				tarr = append(tarr, v.Right)
			}
		}
		rlts = append(rlts, arr)
		dfs(tarr)
	}
	dfs([]*TreeNode{root})
	return rlts
}

// offer-32 III
func levelOrderoffer32III(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	rlts := make([][]int, 0)
	var dfs func(level int, list []*TreeNode)
	dfs = func(level int, list []*TreeNode) {
		if len(list) == 0 {
			return
		}
		tarr := make([]*TreeNode, 0)
		arr := make([]int, 0)
		for _, v := range list {
			arr = append(arr, v.Val)
			if v.Left != nil {
				tarr = append(tarr, v.Left)
			}
			if v.Right != nil {
				tarr = append(tarr, v.Right)
			}
		}
		//
		if level%2 == 0 {
			rlts = append(rlts, arr)
		} else {

			l, r := 0, len(arr)-1
			for l < r {
				arr[l], arr[r] = arr[r], arr[l]
				l++
				r--
			}
			rlts = append(rlts, arr)
		}
		dfs(level+1, tarr)
	}
	dfs(0, []*TreeNode{root})
	return rlts
}

// offer--35
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	midx := make(map[*Node]int, 0)
	sidx := make([]*Node, 0)
	nidx := make([]*Node, 0)
	//
	h := &Node{}
	hs := h
	//
	s := head
	i := 0
	for s != nil {
		sidx = append(sidx, s.Random)
		midx[s] = i
		//
		c := &Node{Val: s.Val}
		nidx = append(nidx, c)
		hs.Next = c
		hs = c
		//
		i++
		s = s.Next
	}
	//
	hs = h.Next
	i = 0
	for hs != nil {
		if sidx[i] != nil {
			hs.Random = nidx[midx[sidx[i]]]
		}
		i++
		hs = hs.Next
	}
	return h.Next
}

// maxSubArray leetcode--42
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxf := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	//
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] += maxf(nums[i-1], 0)
		max = maxf(nums[i], max)
	}
	return max
}

// offer--46
func translateNum(num int) int {
	s := fmt.Sprintf("%d", num)
	arrs := []byte(s)
	//
	checkNum :=func(a string)bool {
		v,_ :=strconv.Atoi(a)
		if a <= 25 {
			return true
		}
		return false
	}
	//
	f := make(map[int]int, 0)
	//
	f[0] = 1
	//
	for i := 1; i < len(arrs); i++ {
		f[i] += f[i-1]
		if i>=1 &&checkNum(string(arrs[i-1:i+1]) {
			f[i] += f[i-2]
		}
	}
	return f[i]
}

// ---offer-47
func maxValue(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	maxf := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	m, n := len(grid), len(grid[0])
	//
	f := make([][]int, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n+1)
	}
	f[0][0] = 0
	f[1][1] = grid[0][0]
	//
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				continue
			}
			if i == 1 {
				f[i][j] = f[i][j-1] + grid[i-1][j-1]
			} else if j == 1 {
				f[i][j] = f[i-1][j] + grid[i-1][j-1]
			} else {
				f[i][j] = maxf(f[i-1][j], f[i][j-1]) + grid[i-1][j-1]
			}
		}
	}
	return f[m][n]
}

// ----0ffer-50
func firstUniqChar(s string) byte {
	if s == "" {
		return ' '
	}
	stack := make([]byte, 0)
	m := make(map[byte]int, 0)
	//
	for _, c := range []byte(s) {
		if _, ok := m[c]; !ok {
			stack = append(stack, c)
		}
		m[c]++
	}
	for _, c := range stack {
		if m[c] == 1 {
			return c
		}
	}
	return ' '
}

// ---offer-53
func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i]-1 != nums[i-1] {
			return nums[i] - 1
		}
	}
	if nums[0] == 0 {
		return nums[len(nums)-1] + 1
	}
	return 0
}

// ---offer-58
func reverseLeftWords(s string, n int) string {
	if n == 0 || n >= len(s) {
		return s
	}
	bs := make([]byte, 0)
	//
	bs = append(bs, []byte(s)[n:]...)
	bs = append(bs, []byte(s)[0:n]...)
	return string(bs)
}

// maxProfit --offer-63
func maxProfit_ans1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := 0
	start := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[start] > max {
			max = prices[i] - prices[start]
		}
		if prices[i] < prices[start] {
			start = i
		}
	}
	return max
}

func maxProfit_ans2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minf := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	maxf := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	//
	f := make(map[int]int, 0)
	f[0] = 0
	//
	max := 0
	cost := 0
	for i := 0; i < len(prices); i++ {
		max = maxf(prices[i]-cost, max)
		cost = minf(prices[i], cost)
	}
	return max
}
