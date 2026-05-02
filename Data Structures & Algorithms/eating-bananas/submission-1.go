func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)

	right := piles[len(piles)-1]
	rs:=0
	left:=1;
	for left<=right && right>0 {
		mid:=left+(right-left)/2;
		consumeHours:= consume(piles, mid)
		if consumeHours>h {
			left=mid+1
		} else {
			rs=mid
			right=mid-1
		}
	}
	return rs
}

func consume(piles []int, bananaPerHour int) int {
	hours:=0;
	for i:=0;i<len(piles);i++ {
		hours+=(piles[i]+bananaPerHour-1)/bananaPerHour
	}
	return hours
}

