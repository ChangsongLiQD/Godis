package main

type listNode struct {
	prev  *listNode
	next  *listNode
	value interface{}
}

type List struct {
	head *listNode
	tail *listNode
	len  int
}
