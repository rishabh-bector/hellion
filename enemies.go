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
	goblinChild.ScaleX = 300
	goblinChild.ScaleY = 300

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

			Hitbox1: NewHitBox(AABB{
				X:      0,
				Y:      0,
				Width:  75,
				Height: 135,
			}, 5),

			aHitbox: AABB{
				X:      60,
				Y:      45,
				Width:  20,
				Height: 60,
			},
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

// --------------------------------------------------
// ENEMY COMPONENTS
// --------------------------------------------------

type Common struct {
	MonsterChild *child.Child2D

	MonsterMaterial *material.BasicMaterial
	CurrentAnim     string

	Hitbox1 Hitbox
	aHitbox AABB

	VXMult   float32
	VYMult   float32
	GravMult float32

	KnockXMult float32
	KnockYMult float32

	AttackTimeout float64

	// Temp state
	AnimBusy bool
	MoveBusy bool
	Knocking bool
	Hitting  bool

	NumJumps int
}

func (c *Common) Update(player Player) {
	hx := c.MonsterChild.X + (c.MonsterChild.ScaleX / 2) - (c.Hitbox1.DAABB.Width / 2)
	hy := c.MonsterChild.Y + (c.MonsterChild.ScaleY / 2) - (c.Hitbox1.LAABB.Height / 2)

	c.aHitbox.X = hx
	c.aHitbox.Y = hy

	top, left, bottom, right, topleft, topright := CheckWorldCollision(c.Hitbox1, c.MonsterChild.VX, c.MonsterChild.VY, hx, hy)

	dx := c.MonsterChild.X - player.PlayerChild.X
	absdx := math.Abs(float64(dx))

	if absdx < 100 {
		if c.AttackTimeout < 0 {
			c.Attack()
			c.AttackTimeout = 10
		}
	}

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

	if !c.MoveBusy && absdx > 100 {
		c.MonsterChild.VX = (BaseSpeedX - 50) * (dx / float32(math.Abs(float64(dx))))
	}

	if (right || topright) && c.MonsterChild.VX < -50 {
		c.MonsterChild.VX = 0
		c.Jump()
	}

	if (left || topleft) && c.MonsterChild.VX > 50 {
		c.MonsterChild.VX = 0
		c.Jump()
	}

	if top && c.MonsterChild.VY > 10 {
		c.MonsterChild.VY = 0
	}

	/*c.MonsterChild.X = Player1.PlayerChild.X
	c.MonsterChild.Y = Player1.PlayerChild.Y
	c.MonsterChild.VX = 0
	c.MonsterChild.VY = 0*/

	c.UpdateAnimations()
	c.AttackTimeout -= Engine.Renderer.DeltaFrameTime
}

func (c *Common) UpdateAnimations() {
	if c.MonsterChild.VX > 0 {
		c.MonsterMaterial.Flipped = 1
	} else if c.MonsterChild.VX < 0 {
		c.MonsterMaterial.Flipped = 0
	}

	if c.AnimBusy {
		return
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
	if c.NumJumps > 0 {
		c.NumJumps--
		c.MonsterChild.VY = BaseSpeedY * c.VYMult
		c.MonsterMaterial.PlayAnimationOnce("jump")
		c.CurrentAnim = "jump"
	}
}

func (c *Common) Attack() {
	c.AnimBusy = true
	c.MoveBusy = true
	c.MonsterChild.VX = 0

	c.MonsterMaterial.PlayAnimationOnceCallback("attack", c.DoneHitting)
	if c.aHitbox.CheckCollision(Player1.FullBox, 0, 0) {
		c.Hitting = true
		Player1.GettingHit = true
		Player1.Hit()
	}
}

func (c *Common) DoneHitting() {
	c.AnimBusy = false
	c.MoveBusy = false
	if c.Hitting {
		c.Hitting = false
		Player1.GettingHit = false
	}
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
