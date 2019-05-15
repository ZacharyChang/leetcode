package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//-----------------------------------------
// 将两个链表依次进行相加，有进位，返回新的链表
//-----------------------------------------
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l3 := &ListNode{}
	cur := l3  // 当前操作Node
	carry := 0 // 进位
	temp := 0
	for l1 != nil {
		if l2 != nil {

			if l1.Val+l2.Val+carry > 9 {
				temp = l1.Val + l2.Val + carry - 10
				carry = 1
			} else {
				temp = l1.Val + l2.Val + carry
				carry = 0
			}
			cur.Next = &ListNode{
				Val: temp,
			}
			cur = cur.Next
		} else { // l1比l2长
			if carry == 1 {
				increase(l1)
			}
			cur.Next = l1
			break
		}
		l1 = l1.Next
		l2 = l2.Next
		// 处理当l1和l2恰好最后一位相加进1的情况
		if carry == 1 && l1 == nil && l2 == nil {
			cur.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
		}
	}
	// l2比l1长
	if l2 != nil {
		if carry == 1 {
			increase(l2)
		}
		cur.Next = l2

	}
	return l3.Next
}

//-------------------------------------------
// 给指定ListNode末位进位，如果高位满则依次进位
//-------------------------------------------
func increase(l *ListNode) {
	l.Val = l.Val + 1
	if l.Val > 9 {
		l.Val = 0
		if l.Next != nil {
			increase(l.Next)
		} else {
			l.Next = &ListNode{
				1,
				nil,
			}
		}
	}
}
