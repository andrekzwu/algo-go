package leetcode

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// convert leetcode-6
func convert(s string, numRows int) string {
	if len(s) == 0 || numRows == 1 {
		return s
	}
	strs := make([]string, numRows)
	for i := 0; i < len(s); i++ {
		rowIndex := 0
		index := i % (2*numRows - 2)
		if index/numRows == 0 {
			rowIndex = index % numRows
		} else {
			index = index % numRows
			rowIndex = numRows - index - 2
		}
		strs[rowIndex] += string(s[i])
	}
	return strings.Join(strs, "")
}

// reverse
// leecode-7
func reverse(x int) int {
	s := int(0)
	for x != 0 {
		s = s*10 + x%10
		x = x / 10
		fmt.Println(s, x)
		if s >= math.MaxInt32 || s <= math.MinInt32 {
			return 0
		}
	}
	return s
}

// isPalindrome leetcode-8
func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}
	v := x
	n := 0
	for v != 0 {
		n = n*10 + v%10
		v = v / 10
	}
	if n != x {
		return false
	}
	return true
}

// myAtoi
// leetcode-8
const (
	START  = "START"
	SIGNED = "SIGNED"
	NUM    = "NUM"
	END    = "END"
)

// '' +/- num other
var myAtoiCheckMap = map[string][]string{
	START:  []string{START, SIGNED, NUM, END},
	SIGNED: []string{END, END, NUM, END},
	NUM:    []string{END, END, NUM, END},
	END:    []string{END, END, END, END},
}

// getCol
func getCol(ch byte) int {
	if ch == ' ' {
		return 0
	}
	if ch == '+' || ch == '-' {
		return 1
	}
	if ch >= '0' && ch <= '9' {
		return 2
	}
	return 3
}

// myAtoi
// leecode-8
func myAtoi(s string) int {
	if s == "" {
		return 0
	}
	state := START
	sign := 1
	ans := 0
	for _, ch := range []byte(s) {
		flag := getCol(ch)
		state = myAtoiCheckMap[state][flag]
		if state == NUM {
			ans = ans*10 + int(ch-'0')
			if ans*sign >= math.MaxInt32 {
				ans = math.MaxInt32
				return ans
			}
			if ans*sign <= math.MinInt32 {
				ans = math.MinInt32
				return ans
			}

		} else if state == SIGNED {
			if ch == '-' {
				sign = -1
			}
		}
	}
	return ans * sign
}

// offer-10
func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	match := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}
	//
	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	//
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if match(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else {
				if match(i, j) {
					f[i][j] = f[i][j] || f[i-1][j-1]
				}
			}
		}
	}
	return f[m][n]
}

// maxArea 双指针法
// leetcode-11
func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	max := int(0)
	i, j := 0, len(height)-1
	for i < j {
		v := int(0)
		if height[i] < height[j] {
			v = height[i] * (j - i)
			i++
		} else {
			v = (j - i) * height[j]
			j--
		}
		if max < v {
			max = v
		}
	}
	return max
}

// 字符          数值
// I             1
// IV			 4
// V             5
// IX 			 9
// X             10
// XL   		 40
// L             50
// XC			 90
// C             100
// CD			 400
// D             500
// CM			 900
// M             1000
// intToRoman
// leetcode 12
func intToRoman(num int) string {
	s := ""
	for num > 0 {
		if num >= 1000 {
			for i := 0; i < num/1000; i++ {
				s += "M"
			}
			num = num % 1000
		} else if num >= 900 {
			s += "CM"
			num = num % 900
		} else if num >= 500 {
			s += "D"
			num = num % 500
		} else if num >= 400 {
			s += "CD"
			num = num % 400
		} else if num >= 100 {
			for i := 0; i < num/100; i++ {
				s += "C"
			}
			num = num % 100
		} else if num >= 90 {
			s += "XC"
			num = num % 90

		} else if num >= 50 {
			s += "L"
			num = num % 50
		} else if num >= 40 {
			s += "XL"
			num = num % 40
		} else if num >= 10 {
			for i := 0; i < num/10; i++ {
				s += "X"
			}
			num = num % 10
		} else if num >= 9 {
			s += "IX"
			num = num % 9
		} else if num >= 5 {
			s += "V"
			num = num % 5
		} else if num >= 4 {
			s += "IV"
			num = num % 4
		} else if num >= 1 {
			for i := 0; i < num/1; i++ {
				s += "I"
			}
			num = num % 1
		}
	}
	return s
}

