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

// JoinFancy joins a string array in a fancy way
func JoinFancy(arr []string) string {
	var r string
	for i, a := range arr {
		if i > 0 && i == len(arr)-1 {
			r += " and "
		} else if len(r) > 0 {
			r += ", "
		}

		r += a
	}

	return r
}

// JoinQuotes joins an array with quotes around each item
func JoinQuotes(arr []string, sep string) string {
	var r string
	for _, a := range arr {
		if len(r) > 0 {
			r += sep
		}

		r += "\"" + a + "\""
	}

	return r
}

// RemoveIndex removes an index from a string array
func RemoveIndex(arr []string, index int) []string {
	arr[len(arr)-1], arr[index] = arr[index], arr[len(arr)-1]
	return arr[:len(arr)-1]
}

// RemoveValue removes a value from a string array
func RemoveValue(arr []string, str string) []string {
	index := -1
	for i, s := range arr {
		if s == str {
			index = i
			break
		}
	}

	if index > -1 {
		return RemoveIndex(arr, index)
	}

	return arr
}
