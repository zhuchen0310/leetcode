func findRepeatNumber(nums []int) int {
    if len(nums) == 0{
        return -1
    }
    // 创建一个cap为100000的数组（桶）
    newNums := make([]int, 100000)
    // 遍历nums, 并增加桶下标n的值
    for _, n := range nums{
        newNums[n] ++
        // 如果桶中下标n的值大于1则为重复数字
        if newNums[n] > 1{
            return n
        }
    }
    return -1
}