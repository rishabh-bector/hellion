package main

import (
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
	SpeedX   float32
	SpeedY   float32
	Gravity  float32
	NumJumps int

	// Temp state
	Crouching bool
	Attacking bool

	// Data
	Health      int
	CurrentAnim string

	// Attacks
	JustPunched bool

	// Collision
	Hitbox1 AABB
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
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_p2"), "punch")
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
		God:            false,
		PlayerChild:    PlayerChild,
		PlayerMaterial: playerMaterial,
		SpeedX:         BaseSpeedX,
		SpeedY:         BaseSpeedY,
		Gravity:        BaseGravity,
		NumJumps:       1,
		CurrentAnim:    "idle",
		Health:         100,
	}

	if Player1.God {
		Player1.NumJumps = 10000
		Player1.SpeedX = 600
	}

	Player1.Hitbox1 = AABB{
		X:      0,
		Y:      0,
		Width:  50,
		Height: 120,

		DirectionOffset: 1,
		MinimumXDist:    47,
		MinimumYDist:    20,
	}
}

func (p *Player) Update(inputs *input.Input) {
	p.UpdateMovement(inputs)
	p.UpdateAnimation()
}

var TCSpeed = float32(11.5 * 30)

var Started = 100

func (p *Player) UpdateMovement(inputs *input.Input) {
	colChild.ScaleX = Player1.Hitbox1.Width
	colChild.ScaleY = Player1.Hitbox1.Height
	colChild.X = Player1.Hitbox1.X
	colChild.Y = Player1.Hitbox1.Y

	p.Hitbox1.X = p.PlayerChild.X + (p.PlayerChild.ScaleX / 2) - (p.Hitbox1.Width / 2) + (p.PlayerChild.VX * float32(Engine.Renderer.DeltaFrameTime))
	p.Hitbox1.Y = p.PlayerChild.Y + (p.PlayerChild.ScaleY / 2) - (p.Hitbox1.Height / 2) + (p.PlayerChild.VY * float32(Engine.Renderer.DeltaFrameTime))

	top, left, bottom, right, topleft, topright := CheckWorldCollision(p.Hitbox1, p.PlayerChild.VX, p.PlayerChild.VY)
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
		if !topleft {
			p.PlayerChild.VY = TCSpeed
			p.NumJumps--
		} else {
			p.PlayerChild.VX = 0
		}
	}
	if (right || topright) && p.PlayerChild.VX < -100 {
		if !topright {
			p.PlayerChild.VY = TCSpeed
			p.NumJumps--
		} else {
			p.PlayerChild.VX = 0
		}
	}
	if top && p.PlayerChild.VY > 0 {
		p.PlayerChild.VY = 0
	}

	// Attacking

	if inputs.Keys["p"] {
		if !p.JustPunched {
			p.PlayerMaterial.PlayAnimationOnceCallback("punch", p.DoneAttack)
			p.CurrentAnim = "punch"
			p.Attacking = true
			p.JustPunched = true
			p.PlayerChild.VX = 0
		}
	} else {
		p.JustPunched = false
	}
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
