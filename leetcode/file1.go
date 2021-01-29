package leetcode

import (
    "fmt"
)

// myAtoi
func myAtoi(str string) int {
    if len(str) == 0 {
        return 0
    }
    result := int(0)
    maxInt32 := int32(^uint32(0) >> 1)
    minInt32 := ^maxInt32
    sign := 1
    start := false
    for _, c := range []byte(str) {
        if c == ' ' && !start {
            continue
        }
        if c == '-' && !start {
            start = true
            sign = -1
            continue
        }
        if c == '+' && !start {
            start = true
            continue
        }
        if c > '9' || c < '0' {
            break
        }
        start = true
        if sign == -1 && result*10+sign*int(c-'0') < int(minInt32) {
            return int(minInt32)
        }
        if sign == 1 && result*10+sign*int(c-'0') > int(maxInt32) {
            return int(maxInt32)
        }
        result = result*10 + sign*int(c-'0')
    }
    return result
}

// isPalindrome
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    s := x
    r := int(0)
    for {
        if s == 0 {
            break
        }
        r = r*10 + s%10
        s /= 10
    }
    return x == r
}

// convert
func convert(s string, numRows int) string {
    if numRows < 2 {
        return s
    }
    rowBytes := make([][]byte, numRows)
    for i := int(0); i < len(s); i += 2*numRows - 2 {
        fmt.Println(i)
        var values []byte
        if len([]byte(s)[i:]) < 2*numRows-2 {
            values = []byte(s)[i:]
        } else {
            values = []byte(s)[i : i+2*numRows-2]
        }
        fmt.Println(string(values))
        for i, ch := range values {
            if i < numRows {
                rowBytes[i] = append(rowBytes[i], ch)
                fmt.Printf("%c ", ch)
            } else {
                rowBytes[2*numRows-i-2] = append(rowBytes[2*numRows-i-2], ch)
                fmt.Printf("%c ", ch)
            }

        }
        fmt.Println("")
    }
    rows := make([]byte, 0)
    for _, bytes := range rowBytes {
        rows = append(rows, bytes...)
    }
    return string(rows)
}
