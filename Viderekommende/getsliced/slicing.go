package getsliced

func GetArrayStartingAt(n int) [10]int {
	var nums [10]int

	for i := 0; i < 10; i++ {
		nums[i] = n
	}

	return nums
}
