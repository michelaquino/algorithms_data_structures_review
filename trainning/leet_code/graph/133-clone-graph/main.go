package main

// https://leetcode.com/problems/clone-graph/

func main() {
	// [[2,4],[1,3],[2,4],[1,3]]

}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	queue := []*Node{}
	queue = append(queue, node)
	processed := map[int]*Node{}

	for {
		if len(queue) == 0 {
			break
		}

		toProcess := queue[0]
		queue = queue[1:]

		if _, alreadyDiscovered := processed[toProcess.Val]; alreadyDiscovered {
			continue
		}

		processed[toProcess.Val] = &Node{
			Val:       toProcess.Val,
			Neighbors: toProcess.Neighbors,
		}

		queue = append(queue, toProcess.Neighbors...)
	}

	for val, node := range processed {

		newNeighbors := []*Node{}
		for _, n := range node.Neighbors {
			neighbor := processed[n.Val]
			newNeighbors = append(newNeighbors, neighbor)
		}

		node.Neighbors = newNeighbors
		processed[val] = node

	}

	return processed[node.Val]
}
