package main

import (
	"time"
)

type Node struct {
	ID      int64
	Name    string
	Attr    map[string][]string
	Edge    []int64
	Created int64 // Seconds
}

func NewNode(id int64, name string) (node *Node) {
	node = new(Node)
	node.ID, node.Name, node.Created = id, name, time.Now().Unix()

	return
}

type Edge struct {
	ID, N1, N2 int64
	name       string
	Weight     float64
	Created    time.Time
}
