package service

func combinationSum(candidates []int, target int) [][]int {
	sort(candidates)

	if target < candidates[0] {
		return nil
	}

	var ans [][]int
	var tmp []int
	lc := len(candidates)
	var trace func(k int)
	trace = func(k int) {
		if k == lc {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), tmp...))
			return
		}
		// 优先跳过
		trace(k + 1)
		if target-candidates[k] >= 0 {
			target -= candidates[k]
			tmp = append(tmp, candidates[k])
			trace(k)
			// 回溯
			tmp = tmp[0 : len(tmp)-1]
			target += candidates[k]
		}
	}
	trace(0)
	return ans
}

func sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
