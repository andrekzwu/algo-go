package sort

import (
    "fmt"
)

const (
    MAOPAO_SORT = "冒泡排序："
    SELECt_SORT = "选择排序："
    INSERT_SORT = "插入排序："
)

// printIntArr
func printIntArr(sortType string, values []int) {
    fmt.Printf("%s", sortType)
    for _, v := range values {
        fmt.Printf("%d ", v)
    }
    fmt.Println("")
}

// maopaoSort
func maopaoSort(values []int) []int {
    n := len(values)
    for i := 1; i < n; i++ {
        for j := n - 1; j > i-1; j-- {
            if values[j] < values[j-1] {
                temp := values[j]
                values[j] = values[j-1]
                values[j-1] = temp
            }
        }
    }
    return values
}

// selectSort
func selectSort(values []int) []int {
    n := len(values)
    for i := 0; i < n-1; i++ {
        min := values[i]
        x := i
        for j := i + 1; j < n; j++ {
            if values[j] < min {
                min = values[j]
                x = j
            }
        }
        values[x] = values[i]
        values[i] = min
    }
    return values
}

// insertSort
func insertSort(values []int) []int {
    n := len(values)
    for i := 1; i < n; i++ {
        for j := i; j > 0 && values[j] < values[j-1]; j-- {
            temp := values[j]
            values[j] = values[j-1]
            values[j-1] = temp
        }
    }
    return values
}
