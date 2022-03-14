package main

type Stack []string

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (string, bool) {
	if len(*s) == 0 {
		return "", false
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, true
}

func (s *Stack) Peek() (string, bool) {
	if len(*s) == 0 {
		return "", false
	}
	return (*s)[len(*s)-1], true
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
