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
