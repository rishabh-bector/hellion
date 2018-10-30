package main

import (
	"math/rand"
	"rapidengine/child"
	"rapidengine/geometry"
	"rapidengine/material"
	"time"

	perlin "github.com/aquilax/go-perlin"
)

//  --------------------------------------------------
//  World Generation Parameters
//  --------------------------------------------------

const WorldHeight = 3000 //4000
const WorldWidth = 2000
const BlockSize = 32
const Flatness = 0.1

const GrassMinimum = 700

const CaveNoiseScalar = 30
const CaveNoiseThreshold = 0.75

const StoneNoiseScalar = 30.0
const StoneTop = 600.0
const StoneTopDeviation = 5
const StoneStartingFrequency = -0.3

//  --------------------------------------------------
//  --------------------------------------------------
//  --------------------------------------------------

var WorldChild child.Child2D
var WorldCopies [WorldWidth][WorldHeight]child.ChildCopy

var NoCollisionChild child.Child2D
var NoCollisionCopies [WorldWidth][WorldHeight]child.ChildCopy

var NatureChild child.Child2D
var NatureCopies [WorldWidth][WorldHeight]child.ChildCopy

var CloudChild child.Child2D
var cloudMaterial material.Material

var p *perlin.Perlin
var WorldMap [WorldWidth + 1][WorldHeight + 1]WorldBlock
var HeightMap [WorldWidth]int

var transparentBlocks = []string{"backdirt"} //"topGrass1", "topGrass2", "topGrass3", "treeRightRoot", "treeLeftRoot", "treeTrunk", "treeBottomRoot", "treeBranchR1", "treeBranchL1", "flower1", "flower2", "flower3", "pebble"}
var natureBlocks = []string{"leaves", "treeRightRoot", "treeLeftRoot", "treeTrunk", "treeBottomRoot", "treeBranchR1", "treeBranchL1", "topGrass1", "topGrass2", "topGrass3", "flower1", "flower2", "flower3", "pebble"}

type WorldBlock struct {
	ID          int
	Orientation string
	Darkness    float32
}

func NewBlock(id string) WorldBlock {
	return WorldBlock{ID: NameMap[id], Orientation: "E", Darkness: 0}
}

func NewOrientBlock(id, orientation string) WorldBlock {
	return WorldBlock{ID: NameMap[id], Orientation: orientation, Darkness: 0}
}

func generateWorld() {
	WorldChild = engine.NewChild2D()
	WorldChild.AttachShader(engine.ShaderControl.GetShader("colorLighting"))
	WorldChild.AttachPrimitive(geometry.NewRectangle(BlockSize, BlockSize, &config))
	WorldChild.AttachTextureCoordsPrimitive()
	WorldChild.EnableCopying()
	WorldChild.AttachCollider(0, 0, BlockSize, BlockSize)

	NoCollisionChild = engine.NewChild2D()
	NoCollisionChild.AttachShader(engine.ShaderControl.GetShader("colorLighting"))
	NoCollisionChild.AttachPrimitive(geometry.NewRectangle(BlockSize, BlockSize, &config))
	NoCollisionChild.AttachTextureCoordsPrimitive()
	NoCollisionChild.EnableCopying()

	NatureChild = engine.NewChild2D()
	NatureChild.AttachShader(engine.ShaderControl.GetShader("colorLighting"))
	NatureChild.AttachPrimitive(geometry.NewRectangle(BlockSize, BlockSize, &config))
	NatureChild.AttachTextureCoordsPrimitive()
	NatureChild.EnableCopying()

	CloudChild = engine.NewChild2D()
	CloudChild.AttachShader(engine.ShaderControl.GetShader("colorLighting"))
	CloudChild.AttachPrimitive(geometry.NewRectangle(300, 145, &config))
	CloudChild.AttachTextureCoordsPrimitive()
	CloudChild.EnableCopying()
	CloudChild.SetSpecificRenderDistance(float32(ScreenWidth/2) + 300)
	engine.TextureControl.NewTexture("./assets/cloud1.png", "cloud1")
	cloudMaterial = material.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	cloudMaterial.BecomeTexture(engine.TextureControl.GetTexture("cloud1"))
	CloudChild.AttachMaterial(&cloudMaterial)

	LoadBlocks()

	rand.Seed(time.Now().UTC().UnixNano())
	p = perlin.NewPerlin(2, 2, 10, int64(rand.Int()))

	// Fill world with 0s
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			b := NewBlock("sky")
			b.Darkness = 0
			WorldMap[x][y] = b
		}
	}

	// Generate heightmap and place grass
	generateHeights()
	for x := 0; x < WorldWidth; x++ {
		WorldMap[x][HeightMap[x]] = NewBlock("grass")
	}

	// Fill everything underneath grass with dirt
	fillHeights()

	// Generate stone based on height
	fillStone()

	// Clean up stone above ground
	cleanStone()

	// Generate caves
	generateCaves()

	// Clean back dirt
	cleanBackDirt()

	// Put grass and some top grass on dirt with air above it
	growGrass()

	// Create clouds
	generateClouds()

	// Place flowers and pebbles above grass
	generateNature()

	// Fix the orientation of blocks in the world
	orientBlocks("dirt", true)
	orientBlocks("grass", true)
	orientBlocks("stone", true)
	orientBlocks("leaves", true)
	orientBlocks("backdirt", true)

	CreateLighting(WorldWidth/2, HeightMap[WorldWidth/2]+5, 0.9)

	// Fix backdirt
	createAllExtraBackdirt()

	Player.SetPosition(float32(WorldWidth*BlockSize/2), float32((HeightMap[WorldWidth/2]+50)*BlockSize))
}

