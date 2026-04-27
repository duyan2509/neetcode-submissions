func maxArea(heights []int) int {
	result :=0;
	left:=0; 
	right:=len(heights)-1;
	for left<right {
		w:=right-left
		h:=0
		if heights[right]>heights[left] {
			h=heights[left]
			left++
		} else {
			h=heights[right]
			right--
		}
		area:= w * h
		if area>result {
			result = area
		}
	}
	return result;
}
