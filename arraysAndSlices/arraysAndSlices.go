package main

func main() {
}

func Sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumAll(slices ...[]int) []int {
	result := make([]int, len(slices))

	for _, nums := range slices {
		result = append(result, Sum(nums))
	}

	return result
}
