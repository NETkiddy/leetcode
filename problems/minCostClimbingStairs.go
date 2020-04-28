package problems

//基于动态规划实现
func minCostClimbingStairs(cost []int) int {
	l:=len(cost)
	if l==0 {
		return 0
	}
	if l==1 {
		return cost[0]
	}
	if l==2 {
		return min(cost[0],cost[1])
	}
	dp := make([]int, l)
	for i:=0;i<=l-2;i++{
		if i==0 || i==1{
			dp[i]=cost[i]
		}else {
			dp[i]=min(dp[i-2],dp[i-1])+cost[i]
		}
	}

	return min(dp[l-3]+cost[l-1],dp[l-2])
}

func min(a int, b int) int{
	m:=a
	if b<a{
		m=b
	}
	return m
}