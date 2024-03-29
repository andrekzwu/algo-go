package tree

import (
    "errors"
    "fmt"
)

type RBCOLOR byte

const (
    RED   = RBCOLOR(0)
    BLACK = RBCOLOR(1)
)

// RBNode
type RBNode struct {
    parent *RBNode // parent node
    left   *RBNode // left node
    right  *RBNode // right node

    value int
    color RBCOLOR
}

// Getleft
func (rbtn *RBNode) Getleft() *RBNode {
    if rbtn == nil || rbtn.left == nil {
        return nil
    }
    return rbtn.left
}

// GetRight
func (rbtn *RBNode) GetRight() *RBNode {
    if rbtn == nil || rbtn.right == nil {
        return nil
    }
    return rbtn.right
}

// GetParent
func (rbtn *RBNode) GetParent() *RBNode {
    if rbtn == nil || rbtn.parent == nil {
        return nil
    }
    return rbtn.parent
}

// GetUncle
func (rbtn *RBNode) GetUncle() *RBNode {
    if rbtn.GetGrandParent() == nil {
        return nil
    }
    if rbtn.GetGrandParent().Getleft() == rbtn.GetParent() {
        return rbtn.GetGrandParent().GetRight()
    }
    if rbtn.GetGrandParent().GetRight() == rbtn.GetParent() {
        return rbtn.GetGrandParent().Getleft()
    }
    return nil
}

// GetGrandParent
func (rbtn *RBNode) GetGrandParent() *RBNode {
    return rbtn.GetParent().GetParent()
}

// isLeft
func (rbtn *RBNode) isLeft() bool {
    return rbtn == rbtn.GetParent().Getleft()
}

// isRight
func (rbtn *RBNode) isRight() bool {
    return rbtn == rbtn.GetParent().GetRight()
}

// LeftRotate
func (rbtn *RBNode) LeftRotate() (*RBNode, error) {
    if rbtn == nil {
        return nil, nil
    }
    if rbtn.right == nil {
        return nil, errors.New("left rotate,but right child node is nil")
    }
    var root *RBNode
    isLeft := false
    //
    parent := rbtn.parent
    if parent != nil {
        isLeft = parent.left == rbtn
    }
    // 以某个结点作为支点(旋转结点)
    right := rbtn.right
    // 右子结点的左子结点变为旋转结点的右子结点
    if right.Getleft() != nil {
        rbtn.right = right.left
        right.left.parent = rbtn
    } else {
        rbtn.right = nil
    }
    // 旋转节点成为右节点的左子节点
    rbtn.parent = right
    right.left = rbtn
    //
    if parent == nil {
        // 如果旋转节点就是根节点，则返回新的根节点
        right.parent = nil
        root = right
    } else {
        // 如果有父节点则更新父节点的子节点信息
        if isLeft {
            parent.left = right
        } else {
            parent.right = right
        }
        right.parent = parent
    }
    return root, nil
}

// RightRotate
func (rbtn *RBNode) RightRotate() (*RBNode, error) {
    if rbtn == nil {
        return nil, nil
    }
    if rbtn.left == nil {
        return nil, errors.New("right rotate,but left child node is nil")
    }
    var root *RBNode
    parent := rbtn.parent
    isRight := false
    if parent != nil {
        isRight = parent.right == rbtn
    }
    // 某个结点作为支点(旋转结点)
    left := rbtn.left
    // 左子结点的右子结点变为旋转结点的左子结点
    if left.GetRight() != nil {
        rbtn.left = left.right
        left.right.parent = rbtn
    } else {
        rbtn.left = nil
    }
    // 旋转接点为其左节点的右节点
    rbtn.parent = left
    left.right = rbtn
    //
    if parent == nil {
        left.parent = nil
        root = left
    } else {
        // 如果有父节点则更新父节点的子节点信息
        if isRight {
            parent.right = left
        } else {
            parent.left = left
        }
        left.parent = parent
    }
    return root, nil
}

// GetBrother
func (rbtn *RBNode) GetBrother() *RBNode {
    if rbtn.isLeft() {
        return rbtn.GetParent().GetRight()
    } else {
        return rbtn.GetParent().Getleft()
    }
}

