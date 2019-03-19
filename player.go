package main

import (
	"fmt"
	"rapidengine/child"
	"rapidengine/geometry"
	"rapidengine/input"
	"rapidengine/material"
)

var Player1 Player

type Player struct {
	God bool

	// Components
	PlayerChild    *child.Child2D
	PlayerMaterial *material.BasicMaterial

	// Movement
	CenterX  float32
	CenterY  float32
	SpeedX   float32
	SpeedY   float32
	Gravity  float32
	NumJumps int

	// Temp state
	Crouching bool
	Attacking bool

	// Data
	Health      float32
	CurrentAnim string

	// Attack info
	PunchDamage float32

	// Timers
	Invincibility float64
	PunchCooldown float64

	// Collision
	Hitbox1   Hitbox
	FullBox   AABB
	AttackBox AABB

	CurrentMiningTimer float64
	lastsnapx          int
	lastsnapy          int
}

func InitializePlayer() {
	Engine.TextureControl.NewTexture("assets/player/idle/1.png", "p_i1", "pixel")
	Engine.TextureControl.NewTexture("assets/player/idle/2.png", "p_i2", "pixel")

	Engine.TextureControl.NewTexture("assets/player/walk/1.png", "p_w1", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/2.png", "p_w2", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/3.png", "p_w3", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/4.png", "p_w4", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/5.png", "p_w5", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/6.png", "p_w6", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/7.png", "p_w7", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/8.png", "p_w8", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/9.png", "p_w9", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/10.png", "p_w10", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/11.png", "p_w11", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/12.png", "p_w12", "pixel")

	Engine.TextureControl.NewTexture("assets/player/jump/1.png", "p_j1", "pixel")
	Engine.TextureControl.NewTexture("assets/player/jump/2.png", "p_j2", "pixel")
	Engine.TextureControl.NewTexture("assets/player/jump/3.png", "p_j3", "pixel")

	Engine.TextureControl.NewTexture("assets/player/fall/1.png", "p_f1", "pixel")
	Engine.TextureControl.NewTexture("assets/player/fall/2.png", "p_f2", "pixel")
	Engine.TextureControl.NewTexture("assets/player/fall/3.png", "p_f3", "pixel")

	Engine.TextureControl.NewTexture("assets/player/crouch/1.png", "p_c1", "pixel")
	Engine.TextureControl.NewTexture("assets/player/crouch/2.png", "p_c2", "pixel")

	Engine.TextureControl.NewTexture("assets/player/punch/1.png", "p_p1", "pixel")
	Engine.TextureControl.NewTexture("assets/player/punch/2.png", "p_p2", "pixel")
	Engine.TextureControl.NewTexture("assets/player/punch/3.png", "p_p3", "pixel")
	Engine.TextureControl.NewTexture("assets/player/punch/4.png", "p_p4", "pixel")
	Engine.TextureControl.NewTexture("assets/player/punch/5.png", "p_p5", "pixel")
	Engine.TextureControl.NewTexture("assets/player/punch/6.png", "p_p6", "pixel")

	playerMaterial := Engine.MaterialControl.NewBasicMaterial()
	playerMaterial.DiffuseLevel = 1
	playerMaterial.DiffuseMap = Engine.TextureControl.GetTexture("p_i1")

	playerMaterial.EnableAnimation()

	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_i1"), "idle")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_i2"), "idle")

	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_w1"), "walk")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_w2"), "walk")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_w3"), "walk")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_w4"), "walk")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_w5"), "walk")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_w6"), "walk")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_w7"), "walk")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_w8"), "walk")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_w9"), "walk")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_w10"), "walk")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_w11"), "walk")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_w12"), "walk")

	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_j1"), "jump")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_j2"), "jump")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_j3"), "jump")

	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_f1"), "fall")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_f2"), "fall")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_f3"), "fall")

	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_c1"), "crouch")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_c2"), "crouch")

	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_p1"), "punch")
	playerMaterial.AddHitFrame(Engine.TextureControl.GetTexture("p_p2"), "punch")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_p3"), "punch")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_p4"), "punch")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_p5"), "punch")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_p6"), "punch")

	playerMaterial.SetAnimationFPS("walk", 25)
	playerMaterial.SetAnimationFPS("idle", 5)
	playerMaterial.SetAnimationFPS("jump", 20)
	playerMaterial.SetAnimationFPS("fall", 10)
	playerMaterial.SetAnimationFPS("crouch", 15)
	playerMaterial.SetAnimationFPS("punch", 15)

	playerMaterial.PlayAnimation("idle")

	PlayerChild := Engine.ChildControl.NewChild2D()
	PlayerChild.AttachMaterial(playerMaterial)
	PlayerChild.AttachMesh(geometry.NewRectangle())
	PlayerChild.ScaleX = 240
	PlayerChild.ScaleY = 240
	PlayerChild.Gravity = 0

	Player1 = Player{
		God: false,

		PlayerChild:    PlayerChild,
		PlayerMaterial: playerMaterial,

		SpeedX: BaseSpeedX,
		SpeedY: BaseSpeedY,

		Gravity: BaseGravity,

		PunchDamage: 10,

		NumJumps:    1,
		CurrentAnim: "idle",
		Health:      100,
	}

	if Player1.God {
		Player1.NumJumps = 10000
		Player1.SpeedX = 600
	}

	original := AABB{
		X:      0,
		Y:      0,
		Width:  50,
		Height: 120,
	}
	Player1.Hitbox1 = NewHitBox(original, 5)

	Player1.FullBox = AABB{0, 0, 50, 120, 0, 0}
	Player1.AttackBox = AABB{
		OffX:   0,
		OffY:   45,
		Width:  55,
		Height: 60,
	}

	V.AddBox(&Player1.Hitbox1)
}

