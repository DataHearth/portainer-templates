package utils

func CheckStringArray(str string, arr []string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}

	return false
}