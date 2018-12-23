package main

import (
	"fmt"
	"rapidengine/material"
	"strconv"
)

var BlockMap map[string]*Block

type Block struct {
	Material *material.BasicMaterial

	SaveColor [3]int

	// Block orientations (16 possible):
	// First character - Left/Right (L/R)
	// Second character - Top/Bottom (T/B)
	// Either can have A (all) or N (none)
	OrientEnabled   bool
	OrientVariation int32
	Orientations    [16]*material.BasicMaterial

	LightBlock float32
}

func (block *Block) GetMaterial(direction string) *material.BasicMaterial {
	if block.OrientEnabled {
		i, err := strconv.Atoi(OrientationsMap[direction])
		if err != nil {
			panic(err)
		}
		return block.Orientations[i]
	}
	return block.Material
}

func (block *Block) CreateOrientations(orientVariation int32) {
	block.OrientEnabled = true
	block.OrientVariation = orientVariation
	for dir := range OrientationsMap {
		newM := *block.Material
		newM.AlphaMap = Engine.TextureControl.GetTexture(fmt.Sprintf("%v%s", orientVariation, dir))
		newM.AlphaMapLevel = 1
		i, _ := strconv.Atoi(OrientationsMap[dir])
		block.Orientations[i] = &newM
	}
}

