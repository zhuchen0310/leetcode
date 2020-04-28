/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := 0
	lenB := 0
	curA := headA
	curB := headB
	// 分别统计A B长度
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}
	// 长度长的先走n步
	if lenA > lenB {
		for i := lenA - lenB; i > 0; i-- {
			headA = headA.Next
		}
	} else if lenA < lenB {
		for i := lenB - lenA; i > 0; i-- {
			headB = headB.Next
		}
	}
	// 同时走
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}