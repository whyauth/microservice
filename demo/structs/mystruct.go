package mystruct

type treeNode struct {
	value       int
	left, right *treeNode
}

func NewTreeNode(value int) *treeNode {
	return &treeNode{value: value}
}

func Start() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = NewTreeNode(2)

	/*nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}*/

	// fmt.Println(root)

}