// GetSuccessor
func (rbtn *RBNode) GetSuccessor() *RBNode {
    if rbtn == nil || rbtn.GetRight() == nil {
        return nil
    }
    successsorNode := rbtn.GetRight()
    //
    for successsorNode.Getleft() != nil {
        successsorNode = successsorNode.Getleft()
    }
    return successsorNode
}

// GetValue
func (rbtn *RBNode) GetValue() int {
    if rbtn == nil {
        return -1
    }
    return rbtn.value
}

// NewRBNode
func NewRBNode(parent *RBNode, color RBCOLOR, value int) *RBNode {
    return &RBNode{
        parent: parent,
        value:  value,
        color:  color,
    }
}

// RBTree
type RBTree struct {
    root *RBNode
}

// Query
func (rbt *RBTree) Query(value int) (*RBNode, bool) {
    rbNode := rbt.query(rbt.root, value)
    if rbNode == nil {
        return nil, false
    }
    return rbNode, true
}

// query
func (rbt *RBTree) query(t *RBNode, value int) *RBNode {
    if t == nil {
        return nil
    }
    if value == t.value {
        return t
    } else if value < t.value {
        return rbt.query(t.left, value)
    } else {
        return rbt.query(t.right, value)
    }
}

// Insert
func (rbt *RBTree) Insert(value int) {
    rbt.insertNode(rbt.root, value)
}

// insertNode
func (rbt *RBTree) insertNode(t *RBNode, value int) {
    // 节点为空，为根节点
    if t == nil {
        rbt.root = NewRBNode(nil, BLACK, value)
        return
    }
    if value <= t.value {
        if t.left != nil {
            rbt.insertNode(t.left, value)
        } else {
            n := NewRBNode(t, RED, value)
            t.left = n
            // insert check
            rbt.insertCheck(n)
        }
    } else {
        if t.right != nil {
            rbt.insertNode(t.right, value)
        } else {
            n := NewRBNode(t, RED, value)
            t.right = n
            // insert check
            rbt.insertCheck(n)
        }
    }
}

// leftRotate
func (rbt *RBTree) leftRotate(t *RBNode) {
    if root, err := t.LeftRotate(); err == nil {
        if root != nil {
            rbt.root = root
        }
    } else {
        // print err
    }
}

// rightRotate
func (rbt *RBTree) rightRotate(t *RBNode) {
    if root, err := t.RightRotate(); err == nil {
        if root != nil {
            rbt.root = root
        }
    } else {
        // print err
    }
}

// insertCheck
func (rbt *RBTree) insertCheck(t *RBNode) {
    // 如果父节点为黑色，则不需要变动
    if t.parent == nil {
        t.color = BLACK
        rbt.root = t
        return
    }
    if t.parent.color == BLACK {
        return
    }
    if t.GetUncle() != nil && t.GetUncle().color == RED {
        // 父节点和叔叔节点都为红色，则都转为黑色
        t.GetParent().color = BLACK
        t.GetUncle().color = BLACK
        // 祖父节点的颜色改为红色
        t.GetGrandParent().color = RED
        // 递归处理祖父节点的颜色变更
        rbt.insertCheck(t.GetGrandParent())
    } else {
        // 父节点为红色，叔叔节点不存在 或者为黑色
        //
        if t.GetParent().isLeft() && t.isLeft() {
            // 左左结构
            // 先变色，父节点变为黑色，祖父节点变为红色
            t.GetParent().color = BLACK
            t.GetGrandParent().color = RED
            // 祖父节点作为旋转节点
            rbt.rightRotate(t.GetGrandParent())
        } else if t.GetParent().isLeft() && t.isRight() {
            // 左右结构
            // 父节点作为旋转节点
            // 先左旋
            rbt.leftRotate(t.GetParent())
            // 再右旋
            rbt.rightRotate(t.GetParent())
            // 变色
            t.color = BLACK
            t.Getleft().color = RED
            t.GetRight().color = RED
        } else if t.GetParent().isRight() && t.isRight() {
            // 右右结构
            // 先变色，父节点变为黑色，祖父节点变为红色
            t.GetParent().color = BLACK
            t.GetGrandParent().color = RED
            // 单左旋
            rbt.leftRotate(t.GetGrandParent())
        } else if t.GetParent().isRight() && t.isLeft() {
            // 右左结构
            // 父节点作为旋转节点
            // 先右旋变成右右结构
            rbt.rightRotate(t.GetParent())
            // 再左旋
            rbt.leftRotate(t.GetParent())
            // 变色
            t.color = BLACK
            t.Getleft().color = RED
            t.GetRight().color = RED
        }
    }
}

