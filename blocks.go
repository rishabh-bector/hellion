package main

import "rapidengine"

var BlockMap map[string]*Block

var NameMap map[string]int
var NameList = []string{"sky", "dirt", "grass", "stone"}

type Block struct {
	Material rapidengine.Material
}

func (block *Block) GetMaterial() *rapidengine.Material {
	return &block.Material
}

func LoadBlocks() {
	engine.TextureControl.NewTexture("./assets/blocks/dirt.png", "dirt")
	engine.TextureControl.NewTexture("./assets/blocks/grass.png", "grass")
	engine.TextureControl.NewTexture("./assets/blocks/stone.png", "stone")

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
			Material: dirtMaterial,
		},
		"grass": &Block{
			Material: grassMaterial,
		},
		"stone": &Block{
			Material: stoneMaterial,
		},
	}

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
