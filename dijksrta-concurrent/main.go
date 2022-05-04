package main

import (
	"fmt"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key int
	adjacent []*Vertex
}


func (g * Graph) addVertex(k int) {
	g.vertices = append(g.vertices, &Vertex{key: k})
}

