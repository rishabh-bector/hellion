package main

// import (
// 	//"rapidengine/child"
// 	//"math"
// )

// // type Monster interface {

// // 	Move func(direction string)

// // 	Jump func()

// // 	Attack func(direction string)

// // 	X int

// // 	Y int

// // 	VelocityX int

// // 	VelocityY int

// // 	Health int

// // 	Texture string

// // 	DropMap map[string]int

// // 	Width int

// // 	Height int

// // 	ContactDamage int
// // }

// type Goblin struct {
// 	Health int
// 	X int
// 	Y int
// 	VelocityX int
// 	VelocityY int
// 	Texture string
// 	Move func(direction string)
// 	Jump func()
// 	Width int
// 	Height int
// 	ContactDamage int

// }

// var g = Goblin{

// 	Move: func(direction string) {
// 		if direction == "left"  && g.VelocityX > 0 {
// 			g.VelocityX *= -1
// 		}
// 		if direction == "right" && g.VelocityX < 0 {
// 			g.VelocityX *= -1
// 		}
// 	},

// 	Jump: func() {
// 		g.VelocityY = 10
// 	},

// 	Health: 100,

// 	Texture: "/assets/player/OrcBoss.png",

// 	Width: 1,

// 	Height: 2,

// 	ContactDamage: 50,

// 	VelocityX: 0,

// 	VelocityY: 0,

// }
