package leetcode

import (
    "fmt"
    "testing"
)

// TestMyAtoi
func TestMyAtoi(t *testing.T) {
    str := "-9223372036854775807"
    fmt.Println(myAtoi(str))
}

// TestIsPalindrome
func TestIsPalindrome(t *testing.T) {
    fmt.Println(isPalindrome(121))
}

// TestConvert
func TestConvert(t *testing.T) {
    fmt.Println(convert("LEETCODEISHIRING", 3))
}

// TestIsMatch
func TestIsMatch(t *testing.T) {
    fmt.Println(isMatch("aaa", "ab*a*c*a"))
}
