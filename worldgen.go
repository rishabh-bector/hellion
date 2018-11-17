package main

import (
	"math/rand"
	"rapidengine/child"
	"rapidengine/procedural"
	"time"
)

func generateWorldTree() {
	// Make sure all blocks are loaded
	loadBlocks()

	// Make sure all children are loaded
	loadWorldChildren()

	// Randomize seed
	randomizeSeed()

	// Create a blank world tree
	WorldMap = NewWorldTree()

	// Generate heightmap and place grass
	generateHeightMap()

	// Fill everything underneath grass with dirt
	generateDirt()

	// Generate stone based on height
	generateStone()

	// Clean up stone above ground
	cleanStone()

	// Generate caves
	//generateCaves()

	// Clean back dirt
	cleanBackDirt()

	// Put grass and some top grass on dirt with air above it
	growGrass()

	// Create clouds
	generateClouds()

	// Place flowers and pebbles above grass
	//generateNature()

	// Generate structure
	generateStructures()

	// Fix the orientation of blocks in the world
	orientBlocks("dirt", true)
	orientBlocks("grass", true)
	orientBlocks("stone", true)
	orientBlocks("leaves", true)

	// Fix backdirt
	createAllExtraBackdirt()
	orientBlocks("backdirt", true)

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

/*func generateCaves() {
	Generator = perlin.NewPerlin(1.5, 2, 3, int64(rand.Int()))
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			n := noise2D(CaveNoiseScalar*float64(x)/WorldWidth*2, CaveNoiseScalar*float64(y)/WorldHeight*4)
			if n > CaveNoiseThreshold && y <= HeightMap[x] {
				WorldMap.RemoveWorldBlock(x, y)
				createBackBlock(x, y, "backdirt")
			}
		}
	}
}*/

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

/*func generateNature() {
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

}*/

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

/*func blockType(name string) string {
	for _, green := range natureBlocks {
		if NameMap[name] == NameMap[green] {
			return "nature"
		}
	}
	return "shit spelling"
}*/

func randomizeSeed() {
	Seed = time.Now().UTC().UnixNano()
	rand.Seed(Seed)
}
