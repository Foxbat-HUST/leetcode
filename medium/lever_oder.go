package medium

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil{
		return [][]int{}
	}

	var levelNodes = []*Node{root}
	var result = [][]int{}
	for {
		levelNodesValue := []int{}
		nextLevelNodes := []*Node{}
		for _, v := range levelNodes {
			levelNodesValue = append(levelNodesValue, v.Val)
			if len(v.Children) > 0 {
				nextLevelNodes = append(nextLevelNodes, v.Children...)
			}
		}
		result = append(result, levelNodesValue)
		if len(nextLevelNodes) == 0 {
			break
		}
		levelNodes = nextLevelNodes
	}

	return result
}
