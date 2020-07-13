package valid_parentheses

import (
	"container/list"
)

func isValid(s string) bool {
	dict := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	li := list.New()
	len := len(s)
	for i := 0; i < len; i++ {
		v := s[i]
		if li.Len() > 0 {
			if x, ok := dict[v]; ok && li.Back().Value == x {
				li.Remove(li.Back())
			} else {
				li.PushBack(v)
			}
		} else {
			li.PushBack(v)
		}
	}
	return li.Len() == 0
}
