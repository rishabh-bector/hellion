package main

import "math/rand"

// NameMap = map[string]int{
// 	"sky":            0,
// 	"dirt":           1,
// 	"grass":          2,
// 	"stone":          3,
// 	"backdirt":       4,
// 	"leaves":         5,
// 	"treeRightRoot":  6,
// 	"treeLeftRoot":   7,
// 	"treeTrunk":      8,
// 	"treeBottomRoot": 9,
// 	"topGrass1":      10,
// 	"topGrass2":      11,
// 	"topGrass3":      12,
// 	"treeBranchR1":   13,
// 	"treeBranchL1":   14,
// 	"flower1":        15,
// 	"flower2":        16,
// 	"flower3":        17,
// 	"pebble":         18,
// 	"torch":          19,
// }

type Structure struct {
	Layout []Building

	PlaceMethod string // fill, mesh, hover, pillars

	XBeginning int
	XEnd       int
	Width      int

	Spacing []int
}

type Building struct {
	Layout [][]int

	FillType string // stilts, none
}

// Goblin Camp

var goblinCampArcherTower = Building{
	Layout: [][]int{
		{20, 4, 0, 0, 0, 0, 4, 20},
		{0, 20, 4, 4, 4, 4, 20, 0},
		{0, 0, 20, 20, 20, 20, 0, 0},
		{0, 0, 20, 0, 0, 20, 0, 0},
		{0, 0, 20, 0, 0, 20, 0, 0},
		{0, 0, 20, 0, 0, 20, 0, 0},
		{0, 0, 20, 0, 0, 20, 0, 0},
		{0, 0, 20, 0, 0, 20, 0, 0},
		{0, 0, 20, 0, 0, 20, 0, 0},
	},
	FillType: "none",
}

var goblinCampBarracks = Building{
	Layout: [][]int{
		{10, 0, 0, 0, 0, 0, 10, 0},
		{10, 10, 8, 8, 8, 10, 10, 0},
		{0, 0, 0, 8, 8, 8, 0, 0},
		{20, 20, 20, 20, 20, 20, 20, 20},
		{20, 4, 4, 4, 4, 4, 4, 20},
		{20, 4, 4, 4, 4, 4, 4, 20},
		{20, 4, 4, 4, 4, 4, 4, 20},
		{20, 4, 4, 4, 4, 4, 4, 20},
	},
	FillType: "none",
}

var goblinCampHall = Building{
	Layout: [][]int{
		{0, 0, 0, 0, 0, 0, 5, 5, 5},
		{0, 0, 0, 0, 0, 0, 5, 5, 5},
		{0, 0, 0, 20, 20, 0, 4, 0, 0},
		{0, 0, 2, 2, 2, 2, 4, 0, 0},
		{0, 2, 2, 2, 2, 2, 2, 0, 0},
		{7, 8, 4, 4, 4, 4, 8, 6, 0},
		{0, 8, 4, 4, 4, 4, 4, 0, 0},
		{0, 8, 4, 4, 4, 4, 4, 0, 0},
	},
	FillType: "none",
}

var goblinCamp = Structure{
	Layout: []Building{
		goblinCampArcherTower,
		goblinCampBarracks,
		goblinCampBarracks,
		goblinCampArcherTower,
		goblinCampHall,
	},
	PlaceMethod: "mesh",
	XBeginning:  300,
	XEnd:        520,
	Width:       60,
	Spacing:     []int{5, 4, 5, 6},
}

// Goblin Fotress

var goblinFortressArcherTowerSmall = Building{
	Layout: [][]int{
		{0, 4, 0, 4, 4, 0, 4, 0},
		{20, 4, 4, 4, 4, 4, 4, 20},
		{0, 20, 20, 20, 20, 20, 20, 0},
		{0, 0, 20, 20, 20, 20, 0, 0},
		{0, 0, 20, 20, 20, 20, 0, 0},
		{0, 0, 20, 20, 4, 20, 0, 0},
		{0, 0, 20, 4, 4, 20, 0, 0},
		{0, 0, 20, 4, 4, 20, 0, 0},
		{0, 0, 20, 4, 4, 20, 0, 0},
	},
	FillType: "none",
}

