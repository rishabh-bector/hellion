package main

import (
	"rapidengine/child"
	"rapidengine/geometry"
	"rapidengine/material"
)

type WorldBlock struct {
	ID          int
	Orientation string
	Darkness    float32
}

func loadWorldChildren() {
	WorldChild = Engine.NewChild2D()
	WorldChild.AttachShader(Engine.ShaderControl.GetShader("colorLighting"))
	WorldChild.AttachPrimitive(geometry.NewRectangle(BlockSize, BlockSize, &Config))
	WorldChild.AttachTextureCoordsPrimitive()
	WorldChild.EnableCopying()
	WorldChild.AttachCollider(0, 0, BlockSize, BlockSize)

	NoCollisionChild = Engine.NewChild2D()
	NoCollisionChild.AttachShader(Engine.ShaderControl.GetShader("colorLighting"))
	NoCollisionChild.AttachPrimitive(geometry.NewRectangle(BlockSize, BlockSize, &Config))
	NoCollisionChild.AttachTextureCoordsPrimitive()
	NoCollisionChild.EnableCopying()

	NatureChild = Engine.NewChild2D()
	NatureChild.AttachShader(Engine.ShaderControl.GetShader("colorLighting"))
	NatureChild.AttachPrimitive(geometry.NewRectangle(BlockSize, BlockSize, &Config))
	NatureChild.AttachTextureCoordsPrimitive()
	NatureChild.EnableCopying()

	CloudChild = Engine.NewChild2D()
	CloudChild.AttachShader(Engine.ShaderControl.GetShader("colorLighting"))
	CloudChild.AttachPrimitive(geometry.NewRectangle(300, 145, &Config))
	CloudChild.AttachTextureCoordsPrimitive()
	CloudChild.EnableCopying()
	CloudChild.SetSpecificRenderDistance(float32(ScreenWidth/2) + 300)
	Engine.TextureControl.NewTexture("./assets/cloud1.png", "cloud1")
	cloudMaterial = material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	cloudMaterial.BecomeTexture(Engine.TextureControl.GetTexture("cloud1"))
	CloudChild.AttachMaterial(&cloudMaterial)
}

func createWorldBlock(x, y int, name string) {
	WorldMap.AddWorldBlock(x, y, &child.ChildCopy{
		X:        float32(x * BlockSize),
		Y:        float32(y * BlockSize),
		Material: GetBlock(name).GetMaterial("NN"),
		Darkness: 1,
		ID:       GetIDFromName(name) + "00",
	})
}

func createBackBlock(x, y int, name string) {
	WorldMap.AddBackBlock(x, y, &child.ChildCopy{
		X:        float32(x * BlockSize),
		Y:        float32(y * BlockSize),
		Material: GetBlock(name).GetMaterial("NN"),
		Darkness: 1,
		ID:       GetIDFromName(name) + "00",
	})
}

func createNatureBlock(x, y int, name string) {
	WorldMap.AddNatureBlock(x, y, &child.ChildCopy{
		X:        float32(x * BlockSize),
		Y:        float32(y * BlockSize),
		Material: GetBlock(name).GetMaterial("NN"),
		Darkness: 1,
		ID:       GetIDFromName(name) + "00",
	})
}

/*func createSingleCopy(x, y int) {

	// Back Blocks
	if isBackBlock(NameList[WorldMap[x][y].ID]) {
		if WorldMap[x][y].Orientation == "E" || WorldMap[x][y].Orientation == "NN" {
			NoCollisionCopies[x][y] = child.ChildCopy{
				X:        float32(x * BlockSize),
				Y:        float32(y * BlockSize),
				Material: GetBlockIndex(WorldMap[x][y].ID).GetMaterial(),
				Darkness: WorldMap[x][y].Darkness,
				ID:       1,
			}
		} else {
			NoCollisionCopies[x][y] = child.ChildCopy{
				X:        float32(x * BlockSize),
				Y:        float32(y * BlockSize),
				Material: GetBlockIndex(WorldMap[x][y].ID).GetOrientMaterial(WorldMap[x][y].Orientation),
				Darkness: WorldMap[x][y].Darkness,
				ID:       1,
			}
		}
		return
	}

	// Nature Blocks
	if blockType(NameList[WorldMap[x][y].ID]) == "nature" {
		if WorldMap[x][y].Orientation == "E" || WorldMap[x][y].Orientation == "NN" {
			NatureCopies[x][y] = child.ChildCopy{
				X:        float32(x * BlockSize),
				Y:        float32(y*BlockSize - 10),
				Material: GetBlockIndex(WorldMap[x][y].ID).GetMaterial(),
				Darkness: WorldMap[x][y].Darkness,
				ID:       1,
			}
		} else {
			NatureCopies[x][y] = child.ChildCopy{
				X:        float32(x * BlockSize),
				Y:        float32(y*BlockSize - 10),
				Material: GetBlockIndex(WorldMap[x][y].ID).GetOrientMaterial(WorldMap[x][y].Orientation),
				Darkness: WorldMap[x][y].Darkness,
				ID:       1,
			}
		}
		return
	}

	// Normal blocks
	if WorldMap[x][y].ID != NameMap["sky"] {

		// No orientation
		if WorldMap[x][y].Orientation == "E" || WorldMap[x][y].Orientation == "NN" {
			WorldCopies[x][y] = child.ChildCopy{
				X:        float32(x * BlockSize),
				Y:        float32(y * BlockSize),
				Material: GetBlockIndex(WorldMap[x][y].ID).GetMaterial(),
				Darkness: WorldMap[x][y].Darkness,
				ID:       1,
			}
		} else {
			// Oriented block
			WorldCopies[x][y] = child.ChildCopy{
				X:        float32(x * BlockSize),
				Y:        float32(y * BlockSize),
				Material: GetBlockIndex(WorldMap[x][y].ID).GetOrientMaterial(WorldMap[x][y].Orientation),
				Darkness: WorldMap[x][y].Darkness,
				ID:       1,
			}
		}
	}
}*/

