package main

import "fmt"

func main() {
	g := newGraph()
	g.add(1)
	g.add(2)
	g.add(3)
	g.add(4)
	g.add(5)
	g.add(6)
	g.add(7)
	g.add(8)

	g.connectMany(1, []int{2, 7, 8})
	g.connectMany(2, []int{1, 7, 3, 5})
	g.connectMany(3, []int{2, 4, 5})
	g.connectMany(4, []int{3, 5})
	g.connectMany(5, []int{2, 3, 4, 6})
	g.connectMany(7, []int{1, 2})

	g.print()
	bfs(g, 1)
}

func bfs(graph graph, startVertexValue int) {
	processed := map[int]uint8{}
	discovered := map[int]uint8{}
	queue := []*vertex{}

	startVertex := graph.vertices[startVertexValue]
	queue = append(queue, startVertex)
	discovered[startVertexValue] = 1

	for {
		if len(queue) == 0 {
			break
		}

		toProcess := queue[0]

		queue = queue[1:]
		process(toProcess)
		processed[toProcess.value] = 1

		neighboor := graph.vertices[toProcess.value]
		for {
			if neighboor == nil {
				break
			}

			if _, exists := discovered[neighboor.value]; !exists {
				discovered[neighboor.value] = 1
				queue = append(queue, neighboor)
			}

			neighboor = neighboor.next
		}

	}
}

func process(vertex *vertex) {
	fmt.Printf("vertex %d processed\n", vertex.value)
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

func (g *graph) connectMany(originValue int, destinations []int) error {
	if _, originExists := g.vertices[originValue]; !originExists {
		return fmt.Errorf("vertex %d doesn't exist", originValue)
	}

	for _, destinationValue := range destinations {
		g.connect(originValue, destinationValue)
	}

	return nil
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
