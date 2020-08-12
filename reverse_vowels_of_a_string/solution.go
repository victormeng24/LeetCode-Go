package reverse_vowels_of_a_string

var vowelMap = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}
func reverseVowels(s string) string {
	arr := []byte(s)
	i, j := 0, len(arr) - 1
	for i < j {
		if !vowelMap[arr[i]] {
			i++
			continue
		}
		if !vowelMap[arr[j]] {
			j--
			continue
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return string(arr)
}