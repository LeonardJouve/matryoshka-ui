package dsl

type Element interface {
	build() *Node
}
