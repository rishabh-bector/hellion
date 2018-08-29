package main

import (
	"rapidengine"
	"rapidengine/configuration"
)

var engine rapidengine.Engine

var child1 rapidengine.Child2D
var child2 rapidengine.Child2D

func main() {
	c := configuration.EngineConfig{
		ScreenWidth:    1920,
		ScreenHeight:   1080,
		WindowTitle:    "game",
		PolygonLines:   false,
		CollisionLines: true,
		Dimensions:     2,
	}

	engine = rapidengine.NewEngine(c, render)
	err := engine.Initialize()
	if err != nil {
		panic(err)
	}

	engine.TextureControl.NewTexture("./krita/blueSword.jpeg", "sword")
	engine.TextureControl.NewTexture("./krita/dirt.jpeg", "dirt")

	child1 = engine.NewChild2D()
	child1.AttachPrimitive(rapidengine.NewRectangle(50, 50, &c))
	child1.AttachTexturePrimitive(engine.TextureControl.GetTexture("dirt"))
	child1.EnableCopying()
	engine.Instance(&child1)

	for x := 0; x < 5000; x += 50 {
		for y := 0; y < 10000; y += 50 {
			child1.AddCopy(rapidengine.ChildCopy{float32(x), float32(y), engine.TextureControl.GetTexture("sword")})
		}
	}

	/*engine.CollisionControl.CreateGroup("children")
	engine.CollisionControl.AddChildToGroup(&child1, "children")
	engine.CollisionControl.AddChildToGroup(&child2, "children")
	engine.CollisionControl.CreateCollision(&child1, "children", cbk)

	child1.SetVelocity(1, 0)
	child1.AttachGravity(0.2)*/

	engine.Renderer.SetRenderDistance(2000)
	engine.Renderer.MainCamera.SetPosition(0, 5000)

	engine.InitializeRenderer()

	engine.StartRenderer()
	<-engine.Done()
}

func render(renderer *rapidengine.Renderer) {
	renderer.RenderChildren()
}
