package Stack

/* Stack follows the principle of "First in - Last out" or FILO, where the first element added to the stack is the
last one to be removed. Example: CD Stack  */

type Stack struct {
	items []int
}

// Push(Inserting into stack)
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop(Releasing from stack)
func (s *Stack) Pop() int {
	poped := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return poped
}
