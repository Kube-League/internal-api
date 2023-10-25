package utils

func ArrStringContains(arr []string, s string) bool {
	for _, sc := range arr {
		if sc == s {
			return true
		}
	}
	return false
}
