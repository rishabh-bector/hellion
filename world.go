package main

import (
	"fmt"
	"math/rand"
	"rapidengine"
	"time"

	perlin "github.com/aquilax/go-perlin"
)

//  --------------------------------------------------
//  World Generation Parameters
//  --------------------------------------------------

const WorldHeight = 4000 //4000
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

var WorldChild rapidengine.Child2D
var NoCollisionChild rapidengine.Child2D
var NatureChild rapidengine.Child2D

var CloudChild rapidengine.Child2D
var cloudMaterial rapidengine.Material

var p *perlin.Perlin
var WorldMap [WorldWidth + 1][WorldHeight + 1]WorldBlock
var HeightMap [WorldWidth]int

var transparentBlocks = []string{"backdirt", "topGrass1", "topGrass2", "topGrass3", "treeRightRoot", "treeLeftRoot", "treeTrunk", "treeBottomRoot", "treeBranchR1", "treeBranchL1"}
var natureBlocks = []string{"leaves", "treeRightRoot", "treeLeftRoot", "treeTrunk", "treeBottomRoot", "treeBranchR1", "treeBranchL1", "topGrass1", "topGrass2", "topGrass3"}

type WorldBlock struct {
	ID          int
	Orientation string
}

func NewBlock(id string) WorldBlock {
	return WorldBlock{ID: NameMap[id], Orientation: "E"}
}

func NewOrientBlock(id, orientation string) WorldBlock {
	return WorldBlock{ID: NameMap[id], Orientation: orientation}
}

func generateWorld() {
	WorldChild = engine.NewChild2D()
	WorldChild.AttachShader(engine.ShaderControl.GetShader("colorLighting"))
	WorldChild.AttachPrimitive(rapidengine.NewRectangle(BlockSize, BlockSize, &config))
	WorldChild.AttachTextureCoordsPrimitive()
	WorldChild.EnableCopying()
	WorldChild.AttachCollider(0, 0, BlockSize, BlockSize)

	NoCollisionChild = engine.NewChild2D()
	NoCollisionChild.AttachShader(engine.ShaderControl.GetShader("colorLighting"))
	NoCollisionChild.AttachPrimitive(rapidengine.NewRectangle(BlockSize, BlockSize, &config))
	NoCollisionChild.AttachTextureCoordsPrimitive()
	NoCollisionChild.EnableCopying()

	NatureChild = engine.NewChild2D()
	NatureChild.AttachShader(engine.ShaderControl.GetShader("colorLighting"))
	NatureChild.AttachPrimitive(rapidengine.NewRectangle(BlockSize, BlockSize, &config))
	NatureChild.AttachTextureCoordsPrimitive()
	NatureChild.EnableCopying()

	CloudChild = engine.NewChild2D()
	CloudChild.AttachShader(engine.ShaderControl.GetShader("colorLighting"))
	CloudChild.AttachPrimitive(rapidengine.NewRectangle(300, 145, &config))
	CloudChild.AttachTextureCoordsPrimitive()
	CloudChild.EnableCopying()
	CloudChild.SetSpecificRenderDistance(float32(ScreenWidth/2) + 300)
	engine.TextureControl.NewTexture("./assets/cloud1.png", "cloud1")
	cloudMaterial = rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	cloudMaterial.BecomeTexture(engine.TextureControl.GetTexture("cloud1"))
	CloudChild.AttachMaterial(&cloudMaterial)

	LoadBlocks()

	rand.Seed(time.Now().UTC().UnixNano())
	p = perlin.NewPerlin(2, 2, 10, int64(rand.Int()))

	// Fill world with 0s
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			WorldMap[x][y] = NewBlock("sky")
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

	// Put grass on dirt with air above it
	growGrass()

	// Place Trees
	generateTrees()

	// Create clouds
	generateClouds()

	// Fix the orientation of blocks in the world
	orientBlock("dirt", true)
	orientBlock("grass", true)
	orientBlock("stone", true)
	orientBlock("leaves", true)
	orientBlock("backdirt", true)

	Player.SetPosition(float32(WorldWidth*BlockSize/2), float32((HeightMap[WorldWidth/2]+50)*BlockSize))
}

