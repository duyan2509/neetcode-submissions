func twoSum(nums []int, target int) []int {
    for i:=0; i < len(nums) - 1 ; i++ {
        for nextI:=i+1; nextI < len(nums) ; nextI ++ {
            if nums[i]+nums[nextI]==target {
                return []int{i, nextI};
            }
        }
    }
    return []int{0, 0};
}
