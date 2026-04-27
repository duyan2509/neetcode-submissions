func isPalindrome(s string) bool {
	s=strings.ToLower(s)
	chars := []byte(s)
	left:=0;
	right:=len(chars)-1;
	for left<right {
		if (chars[left]<'a' || chars[left] >'z') && (chars[left]<'0' || chars[left] >'9') {
			left++;
			continue;
		} 
		if (chars[right]<'a' || chars[right] >'z') && (chars[right]<'0' || chars[right] >'9') {
			right--;
			continue;
		} 

		if chars[left]!=chars[right] {
			return false
		} else {
			left++;
			right--;
		}
	}


	return true;
}


