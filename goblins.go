package main

//"rapidengine/child"
//"math"

import (
	"rapidengine/child"
	//"math/rand"
	"rapidengine/material"
	"fmt"
)

type Goblin struct {
	MonsterChild *child.Child2D
	Health int
	CurrentAnim string
	MonsterMaterial *material.BasicMaterial
}

type Monster interface {
	Move(direction string)

	Jump()

	Attack(direction string)
}

func (g *Goblin) Move(direction string) {
	if direction == "left" && g.MonsterChild.VX > 0 {
		g.MonsterChild.VX *= -1
	}
	if direction == "right" && g.MonsterChild.VX < 0 {
		g.MonsterChild.VX *= -1
	}
}

func (g *Goblin) Jump() {
	g.MonsterChild.VY = 10
}

func (g *Goblin) Attack(direction string) {
	if direction == "left" {
		g.MonsterChild.VX *= 2
		//play attack animation
	}
	if direction == "right" {
		g.MonsterChild.VX *= 2
		//play attack animation
	}
}

func NewGoblin() {
	Engine.TextureControl.NewTexture("assets/player/idle/f1.png", "default", "pixel")
	goblinMaterial := Engine.MaterialControl.NewBasicMaterial()
	goblinMaterial.DiffuseLevel = 1
	goblinMaterial.DiffuseMap = Engine.TextureControl.GetTexture("default")

	monsterChild := Engine.ChildControl.NewChild2D()
	monsterChild.AttachMaterial(goblinMaterial)
	monsterChild.ScaleX = 600
	monsterChild.ScaleY = 1200
	monsterChild.X = Player1.PlayerChild.X //+ float32(rand.Intn(10) + 10)
	monsterChild.Y = float32(HeightMap[int32(monsterChild.X)])

	var g = Goblin{
		MonsterChild: monsterChild,
		MonsterMaterial: goblinMaterial,
		Health: 100,
	}
	EnemyChildList = append(EnemyChildList, g)
	fmt.Print(g.MonsterChild.X)
	fmt.Print(Player1.PlayerChild.X)
}