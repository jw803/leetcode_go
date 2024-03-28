package linkedlist

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

func reverseKGroup(head *ListNode, k int) *ListNode {
	sentinel := &ListNode{
		Next: head,
	}

	last := sentinel

	for {
		// 1. 往後找 k - 1 步，找出一個Group的開頭與結尾
		groupEnd := getGroupEnd(head, k)
		if groupEnd == nil {
			break
		}

		nextGroupHead := groupEnd.Next
		// 2. 將Group內的element進行反轉
		reverse(head, nextGroupHead)

		// 3. 更新group前後和其他group的邊界
		last.Next = groupEnd
		head.Next = nextGroupHead

		// 4. 往下一個Group處理
		last = head
		head = nextGroupHead
	}

	return sentinel.Next
}

func getGroupEnd(head *ListNode, k int) *ListNode {
	for head != nil {
		k--
		if k == 0 {
			return head
		}
		head = head.Next
	}
	return nil
}

func reverse(start, end *ListNode) {
	last := start
	head := start.Next
	for head != end {
		next := head.Next
		head.Next = last
		last = head
		head = next
	}
}
