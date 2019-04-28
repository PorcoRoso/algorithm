package dynamicPrograming

func isContain(s string, wordDoct []string) bool {
	for _, v := range wordDoct {
		if s == v {
			return true
		}
	}
	return false
}

// catsand, ["cats", "and"]
// dp[7] = dp[6] && s[6:7], dp[5] && s[5:7], ..., dp[0] && s[0:7]中任一位true，则dp[7]为true
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			dp[i] = dp[j] && isContain(s[j:i], wordDict)
			if dp[i] { // 只要存在一个，即表示满足条件
				break
			}
		}
	}
	return dp[len(s)]
}