func createCopies() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			createSingleCopy(x, y)
		}
	}
}

func createSingleCopy(x, y int) {

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
}

func createAllExtraBackdirt() {
	for x := 2; x < WorldWidth-2; x++ {
		for y := 2; y < WorldHeight-2; y++ {
			createSingleExtraBackdirt(x, y)
		}
	}
}

func createSingleExtraBackdirt(x, y int) {
	if WorldMap[x][y].Orientation != "E" && WorldMap[x][y].Orientation != "NN" && WorldMap[x][y].ID != NameMap["sky"] {
		if WorldMap[x+1][y].ID == NameMap["sky"] || WorldMap[x-1][y].ID == NameMap["sky"] || WorldMap[x][y+1].ID == NameMap["sky"] || WorldMap[x][y-1].ID == NameMap["sky"] {
			if y <= HeightMap[x] {
				NoCollisionCopies[x][y] = child.ChildCopy{
					X:        float32(x * BlockSize),
					Y:        float32(y * BlockSize),
					Material: GetBlockName("backdirt").GetOrientMaterial(GetSingleBlockOrientation("backdirt", NameMap["backdirt"], true, x, y)),
					Darkness: WorldMap[x][y].Darkness,
					ID:       2,
				}
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
}

func orientBlocks(name string, topBlock bool) {
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
		WorldMap[x][y].Orientation = GetOrientationLetter(left, right, under, above, topBlock)
	}
}

func GetSingleBlockOrientation(name string, block int, topBlock bool, x, y int) string {
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
	return GetOrientationLetter(left, right, under, above, topBlock)

}

func GetOrientationLetter(left, right, under, above, topBlock bool) string {
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
	return "E"
}

//   --------------------------------------------------
//   Lighting
//   --------------------------------------------------

func CreateLighting(x, y int, light float32) {
	if !IsValidPosition(x, y) {
		return
	}
	newLight := light - GetLightBlockAmount(x, y)
	if newLight <= GetLightAt(x, y) {
		return
	}
	WorldMap[x][y].Darkness = newLight
	CreateLighting(x+1, y, newLight)
	CreateLighting(x, y+1, newLight)
	CreateLighting(x-1, y, newLight)
	CreateLighting(x, y-1, newLight)
}

func FixLightingAt(x, y int) {
	maxLight := float32(0)
	if l := GetLightAt(x+1, y); l > maxLight {
		maxLight = l
	}
	if l := GetLightAt(x, y+1); l > maxLight {
		maxLight = l
	}
	if l := GetLightAt(x-1, y); l > maxLight {
		maxLight = l
	}
	if l := GetLightAt(x, y-1); l > maxLight {
		maxLight = l
	}
	WorldMap[x][y].Darkness = maxLight - GetLightBlockAmount(x, y)
}

func GetLightAt(x, y int) float32 {
	return WorldMap[x][y].Darkness
}

func GetLightBlockAmount(x, y int) float32 {
	return BlockMap[NameList[WorldMap[x][y].ID]].LightBlock
}

func IsValidPosition(x, y int) bool {
	if x > 0 && x < WorldWidth {
		if y > 0 && y < WorldHeight {
			return true
		}
	}
	return false
}
