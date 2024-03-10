package helpers

func IsEmpty(str string) bool {
	if str == "" || len(str) == 0 {
		return true
	}
	return false
}
