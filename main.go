package main

import (
	"math/rand"
	"rapidengine"
	"rapidengine/configuration"
	"runtime"

	perlin "github.com/aquilax/go-perlin"
)

var engine rapidengine.Engine
var config configuration.EngineConfig

var child1 rapidengine.Child2D
var child2 rapidengine.Child2D

var blocks []string

func init() {
	runtime.LockOSThread()
}

func main() {
	config = rapidengine.NewEngineConfig(1920, 1080, 2)
	engine = rapidengine.NewEngine(config, render)
	err := engine.Initialize()
	if err != nil {
		panic(err)
	}

	engine.TextureControl.NewTexture("./sprites/sky.png", "sky")
	engine.TextureControl.NewTexture("./sprites/dirt.png", "dirt")
	engine.TextureControl.NewTexture("./sprites/grass.png", "grass")

	blocks = append(blocks, "sky")
	blocks = append(blocks, "dirt")
	blocks = append(blocks, "grass")

	p = perlin.NewPerlin(2, 2, 10, int64(rand.Int()))

	engine.Renderer.SetRenderDistance(1000)
	engine.Renderer.MainCamera.SetPosition(100, 100)
	engine.Renderer.MainCamera.SetSpeed(0.2)

	child1 = engine.NewChild2D()
	child1.AttachPrimitive(rapidengine.NewRectangle(BlockSize, BlockSize, &config))
	child1.AttachTexturePrimitive(engine.TextureControl.GetTexture("grass"))
	child1.EnableCopying()
	child1.AttachCollider(0, 0, 25, 25)

	engine.Config.Logger.Info("Generating world...")
	generateWorld()
	createCopies()

	engine.CollisionControl.CreateGroup("ground")
	engine.CollisionControl.AddChildToGroup(&child1, "ground")

	child2 = engine.NewChild2D()
	child2.AttachPrimitive(rapidengine.NewRectangle(50, 50, &config))
	child2.AttachTexturePrimitive(engine.TextureControl.GetTexture("sky"))
	child2.SetPosition(1000, 20000)
	child2.AttachCollider(0, 0, 100, 100)
	child2.SetGravity(1)

	engine.Instance(&child2)
	engine.Instance(&child1)
	engine.InitializeRenderer()

	engine.StartRenderer()
	<-engine.Done()
}

func render(renderer *rapidengine.Renderer) {
	renderer.RenderChildren()
	renderer.MainCamera.SetPosition(child2.X, child2.Y)
}
