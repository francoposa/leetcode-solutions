from typing import Dict, List, Optional


class Solution:
    """https://leetcode.com/problems/two-sum/"""

    def twoSumNaive(self, nums: List[int], target: int) -> List[int]:
        for i in range(len(nums)):
            for j in range(i + 1, len(nums)):
                if nums[i] + nums[j] == target:
                    return [i, j]

    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # building it ourselves appears to be faster than using defaultdict,
        # although I have not checked if that slowdown comes from the import
        nums_item_indices: Dict[int, List[int]] = {}
        for i, num in enumerate(nums):
            if num in nums_item_indices:
                nums_item_indices[num].append(i)
            else:
                nums_item_indices[num] = [i]

        for num in nums:
            num_target_difference = target - num
            num_target_difference_indices: Optional[List[int]] = nums_item_indices.get(
                num_target_difference
            )
            if num_target_difference_indices:
                # at least one member of the original array is equal to `target - num`
                # that is, at least one value in the original array can be added to
                # the current `num` value to produce the target. This would be enough
                # to declare we have a solution, except for one special case:
                if num_target_difference == num:
                    # `num` value is equal to `num_target_difference` value
                    # check that the value actually occurred twice in the array
                    if len(num_target_difference_indices) > 1:
                        return num_target_difference_indices
                    continue
                return [nums_item_indices[num][0], num_target_difference_indices[0]]