/*func createAllExtraBackdirt() {
	for x := 2; x < WorldWidth-2; x++ {
		for y := 2; y < WorldHeight-2; y++ {
			createSingleExtraBackdirt(x, y)
		}
	}
}

func createSingleExtraBackdirt(x, y int) {
	orient := WorldMap.GetBlockOrientation(x, y)
	if orient != "E" && orient != "NN" && WorldMap.GetBlockID(x, y) != "00000" {
		if WorldMap.GetBlockID(x+1, y) == "00000" ||
			WorldMap.GetBlockID(x-1, y) == "00000" ||
			WorldMap.GetBlockID(x, y+1) == "00000" ||
			WorldMap.GetBlockID(x, y-1) == "00000" {
			if y <= HeightMap[x] {
				updateBackBlock(x, y, "backdirt")
			}
		} else {
			NoCollisionCopies[x][y] = child.ChildCopy{
				X:        float32(x * BlockSize),
				Y:        float32(y * BlockSize),
				Material: GetBlockName("backdirt").GetMaterial(),
				Darkness: WorldMap[x][y].Darkness,
				ID:       2,
			}
		}
	}
}*/

/*func orientBlocks(name string, topBlock bool) {
	block := NameMap[name]
	for x := 1; x < WorldWidth-1; x++ {
		for y := 1; y < WorldHeight-1; y++ {
			orientSingleBlock(name, block, topBlock, x, y)
		}
	}
}

func orientSingleBlock(name string, block int, topBlock bool, x, y int) {
	if WorldMap[x][y].ID == block {
		above := false
		under := false
		left := false
		right := false
		if WorldMap[x-1][y].ID == NameMap["sky"] || (isBackBlock(NameList[WorldMap[x-1][y].ID]) && !isBackBlock(name)) {
			left = true
		}
		if WorldMap[x+1][y].ID == NameMap["sky"] || (isBackBlock(NameList[WorldMap[x+1][y].ID]) && !isBackBlock(name)) {
			right = true
		}
		if WorldMap[x][y-1].ID == NameMap["sky"] || (isBackBlock(NameList[WorldMap[x][y-1].ID]) && !isBackBlock(name)) {
			under = true
		}
		if WorldMap[x][y+1].ID == NameMap["sky"] || (isBackBlock(NameList[WorldMap[x][y+1].ID]) && !isBackBlock(name)) {
			above = true
		}
		WorldMap[x][y].Orientation[4:] = OrientationsMap[getOrientationLetter(left, right, under, above, topBlock)]
	}
}

func getSingleBlockOrientation(name string, block int, topBlock bool, x, y int) string {
	above := false
	under := false
	left := false
	right := false
	if WorldMap[x-1][y].ID == NameMap["sky"] || (isBackBlock(NameList[WorldMap[x-1][y].ID]) && !isBackBlock(name)) {
		left = true
	}
	if WorldMap[x+1][y].ID == NameMap["sky"] || (isBackBlock(NameList[WorldMap[x+1][y].ID]) && !isBackBlock(name)) {
		right = true
	}
	if WorldMap[x][y-1].ID == NameMap["sky"] || (isBackBlock(NameList[WorldMap[x][y-1].ID]) && !isBackBlock(name)) {
		under = true
	}
	if WorldMap[x][y+1].ID == NameMap["sky"] || (isBackBlock(NameList[WorldMap[x][y+1].ID]) && !isBackBlock(name)) {
		above = true
	}
	return getOrientationLetter(left, right, under, above, topBlock)

}

func getOrientationLetter(left, right, under, above, topBlock bool) string {
	if left && right && under && above {
		return "AA"
	}
	if left && right && !under && !above {
		return "AN"
	}
	if !left && !right && under && above {
		return "NA"
	}
	if left && !right && under && above {
		return "LA"
	}
	if !left && right && under && above {
		return "RA"
	}
	if left && right && !under && above {
		return "AT"
	}
	if left && right && under && !above {
		return "AB"
	}
	if left && !right && !under && !above {
		return "LN"
	}
	if !left && right && !under && !above {
		return "RN"
	}
	if !left && !right && !under && above && topBlock {
		return "NT"
	}
	if !left && !right && under && !above {
		return "NB"
	}
	if !left && right && under && !above {
		return "RB"
	}
	if left && !right && under && !above {
		return "LB"
	}
	if !left && !right && !under && !above {
		return "NN"
	}
	if !left && right && !under && above {
		return "RT"
	}
	if left && !right && !under && above {
		return "LT"
	}
	return "NN"
}
*/
