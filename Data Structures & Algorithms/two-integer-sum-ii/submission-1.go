func twoSum(numbers []int, target int) []int {
	right := len(numbers)-1;
	left :=0;
	for left<right {
		sum := numbers[left]+numbers[right];
		if sum > target {
			right--;
		} else if sum < target {
			left++;
		} else {
			return []int{left+1,right+1}
		}
	}
	return []int{}
}
