from typing import Dict, List


class Solution:
    """https://leetcode.com/problems/two-sum/"""

    def twoSumNaive(self, nums: List[int], target: int) -> List[int]:
        for i in range(len(nums)):
            for j in range(i + 1, len(nums)):
                if nums[i] + nums[j] == target:
                    return [i, j]

    def twoSum(self, nums: List[int], target: int) -> List[int]:
        num_target_difference_index_map: Dict[int, int] = {}
        for i, num in enumerate(nums):
            num_target_difference = target - num
            if num in num_target_difference_index_map:
                # num is equal to the target difference of a previous num in the list
                return [num_target_difference_index_map[num], i]
            else:
                num_target_difference_index_map[num_target_difference] = i