// romanToInt leetcode-13
func romanToInt(s string) int {
	num := int(0)
	chs := []byte(s)
	i := 0
	for i < len(chs) {
		switch chs[i] {
		case 'I':
			if i+1 < len(chs) && chs[i+1] == 'V' {
				num += 4
				i++
				break
			}
			if i+1 < len(chs) && chs[i+1] == 'X' {
				num += 9
				i++
				break
			}
			num += 1

		// I IV IX
		case 'V':
			num += 5
		case 'X':
			// X XL XC
			if i+1 < len(chs) && chs[i+1] == 'L' {
				num += 40
				i++
				break
			}
			if i+1 < len(chs) && chs[i+1] == 'C' {
				num += 90
				i++
				break
			}
			num += 10
		case 'L':
			num += 50
		case 'C':
			// C CD CM
			if i+1 < len(chs) && chs[i+1] == 'D' {
				num += 400
				i++
				break
			}
			if i+1 < len(chs) && chs[i+1] == 'M' {
				num += 900
				i++
				break
			}
			num += 100
		case 'D':
			num += 500
		case 'M':
			num += 1000
		}
		i++
	}
	return num
}

// longestCommonPrefix leecode-14
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := strs[0]
	if prefix == "" {
		return ""
	}
	//
	minlen := func(str1, str2 string) int {
		if len(str1) < len(str2) {
			return len(str1)
		}
		return len(str2)
	}

	for i := 1; i < len(strs); i++ {
		if strs[i] == "" {
			return ""
		}
		n := minlen(prefix, strs[i])
		for j := 0; j < n; j++ {
			if prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			} else if j == n-1 {
				if len(strs[i]) > n {
					prefix = strs[i][:n]
				}
				if len(prefix) > n {
					prefix = prefix[:n]
				}
			}
		}
	}
	return prefix
}

// threeSum leecode-15
func threeSum(nums []int) [][]int {
	results := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	//
	for i := 0; i < n-2 && nums[i]+nums[i+1]+nums[i+2] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-1]+nums[n-2] < 0 {
			continue
		}
		for left, rignt := i+1, n-1; left < rignt; {
			sum := nums[i] + nums[left] + nums[rignt]
			if sum == 0 {
				results = append(results, []int{nums[i], nums[left], nums[rignt]})
				//
				for left++; left < rignt && nums[left] == nums[left-1]; left++ {
				}
				//
				for rignt--; left < rignt && nums[rignt] == nums[rignt+1]; rignt-- {
				}
			} else if sum < 0 {
				left++
			} else {
				rignt--
			}
		}
	}
	return results
}

// threeSumClosest--leetcode 16
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	n := len(nums)
	best := math.MaxInt32
	//
	abs := func(x int) int {
		if x < 0 {
			return x * -1
		}
		return x
	}
	update := func(cur int) {
		if abs(target-cur) < abs(target-best) {
			best = cur
		}
	}
	//
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left, right := i+1, n-1; left < right; {
			cur := nums[i] + nums[left] + nums[right]
			if cur == target {
				return cur
			}
			update(cur)
			//
			if cur >= target {
				for right--; left < right && nums[right] == nums[right+1]; right-- {
				}
			} else {
				for left++; left < right && nums[left] == nums[left-1]; left++ {
				}
			}
		}
	}
	return best
}

// letterCombinations leetcode-17
var combinations = []string{}

var digit2chsMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	//
	backtrck(digits, 0, "")
	return combinations
}

func backtrck(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		for _, ch := range digit2chsMap[string(digits[index])] {
			backtrck(digits, index+1, combination+string(ch))
		}
	}
}

// fourSum leetcode-18
func fourSum(nums []int, target int) [][]int {
	results := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	//
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i-1] == nums[i] || nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j-1] == nums[j] || nums[i]+nums[j]+nums[n-1]+nums[n-2] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				//
				if sum == target {
					results = append(results, []int{nums[i], nums[j], nums[left], nums[right]})
					//
					for left++; left < right && nums[left-1] == nums[left]; left++ {
					}
					//
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}

				} else if sum < target {
					left++
				} else {
					right--
				}
			}

		}
	}
	return results
}

// removeNthFromEnd
// leetcode-19
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	list := make([]*ListNode, 0)
	s := head
	for s != nil {
		list = append(list, s)
		s = s.Next
	}
	//
	l := len(list)
	if l < n {
		return nil
	}
	if l == n {
		return head.Next
	}

	list[l-n-1].Next = list[l-n].Next
	return head
}

