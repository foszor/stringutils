package stringutils

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
