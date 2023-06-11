package main

/*
*
  - Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	seen := map[int]*ListNode{}                        //哈希表
	prefix := 0                                        // 前缀和
	for node := dummy; node != nil; node = node.Next { //首先先将所有node对应的前缀和进行计算。 如果这里有前缀和相同的，将直接替换 例如1 -2 2 3 -2：-2 的前缀和 是-1 ，2 的前缀和是 1 ，3 的前缀和是4 ，这里就出现了前缀和的矛盾。所以1 的值变成了第三个节点
		prefix += node.Val
		seen[prefix] = node
	}
	prefix = 0                                         //重置前缀和
	for node := dummy; node != nil; node = node.Next { //再次遍历
		prefix += node.Val
		node.Next = seen[prefix].Next //直接计算结果
	}
	return dummy.Next
}

func main() {

}
