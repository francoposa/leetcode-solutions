package main

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}

func twoSum(nums []int, target int) []int {
	targetDifferenceIndices := map[int]int{}

	for i, num := range nums {
		targetDifference := target - num
		if targetDifferenceIndex, ok := targetDifferenceIndices[num]; ok {
			// num is equal to the target difference of a previous num in the slice
			return []int{targetDifferenceIndex, i}
		} else {
			targetDifferenceIndices[targetDifference] = i
		}
	}
	return []int{}
}
