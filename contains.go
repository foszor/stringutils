package stringutils

// Contains checks if a string array has a value
func Contains(arr []string, value string) bool {
	contains := false
	for _, a := range arr {
		if a == value {
			contains = true
			break
		}
	}

	return contains
}
