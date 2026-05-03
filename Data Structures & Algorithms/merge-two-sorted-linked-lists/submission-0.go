/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	
	if list1 == nil {
    	return list2
	}
	if list2 == nil {
		return list1
	}
	var rs *ListNode;
	if list1.Val>list2.Val {
		rs=list2
		list2=list2.Next
	} else {
		rs=list1
		list1=list1.Next
	}
	head := rs

	for list1 != nil && list2 != nil {
		if list1.Val>list2.Val {
			rs.Next=list2
			list2=list2.Next
		} else {
			rs.Next=list1
			list1=list1.Next
		}
		rs = rs.Next

	}

	if list1 != nil {
		rs.Next = list1
	} else {
		rs.Next = list2
	}
	return head
}
