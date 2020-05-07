class Solution:
    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        def get_first_max_num(num):
            nums2_index = nums2.index(num)
            max_nums = [x for x in nums2[nums2_index:] if x > num]
            return max_nums[0] if max_nums else -1
        ret = []
        for i in nums1:
            ret.append(get_first_max_num(i))
        return ret