// PreOrder
func (rbt *RBTree) PreOrder() {
    fmt.Println("---------------前序遍历")
    rbt.preOrder(rbt.root)
    fmt.Println("")
}

// preOrder
func (rbt *RBTree) preOrder(t *RBNode) {
    if t == nil {
        return
    }
    // print myself
    fmt.Printf("%d ", t.value)
    // print left
    rbt.preOrder(t.Getleft())
    // print right
    rbt.preOrder(t.GetRight())
}

// InOrder
func (rbt *RBTree) InOrder() {
    fmt.Println("---------------中序遍历")
    rbt.inOrder(rbt.root)
    fmt.Println("")
}

// inOrder
func (rbt *RBTree) inOrder(t *RBNode) {
    if t == nil {
        return
    }
    // print left
    rbt.inOrder(t.Getleft())
    // print myself
    fmt.Printf("%d ", t.value)
    // print right
    rbt.inOrder(t.GetRight())
}

// PostOrder
func (rbt *RBTree) PostOrder() {
    fmt.Println("---------------后序遍历")
    rbt.postOrder(rbt.root)
    fmt.Println("")
}

// postOrder
func (rbt *RBTree) postOrder(t *RBNode) {
    if t == nil {
        return
    }
    // print left
    rbt.postOrder(t.Getleft())
    // print right
    rbt.postOrder(t.GetRight())
    // print myself
    fmt.Printf("%d ", t.value)
}

// LevelOrder
func (rbt *RBTree) LevelOrder() {
    fmt.Println("---------------层级遍历")
    if rbt.root == nil {
        return
    }
    rbt.levelOrder([]*RBNode{rbt.root})
    fmt.Println("")
}

// levelOrder
func (rbt *RBTree) levelOrder(nodes []*RBNode) {
    if len(nodes) == 0 {
        return
    }
    subNodes := make([]*RBNode, 0)
    for _, t := range nodes {
        // print node
        t.printNode()
        //
        if t.Getleft() != nil {
            subNodes = append(subNodes, t.Getleft())
        }
        if t.GetRight() != nil {
            subNodes = append(subNodes, t.GetRight())
        }
    }
    fmt.Println("")
    rbt.levelOrder(subNodes)
}

// printNode
func (rbtn *RBNode) printNode() {
    branch := ""
    color := ""
    if rbtn.GetParent() == nil {
        branch = "根"
        color = "黑"
    } else {
        if rbtn.isLeft() {
            branch = "左"
        } else {
            branch = "右"
        }
        //
        if rbtn.color == BLACK {
            color = "黑"
        } else {
            color = "红"
        }
    }
    //
    fmt.Printf("[%d-%s-%s-%d] ", rbtn.GetParent().GetValue(), branch, color, rbtn.value)
}

// Delete
func (rbt *RBTree) Delete(value int) bool {
    // 查询节点
    destNode, ok := rbt.Query(value)
    if !ok {
        return false
    }
    rbt.delete(destNode)
    return true
}

// delete
func (rbt *RBTree) delete(t *RBNode) {
    // 该删除节点存在两个子节点
    if t.Getleft() != nil && t.GetRight() != nil {
        // 则寻找其直接后继节点，删除后继节点
        successorNode := t.GetSuccessor()
        // 将后继节点值更新为删除节点
        t.value = successorNode.value
        // 删除后继节点
        rbt.delete(successorNode)
        return
    }
    // 该删除节点无子节点
    if t.Getleft() == nil && t.GetRight() == nil {
        rbt.deleteEmpty(t)
        return
    }
    // 该删除节点存在一个子节点
    var child *RBNode
    if t.Getleft() != nil {
        child = t.Getleft()
    } else {
        child = t.GetRight()
    }
    // 将子节点替换改删除节点
    t.value = child.value
    // 删除改子节点
    rbt.delete(child)
    return
}

// deleteEmpty
func (rbt *RBTree) deleteEmpty(t *RBNode) {
    // 删除节点是根节点
    if t.GetParent() == nil {
        rbt.root = nil
        return
    }
    // 删除节点为非根节点
    // 删除检查
    rbt.deleteCheck(t)
    if t.isLeft() {
        t.GetParent().left = nil
    } else {
        t.GetParent().right = nil
    }
}

