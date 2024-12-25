package functions

// Count the number of times that <value> appears in <slice>
func Count_in_slice(value int, slice []int) int {
	count := 0
	for i := range len(slice) {
		if slice[i] == value {
			count += 1
		}
	}
	return count
}

// Check if string e is one of the strings in slice s
func StringInSlice(s []string, e string) bool {
	for i := range len(s) {
		if s[i] == e {
			return true
		}
	}
	return false
}
