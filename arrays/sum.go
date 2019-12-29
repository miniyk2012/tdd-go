package arrays

import "fmt"

func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int)  {
	sums = make([]int, len(numbersToSum))
	fmt.Print(numbersToSum)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}