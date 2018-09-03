package main

import (
	"rapidengine"
	"rapidengine/configuration"
	"runtime"
)

var engine rapidengine.Engine
var config configuration.EngineConfig

var WorldChild rapidengine.Child2D
var Player rapidengine.Child2D

var blocks []string

func init() {
	runtime.LockOSThread()
}

func main() {
	config = rapidengine.NewEngineConfig(1920, 1080, 2)
	engine = rapidengine.NewEngine(config, render)

	engine.Renderer.SetRenderDistance(1000)
	engine.Renderer.MainCamera.SetPosition(100, 100)
	engine.Renderer.MainCamera.SetSpeed(0.2)

	engine.TextureControl.NewTexture("./sprites/sky.png", "sky")
	engine.TextureControl.NewTexture("./sprites/dirt.png", "dirt")
	engine.TextureControl.NewTexture("./sprites/grass.png", "grass")

	blocks = append(blocks, "sky")
	blocks = append(blocks, "dirt")
	blocks = append(blocks, "grass")

	WorldChild = engine.NewChild2D()
	WorldChild.AttachPrimitive(rapidengine.NewRectangle(BlockSize, BlockSize, &config))
	WorldChild.AttachTexturePrimitive(engine.TextureControl.GetTexture("grass"))
	WorldChild.EnableCopying()
	WorldChild.AttachCollider(0, 0, BlockSize, BlockSize)

	engine.Config.Logger.Info("Generating world...")
	generateWorld()
	createCopies()

	engine.CollisionControl.CreateGroup("ground")
	engine.CollisionControl.AddChildToGroup(&WorldChild, "ground")

	Player = engine.NewChild2D()
	Player.AttachPrimitive(rapidengine.NewRectangle(30, 50, &config))
	Player.AttachTexturePrimitive(engine.TextureControl.GetTexture("sky"))
	Player.SetPosition(1000, 20000)
	Player.AttachCollider(0, 0, 30, 50)
	Player.SetGravity(1)

	engine.Instance(&Player)
	engine.Instance(&WorldChild)

	err := engine.Initialize()
	if err != nil {
		panic(err)
	}

	engine.StartRenderer()
	<-engine.Done()
}

func render(renderer *rapidengine.Renderer, keys map[string]bool) {
	renderer.RenderChildren()
	renderer.MainCamera.SetPosition(Player.X, Player.Y)
	movePlayer(keys)
}

func movePlayer(keys map[string]bool) {
	if keys["w"] {
		Player.SetVelocityY(10)
	}
	if keys["a"] {
		Player.SetVelocityX(4)
	} else if keys["d"] {
		Player.SetVelocityX(-4)
	} else {
		Player.SetVelocityX(0)
	}
}
