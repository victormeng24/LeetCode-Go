package to_lower_case

/**
将字符串中的大写字母转换为小写字母
 */
func toLowerCase(str string) string {
	res := []byte(str)
	for i := 0; i < len(res); i++ {
		if res[i] >= 'A' && res[i] <= 'Z' {
			res[i] += 32
		}
	}
	return string(res)
}
