package stringutils

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
