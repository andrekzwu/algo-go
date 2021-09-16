package skip_list

import (
	"andrezhz/algo-go/list"
)

type SkipListNode struct {
	head list.List_head
	down *SkipListNode
}
