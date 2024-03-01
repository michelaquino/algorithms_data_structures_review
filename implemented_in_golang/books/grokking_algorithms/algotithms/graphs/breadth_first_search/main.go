package main

import "fmt"

func main() {
	graph := map[string][]string{
		"city-1":  {"city-2", "city-3", "city-4"},
		"city-2":  {"city-5", "city-6"},
		"city-3":  {"city-6"},
		"city-4":  {"city-7", "city-8"},
		"city-5":  {},
		"city-6":  {"city-9"},
		"city-7":  {"city-10"},
		"city-8":  {},
		"city-9":  {"city-10"},
		"city-10": {},
	}

	path := breadthFirstSearch(graph, "city-1", "city-8")
	fmt.Println("path: ", path)
}

func breadthFirstSearch(graph map[string][]string, start, end string) []string {
	path := []string{}

	fathers := map[string][]string{}
	queueToSearch := []string{start}

	nodeProcessed := map[string]bool{}
	for true {
		if len(queueToSearch) == 0 {
			return path
		}

		nodeToSearch := queueToSearch[0]
		if nodeToSearch == end { // found
			return buildPath(fathers, nodeToSearch)
		}

		if alreadyProcessed := nodeProcessed[nodeToSearch]; alreadyProcessed {
			queueToSearch = queueToSearch[1:]
			continue
		}

		nodeProcessed[nodeToSearch] = true
		neighbors := graph[nodeToSearch]
		for _, n := range neighbors {
			fathers[n] = append(fathers[n], nodeToSearch)
		}

		queueToSearch = append(queueToSearch[1:], neighbors...)
	}

	return path
}

func buildPath(fathers map[string][]string, finalNode string) []string {
	path := []string{}

	// montar path
	pathNode := finalNode

	path = append(path, pathNode)
	for true {
		previousNeighboors := fathers[pathNode]

		if len(previousNeighboors) == 0 {
			break
		}

		previous := previousNeighboors[0]
		path = append([]string{previous}, path...)
		pathNode = previous
	}

	return path
}
