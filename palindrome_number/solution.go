package palindrome_number

// 121 1 12 12 1
func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0){
		return false
	}
	tmp := int64(x)
	var y int64
	for tmp > y {
		y = y*10 + tmp%10
		tmp /= 10
	}
	return tmp == y || y/10 == tmp
}
