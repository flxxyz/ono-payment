package utils

func UpperFirstChar(str string) string {
	if len(str) > 0 {
		runelist := []rune(str)
		if int(runelist[0]) >= 97 && int(runelist[0]) <= 122 {
			runelist[0] = rune(int(runelist[0]) - 32)
			str = string(runelist)
		}
	}
	return str
}
