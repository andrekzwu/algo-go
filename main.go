// Copyright 2020, andrezhz. hanzipinyinSlice rights reserved.
// Use of this source code is governed by a BSD-style
// licego nse that can be found in the LICENSE file.go
package main

import (
    "fmt"
)

func main() {
    defer_call()
}

func defer_call() {
    value := uint32(1)
    defer func(x uint32) { fmt.Println("打印前", x) }(value)
    value++
    defer func(x uint32) { fmt.Println("打印中", x) }(value)
    value++
    defer func(x uint32) { fmt.Println("打印后", x) }(value)

    panic("触发异常")
}
