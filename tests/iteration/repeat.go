package iteration

const repeatCount = 5

func Repeat(ch string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += ch
	}
	return repeated
}
