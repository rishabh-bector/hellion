package main

import (
	"fmt"
	"rapidengine/material"
)

var BlockMap map[string]*Block

var NameMap map[string]int
var NameList = []string{
	"sky",
	"dirt",
	"grass",
	"stone",
	"backdirt",
	"leaves", "treeRightRoot", "treeLeftRoot", "treeTrunk", "treeBottomRoot",
	"topGrass1", "topGrass2", "topGrass3",
	"treeBranchR1", "treeBranchL1",
	"flower1", "flower2", "flower3", "pebble",
	"torch",
}

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

func (block *Block) GetMaterial() *material.Material {
	return &block.Material
}

func (block *Block) GetOrientMaterial(direction string) *material.Material {
	return &(block.Orientations[OrientationsMap[direction]])
}

func (block *Block) CreateOrientations(orientVariation int32) {
	block.OrientEnabled = true
	block.OrientVariation = orientVariation
	for dir, ind := range OrientationsMap {
		newM := block.Material
		newM.AttachTransparency(Engine.TextureControl.GetTexture(fmt.Sprintf("%v%s", orientVariation, dir)))
		block.Orientations[ind] = newM
	}
}

func loadBlocks() {
	// Main Blocks
	Engine.TextureControl.NewTexture("./assets/blocks/dirt/dirt.png", "dirt")
	Engine.TextureControl.NewTexture("./assets/blocks/grass/grass.png", "grass")
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
		"treeBranchR1":   13,
		"treeBranchL1":   14,
		"flower1":        15,
		"flower2":        16,
		"flower3":        17,
		"pebble":         18,
		"torch":          19,
	}
}

func GetBlockIndex(i int) *Block {
	return BlockMap[NameList[i]]
}

func GetBlockName(name string) *Block {
	return BlockMap[name]
}
