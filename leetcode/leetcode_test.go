package leetcode

import (
	"fmt"
	"math"
	"testing"
)

// TestLru
func TestLru(t *testing.T) {
	lruCache := NewLRUCache(3, func(key, value interface{}) { fmt.Println("callback", key, value) })

	lruCache.Add("a", "1")
	lruCache.Add("b", "2")
	lruCache.Add("c", "3")
	lruCache.Add("a", "4")
	lruCache.Add("d", "5")
}

// Testconvert
func TestConvert(T *testing.T) {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

// TestReverse
// leecode-7
func TestReverse(t *testing.T) {
	fmt.Println(reverse(-143))
}

// TestMyAtoi
// leecode 8
func TestMyAtoi(t *testing.T) {
	fmt.Println(myAtoi("42"))
	fmt.Println(math.MinInt32, math.MaxInt32)
}

// TestIsMatch
func TestIsMatch(t *testing.T) {
	fmt.Println(isMatch("ab", "a*"))
}

// TestIntToRoman leetcode-12
func TestIntToRoman(t *testing.T) {
	fmt.Println(intToRoman(1994))
}

// TestRomanToInt leetcode-13
func TestRomanToInt(t *testing.T) {
	fmt.Println(romanToInt("IV"))
}

// TestThreeSumClosest leetcode-16
func TestThreeSumClosest(t *testing.T) {
	fmt.Println(threeSumClosest([]int{0, 2, 1, -3}, 1))
}

// TestLetterCombinations
func TestLetterCombinations(t *testing.T) {
	fmt.Println(letterCombinations("2"))
}

// TestMaxArea leetcode-17
func TestMaxArea(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

// TestFourSum leetcode-19
func TestFourSum(t *testing.T) {
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}

// TestGenerateParenthesis
func TestGenerateParenthesis(t *testing.T) {
	fmt.Println(generateParenthesis(1))
}

// TestswapPairs leetcode 24
func TestSwapPairs(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l1.Next = l2
	l3 := &ListNode{Val: 3}
	l2.Next = l3
	l4 := &ListNode{Val: 4}
	l3.Next = l4
	//
	h := swapPairs(l1)
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next

	}
}

// TestReverseKGroup
func TestReverseKGroup(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l1.Next = l2
	l3 := &ListNode{Val: 3}
	l2.Next = l3
	l4 := &ListNode{Val: 4}
	l3.Next = l4
	l5 := &ListNode{Val: 5}
	l4.Next = l5
	//
	printListNode(l1)
	h := reverseKGroup(l1, 2)
	printListNode(h)
}

// TestStrStr leetcode-28
func TestStrStr(t *testing.T) {
	fmt.Println(strStr("mississippi", "issip"))
}

// TestDivide leetcode-29
func TestDivide(t *testing.T) {
	fmt.Println(divide(9, 3))
}

// TestNextPermutation leetcode-31
func TestNextPermutation(t *testing.T) {
	nextPermutation([]int{1, 3, 2})
}

// TestsearchInsert leetcode-36
func TestSearchInsert(t *testing.T) {
	fmt.Println(searchInsert([]int{1, 3}, 2))
}

// TestCountAndSay leecode-38
func TestCountAndSay(t *testing.T) {
	fmt.Println(countAndSay(5))
}

// TestcombinationSum leecode-39
func TestCombinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

// TestCombinationSum2
func TestCombinationSum2(t *testing.T) {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 7, 6, 1, 5}, 8))
}

// TestMyPow
func TestMyPow(t *testing.T) {
	fmt.Println(myPow(2, -3))
}

// leetcode 206
func TestReverseList(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	//
	reverse := reverseList(head)
	for reverse != nil {
		fmt.Printf("%d ", reverse.Val)
		reverse = reverse.Next
	}
}

// leecode-523
func TestCheckSubarraySum(t *testing.T) {
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 6}, 7))
}
