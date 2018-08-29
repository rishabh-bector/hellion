package main

import r "RapidEngine"

var child1 r.Child2D
var child2 r.Child2D

func main() {
	c := r.EngineConfig{
		ScreenWidth:    1920,
		ScreenHeight:   1080,
		WindowTitle:    "game",
		PolygonLines:   false,
		CollisionLines: true,
		Dimensions:     2,
	}

	engine := r.NewEngine(c, render)
	err := engine.Initialize()
	if err != nil {
		panic(err)
	}

	child1 = engine.NewChild2D()
	child1.AttachPrimitive(r.NewRectangle(100, 100, &c))
	child1.AttachTexturePrimitive("./krita/blueSword.jpeg")
	child1.AttachCollider(0, 0, 10, 10)
	child1.SetPosition(500, 500)

	child2 = engine.NewChild2D()
	child2.AttachPrimitive(r.NewRectangle(100, 100, &c))
	child2.AttachTexturePrimitive("./krita/goldSword.jpeg")
	child2.AttachCollider(0, 0, 10, 10)
	child2.SetPosition(600, 500)

	engine.Instance(&child1)
	engine.Instance(&child2)

	engine.CollisionControl.CreateGroup("children")
	engine.CollisionControl.AddChildToGroup(&child1, "children")
	engine.CollisionControl.AddChildToGroup(&child2, "children")
	engine.CollisionControl.CreateCollision(&child1, "children", cbk)

	child1.SetVelocity(1, 0)
	child1.AttachGravity(0.2)

	engine.StartRenderer()
	<-engine.Done()
}

func render(renderer *r.Renderer) {
	renderer.RenderChildren()
}

func cbk() {
	println("YEETED")
}
