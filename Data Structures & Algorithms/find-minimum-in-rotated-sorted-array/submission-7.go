func findMin(nums []int) int {
	if nums[len(nums)-1]>nums[0] || len(nums)==1 {
		return nums[0]
	}
	left:=0;
	right:=len(nums)-1
	for left<=right {

		mid := left + (right-left)/2
		if left==mid && mid==right {
			return nums[left+1]
		}
		if nums[left]>nums[mid] {
			right=mid-1
			if right>1 && nums[right]<nums[right-1] {
				return nums[right]
			}
		} else if nums[mid]>nums[right] {
			left=mid+1
			if nums[left-1]>nums[left] {
				return nums[left]
			}
		}
		if mid>0 && nums[mid]<nums[mid-1]  {
			return nums[mid]
		}
	}
	return nums[left]
}


