/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reversePrint(head *ListNode) []int {
    if head == nil{
        return nil
    }
    // 从头到尾打印链表
    ret := []int{}
    for head != nil{
        ret = append(ret, head.Val)
        head = head.Next
    }
    // 反转数组
    i := 0
    j := len(ret)-1
    for i < j {
        ret[i], ret[j] = ret[j], ret[i]
        i ++
        j --
    } 
    return ret 
}