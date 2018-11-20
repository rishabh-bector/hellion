package main

import (
	"math/rand"
)

func generateCaves() {
	randomizeSeed()

	CaveMap = make([][]bool, WorldWidth)
	for x := range CaveMap {
		CaveMap[x] = make([]bool, WorldHeight)
	}

	// Create caves with varying simplex noise
	for x := 0; x < WorldWidth; x++ {
		thresh := float32(CaveStartingThreshold)
		for y := HeightMap[x]; y >= 0; y-- {
			if n := rand.Float32(); n < thresh {
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
		caveSimulationStep(CaveBirthLimit, CaveDeathLimit)
	}

	// Randomly kill some caves
	for i := 0; i < 2000; i++ {
		removeRandomCave()
	}

	// Do second round of cellular automata simulations (enlarge)
	for i := 0; i < SecondCaveIterations; i++ {
		caveSimulationStep(SecondCaveBirthLimit, SecondCaveDeathLimit)
	}

	// Kill tiny caves
	//removeSmallCaves()

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

func caveSimulationStep(birthLim, deathLim int) {
	newMap := make([][]bool, WorldWidth)
	for x := range newMap {
		newMap[x] = make([]bool, WorldHeight)
	}

	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			nbs := getAliveNeighbors(x, y)

			if CaveMap[x][y] {
				if nbs < deathLim {
					newMap[x][y] = false
				} else {
					newMap[x][y] = true
				}
			} else {
				if nbs > birthLim {
					newMap[x][y] = true
				} else {
					newMap[x][y] = false
				}
			}
		}
	}

	CaveMap = newMap
}

func removeRandomCave() {
	for i := 0; i < 100; i++ {
		x := rand.Intn(WorldWidth)
		y := rand.Intn(WorldHeight)
		if CaveMap[x][y] {
			killCave(x, y)
			return
		}
	}
}

func killCave(x, y int) {
	if x < 0 || x >= WorldWidth || y < 0 || y >= WorldHeight {
		return
	}
	if CaveMap[x][y] == false {
		return
	}

	CaveMap[x][y] = false

	killCave(x+1, y)
	killCave(x-1, y)
	killCave(x, y+1)
	killCave(x, y-1)
}

func removeSmallCaves() {
	for x := 0; x < WorldWidth; x++ {
		for y := HeightMap[x]; y >= 0; y-- {
			if CaveMap[x][y] {
				if r := recurseSize(x, y, 6); r < 12 {
					killCave(x, y)
				}
			}
		}
	}
}

func recurseSize(x, y, limit int) int {
	if limit < 0 {
		return 0
	}
	if x < 0 || x >= WorldWidth || y < 0 || y >= WorldHeight {
		return 0
	}
	if CaveMap[x][y] == false {
		return 0
	}
	total := 1
	total += recurseSize(x+1, y, limit-1)
	total += recurseSize(x-1, y, limit-1)
	total += recurseSize(x, y+1, limit-1)
	total += recurseSize(x, y-1, limit-1)
	return total
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
