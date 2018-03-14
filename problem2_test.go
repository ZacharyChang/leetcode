package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		2,
		&ListNode{
			4,
			&ListNode{
				3,
				nil,
			},
		},
	}
	l2 := &ListNode{
		5,
		&ListNode{
			6,
			&ListNode{
				4,
				nil,
			},
		},
	}
	printList(addTwoNumbers(l1, l2))
}

func TestAddTwoNumbers2(t *testing.T) {
	l1 := &ListNode{
		1,
		&ListNode{
			8,
			nil,
		},
	}
	l2 := &ListNode{
		0,
		nil,
	}
	printList(addTwoNumbers(l1, l2))
	printList(addTwoNumbers(l2, l1))
}

func TestAddTwoNumbers3(t *testing.T) {
	l1 := &ListNode{
		5,
		nil,
	}
	l2 := &ListNode{
		5,
		nil,
	}
	printList(addTwoNumbers(l1, l2))
}

func TestAddTwoNumbers4(t *testing.T) {
	l1 := &ListNode{
		9,
		&ListNode{
			8,
			nil,
		},
	}
	l2 := &ListNode{
		1,
		nil,
	}
	printList(addTwoNumbers(l1, l2))
}

func TestAddTwoNumbers5(t *testing.T) {
	l1 := &ListNode{
		9,
		&ListNode{
			9,
			nil,
		},
	}
	l2 := &ListNode{
		1,
		nil,
	}
	//printList(addTwoNumbers(l1,l2))
	printList(addTwoNumbers(l2, l1))
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d->", node.Val)
		node = node.Next
	}
	fmt.Println()
}
