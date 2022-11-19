package main

import "fmt"

func main() {
	g := graph{}
	g.add(5)
	g.add(1)
	g.add(3)
	g.add(2)
	g.add(4)

	g.connect(1, 2)
	g.connect(1, 5)
	g.connect(2, 1)
	g.connect(2, 5)
	g.connect(2, 4)
	g.connect(2, 3)
	g.connect(3, 2)
	g.connect(3, 4)
	g.connect(4, 5)
	g.connect(4, 2)
	g.connect(4, 3)
	g.connect(5, 1)
	g.connect(5, 2)
	g.connect(5, 4)

	g.print()
	values := dfs(g)
	fmt.Println("values: ", values)
}

func dfs(graph graph) []int {
	values := []int{}

	visited := map[int]bool{}
	for _, v := range graph.vertices {
		toVisit := v

		for {
			if toVisit == nil {
				break
			}

			if visited[toVisit.value] == true {
				toVisit = toVisit.next
				continue
			}

			values = append(values, toVisit.value)
			visited[toVisit.value] = true
			toVisit = toVisit.next
		}
	}

	return values
}

type vertex struct {
	value int
	next  *vertex
}

func (v *vertex) connect(value int) {
	newVertex := vertex{
		value: value,
	}

	lastVertex := v
	for {

		if lastVertex.next == nil {
			lastVertex.next = &newVertex
			return
		}

		lastVertex = lastVertex.next
	}
}

func (v vertex) print() {
	fmt.Printf("*%d ->", v.value)

	lastVertex := v.next
	for {
		if lastVertex == nil {
			break
		}

		fmt.Printf("%d ->", lastVertex.value)
		lastVertex = lastVertex.next
	}

	fmt.Println()
}

type graph struct {
	vertices []*vertex
}

func (g *graph) add(value int) {
	newVertex := vertex{
		value: value,
	}

	g.vertices = append(g.vertices, &newVertex)
}

func (g *graph) connect(vertex1, vertex2 int) {
	for _, v := range g.vertices {
		if v.value == vertex1 {
			v.connect(vertex2)
		}
	}
}

func (g graph) print() {
	for _, v := range g.vertices {
		v.print()
	}
}