// isValid leetcode-20
func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}
	pair := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	//
	for i := 0; i < len(s); i++ {
		if pair[s[i]] > 0 {
			if len(stack) == 0 || pair[s[i]] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]

		} else {
			stack = append(stack, s[i])
		}
	}
	return true
}

// mergeTwoLists
// leetcode-21
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	s := &ListNode{}
	h := s
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			s.Next = l2
			s = l2
			//
			l2 = l2.Next
		} else {
			s.Next = l1
			s = l1
			//
			l1 = l1.Next
		}
	}
	if l1 != nil {
		s.Next = l1
	}
	if l2 != nil {
		s.Next = l2
	}
	return h.Next
}

// generateParenthesis leetcode-22
var parenthesisResults = []string{}
var stack = make([]byte, 0)

func generateParenthesis(n int) []string {
	parenthesisBacktrack(0, 0, n)
	return parenthesisResults
}

// parenthesisBacktrack
func parenthesisBacktrack(open, closei, n int) {
	if len(stack) == 2*n {
		parenthesisResults = append(parenthesisResults, string(stack))
		return
	}
	if open < n {
		stack = append(stack, '(')
		parenthesisBacktrack(open+1, closei, n)
		stack = stack[:len(stack)-1]
	}
	if closei < open {
		stack = append(stack, ')')
		parenthesisBacktrack(open, closei+1, n)
		stack = stack[:len(stack)-1]
	}
}

// swapPairs leetcode-24
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//
	pre := &ListNode{Next: head}
	h := pre
	p := head
	//
	for p != nil && p.Next != nil {
		s := p.Next.Next
		//
		n := p.Next
		//
		p.Next = s
		n.Next = p
		pre.Next = n
		//
		pre = p
		p = s
	}
	return h.Next
}

// mergeKLists -leetcode 25
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	//
	var p *ListNode
	for _, l := range lists {
		if p == nil {
			p = l
			continue
		}
		s := &ListNode{}
		h := s
		//
		for p != nil && l != nil {
			if p.Val > l.Val {
				s.Next = l
				l = l.Next
				s = s.Next
			} else {
				s.Next = p
				p = p.Next
				s = s.Next
			}
		}
		//
		if p != nil {
			s.Next = p
		}
		if l != nil {
			s.Next = l
		}
		p = h.Next
	}
	return p
}

// reverseKGroup leetcode-25
func reverseKGroup(head *ListNode, k int) *ListNode {
	reverse := func(l, r *ListNode) {
		var p *ListNode
		rnext := r.Next
		for l != rnext {
			s := l.Next
			l.Next = p
			p = l
			//
			l = s
		}
	}
	//
	index := 0
	p := head
	var l, r, s *ListNode
	l = head
	h := &ListNode{}
	s = h
	for p != nil {
		index++
		index = index % k
		//
		pnext := p.Next
		if index == 0 {
			// reverse
			r = p
			reverse(l, r)
			//
			s.Next = r
			s = l
			l = pnext
		}
		p = pnext
		fmt.Printf("index:%d,s:%v,p:%v,l:%v,r:%v\n", index, s, p, l, r)
	}
	if index%k > 0 {
		s.Next = l
	}
	return h.Next
}

// removeDuplicates leetcode-26
func removeDuplicates(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	//
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			nums = append(nums[0:i], nums[i+1:]...)
			i--
			n--
		}
	}
	return len(nums)
}

// removeElement leetcode-27
func removeElement(nums []int, val int) int {
	sort.Ints(nums)
	n := len(nums)
	if len(nums) == 0 {
		return 0
	}
	l, r := -1, -1
	for i := 0; i < n; i++ {
		if nums[i] == val {
			if l == -1 {
				l = i
				r = i
			} else {
				r++
			}
		}
	}
	if l >= 0 {
		nums = append(nums[:l], nums[r+1:]...)
	}
	return len(nums)
}

// strStr leetcode-28
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	i, j := 0, 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
		if j == len(needle) {
			return i - j
		}
	}
	return -1
}

// divide leetcode-29
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > math.MinInt32 {
			return -1 * dividend
		}
		return math.MaxInt32
	}

	abs := func(x int) (int, int) {
		if x < 0 {
			return -1 * x, -1
		}
		return x, 1
	}
	//
	x, xsign := abs(dividend)
	y, ysign := abs(divisor)
	sign := xsign * ysign
	res := div(x, y)
	return res * sign
}

