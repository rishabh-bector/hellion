package main

import (
	"fmt"
	"math/rand"
	"rapidengine/child"

	perlin "github.com/aquilax/go-perlin"
)

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

func cleanBackDirt() {
	for i := 0; i < 2; i++ {
		for x := 1; x < WorldWidth-1; x++ {
			for y := WorldHeight - 2; y > 0; y-- {
				if WorldMap[x][y].ID == NameMap["backdirt"] {
					if WorldMap[x][y+1].ID == NameMap["sky"] {
						WorldMap[x][y] = NewBlock("sky")
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
							WorldMap[ccx][dy] = NewBlock("sky")
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
							WorldMap[ccx][dy] = NewBlock("sky")
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
						WorldMap[x][cy] = NewBlock("sky")
					}
				}
				if WorldMap[x][y].ID == NameMap["backdirt"] {
					if WorldMap[x-1][y].ID == NameMap["sky"] && WorldMap[x+1][y].ID == NameMap["sky"] {
						WorldMap[x][y] = NewBlock("sky")
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
				WorldMap[x][y] = NewBlock("grass")
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
						WorldMap[x-1][HeightMap[x]+i+2] = NewBlock("treeBranchL1")
					}
					if rand.Intn(4) == 0 && i < height-2 && i > 0 {
						WorldMap[x+1][HeightMap[x]+i+2] = NewBlock("treeBranchR1")
					}
					WorldMap[x][HeightMap[x]+2+i] = NewBlock("treeTrunk")
				}
				WorldMap[x-1][HeightMap[x]+height+1] = NewBlock("leaves") // TL
				WorldMap[x][HeightMap[x]+height+1] = NewBlock("leaves")   // TM
				WorldMap[x+1][HeightMap[x]+height+1] = NewBlock("leaves") // TR
				WorldMap[x-1][HeightMap[x]+height] = NewBlock("leaves")   // ML
				WorldMap[x][HeightMap[x]+height] = NewBlock("leaves")     // MM
				WorldMap[x+1][HeightMap[x]+height] = NewBlock("leaves")   // MR
				WorldMap[x-1][HeightMap[x]+height-1] = NewBlock("leaves") //BL
				WorldMap[x][HeightMap[x]+height-1] = NewBlock("leaves")   // BM
				WorldMap[x+1][HeightMap[x]+height-1] = NewBlock("leaves") //BL
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
	return (p.Noise2D(x, y) + 0.4) / 0.8
}

func noise1D(x float64) float64 {
	return (p.Noise1D(x) + 0.4) / 0.8
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
