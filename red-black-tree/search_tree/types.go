package search_tree

type Color int

const (
	Red   Color = 0
	Black Color = 1
)

type Node struct {
	Value  interface{}
	color  Color
	parent *Node
	left   *Node
	right  *Node
}
