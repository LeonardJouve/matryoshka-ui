package dsl2

// Anything than can become a Node
type Element interface{ build() *Node }