func createCopies() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {

			if isBackBlock(NameList[WorldMap[x][y].ID]) {
				if WorldMap[x][y].Orientation == "E" || WorldMap[x][y].Orientation == "NN" {
					NoCollisionChild.AddCopy(rapidengine.ChildCopy{
						X:        float32(x * BlockSize),
						Y:        float32(y * BlockSize),
						Material: GetBlockIndex(WorldMap[x][y].ID).GetMaterial(),
					})
				} else {
					NoCollisionChild.AddCopy(rapidengine.ChildCopy{
						X:        float32(x * BlockSize),
						Y:        float32(y * BlockSize),
						Material: GetBlockIndex(WorldMap[x][y].ID).GetOrientMaterial(WorldMap[x][y].Orientation),
					})
				}
				continue
			}

			// Nature Blocks
			if blockType(NameList[WorldMap[x][y].ID]) == "nature" {
				if WorldMap[x][y].Orientation == "E" || WorldMap[x][y].Orientation == "NN" {
					NatureChild.AddCopy(rapidengine.ChildCopy{
						X:        float32(x * BlockSize),
						Y:        float32(y * BlockSize),
						Material: GetBlockIndex(WorldMap[x][y].ID).GetMaterial(),
					})
				} else {
					NatureChild.AddCopy(rapidengine.ChildCopy{
						X:        float32(x * BlockSize),
						Y:        float32(y * BlockSize),
						Material: GetBlockIndex(WorldMap[x][y].ID).GetOrientMaterial(WorldMap[x][y].Orientation),
					})
				}
				continue
			}

			// Normal blocks
			if WorldMap[x][y].ID != NameMap["sky"] {
				if WorldMap[x][y].Orientation == "E" || WorldMap[x][y].Orientation == "NN" {
					WorldChild.AddCopy(rapidengine.ChildCopy{
						X:        float32(x * BlockSize),
						Y:        float32(y * BlockSize),
						Material: GetBlockIndex(WorldMap[x][y].ID).GetMaterial(),
					})
				} else {
					WorldChild.AddCopy(rapidengine.ChildCopy{
						X:        float32(x * BlockSize),
						Y:        float32(y * BlockSize),
						Material: GetBlockIndex(WorldMap[x][y].ID).GetOrientMaterial(WorldMap[x][y].Orientation),
					})

					if y < HeightMap[x]-1 {
						NoCollisionChild.AddCopy(rapidengine.ChildCopy{
							X:        float32(x * BlockSize),
							Y:        float32(y * BlockSize),
							Material: GetBlockName("backdirt").GetMaterial(),
						})
					}
				}
			}
		}
	}
}

func generateCaves() {
	p = perlin.NewPerlin(1.5, 2, 3, int64(rand.Int()))
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			n := noise2D(CaveNoiseScalar*float64(x)/WorldWidth*2, CaveNoiseScalar*float64(y)/WorldHeight*4)
			if n > CaveNoiseThreshold && y <= HeightMap[x] && WorldMap[x][y].Orientation == "E" {
				WorldMap[x][y] = NewBlock("backdirt")
			}
		}
	}
}

func generateHeights() {
	for x := 0; x < WorldWidth; x++ {
		HeightMap[x] = GrassMinimum + int(Flatness*noise1D(float64(x)/(WorldWidth/2))*WorldHeight)
	}
}

func fillHeights() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight-1; y++ {
			WorldMap[x][y] = NewOrientBlock("dirt", "E")
			if WorldMap[x][y+1].ID == NameMap["grass"] {
				break
			}
		}
	}
}

func fillStone() {
	p = perlin.NewPerlin(1.2, 2, 2, int64(rand.Int()))
	stoneFrequency := StoneStartingFrequency
	for y := 0; y < WorldHeight; y++ {
		for x := 0; x < WorldWidth; x++ {
			n := noise2D(StoneNoiseScalar*float64(x)/WorldWidth*2, StoneNoiseScalar*float64(y)/WorldHeight*4)
			if n > stoneFrequency {
				WorldMap[x][y] = NewBlock("stone")
			}
		}
		stoneFrequency += (1 / StoneTop)
	}
}

func cleanStone() {
	for x := 0; x < WorldWidth; x++ {
		grassHeight := HeightMap[x]
		if WorldMap[x][grassHeight].ID == NameMap["stone"] {
			for y := grassHeight + StoneTopDeviation; y < WorldHeight; y++ {
				WorldMap[x][y] = NewBlock("sky")
			}
		} else {
			for y := grassHeight + 1; y < WorldHeight; y++ {
				WorldMap[x][y] = NewBlock("sky")
			}
		}
	}
}

func growGrass() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			if WorldMap[x][y].ID == NameMap["dirt"] && (WorldMap[x][y+1].ID == NameMap["sky"] || WorldMap[x][y+1].ID == NameMap["backdirt"]) {
				WorldMap[x][y] = NewBlock("grass")
			}
			if WorldMap[x][y].ID == NameMap["grass"] && (WorldMap[x][y+1].ID == NameMap["sky"] || WorldMap[x][y+1].ID == NameMap["backdirt"]) {
				if rand.Intn(4) == 0 {
					grassRand := rand.Intn(3) + 1
					grassType := fmt.Sprintf("topGrass%d", grassRand)
					NatureChild.AddCopy(rapidengine.ChildCopy{
						X:        float32(x * BlockSize),
						Y:        float32((y + 1) * BlockSize),
						Material: GetBlockIndex(NameMap[grassType]).GetMaterial(),
					})
				}
			}
		}
	}
}

