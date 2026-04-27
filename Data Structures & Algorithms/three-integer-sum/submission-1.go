func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result :=[][]int{}
	for i:=0;i<len(nums);i++{
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target:=-nums[i];
		left:=i+1;
		right:=len(nums)-1;
		for left<right {


			sum:=nums[left]+nums[right];
			if sum < target {
				left++;
			} else if sum > target {
				right--;
			} else {
				result = append(result,[]int{nums[i],nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return result
}
