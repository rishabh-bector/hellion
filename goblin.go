package main

//"rapidengine/child"
//"math"

import (

	//"math/rand"

	"rapidengine/material"
)

type Goblin struct {
	Health int

	common    *Common
	activator Activator
}

func (g *Goblin) Update() {
	g.common.Update()
	Engine.Renderer.RenderChild(g.common.MonsterChild)
}

func (g *Goblin) Damage(amount float32) {

}

func (g *Goblin) Activator() *Activator {
	return &g.activator
}

func LoadGoblinTextures() {
	Engine.TextureControl.NewTexture("./assets/enemies/doof/idle/1.png", "goblin_i1", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/idle/2.png", "goblin_i2", "pixel")

	Engine.TextureControl.NewTexture("./assets/enemies/doof/attack/1.png", "goblin_a1", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/attack/2.png", "goblin_a2", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/attack/3.png", "goblin_a3", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/attack/4.png", "goblin_a4", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/attack/5.png", "goblin_a5", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/attack/6.png", "goblin_a6", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/attack/7.png", "goblin_a7", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/attack/8.png", "goblin_a8", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/attack/9.png", "goblin_a9", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/attack/10.png", "goblin_a10", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/attack/11.png", "goblin_a11", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/attack/12.png", "goblin_a12", "pixel")

	Engine.TextureControl.NewTexture("./assets/enemies/doof/walk/1.png", "goblin_w1", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/walk/2.png", "goblin_w2", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/walk/3.png", "goblin_w3", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/walk/4.png", "goblin_w4", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/walk/5.png", "goblin_w5", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/walk/6.png", "goblin_w6", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/walk/7.png", "goblin_w7", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/walk/8.png", "goblin_w8", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/walk/9.png", "goblin_w9", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/walk/10.png", "goblin_w10", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/walk/11.png", "goblin_w11", "pixel")
	Engine.TextureControl.NewTexture("./assets/enemies/doof/walk/12.png", "goblin_w12", "pixel")
}

func NewGoblinMaterial() *material.BasicMaterial {
	goblinMaterial := Engine.MaterialControl.NewBasicMaterial()
	goblinMaterial.DiffuseLevel = 1
	goblinMaterial.DiffuseMap = Engine.TextureControl.GetTexture("goblin_w1")

	goblinMaterial.EnableAnimation()

	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_i1"), "idle")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_i2"), "idle")

	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_i1"), "jump")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_i2"), "jump")

	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a1"), "attack")
	goblinMaterial.AddHitFrame(Engine.TextureControl.GetTexture("goblin_a2"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a3"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a4"), "attack")
	goblinMaterial.AddHitFrame(Engine.TextureControl.GetTexture("goblin_a5"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a6"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a7"), "attack")
	goblinMaterial.AddHitFrame(Engine.TextureControl.GetTexture("goblin_a8"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a9"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a10"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a11"), "attack")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_a12"), "attack")

	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_w1"), "walk")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_w2"), "walk")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_w3"), "walk")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_w4"), "walk")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_w5"), "walk")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_w6"), "walk")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_w7"), "walk")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_w8"), "walk")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_w9"), "walk")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_w10"), "walk")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_w11"), "walk")
	goblinMaterial.AddFrame(Engine.TextureControl.GetTexture("goblin_w12"), "walk")

	goblinMaterial.SetAnimationFPS("walk", 20)
	goblinMaterial.SetAnimationFPS("idle", 5)
	goblinMaterial.SetAnimationFPS("jump", 5)
	goblinMaterial.SetAnimationFPS("hit", 5)
	goblinMaterial.SetAnimationFPS("attack", 10)

	goblinMaterial.PlayAnimation("idle")

	return goblinMaterial
}
