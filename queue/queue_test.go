package queue

import (
	"andrezhz/algo-go/list"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type QueueElem struct {
	Value int
	Node  list.List_head
}

// NewQueueElem
func NewQueueElem(value int) *QueueElem {
	return &QueueElem{
		Value: value,
		Node:  list.List_head{},
	}
}

var (
	// offSetElem              = QueueElem{}
	// QueueElemOffSet uintptr = uintptr(unsafe.Pointer(&offSetElem.Node)) - uintptr(unsafe.Pointer(&offSetElem))
	QueueElemOffSet uintptr = getQueueElemOffSet()
)

// getQueueElemOffSet
func getQueueElemOffSet() uintptr {
	t := reflect.TypeOf((*QueueElem)(nil)).Elem()
	fi, ok := t.FieldByName("Node")
	if !ok {
		return 0
	}
	return uintptr(fi.Offset)
}

// ListHead2QueueElem
func ListHead2QueueElem(pnode *list.List_head) *QueueElem {
	return (*QueueElem)(unsafe.Pointer(uintptr(unsafe.Pointer(pnode)) - QueueElemOffSet))
}

// TestQueueElem
func TestQueueElem(t *testing.T) {
	fmt.Println("QueueElemOffSet:", QueueElemOffSet, "  getQueueElemOffSet:", getQueueElemOffSet())

	//
	values := []int{9, 1, 2, 3, 4, 5, 6, 7}
	queueImpl := NewQueueImpl()
	//
	// enter
	for _, v := range values {
		elem := NewQueueElem(v)
		queueImpl.EnQueue(&elem.Node)
	}
	//
	fmt.Println("queue size:", queueImpl.Size())
	// dl
	for !queueImpl.Empty() {
		node := queueImpl.DlQueue()
		elem := ListHead2QueueElem(node)
		fmt.Println(elem)
	}
}