// deleteCheck
// 删除节点N，父节点P，兄弟节点S，兄弟节点左右子节点SL/SR
func (rbt *RBTree) deleteCheck(t *RBNode) {
    // 删除检查节点为根节点 或者删除检查节点为红色，无需处理直接返回
    if t.GetParent() == nil || t.color == RED {
        return
    }
    // 删除节点为非根节点，且为黑色
    //
    if t.GetBrother().color == RED {
        // 兄弟节点为红色，则父节点一定为黑色，且兄弟节点有两个黑子节点
        if t.isLeft() {
            // 删除节点为左节点
            // 旋转前兄弟节点设置为黑色，父节点设置为红色
            t.GetParent().color = RED
            t.GetBrother().color = BLACK
            // 以父节点为支点左旋
            rbt.leftRotate(t.GetParent())
            // 旋转后，颜色变成，兄弟节点为黑色，父节点为红色，转换场景继续检查
            rbt.deleteCheck(t)
        } else {
            // 删除节点为右节点
            // 宣钻前兄弟节点设置为黑色，父节点设置为红色
            t.GetParent().color = RED
            t.GetBrother().color = BLACK
            // 以父节点为支点右旋
            rbt.rightRotate(t.GetParent())
            // 旋转后，颜色变成，兄弟节点为黑色，父节点为红色，转换场景继续检查
            rbt.deleteCheck(t)
        }
    } else {
        // 兄弟节点为黑色，父节点颜色未知
        // 兄弟节点左右子节点不存在或都为黑色
        if (t.GetBrother().Getleft() == nil || t.GetBrother().Getleft().color == BLACK) &&
            (t.GetBrother().GetRight() == nil || t.GetBrother().GetRight().color == BLACK) {
            if t.GetParent().color == RED {
                // 父节点为红色，将兄弟节点设置为红色，父节点设置为黑色
                t.GetParent().color = BLACK
                t.GetBrother().color = RED
            } else {
                // 父节点为黑色，将兄弟节点设置为红色，保持内部平衡，删除检查转移至父节点
                t.GetBrother().color = RED
                rbt.insertCheck(t.GetParent())
            }
            return
        }
        if t.isLeft() { // 删除节点为左节点
            if t.GetBrother().GetRight() != nil && t.GetBrother().GetRight().color == RED {
                // 兄弟节点的右子节点存在且为红色
                // 旋转前，设置兄弟节点为父节点的颜色后，父节点设置为黑色，兄弟节点的右节点设置为黑色
                t.GetBrother().color = t.GetParent().color
                t.GetParent().color = BLACK
                t.GetBrother().GetRight().color = BLACK
                // 以父节点为支点进行左旋
                rbt.leftRotate(t.GetParent())
            } else {
                // 兄弟节点的左子节点存在且为红色
                // 旋转前，将兄弟节点的左子节点的颜色设置为父节点的颜色，将父节点设置为黑色
                t.GetBrother().Getleft().color = t.GetParent().color
                t.GetParent().color = BLACK
                // 先将兄弟节点作为支点右旋
                rbt.rightRotate(t.GetBrother())
                // 再以父节点为支点进行左旋
                rbt.leftRotate(t.GetParent())
            }
        } else { // 删除节点为右节点
            if t.GetBrother().Getleft() != nil && t.GetBrother().Getleft().color == RED {
                // 兄弟节点左子节点存在且为红色
                // 旋转前，设置兄弟节点为父节点的颜色后，父节点设置为黑色，兄弟节点的左子节点设置为黑色
                t.GetBrother().color = t.GetParent().color
                t.GetParent().color = BLACK
                t.GetBrother().Getleft().color = BLACK
                // 以父节点为支点进行右旋
                rbt.rightRotate(t.GetParent())
            } else {
                // 兄弟节点右子节点存在且为红色
                // 旋转前，将兄弟节点的右子节点设置为父节点的颜色，父节点设置为黑色
                t.GetBrother().GetRight().color = t.GetParent().color
                t.GetParent().color = BLACK
                // 先将兄弟节点作为支点左旋
                rbt.leftRotate(t.GetBrother())
                //  再以父节点为支点进行右旋
                rbt.rightRotate(t.GetParent())
            }
        }
    }
}
