func isValid(s string) bool {
	chars :=[]byte(s)
    st := Stack{}
	for i:=0;i<len(chars);i++{
		if chars[i]=='(' || chars[i]=='[' || chars[i]=='{' {
			st.Push(chars[i])
		}
		if chars[i]==')' || chars[i]==']' || chars[i]=='}' {
			if st.Empty() {
				return false
			}
			current:= st.Pop()
			if chars[i]==')' && current!='('{
				return false
			}
			if chars[i]==']' && current!='['{
				return false
			}
			if chars[i]=='}' && current!='{'{
				return false
			}
		}
	}
	return st.Empty();
}
type Stack struct {
	data []byte
}

func (s *Stack) Push(x byte) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() byte {
	n := len(s.data)

	if n == 0 {
		return 0
	}

	val := s.data[n-1]
	s.data = s.data[:n-1]

	return val
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

