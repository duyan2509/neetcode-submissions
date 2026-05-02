type TimeMap struct {
	data map[string][]Pair
}
type Pair struct {
	timestamp int
	value     string
}
func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]Pair),
	}
}


func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.data[key] = append(this.data[key], Pair{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	pairs := this.data[key]
	l, r:=0, len(pairs)-1
	rs:=""
	for l<=r {
		m:= l+(r-l)/2
		if pairs[m].timestamp>timestamp {
			r=m-1
		} else if pairs[m].timestamp<=timestamp {
			rs=pairs[m].value
			l=m+1
		} 
	}

	return rs
}
