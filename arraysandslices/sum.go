package arraysandslices

func Sum(nums []int) int {
	sum := 0
	for _, number := range nums {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthofNumbers := len(numbersToSum)
	sums := make([]int, lengthofNumbers)
	return nil
}
