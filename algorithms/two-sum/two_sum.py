from typing import List


class Solution:
    def twoSumNaive(self, nums: List[int], target: int) -> List[int]:
        for i, num1 in enumerate(nums):
            if num1 > target:
                continue
            for j, num2 in enumerate(nums):
                if num1 + num2 == target:
                    return [i, j]

    def twoSum(self, nums: List[int], target: int) -> List[int]:
        num_indices = {num: i for i, num in enumerate(nums)}
        diffs = [target - num for num in nums]
        for diff in diffs:
            if diff in num_indices:
                return [
                    num_indices[target - diff],
                    num_indices[diff],
                ]
