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

	PlayerChild    *child.Child2D
	PlayerMaterial *material.BasicMaterial

	SpeedX float32
	SpeedY float32

	NumJumps int

	CurrentAnim string

	Gravity float32
}

func InitializePlayer() {
	Engine.TextureControl.NewTexture("assets/player/idle/f1.png", "p_i1", "pixel")
	Engine.TextureControl.NewTexture("assets/player/idle/f2.png", "p_i2", "pixel")

	Engine.TextureControl.NewTexture("assets/player/walk/f1.png", "p_w1", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/f2.png", "p_w2", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/f3.png", "p_w3", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/f4.png", "p_w4", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/f5.png", "p_w5", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/f6.png", "p_w6", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/f7.png", "p_w7", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/f8.png", "p_w8", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/f9.png", "p_w9", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/f10.png", "p_w10", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/f11.png", "p_w11", "pixel")
	Engine.TextureControl.NewTexture("assets/player/walk/f12.png", "p_w12", "pixel")

	Engine.TextureControl.NewTexture("assets/player/jump/f1.png", "p_j1", "pixel")
	Engine.TextureControl.NewTexture("assets/player/jump/f2.png", "p_j2", "pixel")

	Engine.TextureControl.NewTexture("assets/player/fall/f1.png", "p_f1", "pixel")
	Engine.TextureControl.NewTexture("assets/player/fall/f2.png", "p_f2", "pixel")

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

	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_f1"), "fall")
	playerMaterial.AddFrame(Engine.TextureControl.GetTexture("p_f2"), "fall")

	playerMaterial.SetAnimationFPS("walk", 30)
	playerMaterial.SetAnimationFPS("idle", 5)
	playerMaterial.SetAnimationFPS("jump", 10)
	playerMaterial.SetAnimationFPS("fall", 10)
	playerMaterial.PlayAnimation("idle")

	PlayerChild := Engine.ChildControl.NewChild2D()
	PlayerChild.AttachMaterial(playerMaterial)
	PlayerChild.AttachMesh(geometry.NewRectangle())
	PlayerChild.ScaleX = 60
	PlayerChild.ScaleY = 120
	PlayerChild.Gravity = 0

	Player1 = Player{
		God:            true,
		PlayerChild:    PlayerChild,
		PlayerMaterial: playerMaterial,
		SpeedX:         5,
		SpeedY:         20,
		Gravity:        1.05,
		NumJumps:       1,
		CurrentAnim:    "idle",
	}

	if Player1.God {
		Player1.NumJumps = 10000
		Player1.SpeedX = 20
	}
}

func (p *Player) Update(inputs *input.Input) {
	p.UpdateMovement(inputs)
	p.UpdateAnimation()
}

func (p *Player) UpdateMovement(inputs *input.Input) {

	top, left, bottom, right := p.CheckWorldCollision()
	if bottom {
		if p.God {
			p.NumJumps = 10000
		} else {
			p.NumJumps = 1
		}
		p.PlayerChild.VY = 0
	} else {
		p.PlayerChild.VY -= p.Gravity
	}

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

	if left && p.PlayerChild.VX > 1 {
		p.PlayerChild.VX = 0
	}
	if right && p.PlayerChild.VX < 1 {
		p.PlayerChild.VX = 0
	}
	if top && p.PlayerChild.VY > 0 {
		p.PlayerChild.VY = 0
	}

}

func (p *Player) UpdateAnimation() {
	if p.PlayerChild.VX > 0 && p.NumJumps > 0 && p.CurrentAnim != "walk" {
		p.PlayerMaterial.PlayAnimation("walk")
		p.CurrentAnim = "walk"
		return
	}
	if p.PlayerChild.VX < 0 && p.NumJumps > 0 && p.CurrentAnim != "walk" {
		p.PlayerMaterial.PlayAnimation("walk")
		p.CurrentAnim = "walk"
		return
	}
	if p.PlayerChild.VX == 0 && p.NumJumps > 0 && p.CurrentAnim != "idle" {
		p.PlayerMaterial.PlayAnimation("idle")
		p.CurrentAnim = "idle"
		return
	}

	if p.PlayerChild.VY < -24 && p.CurrentAnim != "fall" {
		p.PlayerMaterial.PlayAnimationOnce("fall")
		p.CurrentAnim = "fall"
		return
	}
}

// top, left, bottom, right
func (p *Player) CheckWorldCollision() (bool, bool, bool, bool) {
	top := false
	left := false
	bottom := false
	right := false

	px := int((p.PlayerChild.X + BlockSize/2) / BlockSize)
	py := int((p.PlayerChild.Y)/BlockSize + 1)

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

	// auto jump
	if right && WorldMap.GetWorldBlockID(px + 1, py + 2) == "00000" && WorldMap.GetWorldBlockID(px + 1, py + 3) == "00000" && WorldMap.GetWorldBlockID(px + 1, py + 4) == "00000" {
		p.PlayerChild.VY = p.SpeedY
	}

	return top, left, bottom, right
}
