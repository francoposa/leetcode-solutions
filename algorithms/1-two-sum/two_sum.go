package main

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}

// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	targetDiffIndices := map[int]int{}
	for i, num := range nums {
		// check for existing target diff for this num
		diff := target - num
		if diffIndex, ok := targetDiffIndices[diff]; ok {
			// a previous num is the diff between num and target
			return []int{i, diffIndex}
		} else {
			targetDiffIndices[num] = i
		}
	}
	return []int{}
}
