func longestConsecutive(nums []int) int {
	if len(nums)==0{
		return 0
	}
	exists := make(map[int]bool)

	for i:=0;i<len(nums);i++{
		exists[nums[i]]=true;
	}
	result :=1;
	for k, _ := range exists {
		tmp:=0;
		current := k;
		for exists[current] {
			tmp++;
			if exists[current-1] {
				current--;
			} else {
				break;
			}
		}

		if tmp>result{
			result=tmp;
		}
	}
	return result
}
