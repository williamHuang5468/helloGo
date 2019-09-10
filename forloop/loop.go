package forloop

func Repeat(letter string) string {
	r := ""
	for i := 0; i <= 5; i++ {
		r += letter
	}
	return r
}
