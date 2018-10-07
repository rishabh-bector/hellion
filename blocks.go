package main

import "rapidengine"

var BlockMap map[string]*Block

var NameMap map[string]int
var NameList = []string{"sky", "dirt", "grass", "stone", "backdirt", "leaves", "treeRightRoot", "treeLeftRoot", "treeTrunk", "treeBottomRoot"}

var OrientationsMap = map[string]int{
	"LN": 0,
	"RN": 1,
	"NT": 2,
	"NB": 3,
	"LA": 4,
	"RA": 5,
	"AT": 6,
	"AB": 7,
	"NN": 8,
	"AA": 9,
	"LT": 10,
	"LB": 11,
	"RT": 12,
	"RB": 13,
	"AN": 14,
	"NA": 15,
}

type Block struct {
	Material rapidengine.Material

	// Block orientations (16 possible):
	// First character - Left/Right (L/R)
	// Second character - Top/Bottom (T/B)
	// Either can have A (all) or N (none)
	OrientEnabled bool
	Orientations  [16]rapidengine.Material
}

func (block *Block) GetMaterial() *rapidengine.Material {
	return &block.Material
}

func (block *Block) GetOrientMaterial(direction string) *rapidengine.Material {
	return &(block.Orientations[OrientationsMap[direction]])
}

func (block *Block) CreateOrientations() {
	block.OrientEnabled = true
	for dir, ind := range OrientationsMap {
		newM := block.Material
		newM.AttachTransparency(engine.TextureControl.GetTexture(dir))
		block.Orientations[ind] = newM
	}
}

func LoadBlocks() {
	// Main Blocks
	engine.TextureControl.NewTexture("./assets/blocks/dirt/dirt.png", "dirt")
	engine.TextureControl.NewTexture("./assets/blocks/grass/grass.png", "grass")
	engine.TextureControl.NewTexture("./assets/blocks/stone/stone.png", "stone")

	// Back-Blocks
	engine.TextureControl.NewTexture("./assets/blocks/backblocks/backdirt.png", "backdirt")

	// Tree
	engine.TextureControl.NewTexture("./assets/blocks/tree/treeTrunk.png", "treeTrunk")
	engine.TextureControl.NewTexture("./assets/blocks/tree/treeLeftRoot.png", "treeLeftRoot")
	engine.TextureControl.NewTexture("./assets/blocks/tree/treeRightRoot.png", "treeRightRoot")
	engine.TextureControl.NewTexture("./assets/blocks/tree/treeBottomRoot.png", "treeBottomRoot")
	engine.TextureControl.NewTexture("./assets/blocks/tree/leaves.png", "leaves")

	// Transparency Maps
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "LN")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/RN.png", "RN")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/NT.png", "NT")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/NB.png", "NB")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LA.png", "LA")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/RA.png", "RA")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/AT.png", "AT")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/AB.png", "AB")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/NN.png", "NN")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/AA.png", "AA")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LT.png", "LT")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LB.png", "LB")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/RT.png", "RT")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/RB.png", "RB")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/AN.png", "AN")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/NA.png", "NA")

	skyMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	skyMaterial.BecomeTexture(engine.TextureControl.GetTexture("back"))

	dirtMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	dirtMaterial.BecomeTexture(engine.TextureControl.GetTexture("dirt"))

	stoneMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	stoneMaterial.BecomeTexture(engine.TextureControl.GetTexture("stone"))

	grassMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	grassMaterial.BecomeTexture(engine.TextureControl.GetTexture("grass"))

	backDirtMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	backDirtMaterial.BecomeTexture(engine.TextureControl.GetTexture("backdirt"))

	leavesMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	leavesMaterial.BecomeTexture(engine.TextureControl.GetTexture("leaves"))

	treeRightRootMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	treeRightRootMaterial.BecomeTexture(engine.TextureControl.GetTexture("treeRightRoot"))

	treeLeftRootMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	treeLeftRootMaterial.BecomeTexture(engine.TextureControl.GetTexture("treeLeftRoot"))

	treeTrunkMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	treeTrunkMaterial.BecomeTexture(engine.TextureControl.GetTexture("treeTrunk"))

	treeBottomRootMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	treeBottomRootMaterial.BecomeTexture(engine.TextureControl.GetTexture("treeBottomRoot"))

	BlockMap = make(map[string]*Block)
	BlockMap = map[string]*Block{
		"sky": &Block{
			Material: skyMaterial,
		},
		"dirt": &Block{
			Material:      dirtMaterial,
			OrientEnabled: true, // Do we still need this?
		},
		"grass": &Block{
			Material: grassMaterial,
		},
		"stone": &Block{
			Material: stoneMaterial,
		},
		"backdirt": &Block{
			Material: backDirtMaterial,
		},
		"leaves": &Block{
			Material: leavesMaterial,
		},
		"treeRightRoot": &Block{
			Material: treeRightRootMaterial,
		},
		"treeLeftRoot": &Block{
			Material: treeLeftRootMaterial,
		},
		"treeTrunk": &Block{
			Material: treeTrunkMaterial,
		},
		"treeBottomRoot": &Block{
			Material: treeBottomRootMaterial,
		},
	}

	BlockMap["dirt"].CreateOrientations()
	BlockMap["grass"].CreateOrientations()
	BlockMap["stone"].CreateOrientations()
	BlockMap["leaves"].CreateOrientations()
	BlockMap["backdirt"].CreateOrientations() // why doesn't this work?

	NameMap = make(map[string]int)
	NameMap = map[string]int{
		"sky":            0,
		"dirt":           1,
		"grass":          2,
		"stone":          3,
		"backdirt":       4,
		"leaves":         5,
		"treeRightRoot":  6,
		"treeLeftRoot":   7,
		"treeTrunk":      8,
		"treeBottomRoot": 9,
	}
}

func GetBlockIndex(i int) *Block {
	return BlockMap[NameList[i]]
}

func GetBlockName(name string) *Block {
	return BlockMap[name]
}
