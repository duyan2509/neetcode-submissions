/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	nextMap:= make(map[*ListNode]int)
	for head!=nil{
		if nextMap[head]==0 {
			nextMap[head]++
		} else {
			return true
		}
		head=head.Next
	}
	return false
}
