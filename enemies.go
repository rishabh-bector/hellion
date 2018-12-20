package main

//"rapidengine/child"
//"math"

type Monster interface {
	Move(direction string)

	Jump()

	Attack(direction string)
}

type Goblin struct {
	Health        int
	X             int
	Y             int
	VelocityX     int
	VelocityY     int
	Texture       string
	Width         int
	Height        int
	ContactDamage int
}

func (g *Goblin) Move(direction string) {
	if direction == "left" && g.VelocityX > 0 {
		g.VelocityX *= -1
	}
	if direction == "right" && g.VelocityX < 0 {
		g.VelocityX *= -1
	}
}

func (g *Goblin) Jump() {
	g.VelocityY = 10
}

func (g *Goblin) Attack(direction string) {
	if direction == "left" {
		//play attack animation
	}
	if direction == "right" {
		//play attack animation
	}
}

var g = Goblin{

	Health: 100,

	Texture: "/assets/player/OrcBoss.png",

	Width: 1,

	Height: 2,

	ContactDamage: 50,

	VelocityX: 0,

	VelocityY: 0,
}
