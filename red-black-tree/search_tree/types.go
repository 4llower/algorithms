package search_tree

type Color int

const (
	Red   Color = 0
	Black Color = 1
)

type Node struct {
	value    int
	color    Color
	ancestor *Node
	left     *Node
	right    *Node
}
