package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

// bellow is first version of addTwoNumbers
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil {
// 		return l2
// 	}
// 	if l2 == nil {
// 		return l1
// 	}
// 	var result ListNode
// 	var current = &result
// 	var node1 = l1
// 	var node2 = l2
// 	var remember int = 0
// 	nodeValue := func(node *ListNode) int {
// 		if node == nil {
// 			return 0
// 		}
// 		return node.Val
// 	}
// 	nextNode := func(node *ListNode) *ListNode {
// 		if node == nil {
// 			return nil
// 		}
// 		return node.Next
// 	}
// 	for ok := true; ok; {
// 		sum := remember + nodeValue(node1) + nodeValue(node2)
// 		remember = 0

// 		if sum > 9 {
// 			remember = 1
// 			sum -= 10
// 		}

// 		current.Val = sum
// 		node1 = nextNode(node1)
// 		node2 = nextNode(node2)
// 		if node1 == nil && node2 == nil {
// 			if remember > 0 {
// 				current.Next = &ListNode{
// 					Val:  remember,
// 					Next: nil,
// 				}
// 			}
// 			break
// 		}
// 		current.Next = &ListNode{}
// 		current = current.Next
// 	}

// 	return &result
// }

// latest version
// result:
// Runtime: 7 ms, faster than 94.44% of Go online submissions for Add Two Numbers.
// Memory Usage: 4.4 MB, less than 93.66% of Go online submissions for Add Two Numbers.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var remember int = 0
	l1Node := l1
	l2Node := l2
	for {
		sum := l1Node.Val + l2Node.Val + remember
		if sum >= 10 {
			sum -= 10
			remember = 1
		} else {
			remember = 0
		}
		l1Node.Val = sum

		if l1Node.Next != nil && l2Node.Next != nil {
			l1Node = l1Node.Next
			l2Node = l2Node.Next
			continue
		}

		if l1Node.Next != nil && l2Node.Next == nil {
			if remember > 0 {
				l1Node = l1Node.Next
				l2Node = &ListNode{
					Val:  remember,
					Next: nil,
				}
				remember = 0
				continue
			}

			return l1
		}

		if l1Node.Next == nil && l2Node.Next != nil {
			if remember > 0 {
				l1Node.Next = &ListNode{
					Val:  remember,
					Next: nil,
				}
				remember = 0
				l1Node = l1Node.Next
				l2Node = l2Node.Next
				continue
			} else {
				l1Node.Next = l2Node.Next
				return l1
			}
		}

		if l1Node.Next == nil && l2Node.Next == nil {
			if remember > 0 {
				l1Node.Next = &ListNode{
					Val:  remember,
					Next: nil,
				}
				return l1
			} else {
				return l1
			}
		}
	}
}
