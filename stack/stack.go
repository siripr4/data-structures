package stack

type Stack []string

func (s *Stack) IsEmpty() bool{
	return len(*s) == 0
}

func (s *Stack) Push(element string) {
	*s = append(*s, element)
}

func (s *Stack) Pop() (string, bool){
	if s.IsEmpty() {
		return "", false
	}
	lastIndex := len(*s) - 1
	element := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return element, true
}
