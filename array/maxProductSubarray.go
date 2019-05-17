package array

import "math"

// 解法1： 传统解法，两个指针作为子数组的首尾，暴力遍历
// 解法2： 动态规划，因为有负数乘以负数的情况，所以会导致突然一个很小的负数，遇到了另一个负数，结果很大，如-2，4，-8
//        dp_max[i] = max(dp_max[i-1] * A[i], dp_min[i-1] * A[i], A[i])
//        dp_min[i] = min(dp_max[i-1] * A[i], dp_min[i-1] * A[i], A[i])
// 解法一的结果耗时88ms,2.7M内存
// 解法二的结果耗时4ms,2.9M内存
func maxProduct(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	max := math.MinInt8

	for i := 0; i < len(nums); i++ {
		curMax := 1
		for j := i; j < len(nums); j++ {
			if i == j {
				curMax = nums[i]
			} else {
				curMax *= nums[j]
			}

			if curMax > max {
				max = curMax
			}
		}

	}
	return max
}

func maxProductDP(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var dp_min []int = make([]int, len(nums))
	var dp_max []int = make([]int, len(nums))
	dp_min[0] = nums[0]
	dp_max[0] = nums[0]
	max_global := nums[0]
	for i := 1; i < len(nums); i++ {
		dp_max[i] = max(max(dp_max[i-1]* nums[i], dp_min[i-1] * nums[i]), nums[i])
		dp_min[i] = min(min(dp_max[i-1]* nums[i], dp_min[i-1] * nums[i]), nums[i])
		max_global = max(dp_max[i], max_global)
	}

	return max_global
}

func max(a , b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a , b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
