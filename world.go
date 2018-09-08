package main

import (
	"math/rand"
	"rapidengine"

	perlin "github.com/aquilax/go-perlin"
)

//  --------------------------------------------------
//  World Generation Parameters
//  --------------------------------------------------

const WorldHeight = 1000
const WorldWidth = 1000
const BlockSize = 25
const Flatness = 0.3

const GrassMinimum = 500

const CaveNoiseScalar = 30
const CaveNoiseThreshold = 0.75

const StoneNoiseScalar = 30.0
const StoneTop = 500.0
const StoneTopDeviation = 5
const StoneStartingFrequency = -0.3

//  --------------------------------------------------
//  --------------------------------------------------
//  --------------------------------------------------

var p *perlin.Perlin
var WorldMap [WorldWidth][WorldHeight]int
var HeightMap [WorldWidth]int

func generateWorld() {
	p = perlin.NewPerlin(2, 2, 10, int64(rand.Int()))

	// Fill world with 0s
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			WorldMap[x][y] = 0
		}
	}

	// Generate heightmap and place grass
	generateHeights()
	for x := 0; x < WorldWidth; x++ {
		WorldMap[x][HeightMap[x]] = 2
	}

	// Fill everything underneath grass with dirt
	fillHeights()

	// Generate stone based on height
	fillStone()

	// Clean up stone above ground
	cleanStone()

	// Generate caves
	generateCaves()
}

func createCopies() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			if WorldMap[x][y] != 0 {
				WorldChild.AddCopy(rapidengine.ChildCopy{
					float32(x * BlockSize),
					float32(y * BlockSize),
					engine.TextureControl.GetTexture(blocks[WorldMap[x][y]]),
				})
			}
		}
	}
}

func generateCaves() {
	p = perlin.NewPerlin(1.5, 2, 3, int64(rand.Int()))
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			n := noise2D(CaveNoiseScalar*float64(x)/WorldWidth, CaveNoiseScalar*float64(y)/WorldHeight)
			if n > CaveNoiseThreshold {
				WorldMap[x][y] = 0
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
		for y := 0; y < WorldHeight; y++ {
			WorldMap[x][y] = 1
			if WorldMap[x][y+1] == 2 {
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
			n := noise2D(StoneNoiseScalar*float64(x)/WorldWidth, StoneNoiseScalar*float64(y)/WorldHeight)
			if n > stoneFrequency {
				WorldMap[x][y] = 3
			}
		}
		stoneFrequency += (1 / StoneTop)
	}
}

func cleanStone() {
	for x := 0; x < WorldWidth; x++ {
		grassHeight := HeightMap[x]
		if WorldMap[x][grassHeight] == 3 {
			if WorldMap[x][grassHeight+StoneTopDeviation] == 0 {
				for y := grassHeight + StoneTopDeviation; y < WorldHeight; y++ {
					WorldMap[x][y] = 0
				}
			}
		} else {
			for y := grassHeight + 1; y < WorldHeight; y++ {
				WorldMap[x][y] = 0
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