func orientBlock(name string, topBlock bool) {
	block := NameMap[name]
	for x := 1; x < WorldWidth-1; x++ {
		for y := 1; y < WorldHeight-1; y++ {
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
				if left && right && under && above {
					WorldMap[x][y].Orientation = "AA"
				}
				if left && right && !under && !above {
					WorldMap[x][y].Orientation = "AN"
				}
				if !left && !right && under && above {
					WorldMap[x][y].Orientation = "NA"
				}
				if left && !right && under && above {
					WorldMap[x][y].Orientation = "LA"
				}
				if !left && right && under && above {
					WorldMap[x][y].Orientation = "RA"
				}
				if left && right && !under && above {
					WorldMap[x][y].Orientation = "AT"
				}
				if left && right && under && !above {
					WorldMap[x][y].Orientation = "AB"
				}
				if left && !right && !under && !above {
					WorldMap[x][y].Orientation = "LN"
				}
				if !left && right && !under && !above {
					WorldMap[x][y].Orientation = "RN"
				}
				if !left && !right && !under && above && topBlock {
					WorldMap[x][y].Orientation = "NT"
				}
				if !left && !right && under && !above {
					WorldMap[x][y].Orientation = "NB"
				}
				if !left && right && under && !above {
					WorldMap[x][y].Orientation = "RB"
				}
				if left && !right && under && !above {
					WorldMap[x][y].Orientation = "LB"
				}
				if !left && !right && !under && !above {
					WorldMap[x][y].Orientation = "NN"
				}
				if !left && right && !under && above {
					WorldMap[x][y].Orientation = "RT"
				}
				if left && !right && !under && above {
					WorldMap[x][y].Orientation = "LT"
				}
			}
		}
	}
}

func generateTrees() {
	for x := 1; x < WorldWidth-1; x++ {
		if rand.Intn(16) == 4 {
			if WorldMap[x][HeightMap[x]].ID == NameMap["grass"] && WorldMap[x-1][HeightMap[x]].ID != NameMap["treeTrunk"] {
				WorldMap[x][(HeightMap[x] + 1)] = NewBlock("treeTrunk")
				height := 4 + rand.Intn(8)
				for i := 0; i < height; i++ {
					if rand.Intn(4) == 0 && i < height-2 && i > 0 {
						WorldMap[x-1][(HeightMap[x] + i + 2)] = NewBlock("treeBranchL1")
					}
					if rand.Intn(4) == 0 && i < height-2 && i > 0 {
						WorldMap[x+1][(HeightMap[x] + i + 2)] = NewBlock("treeBranchR1")
					}
					WorldMap[x][(HeightMap[x] + 2 + i)] = NewBlock("treeTrunk")
				}
				WorldMap[x-1][(HeightMap[x] + height + 1)] = NewBlock("leaves") // TL
				WorldMap[x][(HeightMap[x] + height + 1)] = NewBlock("leaves")   // TM
				WorldMap[x+1][(HeightMap[x] + height + 1)] = NewBlock("leaves") // TR
				WorldMap[x-1][(HeightMap[x] + height)] = NewBlock("leaves")     // ML
				WorldMap[x][(HeightMap[x] + height)] = NewBlock("leaves")       // MM
				WorldMap[x+1][(HeightMap[x] + height)] = NewBlock("leaves")     // MR
				WorldMap[x-1][(HeightMap[x] + height - 1)] = NewBlock("leaves") //BL
				WorldMap[x][(HeightMap[x] + height - 1)] = NewBlock("leaves")   // BM
				WorldMap[x+1][(HeightMap[x] + height - 1)] = NewBlock("leaves") //BL
			}
		}
	}
}

func generateClouds() {
	for x := 0; x < WorldWidth; x++ {
		if rand.Float32() < 0.4 {
			CloudChild.AddCopy(
				rapidengine.ChildCopy{
					X:        float32(x * BlockSize),
					Y:        float32((rand.Intn(20) + HeightMap[x] + 15) * BlockSize),
					Material: &cloudMaterial,
				},
			)
			x += 400 / BlockSize
		}
	}
}

func isBackBlock(name string) bool {
	for _, transparent := range transparentBlocks {
		if NameMap[name] == NameMap[transparent] {
			return true
		}
	}
	return false
}
func blockType(name string) string {
	for _, green := range natureBlocks {
		if NameMap[name] == NameMap[green] {
			return "nature"
		}
	}
	return "idk"
}

func noise2D(x, y float64) float64 {
	return (p.Noise2D(x, y) + 0.4) / 0.8
}

func noise1D(x float64) float64 {
	return (p.Noise1D(x) + 0.4) / 0.8
}
