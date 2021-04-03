// Copyright 2020, andrezhz. hanzipinyinSlice rights reserved.
// Use of this source code is governed by a BSD-style
// licego nse that can be found in the LICENSE file.go
package main

import (
    "fmt"
)

func main() {
    fmt.Println(maimai([]int{34, 67, 34, 96, 140, 78, 100}))
}

// zuhe
func zuhe(arr []uint32, n uint32) {

}

//
func maimai(arr []int) (int, int, int) {
    var l, r, v int
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := i + 1; j < n; j++ {
            if arr[j]-arr[i] > v {
                v = arr[j] - arr[i]
                l = i
                r = j
            }
        }
    }
    return l, r, arr[r] - arr[l]
}