func (p *Player) Update(inputs *input.Input) {
	p.UpdateMovement(inputs)
	p.UpdateAnimation()
}

var TCSpeed = float32(11.5 * 30)

var Started = 100

func (p *Player) UpdateMovement(inputs *input.Input) {
	/*colChild.ScaleX = Player1.Hitbox1.Width
	colChild.ScaleY = Player1.Hitbox1.Height
	colChild.X = Player1.Hitbox1.X
	colChild.Y = Player1.Hitbox1.Y*/

	p.PlayerChild.X += p.PlayerChild.VX * -float32(Engine.Renderer.DeltaFrameTime)
	p.PlayerChild.Y += p.PlayerChild.VY * float32(Engine.Renderer.DeltaFrameTime)

	p.CenterX = p.PlayerChild.X + (p.PlayerChild.ScaleX / 2) - (p.Hitbox1.DAABB.Width / 2)
	p.CenterY = p.PlayerChild.Y + (p.PlayerChild.ScaleY / 2) - (p.Hitbox1.LAABB.Height / 2)

	p.FullBox.X = p.CenterX
	p.FullBox.Y = p.CenterY

	p.Hitbox1.X = p.CenterX
	p.Hitbox1.Y = p.CenterY

	flip := p.PlayerMaterial.Flipped
	if flip == 0 {
		flip = 1
	} else {
		flip = 0
	}
	p.AttackBox.X = (p.CenterX + (p.Hitbox1.DAABB.Width / 2)) + (p.AttackBox.OffX+p.AttackBox.Width)*float32(flip-1)
	p.AttackBox.Y = p.CenterY + p.AttackBox.OffY

	top, left, bottom, right, topleft, topright := CheckWorldCollision(p.Hitbox1, p.PlayerChild.VX, p.PlayerChild.VY, p.CenterX, p.CenterY)

	if bottom {
		if p.God {
			p.NumJumps = 10000
		} else {
			p.NumJumps = 1
		}
		p.PlayerChild.VY = 0
	} else if Started < 0 {
		p.PlayerChild.VY -= p.Gravity * float32(Engine.Renderer.DeltaFrameTime)
		p.NumJumps--
	} else {
		Started--
	}

	// Basic movement

	if !p.Crouching && !p.Attacking {
		if inputs.Keys["w"] && p.NumJumps > 0 {
			p.PlayerChild.VY = p.SpeedY
			p.PlayerMaterial.PlayAnimationOnce("jump")
			p.CurrentAnim = "jump"
			p.NumJumps--
		}

		if inputs.Keys["a"] {
			p.PlayerChild.VX = p.SpeedX
			p.PlayerMaterial.Flipped = 1
		} else if inputs.Keys["d"] {
			p.PlayerChild.VX = -1 * p.SpeedX
			p.PlayerMaterial.Flipped = 0
		} else {
			p.PlayerChild.VX = 0
		}
	}

	if inputs.Keys["s"] {
		if !p.Crouching {
			p.PlayerMaterial.PlayAnimationOnce("crouch")
			p.Crouching = true
			p.CurrentAnim = "crouch"
			p.PlayerChild.VX = 0
		}
	} else {
		p.Crouching = false
	}

	// Movement collision

	if (left || topleft) && p.PlayerChild.VX > 100 {
		p.PlayerChild.VX = 0
	}
	if (right || topright) && p.PlayerChild.VX < -100 {
		p.PlayerChild.VX = 0
	}

	if top && p.PlayerChild.VY > 0 {
		p.PlayerChild.VY = 0
	}
	if bottom && p.PlayerChild.VY < 0 {
		p.PlayerChild.VY = 0
	}

	// Attacking
	if inputs.Keys["p"] {
		if p.PunchCooldown <= 0 {
			p.Punch()
		}
	}

	// Timers
	if p.Invincibility > 0 {
		p.Invincibility -= Engine.Renderer.DeltaFrameTime
	}
	if p.PunchCooldown > 0 {
		p.PunchCooldown -= Engine.Renderer.DeltaFrameTime
	}
}

