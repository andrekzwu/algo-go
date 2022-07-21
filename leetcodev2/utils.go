package leetcodev2

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *Node) val() int {
	if n == nil {
		return -1
	}
	return n.Val
}

// min 最小值
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