var goblinFortressArcherTowerMedium = Building{
	Layout: [][]int{
		{0, 20, 4, 0, 4, 0, 4, 0, 4, 20, 0},
		{20, 20, 4, 4, 4, 4, 4, 4, 4, 20, 20},
		{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{0, 20, 20, 20, 20, 20, 20, 20, 20, 20, 0},
		{0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0},
		{0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0},
		{0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0},
		{0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0},
		{0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0},
		{0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0},
		{0, 0, 0, 20, 20, 4, 4, 20, 0, 0, 0},
		{0, 0, 0, 20, 4, 4, 4, 20, 0, 0, 0},
		{0, 0, 0, 20, 4, 4, 4, 20, 0, 0, 0},
		{0, 0, 0, 20, 4, 4, 4, 20, 0, 0, 0},
	},
	FillType: "none",
}

var goblinFortressArcherTowerLarge = Building{
	Layout: [][]int{
		{0, 20, 4, 4, 0, 4, 4, 4, 0, 4, 4, 20, 0},
		{20, 20, 4, 4, 4, 4, 4, 4, 4, 4, 4, 20, 0},
		{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{0, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 0},
		{0, 0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 20, 20, 20, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 20, 4, 4, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 20, 0, 0, 0, 0},
	},
	FillType: "none",
}

var goblinFortressTower = Building{
	Layout: [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 5, 5, 5, 5, 5, 5},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 5, 5, 5, 5, 5, 5},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 5, 5, 5, 5, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0},
		{0, 20, 0, 0, 0, 20, 20, 0, 0, 20, 20, 0, 0, 20, 20, 0, 0, 0, 20, 20, 0},
		{20, 20, 0, 0, 0, 20, 20, 0, 0, 20, 20, 0, 0, 20, 20, 0, 0, 0, 20, 20, 20},
		{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{20, 20, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 20, 20},
		{20, 20, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 20, 20},
		{20, 20, 20, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 20, 20, 20},
		{20, 20, 20, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 20, 20, 20},
		{0, 20, 20, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 20, 20, 0},
		{0, 20, 20, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 20, 20, 0},
		{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{0, 20, 20, 20, 20, 4, 4, 4, 4, 4, 4, 20, 4, 4, 4, 20, 20, 20, 20, 20, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 4, 4, 20, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 4, 4, 4, 20, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 4, 20, 20, 20, 20, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 20, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 20, 4, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 20, 4, 4, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 20, 4, 4, 4, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 20, 4, 4, 4, 4, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 20, 20, 20, 20, 4, 4, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 20, 4, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 20, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 4, 20, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 4, 4, 20, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 4, 4, 4, 20, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 4, 20, 20, 20, 20, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 20, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 20, 4, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 20, 4, 4, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 20, 4, 4, 4, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 20, 4, 4, 4, 4, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 20, 20, 20, 20, 4, 4, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 20, 4, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 20, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 4, 20, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 4, 4, 20, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 4, 4, 4, 20, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 4, 20, 20, 20, 20, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 20, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 4, 20, 4, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 20, 4, 4, 4, 20, 4, 4, 4, 4, 4, 4, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 4, 4, 4, 20, 4, 4, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 4, 4, 20, 4, 4, 4, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0},
	},
	FillType: "none",
}

var goblinFortressBarracks = Building{
	Layout: [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 20, 20, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 20, 20, 20, 20, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 20, 20, 20, 20, 20, 20, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 20, 20, 20, 20, 20, 20, 20, 20, 0, 0},
		{0, 0, 0, 0, 0, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 0},
		{0, 0, 0, 0, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		{0, 0, 0, 0, 0, 20, 0, 0, 0, 0, 0, 0, 0, 0, 20, 0},
		{0, 0, 0, 0, 0, 20, 0, 0, 0, 0, 0, 0, 0, 0, 20, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 20, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 20, 0},
		{0, 0, 0, 0, 20, 20, 20, 0, 0, 0, 0, 0, 0, 0, 20, 0},
		{0, 0, 0, 20, 0, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 0},
		{0, 0, 20, 0, 0, 20, 0, 0, 0, 0, 0, 0, 0, 0, 20, 0},
		{0, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 20, 0},
		{20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 20, 0},
	},
	FillType: "none",
}

var goblinFortress = Structure{
	Layout: []Building{
		goblinFortressArcherTowerSmall,
		goblinFortressArcherTowerMedium,
		goblinFortressArcherTowerLarge,
		goblinFortressTower,
		goblinFortressBarracks,
	},
	PlaceMethod: "mesh",
	XBeginning:  650,
	XEnd:        840,
	Width:       120,
	Spacing:     []int{16, 9, 9, 0},
}

var structures = []Structure{
	goblinCamp, goblinFortress,
}

func generateStructures() {
	for _, current := range structures {
		RightStartx := rand.Intn(current.Width/2) + current.XBeginning + WorldWidth/2
		LeftStartx := WorldWidth/2 - rand.Intn(current.Width/2) - current.XBeginning
		currentX := RightStartx
		if current.PlaceMethod == "mesh" {
			for i, building := range current.Layout {
				if i != 0 {
					currentX += current.Spacing[i-1] + 10
				}
				lowestY := HeightMap[currentX]
				for x := currentX; x < currentX+len(building.Layout[0]); x++ {
					if HeightMap[x] < lowestY {
						lowestY = HeightMap[x]
					}
				}

				height := len(building.Layout)

				for y := 0; y < height; y++ {
					for x := 0; x < len(building.Layout[y]); x++ {
						if WorldMap.GetWorldBlockID(currentX+x, lowestY+(height-y)) == "00000" {
							if bname := GetBlockName(building.Layout[y][x]); bname == "backdirt" {
								createBackBlock(currentX+x, lowestY+(height-y), bname)
							} else {
								WorldMap.RemoveNatureBlock(currentX+x, lowestY+(height-y))
								createWorldBlock(currentX+x, lowestY+(height-y), bname)
							}
						}
					}
				}

				//currentX += len(building.Layout)
			}
			currentX = LeftStartx
			for i := len(current.Layout) - 1; i > -1; i-- {
				building := current.Layout[i]
				if i != 0 {
					currentX -= current.Spacing[i-1] + 10
				}
				lowestY := HeightMap[currentX]
				for x := currentX; x > currentX-len(building.Layout[0]); x-- {
					if HeightMap[x] < lowestY {
						lowestY = HeightMap[x]
					}
				}
				height := len(building.Layout)
				leftLayout := flipMatrix(building.Layout)
				for y := 0; y < height; y++ {
					for x := 0; x < len(leftLayout[y]); x++ {
						if WorldMap.GetWorldBlockID(currentX+x, lowestY+(height-y)) == "00000" {
							if bname := GetBlockName(leftLayout[y][x]); bname == "backdirt" {
								createBackBlock(currentX+x, lowestY+(height-y), bname)
							} else {
								WorldMap.RemoveNatureBlock(currentX+x, lowestY+(height-y))
								createWorldBlock(currentX+x, lowestY+(height-y), bname)
							}
						}
					}
				}
			}
		}
	}
}

func fillStructureFloor(startx int, starty int, endx int, endy int) {

}

func generateStilts(x int, y int) {

}

func flipMatrix(mat [][]int) [][]int {
	fin := make([][]int, len(mat))
	for item := range fin {
		fin[item] = make([]int, len(mat[item]))
	}

	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[x]); y++ {
			fin[x][y] = mat[x][len(mat[x])-y-1]
		}
	}

	return fin
}
