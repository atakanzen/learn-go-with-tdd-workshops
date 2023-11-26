package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	// for i := 0; i < len(numbers); i++ {
	// 	sum += numbers[i]
	// }

	return sum
}
