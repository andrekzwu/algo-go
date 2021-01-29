package leetcode

import (
    "fmt"
)

// isMatch
func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    matches := func(i, j int) bool {
        if i == 0 {
            return false
        }
        if p[j-1] == '.' {
            return true
        }
        return s[i-1] == p[j-1]
    }

    f := make([][]bool, m+1)
    for i := 0; i < len(f); i++ {
        f[i] = make([]bool, n+1)
    }
    f[0][0] = true
    for i := 0; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if p[j-1] == '*' {
                f[i][j] = f[i][j] || f[i][j-2]
                if matches(i, j-1) {
                    f[i][j] = f[i][j] || f[i-1][j]
                }
            } else if matches(i, j) {
                f[i][j] = f[i][j] || f[i-1][j-1]
            }
            fmt.Println(i, j, f[i][j])
        }
    }
    return f[m][n]
}

// [9] 三数之和，排序 + 双指针
// 双指针用于处理边界关系，缩小边界，得到一个结果
func threeSum(nums []int) [][]int {
    n := len(nums)
    if len(nums) < 3 {
        return nil
    }
    // sort nums
    sort.Ints(nums)
    ans := make([][]int, 0)
    //
    for i := int(0); i < n; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        k := n - 1
        target := -1 * nums[i]
        for j := i + 1; j < n; j++ {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            for k > j && nums[j]+nums[k] > target {
                k--
            }
            if k == j {
                break
            }
            if nums[j]+nums[k] == target {
                ans = append(ans, []int{nums[i], nums[j], nums[k]})
            }
        }
    }
    return ans
}

// maxArea
func maxArea(height []int) int {
    n := len(height)
    if len(height) < 2 {
        return 0
    }
    minF := func(x, y int) int {
        if x > y {
            return y
        }
        return x
    }
    maxF := func(x, y int) int {
        if x > y {
            return x
        }
        return y
    }
    //
    l, h := 0, n-1
    ans := int(0)
    //
    for l < h {
        x := (h - l) * minF(height[l], height[h])
        if height[l] < height[h] {
            l++
        } else {
            h--
        }
        ans = maxF(x, ans)
    }
    return ans
}
