package main

import (
	"math"
	"math/rand"
	"rapidengine/child"
	"rapidengine/material"
	"rapidengine/ui"
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
	for i, enemy := range em.AllEnemies {
		if enemy.Activator().IsActive() {
			enemy.Update()
		}
		if enemy.GetCommon().Health <= 0 {
			enemy.GetCommon().Kill()
			enemy.Activator().Deactivate()
			delete(em.AllEnemies, i)
		}
	}
}

func (em *EnemyManager) CheckPlayerCollision() Enemy {
	for _, enemy := range em.AllEnemies {
		pdist := Distance(
			Player1.CenterX, Player1.CenterY,
			enemy.GetChild().X, enemy.GetChild().Y,
		)

		if pdist > 200 {
			continue
		}

		if Player1.CheckEnemyCollision(enemy) {
			return enemy
		}
	}

	return nil
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
		common: &Common{

			Health:    100,
			MaxHealth: 100,

			MonsterChild:    goblinChild,
			MonsterMaterial: mat,

			VXMult:   (rand.Float32()*2-1.0)*0.2 + 0.8,
			VYMult:   (rand.Float32()*2-1.0)*0.2 + 1.0,
			GravMult: 1,

			NumJumps: 1,

			Hitbox1: NewHitBox(AABB{
				X:      0,
				Y:      0,
				Width:  60,
				Height: 135,
			}, 5),

			aHitbox: AABB{
				OffX:   0,
				OffY:   45,
				Width:  55,
				Height: 60,
			},

			State: "normal",
		},

		activator: Activator{},
	}

	g.common.HealthBar = Engine.UIControl.NewProgressBar()
	g.common.HealthBar.SetDimensions(50, 10)
	g.common.HealthBar.BackChild.Static = true
	g.common.HealthBar.BarChild.Static = true
	g.common.HOffsetY = -30

	g.Activator().Activate()

	V.AddBox(&g.common.Hitbox1)
	V.AddAABB(&g.common.aHitbox)

	em.AllEnemies[len(em.AllEnemies)-1] = &g
}

type Enemy interface {
	Update()

	Damage(amount float32)

	GetChild() *child.Child2D
	GetCommon() *Common

	Activator() *Activator
}

// --------------------------------------------------
// ENEMY COMPONENTS
// --------------------------------------------------

type Common struct {
	// Engine components
	MonsterChild    *child.Child2D
	MonsterMaterial *material.BasicMaterial

	// Monster Data
	Damage    float32
	Health    float32
	MaxHealth float32

	// Hitboxes
	Hitbox1 Hitbox
	aHitbox AABB

	// Movement
	VXMult   float32
	VYMult   float32
	GravMult float32

	// Temp state
	State         string
	CurrentAnim   string
	AttackTimeout float64

	HealthBar *ui.ProgressBar
	HOffsetX  float32
	HOffsetY  float32

	NumJumps int
}

func (c *Common) Update() {
	c.UpdateMovement()
	c.UpdateAttacks()
	c.UpdateAnimations()

	c.HealthBar.SetPercentage(c.Health / c.MaxHealth * 100)
	c.HealthBar.Update(nil)

	Engine.Renderer.RenderChild(c.HealthBar.BackChild)
	Engine.Renderer.RenderChild(c.HealthBar.BarChild)
}

func (c *Common) UpdateMovement() {
	// Update position
	c.MonsterChild.X += c.MonsterChild.VX * -float32(Engine.Renderer.DeltaFrameTime)
	c.MonsterChild.Y += c.MonsterChild.VY * float32(Engine.Renderer.DeltaFrameTime)

	// Update collision data
	hx := c.MonsterChild.X + (c.MonsterChild.ScaleX / 2) - (c.Hitbox1.DAABB.Width / 2)
	hy := c.MonsterChild.Y + (c.MonsterChild.ScaleY / 2) - (c.Hitbox1.LAABB.Height / 2)
	c.Hitbox1.X = hx
	c.Hitbox1.Y = hy
	top, left, bottom, right, topleft, topright := CheckWorldCollision(c.Hitbox1, c.MonsterChild.VX, c.MonsterChild.VY, hx, hy)

	camX, camY, _ := Engine.Renderer.MainCamera.GetPosition()
	c.HealthBar.SetPosition(hx+c.HOffsetX-camX+(float32(Engine.Config.ScreenWidth)/2), hy+c.HOffsetY-camY+(float32(Engine.Config.ScreenHeight)/2))

	// Attack hitboxes
	flip := c.MonsterMaterial.Flipped
	if flip == 0 {
		flip = 1
	} else {
		flip = 0
	}
	c.aHitbox.X = (hx + (c.Hitbox1.DAABB.Width / 2)) + (c.aHitbox.OffX+c.aHitbox.Width)*float32(flip-1)
	c.aHitbox.Y = hy + c.aHitbox.OffY

	// Distance from player
	dx := c.MonsterChild.X - Player1.PlayerChild.X
	absdx := math.Abs(float64(dx))

	if c.State == "normal" && absdx > 100 {
		c.MonsterChild.VX = (BaseSpeedX - 50) * (dx / float32(math.Abs(float64(dx)))) * c.VXMult
	}

	if c.State == "attacking" {
		c.MonsterChild.VX = 0
		return
	}

	if c.State == "damaged" {

	}

	// World collision
	if bottom {
		c.NumJumps = 1
		c.MonsterChild.VY = 0
	} else {
		c.MonsterChild.VY -= BaseGravity * float32(Engine.Renderer.DeltaFrameTime) * c.VYMult
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
}

func (c *Common) UpdateAttacks() {
	dx := c.MonsterChild.X - Player1.PlayerChild.X
	absdx := math.Abs(float64(dx))

	if absdx < 100 && c.State == "normal" {
		if c.AttackTimeout < 0 {
			c.Attack()
			c.AttackTimeout = 10
		}
	}

	c.AttackTimeout -= Engine.Renderer.DeltaFrameTime
}

func (c *Common) UpdateAnimations() {
	if c.MonsterChild.VX > 0 {
		c.MonsterMaterial.Flipped = 1
	} else if c.MonsterChild.VX < 0 {
		c.MonsterMaterial.Flipped = 0
	}

	if c.State == "attacking" {
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
	c.State = "attacking"
	c.CurrentAnim = "attacking"

	c.MonsterMaterial.PlayAnimationOnceCallback("attack", c.DoneHitting, c.AttackHitFrame)
}

func (c *Common) AttackHitFrame() {
	if c.CheckPlayerCollision() {
		Player1.Hit(25)
	}
}

func (c *Common) Kill() {

}

func (c *Common) DoneHitting() {
	c.State = "normal"
}

func (c *Common) CheckPlayerCollision() bool {
	if c.aHitbox.CheckCollision(Player1.FullBox, 0, 0) {
		return true
	}
	return false
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

func Distance(x1, y1, x2, y2 float32) float32 {
	return float32(math.Sqrt(
		float64((x2-x1)*(x2-x1) + (y2-y2)*(y2-y1)),
	))
}
