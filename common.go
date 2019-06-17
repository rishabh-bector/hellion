package main

import (
	"math"
	"math/rand"
	"rapidengine/child"
	"rapidengine/material"
	"rapidengine/ui"
)

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
	Dead      bool

	// Hitboxes
	Hitbox1 Hitbox
	aHitbox AABB

	// Movement
	VXMult    float32
	VYMult    float32
	GravMult  float32
	Direction int
	TargetX   float32
	TargetY   float32

	// Temp state
	State         string
	CurrentAnim   string
	AttackTimeout float64

	// Health bar
	HealthBar *ui.ProgressBar
	HOffsetX  float32
	HOffsetY  float32

	NumJumps int
}

func (c *Common) Update() {
	c.UpdateState()
	c.UpdateMovement()
	c.UpdateAnimations()

	c.HealthBar.SetPercentage(c.Health / c.MaxHealth * 100)
	c.HealthBar.Update(nil)

	Engine.Renderer.RenderChild(c.HealthBar.BackChild)
	Engine.Renderer.RenderChild(c.HealthBar.BarChild)
}

// UpdateState handles state changes
func (c *Common) UpdateState() {
	// State updates
	if c.State == "idleing" {

	}
	if c.State == "chasing" {

	}
	if c.State == "attacking" {

	}
	if c.State == "retreating" {

	}

	// State changes
	dist := Distance(c.MonsterChild.X, c.MonsterChild.Y, Player1.PlayerChild.X, Player1.PlayerChild.Y)

	if dist < 100 && c.State == "idleing" {
		c.State = "chasing"
	}
	if dist > 500 {
		c.State = "idleing"
	}
}

// UpdateMovement handles world collision detection
// and moving the mob to a target point, as well as
// the health bar
func (c *Common) UpdateMovement() {
	// Update position
	c.MonsterChild.X += c.MonsterChild.VX * -float32(Engine.Renderer.DeltaFrameTime)
	c.MonsterChild.Y += c.MonsterChild.VY * float32(Engine.Renderer.DeltaFrameTime)

	// Update collision data
	hx := c.MonsterChild.X + (c.MonsterChild.ScaleX / 2) - (c.Hitbox1.DAABB.Width / 2) + c.Hitbox1.OffX
	hy := c.MonsterChild.Y + (c.MonsterChild.ScaleY / 2) - (c.Hitbox1.LAABB.Height / 2) + c.Hitbox1.OffY
	c.Hitbox1.X = hx
	c.Hitbox1.Y = hy
	top, left, bottom, right, topleft, topright := CheckWorldCollision(c.Hitbox1, c.MonsterChild.VX, c.MonsterChild.VY, hx, hy)

	// Update health bar
	camX, camY, _ := Engine.Renderer.MainCamera.GetPosition()
	c.HealthBar.SetPosition(hx+c.HOffsetX-camX+(float32(Engine.Config.ScreenWidth)/2), hy+c.HOffsetY-camY+(float32(Engine.Config.ScreenHeight)/2))

	// Update attack hitboxes
	flip := c.MonsterMaterial.Flipped
	if flip == 0 {
		flip = 1
	} else {
		flip = 0
	}
	c.aHitbox.X = (hx + (c.Hitbox1.DAABB.Width / 2)) + (c.aHitbox.OffX+c.aHitbox.Width)*float32(flip-1)
	c.aHitbox.Y = hy + c.aHitbox.OffY

	// Move toward target
	dx := c.TargetX - c.MonsterChild.X
	//dy := c.TargetY - c.MonsterChild.Y
	if dx > 25 || dx < -25 {
		if dx > 25 {
			c.Direction = 1
		} else {
			c.Direction = -1
		}
		c.MonsterChild.VX = float32(c.Direction) * BaseSpeedX * c.VXMult
	}

	// World collision
	if bottom {
		c.NumJumps = 1
		c.MonsterChild.VY = 0
	} else {
		c.MonsterChild.VY -= BaseGravity * float32(Engine.Renderer.DeltaFrameTime) * c.VYMult
	}
	if top && c.MonsterChild.VY > 10 {
		c.MonsterChild.VY = 0
	}

	// Auto jumping
	if (right || topright) && c.MonsterChild.VX < -50 {
		c.MonsterChild.VX = 0
		c.Jump()
	}
	if (left || topleft) && c.MonsterChild.VX > 50 {
		c.MonsterChild.VX = 0
		c.Jump()
	}

	// Testing movement
	/*if Inputs.Keys["up"] && c.NumJumps > 0 {
		c.Jump()
	}
	if Inputs.Keys["left"] {
		c.MonsterChild.VX = 50
	} else if Inputs.Keys["right"] {
		c.MonsterChild.VX = -50
	} else {
		c.MonsterChild.VX = 0
	}*/

	if Inputs.Keys["i"] {
		c.Hitbox1.OffY += 1
		println(c.Hitbox1.OffX, c.Hitbox1.OffY)
	}
	if Inputs.Keys["k"] {
		c.Hitbox1.OffY -= 1
		println(c.Hitbox1.OffX, c.Hitbox1.OffY)
	}
	if Inputs.Keys["j"] {
		c.Hitbox1.OffX += 1
		println(c.Hitbox1.OffX, c.Hitbox1.OffY)
	}
	if Inputs.Keys["l"] {
		c.Hitbox1.OffX -= 1
		println(c.Hitbox1.OffX, c.Hitbox1.OffY)
	}
}

func (c *Common) UpdateAttacking() {
	dx := c.MonsterChild.X - Player1.PlayerChild.X
	absdx := math.Abs(float64(dx))

	if absdx < 100 && c.State == "normal" {
		if c.AttackTimeout < 0 {
			c.FacePlayer()
			c.Attack()
			c.AttackTimeout = 10
		}
	}

	c.AttackTimeout -= Engine.Renderer.DeltaFrameTime
}

// UpdateAnimations plays the appropriate animation
// based on the state of the mob
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

//  --------------------------------------------------
//  Helper functions
//  --------------------------------------------------

func (c *Common) MoveTo(x, y float32) {
	c.TargetX = x
	c.TargetY = y
}

func (c *Common) Jump() {
	if c.NumJumps > 0 {
		c.NumJumps--
		c.MonsterChild.VY = BaseSpeedY * c.VYMult
		c.MonsterMaterial.PlayAnimationOnce("jump")
		c.CurrentAnim = "jump"
	}
}

func (c *Common) FacePlayer() {
	dx := c.MonsterChild.X - Player1.PlayerChild.X
	if dx < 0 {

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
	Player1.Money += rand.Intn(4) + 2
	// c.MonsterMaterial.PlayAnimationOnceCallback("die", c.SetDead, nil) when we have dying anim's in
	c.SetDead() // temporary

}

func (c *Common) SetDead() {
	c.Dead = true
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

//  --------------------------------------------------
//  Move this somewhere else eventually
//  --------------------------------------------------

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
