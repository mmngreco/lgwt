package sum

func Sum(numbers []int) (total int) {

	for _, number := range numbers {
		total += number
	}
	return total
}

// func SumAll(numbersToSum ...[]int) []int {
//
// 	lengthOfNumbers := len(numbersToSum)
// 	sums := make([]int, lengthOfNumbers)
//
// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}
// 	return sums
// }

func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
