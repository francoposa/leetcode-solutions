from typing import List


class Solution:
    """https://leetcode.com/problems/median-of-two-sorted-arrays/"""

    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        len_nums1 = len(nums1)
        len_nums2 = len(nums2)
        merged_max_index = len_nums1 + len_nums2 - 1

        # These two values will be equal for an even merged_max_index
        median_index_lower = merged_max_index // 2
        median_index_upper = median_index_lower + (merged_max_index % 2)

        merged_nums = []
        i, j, k = 0, 0, 0
        while k <= median_index_upper:
            if i == len_nums1:
                # We have reached the end of nums1
                merged_nums.append(nums2[j])
                j += 1
            elif j == len_nums2:
                # We have reached the end of nums2
                merged_nums.append(nums1[i])
                i += 1
            elif nums1[i] < nums2[j]:
                merged_nums.append(nums1[i])
                i += 1
            else:
                merged_nums.append(nums2[j])
                j += 1

            k += 1

        return (merged_nums[median_index_lower] + merged_nums[median_index_upper]) / 2
