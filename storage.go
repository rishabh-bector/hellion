package main

import (
	"rapidengine/child"
)

// WorldTree contains the entire world map
type WorldTree struct {
	blockNodes [WorldWidth][WorldHeight]BlockNode
}

// BlockNode contains all the data for one tile on the map
type BlockNode struct {
	worldBlock  *child.ChildCopy
	backBlock   *child.ChildCopy
	natureBlock *child.ChildCopy
}

// NewWorldTree returns an empty WorldTree
func NewWorldTree() WorldTree {
	w := WorldTree{}
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			w.AddWorldBlock(x, y, &child.ChildCopy{
				ID: "00000",
			})
			w.AddBackBlock(x, y, &child.ChildCopy{
				ID: "00000",
			})
			w.AddNatureBlock(x, y, &child.ChildCopy{
				ID: "00000",
			})
		}
	}
	return w
}

func (tree *WorldTree) AddNode(x, y int, node BlockNode) {
	tree.blockNodes[x][y] = node
}

func (tree *WorldTree) AddWorldBlock(x, y int, cpy *child.ChildCopy) {
	tree.blockNodes[x][y].worldBlock = cpy
}

func (tree *WorldTree) AddBackBlock(x, y int, cpy *child.ChildCopy) {
	tree.blockNodes[x][y].backBlock = cpy
}

func (tree *WorldTree) AddNatureBlock(x, y int, cpy *child.ChildCopy) {
	tree.blockNodes[x][y].natureBlock = cpy
}

func (tree *WorldTree) GetWorldBlock(x, y int) *child.ChildCopy {
	return tree.blockNodes[x][y].worldBlock
}

func (tree *WorldTree) GetBackBlock(x, y int) *child.ChildCopy {
	return tree.blockNodes[x][y].backBlock
}

func (tree *WorldTree) GetNatureBlock(x, y int) *child.ChildCopy {
	return tree.blockNodes[x][y].natureBlock
}

func (tree *WorldTree) GetBlockName(x, y int) string {
	return GetNameFromID(tree.blockNodes[x][y].worldBlock.ID[:4])
}

func (tree *WorldTree) GetBlockID(x, y int) string {
	return tree.blockNodes[x][y].worldBlock.ID
}

func (tree *WorldTree) GetBlockOrientation(x, y int) string {
	return GetOrientationFromID(tree.blockNodes[x][y].worldBlock.ID[4:])
}

func (tree *WorldTree) SetBlockOrientation(x, y int, orient string) {
	tree.blockNodes[x][y].worldBlock.ID = tree.blockNodes[x][y].worldBlock.ID[:4] + orient
}

func (tree *WorldTree) SetBackBlockOrientation(x, y int, orient string) {
	tree.blockNodes[x][y].backBlock.ID = tree.blockNodes[x][y].backBlock.ID[:4] + orient
}

func (tree *WorldTree) GetDarkness(x, y int) float32 {
	return tree.blockNodes[x][y].worldBlock.Darkness
}

func (tree *WorldTree) SetDarkness(x, y int, darkness float32) {
	tree.blockNodes[x][y].worldBlock.Darkness = darkness
	tree.blockNodes[x][y].backBlock.Darkness = darkness
	tree.blockNodes[x][y].natureBlock.Darkness = darkness
}
