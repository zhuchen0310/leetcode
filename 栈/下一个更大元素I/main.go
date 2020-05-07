func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]int)
	for _, n1 := range nums1 {
		nums1Map[n1] = -1
	}
	// 遍历nums2 如果元素在nums1Map 中 则判断是否有下一个
	for i, n2 := range nums2 {
		if _, ok := nums1Map[n2]; ok {
			for _, n := range nums2[i:] {
				if n > n2 {
					nums1Map[n2] = n
					break
				}
			}
		}
	}
	ret := []int{}
	for _, k := range nums1 {
		ret = append(ret, nums1Map[k])
	}
	return ret
}