from typing import List


class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:

        len_nums1 = len(nums1)
        len_nums2 = len(nums2)
        merged_len = len_nums1 + len_nums2 - 1

        medianIndexLower = merged_len // 2
        medianIndexUpper = medianIndexLower + (merged_len % 2)

        merged_nums = []
        i, j, k = 0, 0, 0
        while k <= medianIndexUpper:
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

        return (merged_nums[medianIndexLower] + merged_nums[medianIndexUpper]) / 2
