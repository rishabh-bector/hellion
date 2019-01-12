package main

import (
	"fmt"
	"math/rand"
	"rapidengine/child"
	"rapidengine/procedural"
	"time"
)

func initializeWorldTree() {
	loadBlocks()
	WorldMap = NewWorldTree()
}

func generateWorldTree() {

	Engine.Logger.Info("Loading blocks...")
	ProgressText.Text = "Loading blocks..."
	ProgressBar.SetPercentage(10)
	updateLoadingScreen()

	// Randomize seed
	randomizeSeed()

	Engine.Logger.Info("Placing dirt...")
	ProgressText.Text = "Placing dirt..."
	ProgressBar.SetPercentage(20)
	updateLoadingScreen()

	// Generate heightmap and place grass
	generateHeightMap()

	// Fill everything underneath grass with dirt
	generateDirt()

	Engine.Logger.Info("Placing stone...")
	ProgressText.Text = "Placing stone..."
	ProgressBar.SetPercentage(30)
	updateLoadingScreen()

	// Generate stone based on height
	generateStone()

	// Clean up stone above ground
	cleanStone()

	Engine.Logger.Info("Generating caves...")
	ProgressText.Text = "Generating caves..."
	ProgressBar.SetPercentage(40)
	updateLoadingScreen()

	// Generate caves
	generateCaves()

	// Clean back dirt
	cleanBackDirt()

	Engine.Logger.Info("Growing grass...")
	ProgressText.Text = "Growing grass..."
	ProgressBar.SetPercentage(50)
	updateLoadingScreen()

	// Put grass and some top grass on dirt with air above it
	growGrass()

	Engine.Logger.Info("Generating nature...")
	ProgressText.Text = "Generating nature..."
	ProgressBar.SetPercentage(60)
	updateLoadingScreen()

	// Create clouds
	generateClouds()

	// Place flowers and pebbles above grass
	generateNature()

	Engine.Logger.Info("Generating structures...")
	ProgressText.Text = "Generating structures..."
	ProgressBar.SetPercentage(70)
	updateLoadingScreen()

	// Generate structure
	generateStructures()
	generateAllDungeons()

	Engine.Logger.Info("Orienting blocks...")
	ProgressText.Text = "Orienting blocks..."
	ProgressBar.SetPercentage(80)
	updateLoadingScreen()

	// Fix the orientation of blocks in the world
	orientBlocks("dirt", true)
	orientBlocks("grass", true)
	orientBlocks("stone", true)
	orientBlocks("leaves", true)

	// Fix backdirt
	createAllExtraBackdirt()
	orientBlocks("backdirt", true)

	Engine.Logger.Info("Creating light...")
	ProgressText.Text = "Creating light..."
	ProgressBar.SetPercentage(90)
	updateLoadingScreen()

	// Light up all blocks
	CreateLighting(WorldWidth/2, HeightMap[WorldWidth/2]+5, 0.9)

	// Save world to image
	WorldMap.writeToImage()

	ProgressBar.SetPercentage(100)
	updateLoadingScreen()

	// Set player starting position
	Player1.PlayerChild.SetPosition(float32(WorldWidth*BlockSize/2), float32((HeightMap[WorldWidth/2]+50)*BlockSize))
	Player1.Hitbox1.X = Player1.PlayerChild.X + 90
	Player1.Hitbox1.Y = Player1.PlayerChild.Y + 60

	Engine.SceneControl.SetCurrentScene(WorldScene)
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
		createWorldBlock(x, HeightMap[x], "dirt")
	}
}

func generateDirt() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight-1; y++ {
			createWorldBlock(x, y, "dirt")
			if WorldMap.GetWorldBlockName(x, y+1) == "dirt" {
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

func cleanBackDirt() {

}

func growGrass() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			if WorldMap.GetWorldBlockName(x, y) == "dirt" && (WorldMap.GetWorldBlockName(x, y+1) == "sky" || WorldMap.GetBackBlockName(x, y+1) == "backdirt") {
				createGrassBlock(x, y, "grasstop")
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
					Material: cloudMaterial,
					Darkness: 1,
				},
			)
			x += 400 / BlockSize
		}
	}
}

func generateNature() {
	for x := 1; x < WorldWidth-1; x++ {
		if WorldMap.GetWorldBlockName(x, HeightMap[x]+1) == "sky" || WorldMap.GetWorldBlockName(x, HeightMap[x]+1) == "backdirt" {
			natureRand := rand.Intn(16)
			if natureRand == 15 && WorldMap.GetWorldBlockName(x-1, HeightMap[x]+2) != "treetrunk" {

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
