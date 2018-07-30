package letter

func Frequency(text string) (my_map map[rune]int) {
	my_map = make(map[rune]int)
	for _, val := range text {
		my_map[val]++
	}
	return
}

func ToChannel(str string, ch chan map[rune]int) {
	ch <- Frequency(str)
}

func ConcurrentFrequency(texts []string) (my_map map[rune]int) {
	my_map = make(map[rune]int)
	ch := make(chan map[rune]int)
	for i := range texts {
		go ToChannel(texts[i], ch)
	}
	for range texts {
		for symbol, number := range <-ch {
			my_map[symbol] += number
		}
	}
	return
}
