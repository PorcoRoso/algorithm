package greedy

// 当前能到达的最远距离为cur，上一次能到达的最远距离为last
func jump(nums []int) int {
	cnt := 0
	cur, last := 0, 0 // 当前能达到的最远位置，上一次能达到的最远位置
	for i := 0; i < len(nums); i++ {
		if i > last { // 当前位置大于last时，计数一次，更新last
			cnt++
			last = cur
		}

		if cur < i + nums[i] { // 一直更新为最大
			cur = i + nums[i]
		}
	}
	return cnt
}
