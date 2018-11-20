package main

import (
	"rapidengine/child"
	"rapidengine/geometry"
	"rapidengine/input"
	"rapidengine/material"
)

var Player1 Player

type Player struct {
	PlayerChild *child.Child2D

	SpeedX float32
	SpeedY float32

	Gravity float32
}

func InitializePlayer() {
	Engine.TextureControl.NewTexture("assets/player/player.png", "player", "pixel")
	playerMaterial := material.NewMaterial(Engine.ShaderControl.GetShader("texture"), &Config)
	playerMaterial.BecomeTexture(Engine.TextureControl.GetTexture("player"))

	PlayerChild := Engine.ChildControl.NewChild2D()
	PlayerChild.AttachMaterial(&playerMaterial)
	PlayerChild.AttachMesh(geometry.NewRectangle())
	PlayerChild.ScaleX = 32
	PlayerChild.ScaleY = 64
	PlayerChild.Gravity = 0

	Player1 = Player{
		PlayerChild: PlayerChild,
		SpeedX:      10,
		SpeedY:      20,
		Gravity:     1.05,
	}
}

func (p *Player) Update(inputs *input.Input) {
	p.UpdateMovement(inputs)
}

func (p *Player) UpdateMovement(inputs *input.Input) {
	if inputs.Keys["w"] {
		p.PlayerChild.VY = p.SpeedY
	}
	if inputs.Keys["a"] {
		p.PlayerChild.VX = p.SpeedX
	} else if inputs.Keys["d"] {
		p.PlayerChild.VX = -1 * p.SpeedX
	} else {
		p.PlayerChild.VX = 0
	}

	p.PlayerChild.VY -= p.Gravity

	top, left, bottom, right := p.CheckWorldCollision()
	if bottom && p.PlayerChild.VY < 0 {
		p.PlayerChild.VY = 0
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

	return top, left, bottom, right
}
