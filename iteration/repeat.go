package iteration

func Repeat(text string) string {
	answer := text
	for i := 0; i < 4; i++ {
		answer += text
	}

	return answer
}
