package stack

import (
    "andrezhz/algo-go/list"
    "fmt"
    "testing"
    "unsafe"
)

type StackElem struct {
    Node  *list.List_head
    Value int
}

func NewStackElem(value int) *StackElem {
    return &StackElem{Value: value, Node: &list.List_head{}}
}

// StackElemOffSet
func StackElemOffSet() uintptr {
    elem := StackElem{}
    return uintptr(unsafe.Pointer(&elem.Node)) - uintptr(unsafe.Pointer(&elem))
}

// ListHead2StackElem
func ListHead2StackElem(pnode **list.List_head) *StackElem {
    return (*StackElem)(unsafe.Pointer(uintptr(unsafe.Pointer(pnode)) - StackElemOffSet()))
}

// TestListHead2StackElem
func TestListHead2StackElem(t *testing.T) {
    elem := NewStackElem(1)
    elem1 := ListHead2StackElem(&elem.Node)

    fmt.Println(elem, elem1)

}

func TestStackImpl(t *testing.T) {
    values := []int{1, 2, 3, 4, 5, 6}
    stackImpl := StackImplInit()
    //
    for _, value := range values {
        elem := NewStackElem(value)
        stackImpl.Push(elem.Node)
    }
    // TOP
    fmt.Println("stack elem size:", stackImpl.Size())
    //
    for {
        node, ok := stackImpl.Pop()
        if ok {
            break
        }
        elem := ListHead2StackElem(&node)
        fmt.Println(elem.Value)
    }
}
