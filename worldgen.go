package main

import (
	"fmt"
	"math/rand"
	"rapidengine/child"
	"rapidengine/procedural"
	"time"
)

func generateWorldTree() {
	Engine.Logger.Info("Loading blocks...")
	// Make sure all blocks are loaded
	loadBlocks()

	// Make sure all children are loaded
	loadWorldChildren()

	// Randomize seed
	randomizeSeed()

	// Create a blank world tree
	WorldMap = NewWorldTree()

	Engine.Logger.Info("Placing dirt...")
	// Generate heightmap and place grass
	generateHeightMap()

	// Fill everything underneath grass with dirt
	generateDirt()

	Engine.Logger.Info("Placing stone...")
	// Generate stone based on height
	generateStone()

	// Clean up stone above ground
	cleanStone()

	Engine.Logger.Info("Generating caves...")
	// Generate caves
	generateCaves()

	// Clean back dirt
	cleanBackDirt()

	Engine.Logger.Info("Growing grass..")
	// Put grass and some top grass on dirt with air above it
	growGrass()

	Engine.Logger.Info("Generating nature...")
	// Create clouds
	generateClouds()

	// Place flowers and pebbles above grass
	generateNature()

	Engine.Logger.Info("Generating structures...")
	// Generate structure
	generateStructures()

	Engine.Logger.Info("Orienting blocks...")
	// Fix the orientation of blocks in the world
	orientBlocks("dirt", true)
	orientBlocks("grass", true)
	orientBlocks("stone", true)
	orientBlocks("leaves", true)

	// Fix backdirt
	createAllExtraBackdirt()
	orientBlocks("backdirt", true)

	Engine.Logger.Info("Creating light...")
	// Light up all blocks
	CreateLighting(WorldWidth/2, HeightMap[WorldWidth/2]+5, 0.9)

	// Save world to image
	WorldMap.writeToImage()

	// Set player starting position
	Player.SetPosition(float32(WorldWidth*BlockSize/2), float32((HeightMap[WorldWidth/2]+25)*BlockSize))
}

//  --------------------------------------------------
//  World Generation Functions
//  --------------------------------------------------

func generateHeightMap() {
	randomizeSeed()
	gen := procedural.NewSimplexGenerator(0.001, 1, 0.5, 8, Seed)

	for x := 0; x < WorldWidth; x++ {
		HeightMap[x] = GrassMinimum + int(Flatness*gen.Noise1D(float64(x))*float64(WorldHeight))
	}
	for x := 0; x < WorldWidth; x++ {
		createWorldBlock(x, HeightMap[x], "grass")
	}
}

func generateDirt() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight-1; y++ {
			createWorldBlock(x, y, "dirt")
			if WorldMap.GetWorldBlockName(x, y+1) == "grass" {
				break
			}
		}
	}
}

func generateStone() {
	randomizeSeed()
	gen := procedural.NewSimplexGenerator(10, 1, 0.5, 5, Seed)

	for x := 0; x < WorldWidth; x++ {
		stoneFrequency := float64(StoneStartingFrequency)
		for y := HeightMap[x]; y >= 0; y-- {
			n := gen.Noise2D(float64(x)/300, float64(y)/300)

			if n < stoneFrequency {
				createWorldBlock(x, y, "stone")
			}

			stoneFrequency += StoneFrequencyDelta
			if stoneFrequency > StoneEndingFrequency {
				stoneFrequency = StoneEndingFrequency
			}
		}
	}
}

func cleanStone() {
	for x := 0; x < WorldWidth; x++ {
		grassHeight := HeightMap[x]
		if WorldMap.GetWorldBlockName(x, grassHeight) == "stone" {
			for y := grassHeight + StoneTopDeviation; y < WorldHeight; y++ {
				createWorldBlock(x, y, "sky")
			}
		} else {
			for y := grassHeight + 1; y < WorldHeight; y++ {
				createWorldBlock(x, y, "sky")
			}
		}
	}
}

func generateCaves() {
	randomizeSeed()

	CaveMap = make([][]bool, WorldWidth)
	for x := range CaveMap {
		CaveMap[x] = make([]bool, WorldHeight)
	}

	// Create caves with varying simplex noise
	gen := procedural.NewSimplexGenerator(100, 1, 0.5, 8, Seed)
	for x := 0; x < WorldWidth; x++ {
		thresh := float64(CaveStartingThreshold)
		for y := HeightMap[x]; y >= 0; y-- {
			n := gen.Noise2D(float64(x)/3000, float64(y)/3000)

			if n < thresh {
				CaveMap[x][y] = true
			}

			thresh += CaveThresholdDelta
			if thresh > CaveEndingThreshold {
				thresh = CaveEndingThreshold
			}
		}
	}

	// Do cellular automata simulations
	for i := 0; i < CaveIterations; i++ {
		caveSimulationStep()
	}

	// Translate to worldmap
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			if CaveMap[x][y] {
				if y <= HeightMap[x] {
					WorldMap.RemoveWorldBlock(x, y)
					createBackBlock(x, y, "backdirt")
				}
			}
		}
	}
}

