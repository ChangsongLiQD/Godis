package main

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

type List struct {
	head *Node
	tail *Node
	len  int
}
