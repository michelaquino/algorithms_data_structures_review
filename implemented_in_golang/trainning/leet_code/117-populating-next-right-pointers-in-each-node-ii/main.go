package main

// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii
func main() {
	// root := &Node{
	// 	Val: 1,
	// 	Left: &Node{
	// 		Val: 2,
	// 		Left: &Node{
	// 			Val: 4,
	// 		},
	// 	},
	// 	Right: &Node{
	// 		Val: 3,
	// 		Right: &Node{
	// 			Val: 5,
	// 		},
	// 	},
	// }
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
				Left: &Node{
					Val: 7,
				},
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Right: &Node{
				Val: 6,
			},
		},
	}

	connect(root)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	queue := []*Node{root, nil}

	for len(queue) > 0 {
		toProcess := queue[0]
		if toProcess == nil {
			break
		}

		queue = queue[1:]

		// linking
		for toProcess != nil {
			if toProcess.Left != nil {
				queue = append(queue, toProcess.Left)
			}

			if toProcess.Right != nil {
				queue = append(queue, toProcess.Right)
			}

			toLink := queue[0]
			queue = queue[1:]
			toProcess.Next = toLink

			toProcess = toLink
		}

		queue = append(queue, nil)
	}

	return root
}