func caveSimulationStep() {
	newMap := make([][]bool, WorldWidth)
	for x := range newMap {
		newMap[x] = make([]bool, WorldHeight)
	}

	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			nbs := getAliveNeighbors(x, y)

			if CaveMap[x][y] {
				if nbs < CaveDeathLimit {
					newMap[x][y] = false
				} else {
					newMap[x][y] = true
				}
			} else {
				if nbs > CaveBirthLimit {
					newMap[x][y] = true
				} else {
					newMap[x][y] = false
				}
			}
		}
	}

	CaveMap = newMap
}

func getAliveNeighbors(x, y int) int {
	count := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			nx := x + i
			ny := y + j

			if i == 0 && j == 0 {

			} else if nx < 0 || nx >= WorldWidth || ny < 0 || ny >= WorldHeight {
				count++
			} else if CaveMap[nx][ny] {
				count++
			}
		}
	}
	return count
}

func cleanBackDirt() {

}

func growGrass() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			if WorldMap.GetWorldBlockName(x, y) == "dirt" && (WorldMap.GetWorldBlockName(x, y+1) == "sky" || WorldMap.GetBackBlockName(x, y+1) == "backdirt") {
				createWorldBlock(x, y, "grass")
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
		if WorldMap.GetWorldBlockName(x, HeightMap[x]) == "grass" && WorldMap.GetWorldBlockName(x, HeightMap[x]+1) == "sky" || WorldMap.GetWorldBlockName(x, HeightMap[x]+1) == "backdirt" {
			natureRand := rand.Intn(16)
			if natureRand == 15 && WorldMap.GetWorldBlockName(x-1, HeightMap[x]+2) != "treetrunk" {
				createNatureBlock(x, HeightMap[x]+1, "treeBottomRoot")
				height := 4 + rand.Intn(8)
				for i := 0; i < height; i++ {
					if rand.Intn(4) == 0 && i < height-2 && i > 0 {
						createNatureBlock(x-1, HeightMap[x]+i+2, "treeBranchL1")
					}
					if rand.Intn(4) == 0 && i < height-2 && i > 0 {
						createNatureBlock(x+1, HeightMap[x]+i+2, "treeBranchR1")
					}
					createNatureBlock(x, HeightMap[x]+i+2, "treeTrunk")
				}

				createNatureBlock(x-1, HeightMap[x]+height+1, "leaves") // TL
				createNatureBlock(x, HeightMap[x]+height+1, "leaves")   // TM
				createNatureBlock(x+1, HeightMap[x]+height+1, "leaves") // TR
				createNatureBlock(x-1, HeightMap[x]+height, "leaves")   // ML
				createNatureBlock(x, HeightMap[x]+height, "leaves")     // MM
				createNatureBlock(x+1, HeightMap[x]+height, "leaves")   // MR
				createNatureBlock(x-1, HeightMap[x]+height-1, "leaves") //BL
				createNatureBlock(x, HeightMap[x]+height-1, "leaves")   // BM
				createNatureBlock(x+1, HeightMap[x]+height-1, "leaves") //BL
			} else if natureRand > 13 {
				floraRand := rand.Intn(4) + 1
				floraType := fmt.Sprintf("flower%d", floraRand)
				if floraRand != 4 {
					createNatureBlock(x, HeightMap[x]+1, floraType)
				} else {
					createNatureBlock(x, HeightMap[x]+1, "pebble")
				}
			} else if natureRand > 9 {
				grassRand := rand.Intn(3) + 1
				grassType := fmt.Sprintf("topGrass%d", grassRand)
				createNatureBlock(x, HeightMap[x]+1, grassType)
			}
		}
	}

}

//  --------------------------------------------------
//  World Generation Helpers
//  --------------------------------------------------

func isBackBlock(name string) bool {
	for _, transparent := range TransparentBlocks {
		if name == transparent {
			return true
		}
	}
	return false
}

func randomizeSeed() {
	Seed = time.Now().UTC().UnixNano()
	rand.Seed(Seed)
}
