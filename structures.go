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
		{3, 4, 0, 0, 0, 0, 4, 3},
		{0, 3, 4, 4, 4, 4, 3, 0},
		{0, 0, 3, 3, 3, 3, 0, 0},
		{0, 0, 3, 0, 0, 3, 0, 0},
		{0, 0, 3, 0, 0, 3, 0, 0},
		{0, 0, 3, 0, 0, 3, 0, 0},
		{0, 0, 3, 0, 0, 3, 0, 0},
		{0, 0, 3, 0, 0, 3, 0, 0},
		{0, 0, 3, 0, 0, 3, 0, 0},
	},
	FillType: "none",
}

var goblinCampBarracks = Building{
	Layout: [][]int{
		{10, 0, 0, 0, 0, 0, 10, 0},
		{10, 10, 8, 8, 8, 10, 10, 0},
		{0, 0, 0, 8, 8, 8, 0, 0},
		{3, 3, 3, 3, 3, 3, 3, 3},
		{3, 4, 4, 4, 4, 4, 4, 3},
		{3, 4, 4, 4, 4, 4, 4, 3},
		{3, 4, 4, 4, 4, 4, 4, 3},
		{3, 4, 4, 4, 4, 4, 4, 3},
	},
	FillType: "none",
}

var goblinCampHall = Building{
	Layout: [][]int{
		{0, 0, 0, 0, 0, 0, 5, 5, 5},
		{0, 0, 0, 0, 0, 0, 5, 5, 5},
		{0, 0, 0, 3, 3, 0, 4, 0, 0},
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
		{3, 4, 4, 4, 4, 4, 4, 3},
		{0, 3, 3, 3, 3, 3, 3, 0},
		{0, 0, 3, 3, 3, 3, 0, 0},
		{0, 0, 3, 3, 3, 3, 0, 0},
		{0, 0, 3, 3, 4, 3, 0, 0},
		{0, 0, 3, 4, 4, 3, 0, 0},
		{0, 0, 3, 4, 4, 3, 0, 0},
		{0, 0, 3, 4, 4, 3, 0, 0},
	},
	FillType: "none",
}

var goblinFortressArcherTowerMedium = Building{
	Layout: [][]int{
		{0, 3, 4, 0, 4, 0, 4, 0, 4, 3, 0},
		{3, 3, 4, 4, 4, 4, 4, 4, 4, 3, 3},
		{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{0, 3, 3, 3, 3, 3, 3, 3, 3, 3, 0},
		{0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0},
		{0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0},
		{0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0},
		{0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0},
		{0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0},
		{0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0},
		{0, 0, 0, 3, 3, 4, 4, 3, 0, 0, 0},
		{0, 0, 0, 3, 4, 4, 4, 3, 0, 0, 0},
		{0, 0, 0, 3, 4, 4, 4, 3, 0, 0, 0},
		{0, 0, 0, 3, 4, 4, 4, 3, 0, 0, 0},
	},
	FillType: "none",
}

var goblinFortressArcherTowerLarge = Building{
	Layout: [][]int{
		{0, 3, 4, 4, 0, 4, 4, 4, 0, 4, 4, 3, 0},
		{3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 3, 0},
		{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{0, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 0},
		{0, 0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 3, 3, 3, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 3, 4, 4, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 3, 0, 0, 0, 0},
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
		{0, 3, 0, 0, 0, 3, 3, 0, 0, 3, 3, 0, 0, 3, 3, 0, 0, 0, 3, 3, 0},
		{3, 3, 0, 0, 0, 3, 3, 0, 0, 3, 3, 0, 0, 3, 3, 0, 0, 0, 3, 3, 3},
		{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 3, 3},
		{3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 3, 3},
		{3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 3, 3, 3},
		{3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 3, 3, 3},
		{0, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 3, 3, 0},
		{0, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 3, 3, 0},
		{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{0, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 3, 4, 4, 4, 3, 3, 3, 3, 3, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 4, 4, 4, 3, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 4, 4, 4, 4, 3, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 4, 3, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 3, 4, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 3, 4, 4, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 3, 4, 4, 4, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 3, 4, 4, 4, 4, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 3, 4, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 4, 3, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 4, 4, 3, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 4, 4, 4, 3, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 4, 4, 4, 4, 3, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 4, 3, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 3, 4, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 3, 4, 4, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 3, 4, 4, 4, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 3, 4, 4, 4, 4, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 3, 4, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 4, 3, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 4, 4, 3, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 4, 4, 4, 3, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 4, 4, 4, 4, 3, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 4, 3, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 4, 3, 4, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 4, 4, 4, 3, 4, 4, 4, 4, 4, 4, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 4, 4, 4, 3, 4, 4, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 4, 4, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0},
	},
	FillType: "none",
}

var goblinFortressBarracks = Building{
	Layout: [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 3, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 3, 3, 3, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 3, 3, 3, 3, 3, 3, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 3, 3, 3, 3, 3, 3, 3, 3, 0, 0},
		{0, 0, 0, 0, 0, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 0},
		{0, 0, 0, 0, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0},
		{0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0},
		{0, 0, 0, 0, 3, 3, 3, 0, 0, 0, 0, 0, 0, 0, 3, 0},
		{0, 0, 0, 3, 0, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 0},
		{0, 0, 3, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0},
		{0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0},
		{3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0},
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
		Startx := rand.Intn(current.Width/2) + current.XBeginning
		currentX := Startx
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
							createWorldBlock(currentX+x, lowestY+(height-y), GetBlockName(building.Layout[y][x]))
						}
					}
				}

				currentX += len(building.Layout)
			}
		}
	}
}

func fillStructureFloor(startx int, starty int, endx int, endy int) {

}

func generateStilts(x int, y int) {

}