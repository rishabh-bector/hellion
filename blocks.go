package main

import (
	"fmt"
	"rapidengine"
)

var BlockMap map[string]*Block

var NameMap map[string]int
var NameList = []string{
	"sky",
	"dirt",
	"grass",
	"stone",
	"backdirt",
	"leaves", "treeRightRoot", "treeLeftRoot", "treeTrunk", "treeBottomRoot", "topGrass1", "topGrass2", "topGrass3"}

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
	OrientEnabled   bool
	OrientVariation int32
	Orientations    [16]rapidengine.Material
}

func (block *Block) GetMaterial() *rapidengine.Material {
	return &block.Material
}

func (block *Block) GetOrientMaterial(direction string) *rapidengine.Material {
	return &(block.Orientations[OrientationsMap[direction]])
}

func (block *Block) CreateOrientations(orientVariation int32) {
	block.OrientEnabled = true
	block.OrientVariation = orientVariation
	for dir, ind := range OrientationsMap {
		newM := block.Material
		newM.AttachTransparency(engine.TextureControl.GetTexture(fmt.Sprintf("%v%s", orientVariation, dir)))
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
	engine.TextureControl.NewTexture("./assets/blocks/tree/treeTrunk3.png", "treeTrunk")
	engine.TextureControl.NewTexture("./assets/blocks/tree/treeLeftRoot.png", "treeLeftRoot")
	engine.TextureControl.NewTexture("./assets/blocks/tree/treeRightRoot.png", "treeRightRoot")
	engine.TextureControl.NewTexture("./assets/blocks/tree/treeBottomRoot.png", "treeBottomRoot")
	engine.TextureControl.NewTexture("./assets/blocks/tree/leaves.png", "leaves")

	// Flora
	engine.TextureControl.NewTexture("./assets/blocks/grass/grass.png", "topGrass1") //fix
	engine.TextureControl.NewTexture("./assets/blocks/grass/topGrass2.png", "topGrass2")
	engine.TextureControl.NewTexture("./assets/blocks/grass/topGrass3.png", "topGrass3")

	// Transparency Maps
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "0LN")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/RN.png", "0RN")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/NT.png", "0NT")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/NB.png", "0NB")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LA.png", "0LA")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/RA.png", "0RA")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/AT.png", "0AT")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/AB.png", "0AB")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/NN.png", "0NN")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/AA.png", "0AA")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LT.png", "0LT")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LB.png", "0LB")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/RT.png", "0RT")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/RB.png", "0RB")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/AN.png", "0AN")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/NA.png", "0NA")

	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "1LN")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/RN.png", "1RN")
	engine.TextureControl.NewTexture("./assets/blocks/grass/transparency/NT.png", "1NT")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/NB.png", "1NB")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LA.png", "1LA")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/RA.png", "1RA")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/AT.png", "1AT")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/AB.png", "1AB")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/NN.png", "1NN")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/AA.png", "1AA")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LT.png", "1LT")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LB.png", "1LB")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/RT.png", "1RT")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/RB.png", "1RB")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/AN.png", "1AN")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/NA.png", "1NA")

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

	topGrass1Material := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	topGrass1Material.BecomeTexture(engine.TextureControl.GetTexture("topGrass1"))

	topGrass2Material := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	topGrass2Material.BecomeTexture(engine.TextureControl.GetTexture("topGrass2"))

	topGrass3Material := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	topGrass3Material.BecomeTexture(engine.TextureControl.GetTexture("topGrass3"))

	BlockMap = make(map[string]*Block)
	BlockMap = map[string]*Block{
		"sky": &Block{
			Material: skyMaterial,
		},
		"dirt": &Block{
			Material: dirtMaterial,
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
		"topGrass1": &Block{
			Material: topGrass1Material,
		},
		"topGrass2": &Block{
			Material: topGrass2Material,
		},
		"topGrass3": &Block{
			Material: topGrass3Material,
		},
	}

	BlockMap["dirt"].CreateOrientations(0)
	BlockMap["grass"].CreateOrientations(1)
	BlockMap["stone"].CreateOrientations(0)
	BlockMap["leaves"].CreateOrientations(0)
	BlockMap["backdirt"].CreateOrientations(0) // why doesn't this work?

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
		"topGrass1":      10,
		"topGrass2":      11,
		"topGrass3":      12,
	}
}

func GetBlockIndex(i int) *Block {
	return BlockMap[NameList[i]]
}

func GetBlockName(name string) *Block {
	return BlockMap[name]
}
