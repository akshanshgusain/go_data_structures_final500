package Stack

import "container/list"

func Stack0009(s string) bool {
	st := &Stack{
		stack: list.New(),
	}

	for _, val := range s {
		if val == '(' || val == '{' || val == '[' {
			st.Push(string(val))
		} else if val == ')' {
			ele, _ := st.Top()
			if ele != "(" {
				return false
			}
			err := st.Pop()
			if err != nil {
				return false
			}
		} else if val == '}' {
			ele, _ := st.Top()
			if ele != "{" {
				return false
			}
			err := st.Pop()
			if err != nil {
				return false
			}
		} else if val == ']' {
			poppedValue, _ := st.Top()
			if poppedValue != "[" {
				return false
			}
			err := st.Pop()
			if err != nil {
				return false
			}
		}
	}
	return st.Size() == 0
}
