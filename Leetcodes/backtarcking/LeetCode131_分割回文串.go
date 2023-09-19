package leetcodes

/**
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。



示例 1：

输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
示例 2：

输入：s = "a"
输出：[["a"]]


提示：

1 <= s.length <= 16
s 仅由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/palindrome-partitioning
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Partition(s string) (res [][]string) {

	var path []string
	var process func(int)
	n := len(s)
	process = func(i int) {
		if i == n {
			res = append(res, append([]string(nil), path...))
			return
		}

		for j := i; j < n; j++ {
			if Checks(s, i, j) {
				path = append(path, s[i:j+1])
				process(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	process(0)
	return
}

func Checks(str string, i int, j int) (bo bool) {

	bo = true
	for i < j {
		if str[i] != str[j] {
			bo = false
			return
		}
		i++
		j--
	}
	return
}

func Check(str string) (bo bool) {
	i := 0
	j := len(str) - 1
	bys := []byte(str)
	bo = true
	for i < j {
		if bys[i] != bys[j] {
			bo = false
			return
		}
		i++
		j--
	}
	return
}

func GetLen(strs []string) int {
	var str string
	for _, v := range strs {
		str += v
	}
	return len(str)
}

func Partitions(s string) (res [][]string) {
	// strs := strings.Split(s, "")
	var nums []string
	var str string
	var process func(int)
	n := len(s)
	process = func(i int) {
		if i == len(s) && Check(str) && n == GetLen(nums) {
			comb := make([]string, len(nums))
			copy(comb, nums)
			res = append(res, comb)
			return
		}
		if i >= len(s) {
			return
		}

		str += s[i : i+1]
		if Check(str) {
			tmp := str
			str = ""
			nums = append(nums, tmp)
			process(i + 1)
			str = tmp
			nums = nums[:len(nums)-1]
		}
		if i < n-1 {

			process(i + 1)
		}

	}
	process(0)
	return
}
