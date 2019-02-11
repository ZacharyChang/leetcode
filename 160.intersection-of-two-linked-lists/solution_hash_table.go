package leetcode

// hash table
func getIntersectionNode_2(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]bool, 0)
	for ; headA != nil; headA = headA.Next {
		nodeMap[headA] = true
	}
	for ; headB != nil; headB = headB.Next {
		if nodeMap[headB] {
			return headB
		}
	}
	return nil
}
