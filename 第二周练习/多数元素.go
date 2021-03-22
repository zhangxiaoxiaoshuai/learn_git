//第一种解法
func majorityElement(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n < 2 {
		return nums[0]
	}

	hash := make(map[int]int)
	for i := 0; i < n; i++ {
		hash[nums[i]]++
		count, _ := hash[nums[i]]
		if count > n/2 {
			return nums[i]
		}
	}
	return -1
}

//第二种解法
func majorityElement(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n < 2 {
		return nums[0]
	}
	value, count := 0, 0

	for i := 0; i < n; i++ {
		if count == 0 {
			value, count = nums[i], 1
		} else if value == nums[i] {
			count++
		} else {
			count--
		}
	}
	return value
}

//第三种解法
func majorityElement(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n < 2 {
		return nums[0]
	}
	sort.Ints(nums)
	return nums[n/2]
}