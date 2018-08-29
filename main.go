package main

import r "RapidEngine"

var engine r.Engine

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

	engine = r.NewEngine(c, render)
	err := engine.Initialize()
	if err != nil {
		panic(err)
	}

	engine.TextureControl.NewTexture("./krita/blueSword.jpeg", "blue-sword")

	child1 = engine.NewChild2D()
	child1.AttachPrimitive(r.NewRectangle(50, 50, &c))
	child1.AttachTexturePrimitive(engine.TextureControl.GetTexture("blue-sword"))
	child1.EnableCopying()
	engine.Instance(&child1)

	for x := 0; x < 5000; x += 50 {
		for y := 0; y < 5000; y += 50 {
			child1.AddCopy(r.ChildCopy{float32(x), float32(y)})
		}
	}

	child2 = engine.NewChild2D()
	child2.AttachPrimitive(r.NewRectangle(100, 100, &c))
	child2.AttachTexturePrimitive(engine.TextureControl.GetTexture("blue-sword"))
	child2.SetPosition(0, 0)
	engine.Instance(&child2)

	/*engine.CollisionControl.CreateGroup("children")
	engine.CollisionControl.AddChildToGroup(&child1, "children")
	engine.CollisionControl.AddChildToGroup(&child2, "children")
	engine.CollisionControl.CreateCollision(&child1, "children", cbk)

	child1.SetVelocity(1, 0)
	child1.AttachGravity(0.2)*/

	engine.Renderer.SetRenderDistance(2000)
	engine.Renderer.MainCamera.SetPosition(0, 5000)

	engine.StartRenderer()
	<-engine.Done()
}

func render(renderer *r.Renderer) {
	renderer.RenderChildren()
}
