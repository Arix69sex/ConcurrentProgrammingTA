package main

import (
	"fmt"
	//"math/rand"
	"time"
	"sync"
)

func randSleep() {
	//time.Sleep(time.Duration(rand.Intn(15000)+5000) * time.Microsecond)
	time.Sleep(time.Duration(5000) * time.Microsecond)
}

var inUse map[int]bool = make(map[int]bool);
var wg sync.WaitGroup
var visited map[int]bool = make(map[int]bool)
var path []int
var queue []*Node 

type Graph struct {
	nodes []*Node
}

type Node struct {
	key int
	adjacent []*Node
}

func (n *Node) ProcessNode() {
	randSleep()
}

func createGraph(n int) *Graph {
	g := new(Graph)
	for i := 0; i < n; i++ {
		node := new(Node)
		node.key = i
		g.nodes = append(g.nodes, node)
	}
	return g
}

func createVisited(n int, visited map[int]bool) map[int]bool{
	for i := 0; i < n; i++ {
		visited[i] = false
	}
	return visited
}

func createInUse(n int, inUse map[int]bool) map[int]bool{
	for i := 0; i < n; i++ {
		inUse[i] = false
	}
	return inUse
}


func (g * Graph) Print() {
	for _,n := range g.nodes{
		fmt.Printf("\nNode %v: ", n.key)
		for _,node := range n.adjacent {
			fmt.Printf("%v ", node.key)
		}
	}
}

func createPath(from int, to int, g * Graph) *Graph{
	n := len(g.nodes)
	if from < n && to < n {
		g.nodes[from].adjacent = append(g.nodes[from].adjacent, g.nodes[to])
	}else {
		fmt.Printf("Nodes doesnt exist.")
	}
	return g
}


func (root *Node) BFS() {
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if !visited[node.key] {
			visited[node.key] = true
			fmt.Printf("CS Node %v ⚠️\n", node.key)
			for inUse[node.key] {
				return
			}
			inUse[node.key] = true
			node.ProcessNode()
			path = append(path, node.key)
			for i := 0; i < len(node.adjacent); i++ {
				queue = append(queue, node.adjacent[i])
			}
		}
		inUse[node.key] = false
	}
	wg.Done()
}

func (root *Node) BFSConcurrent() {
	path = append(path, root.key)
	fmt.Printf("CS Node %v ⚠️\n", root.key)
	visited[root.key] = true
	root.ProcessNode()
	for i := 0; i < len(root.adjacent); i++ {
		wg.Add(1)
		go root.adjacent[i].BFS()
	}
}

func main() {
	start := time.Now()
	g := createGraph(5)
	visited = createVisited(len(g.nodes), visited)
	g = createPath(0, 1, g)
	g = createPath(0, 2, g)
	g = createPath(1, 0, g)
	g = createPath(1, 2, g)
	g = createPath(1, 4, g)
	g = createPath(2, 0, g)
	g = createPath(2, 1, g)
	g = createPath(2, 3, g)
	g = createPath(3, 2, g)
	g = createPath(3, 4, g)
	g = createPath(4, 1, g)
	g = createPath(4, 3, g)
	g.Print()
	inUse = createInUse(len(g.nodes), inUse)
	g.nodes[3].BFSConcurrent()
	wg.Wait()
	fmt.Printf("Path %v.", path)

	fmt.Printf("\nTime elapsed: %v\n\n", time.Since(start))
}