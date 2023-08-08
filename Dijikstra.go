package main

import (
	"fmt"
	"math"
)

type Edge struct {
	To     int
	Weight int
}

func Dijkstra(graph [][]Edge, start int) []int {
	numVertices := len(graph)
	dist := make([]int, numVertices)
	visited := make([]bool, numVertices)

	for i := 0; i < numVertices; i++ {
		dist[i] = math.MaxInt32
		visited[i] = false
	}
	dist[start] = 0

	for count := 0; count < numVertices-1; count++ {
		u := -1
		for v := 0; v < numVertices; v++ {
			if !visited[v] && (u == -1 || dist[v] < dist[u]) {
				u = v
			}
		}

		visited[u] = true
		for _, edge := range graph[u] {
			if dist[u] != math.MaxInt32 && dist[u]+edge.Weight < dist[edge.To] {
				dist[edge.To] = dist[u] + edge.Weight
			}
		}
	}

	return dist
}

func main() {
	graph := [][]Edge{
		{{1, 4}, {2, 1}},
		{{3, 2}},
		{{1, 2}, {3, 5}},
		{{4, 3}},
		{{2, 3}},
	}

	startNode := 0
	distances := Dijkstra(graph, startNode)

	fmt.Println("Shortest distances from node", startNode)
	for i, distance := range distances {
		fmt.Printf("Node %d: %d\n", i, distance)
	}
}
