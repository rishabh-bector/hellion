package main

import (
	"rapidengine/child"
)

// WorldTree contains the entire world map
type WorldTree struct {
	BlockNodes [WorldWidth][WorldHeight]BlockNode
}

// WorldNode contains all the data for one tile on the map
type BlockNode struct {
	WorldBlock  *child.ChildCopy
	BackBlock   *child.ChildCopy
	NatureBlock *child.ChildCopy
}

// NewWorldTree returns an empty WorldTree
func NewWorldTree() WorldTree {
	return WorldTree{}
}

func (tree *WorldTree) AddNode(x, y int, node BlockNode) {
	tree.BlockNodes[x][y] = node
}

func (tree *WorldTree) AddWorldBlock(x, y int, cpy *child.ChildCopy) {
	tree.BlockNodes[x][y].WorldBlock = cpy
}

func (tree *WorldTree) AddBackBlock(x, y int, cpy *child.ChildCopy) {
	tree.BlockNodes[x][y].BackBlock = cpy
}

func (tree *WorldTree) AddNatureBlock(x, y int, cpy *child.ChildCopy) {
	tree.BlockNodes[x][y].NatureBlock = cpy
}
