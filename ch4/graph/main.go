package main

var graph = make(map[string]map[string]bool)

func main() {
	addEdge("a", "b")
	addEdge("a", "b")
	addEdge("a", "b")
	addEdge("a", "b")
}

func hasEdges(from, to string) bool {
	return graph[from][to]
}

func addEdges(from, to string) {
	edge := graph[from] // sets value
	if edge == nil {		// checks value
		edge = make(map[string]bool)  // sets value 
		graph[from] = edge				
	}
	edge[to] = true
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}