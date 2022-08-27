package linked_list_algorithms

/*ListNode given by leetcode.*/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*hasCycle
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again
by continuously following the next pointer. Internally, pos is used to denote the index of
the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.
*/
func hasCycle(head *ListNode) bool {
	slowPtr := head
	fastPtr := head
	for slowPtr != nil && fastPtr != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
		if slowPtr == fastPtr {
			return true
		}
	}
	return false
}

/*mergeTwoLists
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by
splicing together the nodes of the first two lists.
Return the head of the merged linked list.
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	chosenList := list2
	otherList := list1
	if list1.Val < list2.Val {
		chosenList = list1
		otherList = list2
	}
	chosenListPrevious := chosenList
	chosenList = chosenList.Next
	for chosenList != nil && otherList != nil {
		if chosenList.Val > otherList.Val {
			chosenListPrevious.Next = otherList
			otherList = otherList.Next
			chosenListPrevious.Next.Next = chosenList
		}
	}
}
