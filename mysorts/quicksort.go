package mysorts

func partition(nums []int, left, right int) int {
	target := nums[left]
	i := left
	j := right
	for i < j {
		for j > i && target <= nums[j] {
			j--
		}
		for i < j && target >= nums[i] {
			i++
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

func quickSort(nums []int, left, right int) {

	if left < right {
		i := partition(nums, left, right)
		quickSort(nums, left, i-1)
		quickSort(nums, i+1, right)
	}

}
