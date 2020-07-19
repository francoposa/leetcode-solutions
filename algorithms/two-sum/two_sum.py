from typing import List


class Solution:
    def twoSumNaive(self, nums: List[int], target: int) -> List[int]:
        for i, num1 in enumerate(nums):
            if num1 > target:
                continue
            for j, num2 in enumerate(nums):
                if num1 + num2 == target:
                    return i, j

    def twoSum(self, nums: List[int], target: int) -> List[int]:
        num_indices = {num: [] for num in nums}
        for i, num in enumerate(nums):
            num_indices[num].append(i)

        diffs = [target - num for num in nums]
        for diff_num in diffs:
            diff_num_indices = num_indices.get(diff_num)
            other_num = target - diff_num
            other_num_indices = num_indices.get(other_num)
            if all((diff_num_indices, other_num_indices)):
                if diff_num == other_num:
                    # both candidate nums are half the target sum.
                    # we need to check that the num actually occurred twice
                    # in the original array before declaring success
                    if len(diff_num_indices) > 1:
                        return diff_num_indices
                    continue  # false positive, the candidate num only appeared once
                else:
                    return [
                        other_num_indices[0],
                        diff_num_indices[0],
                    ]
