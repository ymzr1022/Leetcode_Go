package contest

func LongestEqualSubarray(nums []int, k int) (res int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		left := i
		target := nums[left]
		ans := 0
		kk := k
		maxx := -1
		for _, v := range nums {
			if target == v {
				ans++
				maxx = max(maxx, ans)
			} else {
				if kk > 0 {
					kk--

				} else {
					if kk < k {
						ans--
						left++
						for left < n && nums[left] == target {
							ans--
							kk++
							left++
						}
					} else {
						break
					}
				}
			}
		}
		res = max(ans, res)
	}
	return
}
