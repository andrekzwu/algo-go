package sort

import (
    "testing"
)

var (
    values = []int{1, 9, 4, 5, 2, 8, 3, 6, 7}
)

// TestMaoaoSort
func TestMaoaoSort(t *testing.T) {
    printIntArr(MAOPAO_SORT, values)
    maopaoSort(values)
    printIntArr(MAOPAO_SORT, values)
}

// TestSelectSort
func TestSelectSort(t *testing.T) {
    printIntArr(SELECt_SORT, values)
    selectSort(values)
    printIntArr(SELECt_SORT, values)
}

// TestInsertSort
func TestInsertSort(t *testing.T) {
    printIntArr(INSERT_SORT, values)
    insertSort(values)
    printIntArr(INSERT_SORT, values)
}
