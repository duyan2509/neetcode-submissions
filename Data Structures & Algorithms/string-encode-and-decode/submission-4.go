type Solution struct{
	sizes []int;
}
//aa bb
func (s *Solution) Encode(strs []string) string {
	s.sizes = []int{}
	result := "";
	currentIndex := 0;
	for i := 0; i<len(strs); i++ {
		s.sizes = append(s.sizes, currentIndex);
		result+=strs[i];
		currentIndex += len(strs[i])
	}
	s.sizes = append(s.sizes,currentIndex)
	return result;
}

func (s *Solution) Decode(encoded string) []string {
	var result []string;
	prev := 0;
	for i:=1;i<len(s.sizes);i++{
		result = append(result, encoded[prev:s.sizes[i]]);
		prev =s.sizes[i];
	}

	return result;
}

