package iteration

// Repeats an entered string n times
func Repeat(character string, n int) string {
	var repeated string

	for i := 0; i < n; i++ {
		repeated += character
	}
	return repeated
}
