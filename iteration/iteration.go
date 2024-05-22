package iteration

func Repeat(character string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}

/*
func RepeatWithRange(times int) {
	for i := range times {
			fmt.Println("range", i)
	}
}
*/
