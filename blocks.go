package main

import (
	"fmt"
	"rapidengine/material"
	"strconv"
)

var BlockMap map[string]*Block

type Block struct {
	Material material.Material

	// Block orientations (16 possible):
	// First character - Left/Right (L/R)
	// Second character - Top/Bottom (T/B)
	// Either can have A (all) or N (none)
	OrientEnabled   bool
	OrientVariation int32
	Orientations    [16]material.Material

	LightBlock float32
}

func (block *Block) GetMaterial(direction string) *material.Material {
	if block.OrientEnabled {
		i, _ := strconv.Atoi(OrientationsMap[direction])
		return &(block.Orientations[i])
	}
	return &block.Material
}

func (block *Block) CreateOrientations(orientVariation int32) {
	block.OrientEnabled = true
	block.OrientVariation = orientVariation
	for dir, _ := range OrientationsMap {
		newM := block.Material
		newM.AttachTransparency(Engine.TextureControl.GetTexture(fmt.Sprintf("%v%s", orientVariation, dir)))

		i, _ := strconv.Atoi(OrientationsMap[dir])
		block.Orientations[i] = newM
	}
}

func loadBlocks() {
	// Main Blocks
	Engine.TextureControl.NewTexture("./assets/blocks/dirt/dirt.png", "dirt")
	Engine.TextureControl.NewTexture("./assets/blocks/grass/grass2.png", "grass")
	Engine.TextureControl.NewTexture("./assets/blocks/stone/stone.png", "stone")

	Engine.TextureControl.NewTexture("./assets/blocks/torch.png", "torch")

	// Back-Blocks
	Engine.TextureControl.NewTexture("./assets/blocks/backblocks/backdirt.png", "backdirt")

	// Tree
	Engine.TextureControl.NewTexture("./assets/blocks/tree/treeTrunk3.png", "treeTrunk")
	Engine.TextureControl.NewTexture("./assets/blocks/tree/treeLeftRoot.png", "treeLeftRoot")
	Engine.TextureControl.NewTexture("./assets/blocks/tree/treeRightRoot.png", "treeRightRoot")
	Engine.TextureControl.NewTexture("./assets/blocks/tree/treeBottomRoot.png", "treeBottomRoot")
	Engine.TextureControl.NewTexture("./assets/blocks/tree/leaves.png", "leaves")
	Engine.TextureControl.NewTexture("./assets/blocks/tree/treeBranchR1.png", "treeBranchR1")
	Engine.TextureControl.NewTexture("./assets/blocks/tree/treeBranchL1.png", "treeBranchL1")

	// Flora
	Engine.TextureControl.NewTexture("./assets/blocks/grass/topGrass1.png", "topGrass1")
	Engine.TextureControl.NewTexture("./assets/blocks/grass/topGrass2.png", "topGrass2")
	Engine.TextureControl.NewTexture("./assets/blocks/grass/topGrass3.png", "topGrass3")
	Engine.TextureControl.NewTexture("./assets/blocks/nature/flower1.png", "flower1")
	Engine.TextureControl.NewTexture("./assets/blocks/nature/flower2.png", "flower2")
	Engine.TextureControl.NewTexture("./assets/blocks/nature/flower3.png", "flower3")
	Engine.TextureControl.NewTexture("./assets/blocks/nature/pebble.png", "pebble")

	// Transparency Maps
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "0LN")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/RN.png", "0RN")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/NT.png", "0NT")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/NB.png", "0NB")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/LA.png", "0LA")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/RA.png", "0RA")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/AT.png", "0AT")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/AB.png", "0AB")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/NN.png", "0NN")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/AA.png", "0AA")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/LT.png", "0LT")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/LB.png", "0LB")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/RT.png", "0RT")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/RB.png", "0RB")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/AN.png", "0AN")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/NA.png", "0NA")

	Engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "1LN")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/RN.png", "1RN")
	Engine.TextureControl.NewTexture("./assets/blocks/grass/transparency/NT.png", "1NT")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/NB.png", "1NB")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/LA.png", "1LA")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/RA.png", "1RA")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/AT.png", "1AT")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/AB.png", "1AB")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/NN.png", "1NN")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/AA.png", "1AA")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/LT.png", "1LT")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/LB.png", "1LB")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/RT.png", "1RT")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/RB.png", "1RB")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/AN.png", "1AN")
	Engine.TextureControl.NewTexture("./assets/blocks/transparency/NA.png", "1NA")

	skyMaterial := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	skyMaterial.BecomeTexture(Engine.TextureControl.GetTexture("back"))

	dirtMaterial := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	dirtMaterial.BecomeTexture(Engine.TextureControl.GetTexture("dirt"))

	stoneMaterial := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	stoneMaterial.BecomeTexture(Engine.TextureControl.GetTexture("stone"))

	grassMaterial := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	grassMaterial.BecomeTexture(Engine.TextureControl.GetTexture("grass"))

	backDirtMaterial := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	backDirtMaterial.BecomeTexture(Engine.TextureControl.GetTexture("backdirt"))

	leavesMaterial := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	leavesMaterial.BecomeTexture(Engine.TextureControl.GetTexture("leaves"))

	treeRightRootMaterial := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	treeRightRootMaterial.BecomeTexture(Engine.TextureControl.GetTexture("treeRightRoot"))

	treeLeftRootMaterial := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	treeLeftRootMaterial.BecomeTexture(Engine.TextureControl.GetTexture("treeLeftRoot"))

	treeTrunkMaterial := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	treeTrunkMaterial.BecomeTexture(Engine.TextureControl.GetTexture("treeTrunk"))

	treeBottomRootMaterial := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	treeBottomRootMaterial.BecomeTexture(Engine.TextureControl.GetTexture("treeBottomRoot"))

	topGrass1Material := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	topGrass1Material.BecomeTexture(Engine.TextureControl.GetTexture("topGrass1"))

	topGrass2Material := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	topGrass2Material.BecomeTexture(Engine.TextureControl.GetTexture("topGrass2"))

	topGrass3Material := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	topGrass3Material.BecomeTexture(Engine.TextureControl.GetTexture("topGrass3"))

	treeBranchR1Material := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	treeBranchR1Material.BecomeTexture(Engine.TextureControl.GetTexture("treeBranchR1"))

	treeBranchL1Material := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	treeBranchL1Material.BecomeTexture(Engine.TextureControl.GetTexture("treeBranchL1"))

	flower1Material := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	flower1Material.BecomeTexture(Engine.TextureControl.GetTexture("flower1"))

	flower2Material := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	flower2Material.BecomeTexture(Engine.TextureControl.GetTexture("flower2"))

	flower3Material := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	flower3Material.BecomeTexture(Engine.TextureControl.GetTexture("flower3"))

	pebbleMaterial := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	pebbleMaterial.BecomeTexture(Engine.TextureControl.GetTexture("pebble"))

	torchMaterial := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	torchMaterial.BecomeTexture(Engine.TextureControl.GetTexture("torch"))

	BlockMap = make(map[string]*Block)
	BlockMap = map[string]*Block{
		"sky": &Block{
			Material:   skyMaterial,
			LightBlock: 0,
		},
		"dirt": &Block{
			Material:   dirtMaterial,
			LightBlock: 0.1,
		},
		"grass": &Block{
			Material:   grassMaterial,
			LightBlock: 0.1,
		},
		"stone": &Block{
			Material:   stoneMaterial,
			LightBlock: 0.15,
		},
		"backdirt": &Block{
			Material:   backDirtMaterial,
			LightBlock: 0.05,
		},
		"leaves": &Block{
			Material:   leavesMaterial,
			LightBlock: 0,
		},
		"treeRightRoot": &Block{
			Material:   treeRightRootMaterial,
			LightBlock: 0,
		},
		"treeLeftRoot": &Block{
			Material:   treeLeftRootMaterial,
			LightBlock: 0,
		},
		"treeTrunk": &Block{
			Material:   treeTrunkMaterial,
			LightBlock: 0,
		},
		"treeBottomRoot": &Block{
			Material:   treeBottomRootMaterial,
			LightBlock: 0,
		},
		"topGrass1": &Block{
			Material:   topGrass1Material,
			LightBlock: 0,
		},
		"topGrass2": &Block{
			Material:   topGrass2Material,
			LightBlock: 0,
		},
		"topGrass3": &Block{
			Material:   topGrass3Material,
			LightBlock: 0,
		},
		"treeBranchR1": &Block{
			Material:   treeBranchR1Material,
			LightBlock: 0,
		},
		"treeBranchL1": &Block{
			Material:   treeBranchL1Material,
			LightBlock: 0,
		},
		"flower1": &Block{
			Material:   flower1Material,
			LightBlock: 0,
		},
		"flower2": &Block{
			Material:   flower2Material,
			LightBlock: 0,
		},
		"flower3": &Block{
			Material:   flower3Material,
			LightBlock: 0,
		},
		"pebble": &Block{
			Material:   pebbleMaterial,
			LightBlock: 0,
		},
		"torch": &Block{
			Material:   torchMaterial,
			LightBlock: 0,
		},
	}

	BlockMap["dirt"].CreateOrientations(0)
	BlockMap["grass"].CreateOrientations(1)
	BlockMap["stone"].CreateOrientations(0)
	BlockMap["leaves"].CreateOrientations(0)
	BlockMap["backdirt"].CreateOrientations(0)
}

func GetBlock(name string) *Block {
	return BlockMap[name]
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

func GetOrientationFromID(id string) string {
	for orientation, bid := range OrientationsMap {
		if bid == id {
			return orientation
		}
	}
	return ""
}
