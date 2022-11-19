package main

import "fmt"

func main() {
	g := newGraph()
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
		if lastVertex.value == value {
			return
		}

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
	vertices map[int]*vertex
}

func newGraph() graph {
	return graph{vertices: make(map[int]*vertex)}
}

func (g *graph) add(value int) {
	newVertex := vertex{
		value: value,
	}

	g.vertices[value] = &newVertex
}

func (g *graph) connect(vertex1Value, vertex2Value int) error {
	vertex1, vertex1Exists := g.vertices[vertex1Value]
	if !vertex1Exists {
		return fmt.Errorf("vertex %d doesn't exist", vertex1Value)
	}

	vertex2, vertex2Exists := g.vertices[vertex2Value]
	if !vertex2Exists {
		return fmt.Errorf("vertex %d doesn't exist", vertex2Value)
	}

	vertex1.connect(vertex2Value)
	vertex2.connect(vertex1Value)
	return nil
}

func (g graph) print() {
	for _, v := range g.vertices {
		v.print()
	}
}