func div(x, y int) int {
	if x < y {
		return 0
	}
	//
	y1 := y
	count := int(1)
	for x >= y<<1 {
		count = 2 * count
		y = y << 1
	}
	return count + div(x-y, y1)
}

// nextPermutation leetcode-31
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	n := len(nums)
	for i := n - 2; i >= 0; i-- {
		index := -1
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				if index == -1 || nums[j] < nums[index] {
					index = j
				}
			}
		}
		if index > 0 {
			nums[i], nums[index] = nums[index], nums[i]
			sort.Ints(nums[i+1:])
			return
		}
	}
	//
	sort.Ints(nums)
}

// search --leetcode 33
func searchleetcode(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

// searchRange leetcode-34
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	//
	start, end := -1, -1
	//
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			if start == -1 {
				start = i
			}
			end = i
		} else if start != -1 {
			return []int{start, end}
		}
	}
	return []int{start, end}
}

// searchInsert leecode-36
func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i < j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	if target > nums[i] {
		return i + 1
	}
	return i
}

// countAndSay --leecode-38
func countAndSay(n int) string {
	var dfs func(level int, s string) string
	dfs = func(level int, s string) string {
		if level == n {
			return s
		}
		dest := make([]byte, 0)
		//
		src := []byte(s)
		//
		idx := 0
		prec := src[0]
		i := 0
		for i = 0; i < len(src); i++ {
			if src[i] == prec {
				continue
			}
			dest = append(dest, '0'+byte(i-idx), prec)
			//
			prec = src[i]
			idx = i
		}
		if i >= idx {
			dest = append(dest, '0'+byte(i-idx), prec)
		}
		//
		return dfs(level+1, string(dest))
	}
	return dfs(1, "1")
}

// combinationSum leetcode-39
// 2 3 6 7
func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	man := make([]int, 0)
	//
	var dfs func(target, idx int)
	dfs = func(target int, index int) {
		if target == 0 {
			ans = append(ans, append([]int(nil), man...))
			return
		}
		if index == len(candidates) {
			return
		}
		//
		dfs(target, index+1)
		//
		if target-candidates[index] >= 0 {
			man = append(man, candidates[index])
			dfs(target-candidates[index], index)
			man = man[:len(man)-1]
		}
	}
	//
	dfs(target, 0)
	return ans

}

// combinationSum2 leecode-40
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	fmt.Println(candidates)
	ans := make([][]int, 0)
	man := make([]int, 0)
	//
	var dfs func(target, idx int)
	dfs = func(target int, index int) {
		if target == 0 {
			ans = append(ans, append([]int(nil), man...))
			return
		}
		if index == len(candidates) {
			return
		}
		//
		for i := index; i < len(candidates) && target >= candidates[i]; i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			man = append(man, candidates[i])
			dfs(target-candidates[i], i+1)
			man = man[:len(man)-1]
		}
	}
	//
	dfs(target, 0)
	return ans
}

// myPow leetcode-50
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	abs := func(x int) (int, bool) {
		if n < 0 {
			return -1 * n, false
		}
		return n, true
	}
	nn, nok := abs(n)
	//
	res := pow(x, nn)
	//
	if nok {
		return res
	} else {
		return 1 / res
	}
}

func pow(x float64, n int) float64 {
	if n <= 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	res := x
	i := 2
	for ; i <= n; i = i * 2 {
		res = res * res
	}

	return res * pow(x, n-i/2)
}

// maxProfit leetcode-121
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	res := 0
	for l, r := 0, 1; r < len(prices); r++ {
		sum := prices[r] - prices[l]
		if sum > res {
			res = sum
		}
		if prices[r] < prices[l] {
			l = r
		}
	}
	return res
}

// reverseList
// leetcode-206
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := (*ListNode)(nil)
	cur := (*ListNode)(nil)
	s := head
	for s != nil {
		cur = s
		s = s.Next
		cur.Next = p
		p = cur
	}
	return p
}

// leecode-523
func checkSubarraySum(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	remainder := 0
	remainder2Index := make(map[int]int, 0)
	remainder2Index[0] = -1
	//
	for i, num := range nums {
		remainder = (remainder + num) % k
		if preindex, ok := remainder2Index[remainder]; ok {
			if i-preindex >= 2 {
				fmt.Println(preindex, i)
				return true
			}

		} else {
			remainder2Index[remainder] = i
		}
	}
	return false
}
