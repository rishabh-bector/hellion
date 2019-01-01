package main

import (
	"math"
	"math/rand"
	"rapidengine/child"
	"rapidengine/material"
)

type EnemyManager struct {
	AllEnemies map[int]Enemy
}

func InitializeEnemyManager() *EnemyManager {
	em := EnemyManager{
		AllEnemies: make(map[int]Enemy),
	}

	LoadGoblinTextures()

	return &em
}

func (em *EnemyManager) Update() {
	for _, enemy := range em.AllEnemies {
		if enemy.Activator().IsActive() {
			enemy.Update(Player1)
		}
	}
}

func (em *EnemyManager) NewGoblin(radius float32) {
	mat := NewGoblinMaterial()

	goblinChild := Engine.ChildControl.NewChild2D()
	goblinChild.AttachMaterial(mat)
	goblinChild.ScaleX = 250
	goblinChild.ScaleY = 250

	screenSide := (rand.Intn(2) * 2) - 1

	goblinChild.X = Player1.PlayerChild.X + float32(screenSide)*((float32(ScreenWidth/2)+100)+radius)
	goblinChild.Y = float32(HeightMap[int(goblinChild.X)/BlockSize]*BlockSize) + 50

	var g = Goblin{
		Health: 100,

		common: &Common{
			MonsterChild:    goblinChild,
			MonsterMaterial: mat,

			VXMult:   0.8,
			VYMult:   1,
			GravMult: 1,

			KnockXMult: 1,
			KnockYMult: 0.8,

			NumJumps: 1,
		},

		activator: Activator{},
	}

	g.Activator().Activate()

	em.AllEnemies[len(em.AllEnemies)-1] = &g
}

type Enemy interface {
	Update(player Player)

	Damage(amount float32)

	Activator() *Activator
}

func CheckWorldCollision(x, y float32) (bool, bool, bool, bool, bool, bool) {
	top := false
	left := false
	bottom := false
	right := false

	topleft := false
	topright := false

	px := int((x + BlockSize/2) / BlockSize)
	py := int((y)/BlockSize + 1)

	if WorldMap.GetWorldBlockID(px, py+1) != "00000" {
		top = true
	}
	if WorldMap.GetWorldBlockID(px, py-1) != "00000" {
		bottom = true
	}

	if WorldMap.GetWorldBlockID(px-1, py) != "00000" || WorldMap.GetWorldBlockID(px-1, py+1) != "00000" {
		left = true
	}
	if WorldMap.GetWorldBlockID(px+1, py) != "00000" || WorldMap.GetWorldBlockID(px+1, py+1) != "00000" {
		right = true
	}

	if WorldMap.GetWorldBlockID(px-1, py+1) != "00000" {
		topleft = true
	}
	if WorldMap.GetWorldBlockID(px+1, py+1) != "00000" {
		topright = true
	}

	return top, left, bottom, right, topleft, topright
}

// --------------------------------------------------
// ENEMY COMPONENTS
// --------------------------------------------------

type Common struct {
	MonsterChild *child.Child2D

	MonsterMaterial *material.BasicMaterial
	CurrentAnim     string

	VXMult   float32
	VYMult   float32
	GravMult float32

	KnockXMult float32
	KnockYMult float32
	Knocking   bool

	NumJumps int
}

func (c *Common) Update(player Player) {
	_, left, bottom, right, topleft, topright := CheckWorldCollision(c.MonsterChild.X, c.MonsterChild.Y)

	dx := c.MonsterChild.X - player.PlayerChild.X

	if c.Knocking {
		c.MonsterChild.VX = -1 * (BaseSpeedX * c.KnockXMult) * (dx / float32(math.Abs(float64(dx))))
		c.MonsterChild.VY = (BaseSpeedY * c.KnockYMult)
		c.NumJumps = 0
		c.MonsterMaterial.PlayAnimationOnce("hit")
		c.CurrentAnim = "hit"
		c.Knocking = false
		return
	}

	// Ground
	if bottom {
		c.NumJumps = 1
		c.MonsterChild.VY = 0
	} else {
		c.MonsterChild.VY -= BaseGravity * float32(Engine.Renderer.DeltaFrameTime)
	}

	if bottom {
		c.MonsterChild.VX = (BaseSpeedX - 50) * (dx / float32(math.Abs(float64(dx))))
	}

	if right && c.MonsterChild.VX > 0 {
		if topright {
			c.MonsterChild.VX = 0
		} else {
			c.Jump()
		}
	}

	if left && c.MonsterChild.VX < 0 {
		if topleft {
			c.MonsterChild.VX = 0
		} else {
			c.Jump()
		}
	}

	c.UpdateAnimations()
}

func (c *Common) UpdateAnimations() {
	if c.MonsterChild.VX > 0 {
		c.MonsterMaterial.Flipped = 1
	} else {
		c.MonsterMaterial.Flipped = 0
	}

	if c.MonsterChild.VX > 0 && c.NumJumps > 0 && c.CurrentAnim != "walk" {
		c.MonsterMaterial.PlayAnimation("walk")
		c.CurrentAnim = "walk"
	}
	if c.MonsterChild.VX < 0 && c.NumJumps > 0 && c.CurrentAnim != "walk" {
		c.MonsterMaterial.PlayAnimation("walk")
		c.CurrentAnim = "walk"
	}
	if c.MonsterChild.VX == 0 && c.NumJumps > 0 && c.CurrentAnim != "idle" {
		c.MonsterMaterial.PlayAnimation("idle")
		c.CurrentAnim = "idle"
	}
}

func (c *Common) Jump() {
	c.MonsterChild.VY = BaseSpeedY * c.VYMult
	c.MonsterMaterial.PlayAnimationOnce("jump")
	c.CurrentAnim = "jump"
}

func (c *Common) Knockback(mult float32) {
	c.Knocking = true
}

type Activator struct {
	active bool
}

func (a *Activator) Activate() {
	a.active = true
}

func (a *Activator) Deactivate() {
	a.active = false
}

func (a *Activator) IsActive() bool {
	return a.active
}
