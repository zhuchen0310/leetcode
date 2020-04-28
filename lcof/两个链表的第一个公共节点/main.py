# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        curA = headA
        curB = headB
        lenA = 0
        lenB = 0
        # 先统计各链表长度
        while curA != None:
            curA = curA.next
            lenA += 1
        while curB != None:
            curB = curB.next
            lenB += 1

        # 如果A 比 B 长， 则A先走n步
        if lenA > lenB:
            for i in range(lenA-lenB):
                headA = headA.next 
        else:
            for i in range(lenB-lenA):
                headB = headB.next
        # 两个链表同时走
        while headA != None and headB != None:
            if headA == headB:
                return headA
            headA = headA.next
            headB = headB.next 
        return None