func (p *Player) Punch() {
	p.PlayerMaterial.PlayAnimationOnceCallback("punch", p.DoneAttack, p.PunchHitFrame)
	p.PunchCooldown = 0.5
	p.CurrentAnim = "punch"
	p.Attacking = true
	p.PlayerChild.VX = 0
}

func (p *Player) PunchHitFrame() {
	if enemy := EM.CheckPlayerCollision(); enemy != nil {
		enemy.Damage(p.PunchDamage)
	}
}

func (p *Player) Hit(damage float32) {
	if p.Invincibility > 0 {
		return
	} else {
		p.Invincibility = 0.75
	}

	Engine.Renderer.MainCamera.Shake(0.3, 0.01)
	p.Health -= damage
	fmt.Printf("Player hit! Health: %v\n", p.Health)
}

func (p *Player) DoneAttack() {
	p.Attacking = false
}

func (p *Player) UpdateAnimation() {
	if p.PlayerChild.VX > 0 && p.NumJumps > 0 && p.CurrentAnim != "walk" && !p.Crouching && !p.Attacking {
		p.PlayerMaterial.PlayAnimation("walk")
		p.CurrentAnim = "walk"
		return
	}
	if p.PlayerChild.VX < 0 && p.NumJumps > 0 && p.CurrentAnim != "walk" && !p.Crouching && !p.Attacking {
		p.PlayerMaterial.PlayAnimation("walk")
		p.CurrentAnim = "walk"
		return
	}
	if p.PlayerChild.VX == 0 && p.NumJumps > 0 && p.CurrentAnim != "idle" && !p.Crouching && !p.Attacking {
		p.PlayerMaterial.PlayAnimation("idle")
		p.CurrentAnim = "idle"
		return
	}
	if p.PlayerChild.VY < -400 && p.CurrentAnim != "fall" {
		p.PlayerMaterial.PlayAnimationOnce("fall")
		p.CurrentAnim = "fall"
		return
	}
}

func (p *Player) CheckEnemyCollision(enemy Enemy) bool {
	if c1, c2, c3, c4 := enemy.GetCommon().Hitbox1.CheckCollisionAABB(p.AttackBox, 0, 0, enemy.GetCommon().Hitbox1.X, enemy.GetCommon().Hitbox1.Y); c1 || c2 || c3 || c4 {
		return true
	}
	return false
}
