package stringutils

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
