type Pair struct {
	Key   int
	Value int
}

func topKFrequent(nums []int, k int) []int {
    frequentMap := make(map[int]int);
    for i:=0;i<len(nums);i++{
        frequentMap[nums[i]]++;
    }
    result := []int{};
	pairs := []Pair{}

    for k, v := range frequentMap {
		pairs = append(pairs, Pair{k, v})
    }
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].Value > pairs[j].Value
    })
    for i:=0;i<len(pairs);i++{
        if i==k{
            return result;
        }
        result = append(result,pairs[i].Key);
    }
    return result;
}

