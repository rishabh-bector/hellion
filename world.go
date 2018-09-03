package main

import (
	"rapidengine"

	perlin "github.com/aquilax/go-perlin"
)

const WorldHeight = 1000
const WorldWidth = 1000
const BlockSize = 25
const Flatness = 0.3

var p *perlin.Perlin
var world [WorldWidth][WorldHeight]int

func generateWorld() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			world[x][y] = 0
		}
	}
	heights := generateHeights()
	for x := 0; x < WorldWidth; x++ {
		world[x][heights[x]] = 2
	}
	fillHeights()
}

func createCopies() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			if world[x][y] != 0 {
				child1.AddCopy(rapidengine.ChildCopy{float32(x * BlockSize), float32(y * BlockSize), engine.TextureControl.GetTexture(blocks[world[x][y]])})
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

func generateHeights() [WorldWidth]int {
	heights := [WorldWidth]int{}
	for x := 0; x < WorldWidth; x++ {
		heights[x] = int(Flatness * noise1D(float64(x)/WorldWidth) * WorldHeight)
	}
	return heights
}

func fillHeights() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			world[x][y] = 1
			if world[x][y+1] == 2 {
				break
			}
		}
	}
}
