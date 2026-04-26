func groupAnagrams(strs []string) [][]string {
    result := [][]string{};
    normalizedMap := make(map[string][]string)
    for i:=0; i<len(strs);i++ {
		currentStr := strs[i]
		nmlStr := normalizedStr(currentStr)
		normalizedMap[nmlStr] = append(normalizedMap[nmlStr], currentStr)
    }

    for _, value :=  range normalizedMap {
        result = append(result,value);
    }
    return result;
}
func normalizedStr(str string) string{
    arr := []rune(str)
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })
    return string(arr);
}
