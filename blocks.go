package main

import "rapidengine"

var BlockMap map[string]*Block

var NameMap map[string]int
var NameList = []string{"sky", "dirt", "grass", "stone"}

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
	Orientations  []rapidengine.Material
}

func (block *Block) GetMaterial() *rapidengine.Material {
	return &block.Material
}

func (block *Block) GetOrientMaterial(direction string) *rapidengine.Material {
	return &(block.Orientations[OrientationsMap[direction]])
}

func (block *Block) CreateOrientations() {
	block.OrientEnabled = true
	for dir := range OrientationsMap {
		newM := block.Material
		newM.AttachTransparency(engine.TextureControl.GetTexture(dir))
		block.Orientations = append(block.Orientations, newM)
	}
}

func LoadBlocks() {
	engine.TextureControl.NewTexture("./assets/blocks/dirt/dirt.png", "dirt")
	engine.TextureControl.NewTexture("./assets/blocks/grass/grass.png", "grass")
	engine.TextureControl.NewTexture("./assets/blocks/stone/stone.png", "stone")

	// Transparency Maps
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "LN")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "RN")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "NT")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "NB")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "LA")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "RA")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "AT")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "AB")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "NN")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "AA")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "LT")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "LB")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "RT")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "RB")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "AN")
	engine.TextureControl.NewTexture("./assets/blocks/transparency/LN.png", "NA")

	skyMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"))
	skyMaterial.BecomeTexture(engine.TextureControl.GetTexture("back"))

	dirtMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"))
	dirtMaterial.BecomeTexture(engine.TextureControl.GetTexture("dirt"))

	stoneMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"))
	stoneMaterial.BecomeTexture(engine.TextureControl.GetTexture("stone"))

	grassMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"))
	grassMaterial.BecomeTexture(engine.TextureControl.GetTexture("grass"))

	BlockMap = make(map[string]*Block)
	BlockMap = map[string]*Block{
		"sky": &Block{
			Material: skyMaterial,
		},
		"dirt": &Block{
			Material:      dirtMaterial,
			OrientEnabled: true,
		},
		"grass": &Block{
			Material: grassMaterial,
		},
		"stone": &Block{
			Material: stoneMaterial,
		},
	}

	BlockMap["dirt"].CreateOrientations()

	NameMap = make(map[string]int)
	NameMap = map[string]int{
		"sky":   0,
		"dirt":  1,
		"grass": 2,
		"stone": 3,
	}
}

func GetBlockIndex(i int) *Block {
	return BlockMap[NameList[i]]
}

func GetBlockName(name string) *Block {
	return BlockMap[name]
}
