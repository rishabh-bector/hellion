package main

import (
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
const BlockSize = 25
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

var p *perlin.Perlin
var WorldMap [WorldWidth][WorldHeight]WorldBlock
var HeightMap [WorldWidth]int

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

	// Fix the orientation of Dirt in the world
	orientBlock("dirt")
}

func createCopies() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			if WorldMap[x][y].ID != NameMap["sky"] {
				if WorldMap[x][y].Orientation == "E" {
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
			if n > CaveNoiseThreshold {
				WorldMap[x][y] = NewBlock("sky")
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
			if WorldMap[x][y].ID == NameMap["dirt"] && WorldMap[x][y+1].ID == NameMap["sky"] {
				WorldMap[x][y] = NewBlock("grass")
			}
		}
	}
}

func orientBlock(name string) {
	block := NameMap[name]
	for x := 1; x < WorldWidth-1; x++ {
		for y := 1; y < WorldHeight-1; y++ {
			if WorldMap[x][y].ID == block {
				above := false
				under := false
				left := false
				right := false
				if WorldMap[x-1][y].ID == NameMap["sky"] {
					left = true
				}
				if WorldMap[x+1][y].ID == NameMap["sky"] {
					right = true
				}
				if WorldMap[x][y-1].ID == NameMap["sky"] {
					under = true
				}
				if WorldMap[x][y+1].ID == NameMap["sky"] {
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
				if !left && !right && !under && above {
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
				if !left && !right && !under && above {
					WorldMap[x][y].Orientation = "LT"
				}
			}
		}
	}
}

func noise2D(x, y float64) float64 {
	return (p.Noise2D(x, y) + 0.4) / 0.8
}

func noise1D(x float64) float64 {
	return (p.Noise1D(x) + 0.4) / 0.8
}