func loadBlocks() {
	// Main Blocks
	Engine.TextureControl.NewTexture("./assets/blocks/dirt/dirt1.png", "dirt", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/grass/grass_g.png", "grass", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/stone/stone1.png", "stone", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/stone/stoneBrick.png", "stoneBrick", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/torch.png", "torch", "pixel")

	// Back-Blocks
	Engine.TextureControl.NewTexture("./assets/blocks/backblocks/backdirt8.png", "backdirt", "pixel")

	// Tree
	Engine.TextureControl.NewTexture("./assets/blocks/tree/treeTrunk3.png", "treeTrunk", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/tree/treeLeftRoot.png", "treeLeftRoot", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/tree/treeRightRoot.png", "treeRightRoot", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/tree/treeTrunk3.png", "treeBottomRoot", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/tree/leaves.png", "leaves", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/tree/treeBranchR1.png", "treeBranchR1", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/tree/treeBranchL1.png", "treeBranchL1", "pixel")

	// Flora
	Engine.TextureControl.NewTexture("./assets/blocks/grass/grasstop.png", "grasstop", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/grass/topGrass1_g.png", "topGrass1", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/grass/topGrass2_g.png", "topGrass2", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/grass/topGrass3_g.png", "topGrass3", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/nature/flower1_o.png", "flower1", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/nature/flower1_y.png", "flower2", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/nature/flower3.png", "flower3", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/nature/pebble.png", "pebble", "pixel")

	// Transparency Maps
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "0LN", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/RN.png", "0RN", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/NT.png", "0NT", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/NB.png", "0NB", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/LA.png", "0LA", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/RA.png", "0RA", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/AT.png", "0AT", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/AB.png", "0AB", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/NN.png", "0NN", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/AA.png", "0AA", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/LT.png", "0LT", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/LB.png", "0LB", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/RT.png", "0RT", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/RB.png", "0RB", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/AN.png", "0AN", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/NA.png", "0NA", "pixel")

	Engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "1LN", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/RN.png", "1RN", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/NT.png", "1NT", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/NB.png", "1NB", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/LA.png", "1LA", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/RA.png", "1RA", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/AT.png", "1AT", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/AB.png", "1AB", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/NN.png", "1NN", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/AA.png", "1AA", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/LT.png", "1LT", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/LB.png", "1LB", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/RT.png", "1RT", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/RB.png", "1RB", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/AN.png", "1AN", "pixel")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/NA.png", "1NA", "pixel")

	skyMaterial := Engine.MaterialControl.NewBasicMaterial()
	skyMaterial.DiffuseLevel = 1
	skyMaterial.DiffuseMap = Engine.TextureControl.GetTexture("back")

	dirtMaterial := Engine.MaterialControl.NewBasicMaterial()
	dirtMaterial.DiffuseLevel = 1
	dirtMaterial.DiffuseMap = Engine.TextureControl.GetTexture("dirt")

	stoneMaterial := Engine.MaterialControl.NewBasicMaterial()
	stoneMaterial.DiffuseLevel = 1
	stoneMaterial.DiffuseMap = Engine.TextureControl.GetTexture("stone")

	stoneBrickMaterial := Engine.MaterialControl.NewBasicMaterial()
	stoneBrickMaterial.DiffuseLevel = 1
	stoneBrickMaterial.DiffuseMap = Engine.TextureControl.GetTexture("stoneBrick")

	grassMaterial := Engine.MaterialControl.NewBasicMaterial()
	grassMaterial.DiffuseLevel = 1
	grassMaterial.DiffuseMap = Engine.TextureControl.GetTexture("grass")

	backDirtMaterial := Engine.MaterialControl.NewBasicMaterial()
	backDirtMaterial.DiffuseLevel = 1
	backDirtMaterial.DiffuseMap = Engine.TextureControl.GetTexture("backdirt")

	leavesMaterial := Engine.MaterialControl.NewBasicMaterial()
	leavesMaterial.DiffuseLevel = 1
	leavesMaterial.DiffuseMap = Engine.TextureControl.GetTexture("leaves")

	treeRightRootMaterial := Engine.MaterialControl.NewBasicMaterial()
	treeRightRootMaterial.DiffuseLevel = 1
	treeRightRootMaterial.DiffuseMap = Engine.TextureControl.GetTexture("treeRightRoot")

	treeLeftRootMaterial := Engine.MaterialControl.NewBasicMaterial()
	treeLeftRootMaterial.DiffuseLevel = 1
	treeLeftRootMaterial.DiffuseMap = Engine.TextureControl.GetTexture("treeLeftRoot")

	treeTrunkMaterial := Engine.MaterialControl.NewBasicMaterial()
	treeTrunkMaterial.DiffuseLevel = 1
	treeTrunkMaterial.DiffuseMap = Engine.TextureControl.GetTexture("treeTrunk")

	treeBottomRootMaterial := Engine.MaterialControl.NewBasicMaterial()
	treeBottomRootMaterial.DiffuseLevel = 1
	treeBottomRootMaterial.DiffuseMap = Engine.TextureControl.GetTexture("treeBottomRoot")

	topGrass1Material := Engine.MaterialControl.NewBasicMaterial()
	topGrass1Material.DiffuseLevel = 1
	topGrass1Material.DiffuseMap = Engine.TextureControl.GetTexture("topGrass1")

	topGrass2Material := Engine.MaterialControl.NewBasicMaterial()
	topGrass2Material.DiffuseLevel = 1
	topGrass2Material.DiffuseMap = Engine.TextureControl.GetTexture("topGrass2")

	topGrass3Material := Engine.MaterialControl.NewBasicMaterial()
	topGrass3Material.DiffuseLevel = 1
	topGrass3Material.DiffuseMap = Engine.TextureControl.GetTexture("topGrass3")

	treeBranchR1Material := Engine.MaterialControl.NewBasicMaterial()
	treeBranchR1Material.DiffuseLevel = 1
	treeBranchR1Material.DiffuseMap = Engine.TextureControl.GetTexture("treeBranchR1")

	treeBranchL1Material := Engine.MaterialControl.NewBasicMaterial()
	treeBranchL1Material.DiffuseLevel = 1
	treeBranchL1Material.DiffuseMap = Engine.TextureControl.GetTexture("treeBranchL1")

	flower1Material := Engine.MaterialControl.NewBasicMaterial()
	flower1Material.DiffuseLevel = 1
	flower1Material.DiffuseMap = Engine.TextureControl.GetTexture("flower1")

	flower2Material := Engine.MaterialControl.NewBasicMaterial()
	flower2Material.DiffuseLevel = 1
	flower2Material.DiffuseMap = Engine.TextureControl.GetTexture("flower2")

	flower3Material := Engine.MaterialControl.NewBasicMaterial()
	flower3Material.DiffuseLevel = 1
	flower3Material.DiffuseMap = Engine.TextureControl.GetTexture("flower3")

	pebbleMaterial := Engine.MaterialControl.NewBasicMaterial()
	pebbleMaterial.DiffuseLevel = 1
	pebbleMaterial.DiffuseMap = Engine.TextureControl.GetTexture("pebble")

	torchMaterial := Engine.MaterialControl.NewBasicMaterial()
	torchMaterial.DiffuseLevel = 1
	torchMaterial.DiffuseMap = Engine.TextureControl.GetTexture("torch")

	grasstopMaterial := Engine.MaterialControl.NewBasicMaterial()
	grasstopMaterial.DiffuseLevel = 1
	grasstopMaterial.DiffuseMap = Engine.TextureControl.GetTexture("grasstop")

	BlockMap = make(map[string]*Block)
	BlockMap = map[string]*Block{
		"sky": &Block{
			Material:   skyMaterial,
			LightBlock: 0,
			SaveColor:  [3]int{107, 185, 240},
		},
		"dirt": &Block{
			Material:   dirtMaterial,
			LightBlock: 0.1,
			SaveColor:  [3]int{112, 85, 74},
		},
		"grass": &Block{
			Material:   grassMaterial,
			LightBlock: 0.1,
			SaveColor:  [3]int{115, 173, 87},
		},
		"stone": &Block{
			Material:   stoneMaterial,
			LightBlock: 0.15,
			SaveColor:  [3]int{116, 116, 116},
		},
		"backdirt": &Block{
			Material:   backDirtMaterial,
			LightBlock: 0.035,
			SaveColor:  [3]int{70, 48, 38},
		},
		"leaves": &Block{
			Material:   leavesMaterial,
			LightBlock: 0,
			SaveColor:  [3]int{91, 141, 68},
		},
		"treeRightRoot": &Block{
			Material:   treeRightRootMaterial,
			LightBlock: 0,
			SaveColor:  [3]int{87, 66, 59},
		},
		"treeLeftRoot": &Block{
			Material:   treeLeftRootMaterial,
			LightBlock: 0,
			SaveColor:  [3]int{87, 66, 59},
		},
		"treeTrunk": &Block{
			Material:   treeTrunkMaterial,
			LightBlock: 0,
			SaveColor:  [3]int{87, 66, 59},
		},
		"treeBottomRoot": &Block{
			Material:   treeBottomRootMaterial,
			LightBlock: 0,
			SaveColor:  [3]int{87, 66, 59},
		},
		"topGrass1": &Block{
			Material:   topGrass1Material,
			LightBlock: 0,
			SaveColor:  [3]int{107, 185, 240},
		},
		"topGrass2": &Block{
			Material:   topGrass2Material,
			LightBlock: 0,
			SaveColor:  [3]int{107, 185, 240},
		},
		"topGrass3": &Block{
			Material:   topGrass3Material,
			LightBlock: 0,
			SaveColor:  [3]int{107, 185, 240},
		},
		"treeBranchR1": &Block{
			Material:   treeBranchR1Material,
			LightBlock: 0,
			SaveColor:  [3]int{107, 185, 240},
		},
		"treeBranchL1": &Block{
			Material:   treeBranchL1Material,
			LightBlock: 0,
			SaveColor:  [3]int{107, 185, 240},
		},
		"flower1": &Block{
			Material:   flower1Material,
			LightBlock: 0,
			SaveColor:  [3]int{107, 185, 240},
		},
		"flower2": &Block{
			Material:   flower2Material,
			LightBlock: 0,
			SaveColor:  [3]int{107, 185, 240},
		},
		"flower3": &Block{
			Material:   flower3Material,
			LightBlock: 0,
			SaveColor:  [3]int{107, 185, 240},
		},
		"pebble": &Block{
			Material:   pebbleMaterial,
			LightBlock: 0,
			SaveColor:  [3]int{107, 185, 240},
		},
		"torch": &Block{
			Material:   torchMaterial,
			LightBlock: 0,
			SaveColor:  [3]int{107, 185, 240},
		},
		"stoneBrick": &Block{
			Material:   stoneBrickMaterial,
			LightBlock: 0.15,
			SaveColor:  [3]int{116, 116, 116},
		},
		"grasstop": &Block{
			Material:   grasstopMaterial,
			LightBlock: 0.15,
			SaveColor:  [3]int{116, 116, 116},
		},
	}

	BlockMap["dirt"].CreateOrientations(0)
	BlockMap["grass"].CreateOrientations(1)
	BlockMap["stone"].CreateOrientations(0)
	//BlockMap["stoneBrick"].CreateOrientations(0)
	BlockMap["leaves"].CreateOrientations(0)
	BlockMap["backdirt"].CreateOrientations(1)

	InverseOrientationMap = make(map[string]string)
	for d, o := range OrientationsMap {
		InverseOrientationMap[o] = d
	}
}

