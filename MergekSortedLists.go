package main


type ListNode struct {
      Val int
      Next *ListNode
 }
 
 func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0{
		return nil
	}
	//全部都是nil
	allNil := true
	var mimNode *ListNode
	var minIndex int
	for index := range lists{
		if allNil && lists[index] != nil{
			mimNode = lists[index]
			allNil = false
			minIndex = index
		}else{
			if lists[index] != nil{
				if lists[index].Val < mimNode.Val{
					mimNode = lists[index]
					minIndex = index
				}
			}
		}
	}
	if allNil{
		return nil
	}else{
		node := new(ListNode)
		node.Val = mimNode.Val
		if mimNode.Next == nil{
			lists[minIndex] = nil
		}else{
			lists[minIndex] = lists[minIndex].Next
		}
		node.Next = mergeKLists(lists)
		return node
	}
}