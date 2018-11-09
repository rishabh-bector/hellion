package main

import (
	"rapidengine/child"
)

//  --------------------------------------------------
//  Storage.go contains the WorldTree, which stores the
//  entire world, and can save/load worlds.
//
//  ID System:
//  Each block in the map has a 5 digit ID, in which the
//  first 3 digits specify the Block ID (see blocks.go)
//  and the last 2 specify the orientation of the block.
//
//  This allows for simple serialization of the world
//	in and out of a file on the disk.
//  --------------------------------------------------

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

//  --------------------------------------------------
//  Node Creation
//  --------------------------------------------------

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

//  --------------------------------------------------
//  Node Retrieval
//  --------------------------------------------------

func (tree *WorldTree) GetWorldBlock(x, y int) *child.ChildCopy {
	return tree.blockNodes[x][y].worldBlock
}

func (tree *WorldTree) GetBackBlock(x, y int) *child.ChildCopy {
	return tree.blockNodes[x][y].backBlock
}

func (tree *WorldTree) GetNatureBlock(x, y int) *child.ChildCopy {
	return tree.blockNodes[x][y].natureBlock
}

func (tree *WorldTree) GetWorldBlockName(x, y int) string {
	return GetNameFromID(tree.blockNodes[x][y].worldBlock.ID[:3])
}

func (tree *WorldTree) GetBackBlockName(x, y int) string {
	return GetNameFromID(tree.blockNodes[x][y].backBlock.ID[:3])
}

func (tree *WorldTree) GetNatureBlockName(x, y int) string {
	return GetNameFromID(tree.blockNodes[x][y].natureBlock.ID[:3])
}

func (tree *WorldTree) GetWorldBlockID(x, y int) string {
	return tree.blockNodes[x][y].worldBlock.ID
}

func (tree *WorldTree) GetBackBlockID(x, y int) string {
	return tree.blockNodes[x][y].worldBlock.ID
}

func (tree *WorldTree) GetNatureBlockID(x, y int) string {
	return tree.blockNodes[x][y].worldBlock.ID
}

func (tree *WorldTree) GetBlockOrientation(x, y int) string {
	return GetOrientationFromID(tree.blockNodes[x][y].worldBlock.ID[3:])
}

func (tree *WorldTree) GetDarkness(x, y int) float32 {
	return tree.blockNodes[x][y].worldBlock.Darkness
}

//  --------------------------------------------------
//  Node Modification
//  --------------------------------------------------

// Updates node materials, for when orientations change
func (tree *WorldTree) UpdateNodeMaterials(x, y int) {
	tree.blockNodes[x][y].worldBlock.Material = GetBlock(tree.GetWorldBlockName(x, y)).GetMaterial(tree.GetBlockOrientation(x, y))
	tree.blockNodes[x][y].natureBlock.Material = GetBlock(tree.GetNatureBlockName(x, y)).GetMaterial(tree.GetBlockOrientation(x, y))
	tree.blockNodes[x][y].backBlock.Material = GetBlock(tree.GetBackBlockName(x, y)).GetMaterial(tree.GetBlockOrientation(x, y))
}

func (tree *WorldTree) SetBlockOrientation(x, y int, orient string) {
	tree.blockNodes[x][y].worldBlock.ID = tree.blockNodes[x][y].worldBlock.ID[:3] + OrientationsMap[orient]
}

func (tree *WorldTree) SetBackBlockOrientation(x, y int, orient string) {
	tree.blockNodes[x][y].backBlock.ID = tree.blockNodes[x][y].backBlock.ID[:3] + OrientationsMap[orient]
}

func (tree *WorldTree) SetDarkness(x, y int, darkness float32) {
	tree.blockNodes[x][y].worldBlock.Darkness = darkness
	tree.blockNodes[x][y].backBlock.Darkness = darkness
	tree.blockNodes[x][y].natureBlock.Darkness = darkness
}
