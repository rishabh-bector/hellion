package main

import r "RapidEngine"

func main() {
	c := r.EngineConfig{
		ScreenWidth:  1920,
		ScreenHeight: 1080,
		WindowTitle:  "game",
		PolygonLines: false,
		Dimensions:   2,
	}

	engine := r.NewEngine(c)
	err := engine.Initialize()
	if err != nil {
		panic(err)
	}
	engine.SetRenderFunc(render)

	child1 := engine.NewChild2D()
	child1.AttachPrimitive(r.NewRectangle(100, 100, &c))
	child1.AttachTexturePrimitive("./texture.png")
	child1.SetPosition(500, 500)

	engine.Instance(&child1)

	engine.StartRenderer()
	<-engine.Done()
}

func render(renderer *r.Renderer) {
	renderer.RenderChildren()
}
