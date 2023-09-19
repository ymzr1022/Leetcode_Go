package dp

/*
输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。

例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。



示例 1：

输入：n = 12
输出：5
示例 2：

输入：n = 13
输出：6


限制：

1 <= n < 2^31
*/

func CountDigitOne(n int) int {

	dp0 := 1
	counts := 0
	for i := 1; i < n; i++ {
		counts = count1(i + 1)
		if counts > 0 {

			dp0 = dp0 + counts
			if counts == 1 && (i+1)%10 == 1 {
				i += 9
			}
		}
	}
	return dp0
}

func count1(i int) (res int) {
	for i != 0 {
		if i%10 == 1 {
			res++
		}
		i /= 10
	}
	return
}
