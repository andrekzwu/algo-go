package leetcode

import (
	"fmt"
	"testing"
)

// TestFindNumberIn2DArray offer-04
func TestFindNumberIn2DArray(t *testing.T) {
	//     [
	//   [1,   4,  7, 11, 15],
	//   [2,   5,  8, 12, 19],
	//   [3,   6,  9, 16, 22],
	//   [10, 13, 14, 17, 24],
	//   [18, 21, 23, 26, 30]
	// ]

	matrix := make([][]int, 0)
	matrix = append(matrix, []int{1, 1})
	// matrix = append(matrix, []int{1, 4, 7, 11, 15})
	// matrix = append(matrix, []int{2, 5, 8, 12, 19})
	// matrix = append(matrix, []int{3, 6, 9, 16, 22})
	// matrix = append(matrix, []int{10, 13, 14, 17, 24})
	// matrix = append(matrix, []int{18, 21, 23, 26, 30})

	fmt.Println(findNumberIn2DArray(matrix, 0))
}

// TestreplaceSpace offer-05
func TestReplaceSpace(t *testing.T) {
	fmt.Println(replaceSpace("a c d "))
}

// TestFib offer-10-I
func TestFib(t *testing.T) {
	fmt.Println(fib(95))
}

// TestNumWays
func TestNumWays(t *testing.T) {
	fmt.Println(numWays(7))
}

// TestCopyRandomList offer-35
func TestCopyRandomList(t *testing.T) {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n1.Next = n2
	n3 := &Node{Val: 3, Random: n1}
	n2.Next = n3
	n1.Random = n3
	//
	printNode(n1)
	h := copyRandomList(n1)
	printNode(h)
}

// TestMaxSubArray
func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

// offer-47
func TestMaxValue(t *testing.T) {
	grid := make([][]int, 0)
	grid = append(grid, []int{1, 3, 1})
	grid = append(grid, []int{1, 5, 1})
	grid = append(grid, []int{4, 2, 1})
	fmt.Println(maxValue(grid))
}

// TestReverseLeftWords 0ffer-58
func TestReverseLeftWords(t *testing.T) {
	fmt.Println(reverseLeftWords("abcdefghi", 3))
}
