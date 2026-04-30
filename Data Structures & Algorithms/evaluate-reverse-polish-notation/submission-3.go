func evalRPN(tokens []string) int {
	st := Stack{ data:[]int{}}
	for i:=0;i<len(tokens);i++ {
		if tokens[i]=="+" || tokens[i]=="-" ||tokens[i]=="*" || tokens[i]=="/" {
			switch tokens[i] {
				case "+":
					// code
					val1 := st.Pop();
					val2 := st.Pop()
					st.Push(val1+val2)
				case "-":
					// code
					val1 := st.Pop();
					val2 := st.Pop()
					st.Push(val2-val1)
				case "*":
					// code
					val1 := st.Pop();
					val2 := st.Pop()
					st.Push(val1*val2)
				case "/":
					// code
					val1 := st.Pop();
					val2 := st.Pop()
					st.Push(val2/val1)
				}
			
		} else {
			val, _ := strconv.Atoi(tokens[i])
			st.Push(val)
		}
	}
	return st.Pop()
}
type Stack struct {
	data []int
}
func (s *Stack) Push(val int){
	s.data = append(s.data, val)
}
func (s *Stack) Pop() int {
	top:=s.data[len(s.data)-1]
	s.data=s.data[:len(s.data)-1]
	return top
}
func (s *Stack) GetSize() int {
	return len(s.data)
}