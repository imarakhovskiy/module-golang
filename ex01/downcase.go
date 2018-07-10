package downcase

func Downcase(str string) (string, error) {
	length := len(str)
	str1 := ""
	
	for i := 0; i < length; i++ {
		if str[i] >= byte('A') && str[i] <= byte('Z') {
			str1 += string(str[i] + uint8(32))
		} else {
			str1 += string(str[i])
		}
	}
	return str1, nil
}
