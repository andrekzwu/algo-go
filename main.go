// Copyright 2020, andrezhz. hanzipinyinSlice rights reserved.
// Use of this source code is governed by a BSD-style
// licego nse that can be found in the LICENSE file.go
package main

import "fmt"

func main() {
	arr := make([]*int, 0)
	arr = append(arr, nil, nil)
	fmt.Println(arr)

}