func GetBlock(name string) *Block {
	return BlockMap[name]
}

func GetBlockName(id int) string {
	for n, index := range NameToID {
		ind, _ := strconv.Atoi(index)
		if ind == id {
			return n
		}
	}
	return "sky"
}

var NameToID = map[string]string{
	"sky":            "000",
	"dirt":           "001",
	"grass":          "002",
	"stone":          "003",
	"backdirt":       "004",
	"leaves":         "005",
	"treeRightRoot":  "006",
	"treeLeftRoot":   "007",
	"treeTrunk":      "008",
	"treeBottomRoot": "009",
	"topGrass1":      "010",
	"topGrass2":      "011",
	"topGrass3":      "012",
	"treeBranchR1":   "013",
	"treeBranchL1":   "014",
	"flower1":        "015",
	"flower2":        "016",
	"flower3":        "017",
	"pebble":         "018",
	"torch":          "019",
	"stoneBrick":     "020",
	"grasstop":       "021",
}

var IDToName = map[string]string{
	"000": "sky",
	"001": "dirt",
	"002": "grass",
	"003": "stone",
	"004": "backdirt",
	"005": "leaves",
	"006": "treeRightRoot",
	"007": "treeLeftRoot",
	"008": "treeTrunk",
	"009": "treeBottomRoot",
	"010": "topGrass1",
	"011": "topGrass2",
	"012": "topGrass3",
	"013": "treeBranchR1",
	"014": "treeBranchL1",
	"015": "flower1",
	"016": "flower2",
	"017": "flower3",
	"018": "pebble",
	"019": "torch",
	"020": "stoneBrick",
	"021": "grasstop",
}

func GetIDFromName(name string) string {
	return NameToID[name]
}

func GetNameFromID(id string) string {
	return IDToName[id]
}

var OrientationsMap = map[string]string{
	"LN": "00",
	"RN": "01",
	"NT": "02",
	"NB": "03",
	"LA": "04",
	"RA": "05",
	"AT": "06",
	"AB": "07",
	"NN": "08",
	"AA": "09",
	"LT": "10",
	"LB": "11",
	"RT": "12",
	"RB": "13",
	"AN": "14",
	"NA": "15",
}

var InverseOrientationMap map[string]string

func GetOrientationFromID(id string) string {
	for orientation, bid := range OrientationsMap {
		if bid == id {
			return orientation
		}
	}
	return "NN"
}
