package main

import (
	"fmt"
	"math/rand"
	"rapidengine/child"
	"time"

	perlin "github.com/aquilax/go-perlin"
)

func generateWorldTree() {
	// Make sure all blocks are loaded
	loadBlocks()

	// Make sure all children are loaded
	loadWorldChildren()

	// Randomize seed
	randomizeSeed()

	// Generate heightmap and place grass
	generateHeights()

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

	// Light up all blocks
	CreateLighting(WorldWidth/2, HeightMap[WorldWidth/2]+5, 0.9)

	// Fix backdirt
	createAllExtraBackdirt()

	// Set player starting position
	Player.SetPosition(float32(WorldWidth*BlockSize/2), float32((HeightMap[WorldWidth/2]+25)*BlockSize))

	// Create block children
	createCopies()
}

//  --------------------------------------------------
//  World Generation Functions
//  --------------------------------------------------

func generateHeights() {
	for x := 0; x < WorldWidth; x++ {
		HeightMap[x] = GrassMinimum + int(Flatness*noise1D(float64(x)/(WorldWidth/2))*WorldHeight)
	}
	for x := 0; x < WorldWidth; x++ {
		WorldMap[x][HeightMap[x]] = newBlock("grass")
	}
}

func fillHeights() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight-1; y++ {
			WorldMap[x][y] = newOrientBlock("dirt", "E")
			if WorldMap[x][y+1].ID == NameMap["grass"] {
				break
			}
		}
	}
}

func fillStone() {
	Generator = perlin.NewPerlin(1.2, 2, 2, int64(rand.Int()))
	stoneFrequency := StoneStartingFrequency
	for y := 0; y < WorldHeight; y++ {
		for x := 0; x < WorldWidth; x++ {
			n := noise2D(StoneNoiseScalar*float64(x)/WorldWidth*2, StoneNoiseScalar*float64(y)/WorldHeight*4)
			if n > stoneFrequency {
				WorldMap[x][y] = newBlock("stone")
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
				WorldMap[x][y] = newBlock("sky")
			}
		} else {
			for y := grassHeight + 1; y < WorldHeight; y++ {
				WorldMap[x][y] = newBlock("sky")
			}
		}
	}
}

func generateCaves() {
	Generator = perlin.NewPerlin(1.5, 2, 3, int64(rand.Int()))
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			n := noise2D(CaveNoiseScalar*float64(x)/WorldWidth*2, CaveNoiseScalar*float64(y)/WorldHeight*4)
			if n > CaveNoiseThreshold && y <= HeightMap[x] && WorldMap[x][y].Orientation == "E" {
				WorldMap[x][y] = newBlock("backdirt")
			}
		}
	}
}

func cleanBackDirt() {
	for i := 0; i < 2; i++ {
		for x := 1; x < WorldWidth-1; x++ {
			for y := WorldHeight - 2; y > 0; y-- {
				if WorldMap[x][y].ID == NameMap["backdirt"] {
					if WorldMap[x][y+1].ID == NameMap["sky"] {
						WorldMap[x][y] = newBlock("sky")
					}
				}
			}
		}
	}
	for x := 3; x < WorldWidth-3; x++ {
		for y := 3; y < WorldHeight-3; y++ {
			if WorldMap[x][y].ID == NameMap["backdirt"] {
				if WorldMap[x][y+1].ID != NameMap["backdirt"] && WorldMap[x+1][y].ID == NameMap["sky"] {
					cx := x
					for dy := y; dy > 0; dy-- {
						if WorldMap[cx][dy].ID != NameMap["backdirt"] {
							break
						}
						ccx := cx
						for {
							if WorldMap[ccx][dy].ID != NameMap["backdirt"] {
								break
							}
							WorldMap[ccx][dy] = newBlock("sky")
							ccx++
						}
						cx--
					}
				}
				if WorldMap[x][y+1].ID != NameMap["backdirt"] && WorldMap[x-1][y].ID == NameMap["sky"] {
					cx := x
					for dy := y; dy > 0; dy-- {
						if WorldMap[cx][dy].ID != NameMap["backdirt"] {
							break
						}
						ccx := cx
						for {
							if WorldMap[ccx][dy].ID != NameMap["backdirt"] {
								break
							}
							WorldMap[ccx][dy] = newBlock("sky")
							ccx--
						}
						cx++
					}
				}
			}
		}
	}
	for i := 0; i < 10; i++ {
		for x := 2; x < WorldWidth-2; x++ {
			for y := 2; y < WorldHeight-2; y++ {
				if WorldMap[x][y].ID == NameMap["backdirt"] && WorldMap[x][y+1].ID == NameMap["sky"] {
					for cy := y; y > 0; y-- {
						if WorldMap[x][cy].ID != NameMap["backdirt"] {
							break
						}
						WorldMap[x][cy] = newBlock("sky")
					}
				}
				if WorldMap[x][y].ID == NameMap["backdirt"] {
					if WorldMap[x-1][y].ID == NameMap["sky"] && WorldMap[x+1][y].ID == NameMap["sky"] {
						WorldMap[x][y] = newBlock("sky")
					}
				}
			}
		}
	}
}

