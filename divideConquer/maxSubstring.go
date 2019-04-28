package divideConquer

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else {
		return maxSubArrayDivide(nums, 0, len(nums)-1)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 分为三部分：左半部分最大值，右半部分最大值，左半部分右 + 右半部分左
// 分析参考：https://blog.csdn.net/HaixWang/article/details/80719804
func maxSubArrayDivide(nums []int, from, to int) int {
	if from == to {
		return nums[from]
	}
	middle := (from+to) / 2
	left := maxSubArrayDivide(nums, from, middle)
	right := maxSubArrayDivide(nums, middle+1, to)

	// 左半部分的最大 + 右半部分最大
	// leftEnd从middle处往左累加，leftSuffix记录这其中的最大值
	leftEnd, leftSuffix := nums[middle], nums[middle]
	rightBegin, rightPrefix := nums[middle+1], nums[middle+1]

	for i := 1;middle - i >= from; i++ {
		leftEnd += nums[middle - i]
		leftSuffix = max(leftSuffix, leftEnd)
	}

	for i := 2;middle + i <= to; i++ {
		rightBegin += nums[middle+i]
		rightPrefix = max(rightPrefix, rightBegin)
	}

	join := leftSuffix + rightPrefix
	return max(left, max(right, join))
}
