package greedy

// 贪心算法
// 收5元：OK
// 收10元：找5元
// 收20元：找5+5+5或者10+5,贪心角度，尽可能多的服务客户，找10+5的
func lemonadeChange(bills []int) bool {
	if len(bills) <= 1 {
		return true
	}

	fiveCnt := 0
	tenCnt := 0
	for _,v := range bills {
		if v == 5 {
			fiveCnt++
		} else if v == 10 {
			if fiveCnt <= 0 {
				return false
			}
			fiveCnt--
			tenCnt++
		} else if v == 20 {
			if tenCnt > 0 && fiveCnt > 0 {
				tenCnt--
				fiveCnt--
			} else if fiveCnt >= 3 {
				fiveCnt--
				fiveCnt--
				fiveCnt--
			} else {
				return false
			}
		}
	}
	return true
}
