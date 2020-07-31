package string_rotation_lcci

import "strings"

func isFlipedString(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1 + s1, s2)
}