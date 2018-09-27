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

const WorldHeight = 4000
const WorldWidth = 2000
const BlockSize = 25
const Flatness = 0.1

const GrassMinimum = 500

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
var WorldMap [WorldWidth][WorldHeight]int
var HeightMap [WorldWidth]int

var blocksMap = map[string]int{
	"sky":             0,
	"dirt":            1,
	"grass":           2,
	"stone":           3,
	"topDirt":         4,
	"topLeftDirt":     5,
	"topRightDirt":    6,
	"bottomRightDirt": 7,
	"bottomLeftDIrt":  8,
	"rightDirt":       9,
	"leftDirt":        10,
}

func generateWorld() {
	rand.Seed(time.Now().UTC().UnixNano())
	p = perlin.NewPerlin(2, 2, 10, int64(rand.Int()))

	// Fill world with 0s
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			WorldMap[x][y] = blocksMap["sky"]
		}
	}

	// Generate heightmap and place grass
	generateHeights()
	for x := 0; x < WorldWidth; x++ {
		WorldMap[x][HeightMap[x]] = blocksMap["grass"]
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
	orientDirt()
}

func createCopies() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			if WorldMap[x][y] != blocksMap["sky"] {
				WorldChild.AddCopy(rapidengine.ChildCopy{
					X:        float32(x * BlockSize),
					Y:        float32(y * BlockSize),
					Material: &blocks[WorldMap[x][y]],
				})
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
				WorldMap[x][y] = blocksMap["sky"]
			}
		}
	}
}

func generateHeights() {
	for x := 0; x < WorldWidth; x++ {
		HeightMap[x] = GrassMinimum + int(Flatness*noise1D(float64(x)/WorldWidth)*WorldHeight)
	}
}

func fillHeights() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight-1; y++ {
			WorldMap[x][y] = blocksMap["dirt"]
			if WorldMap[x][y+1] == blocksMap["grass"] {
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
				WorldMap[x][y] = blocksMap["stone"]
			}
		}
		stoneFrequency += (1 / StoneTop)
	}
}

func cleanStone() {
	for x := 0; x < WorldWidth; x++ {
		grassHeight := HeightMap[x]
		if WorldMap[x][grassHeight] == blocksMap["stone"] {
			for y := grassHeight + StoneTopDeviation; y < WorldHeight; y++ {
				WorldMap[x][y] = blocksMap["sky"]
			}
		} else {
			for y := grassHeight + 1; y < WorldHeight; y++ {
				WorldMap[x][y] = blocksMap["sky"]
			}
		}
	}
}

func growGrass() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			if WorldMap[x][y] == blocksMap["dirt"] && WorldMap[x][y+1] == blocksMap["sky"] {
				WorldMap[x][y] = blocksMap["grass"]
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

func orientDirt() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			if x > 1 && x < WorldWidth-1 && y > 1 && y < WorldHeight-1 {
				if WorldMap[x][y] == blocksMap["dirt"] {
					if WorldMap[x+1][y] == blocksMap["sky"] { //Right
						if WorldMap[x][y-1] == blocksMap["sky"] {
							WorldMap[x][y] = blocksMap["bottomRightDirt"]
						} else if WorldMap[x][y+1] == blocksMap["sky"] {
							WorldMap[x][y] = blocksMap["topRightDirt"]
						} else {
							WorldMap[x][y] = blocksMap["rightDirt"]
						}
					} else if WorldMap[x-1][y] == blocksMap["sky"] { //Left
						if WorldMap[x][y-1] == blocksMap["sky"] {
							WorldMap[x][y] = blocksMap["bottomLeftDirt"]
						} else if WorldMap[x][y+1] == blocksMap["sky"] {
							WorldMap[x][y] = blocksMap["topLefttDirt"]
						} else {
							WorldMap[x][y] = blocksMap["leftDirt"]
						}
					} else if WorldMap[x][y+1] == blocksMap["sky"] { //Top
						WorldMap[x][y] = blocksMap["topRightDirt"]
					}
				}
			}
		}
	}
}