func growGrass() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			if WorldMap[x][y].ID == NameMap["dirt"] && (WorldMap[x][y+1].ID == NameMap["sky"] || WorldMap[x][y+1].ID == NameMap["backdirt"]) {
				WorldMap[x][y] = newBlock("grass")
			}
		}
	}
}

func generateClouds() {
	for x := 0; x < WorldWidth; x++ {
		if rand.Float32() < 0.4 {
			CloudChild.AddCopy(
				child.ChildCopy{
					X:        float32(x * BlockSize),
					Y:        float32((rand.Intn(20) + HeightMap[x] + 15) * BlockSize),
					Material: &cloudMaterial,
					Darkness: 1,
				},
			)
			x += 400 / BlockSize
		}
	}
}

func generateNature() {
	for x := 1; x < WorldWidth-1; x++ {
		if WorldMap[x][HeightMap[x]].ID == NameMap["grass"] && (WorldMap[x][HeightMap[x]+1].ID == NameMap["sky"] || WorldMap[x][HeightMap[x]+1].ID == NameMap["backdirt"]) {
			natureRand := rand.Intn(16)
			if natureRand == 15 && WorldMap[x-1][HeightMap[x]+2].ID != NameMap["treeTrunk"] {
				hardAddCopy(x, HeightMap[x]+1, "treeTrunk", "nature", 0.9)
				height := 4 + rand.Intn(8)
				for i := 0; i < height; i++ {
					if rand.Intn(4) == 0 && i < height-2 && i > 0 {
						WorldMap[x-1][HeightMap[x]+i+2] = newBlock("treeBranchL1")
					}
					if rand.Intn(4) == 0 && i < height-2 && i > 0 {
						WorldMap[x+1][HeightMap[x]+i+2] = newBlock("treeBranchR1")
					}
					WorldMap[x][HeightMap[x]+2+i] = newBlock("treeTrunk")
				}
				WorldMap[x-1][HeightMap[x]+height+1] = newBlock("leaves") // TL
				WorldMap[x][HeightMap[x]+height+1] = newBlock("leaves")   // TM
				WorldMap[x+1][HeightMap[x]+height+1] = newBlock("leaves") // TR
				WorldMap[x-1][HeightMap[x]+height] = newBlock("leaves")   // ML
				WorldMap[x][HeightMap[x]+height] = newBlock("leaves")     // MM
				WorldMap[x+1][HeightMap[x]+height] = newBlock("leaves")   // MR
				WorldMap[x-1][HeightMap[x]+height-1] = newBlock("leaves") //BL
				WorldMap[x][HeightMap[x]+height-1] = newBlock("leaves")   // BM
				WorldMap[x+1][HeightMap[x]+height-1] = newBlock("leaves") //BL
			} else if natureRand > 13 {
				floraRand := rand.Intn(4) + 1
				floraType := fmt.Sprintf("flower%d", floraRand)
				if floraRand != 4 {
					hardAddCopy(x, HeightMap[x]+1, floraType, "nature", 1)
				} else {
					hardAddCopy(x, HeightMap[x]+1, "pebble", "nature", 1)
				}
			} else if natureRand > 9 {
				grassRand := rand.Intn(3) + 1
				grassType := fmt.Sprintf("topGrass%d", grassRand)
				hardAddCopy(x, HeightMap[x]+1, grassType, "nature", 1)
			}
		}
	}

}

//  --------------------------------------------------
//  World Generation Helpers
//  --------------------------------------------------

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
	return "shit spelling"
}

func noise2D(x, y float64) float64 {
	return (Generator.Noise2D(x, y) + 0.4) / 0.8
}

func noise1D(x float64) float64 {
	return (Generator.Noise1D(x) + 0.4) / 0.8
}

func randomizeSeed() {
	rand.Seed(time.Now().UTC().UnixNano())
	Generator = perlin.NewPerlin(2, 2, 10, int64(rand.Int()))
}

func hardAddCopy(x int, y int, name string, c string, dark float32) {
	if c == "nature" {
		NatureCopies[x][y] = child.ChildCopy{
			X:        float32(x * BlockSize),
			Y:        float32((y)*BlockSize - 5),
			Material: GetBlockIndex(NameMap[name]).GetMaterial(),
			Darkness: dark,
			ID:       1,
		}
	}
}
