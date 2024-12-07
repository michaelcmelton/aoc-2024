package day01

func ProcessAnswerPart2(path string) int {
	var answer int
	inputs, _ := readInput(path)
	freqs := make(map[int64]int)

	for _, num := range inputs[0] {
		if _, ok := freqs[num]; !ok {
			freqs[num] = 0
		}
	}

	for _, num := range inputs[1] {
		if _, ok := freqs[num]; ok {
			freqs[num] = freqs[num] + 1
		}
	}

	for _, num := range inputs[0] {
		if val, ok := freqs[num]; ok {
			answer += int(num) * val
		}
	}

	return answer
}
