package lab2

type Stack []float64

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(item float64) {
	*s = append(*s, item)
}

func (s *Stack) Pop() (float64, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
