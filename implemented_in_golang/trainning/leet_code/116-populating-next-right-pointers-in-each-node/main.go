package main

// https://leetcode.com/problems/populating-next-right-pointers-in-each-node
func main() {

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
