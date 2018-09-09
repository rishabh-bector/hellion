package main

import (
	_ "net/http/pprof"
	"rapidengine"
	"rapidengine/configuration"
	"rapidengine/input"
	"runtime"
)

var engine rapidengine.Engine
var config configuration.EngineConfig

var WorldChild rapidengine.Child2D
var Player rapidengine.Child2D
var SkyChild rapidengine.Child2D

var blocks []string

func init() {
	runtime.LockOSThread()
}

func main() {
	config = rapidengine.NewEngineConfig(1920, 1080, 2)
	engine = rapidengine.NewEngine(config, render)

	engine.Renderer.SetRenderDistance(1000)
	engine.Renderer.MainCamera.SetPosition(100, 100, 0)
	engine.Renderer.MainCamera.SetSpeed(0.2)

	engine.TextureControl.NewTexture("./assets/player/player.png", "player")
	engine.TextureControl.NewTexture("./assets/player/playerWalking3.png", "player3")
	//engine.TextureControl.NewTexture("./assets/player/player.png", "player")

	engine.TextureControl.NewTexture("./assets/blocks/dirt.png", "dirt")
	engine.TextureControl.NewTexture("./assets/blocks/grass.png", "grass")
	engine.TextureControl.NewTexture("./assets/blocks/stone.png", "stone")
	engine.TextureControl.NewTexture("./assets/back.png", "back")

	blocks = append(blocks, "player")
	blocks = append(blocks, "dirt")
	blocks = append(blocks, "grass")
	blocks = append(blocks, "stone")

	WorldChild = engine.NewChild2D()
	WorldChild.AttachShader(engine.ShaderControl.GetShader("texture"))
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
	Player.AttachShader(engine.ShaderControl.GetShader("texture"))
	Player.AttachPrimitive(rapidengine.NewRectangle(30, 50, &config))
	Player.AttachTexturePrimitive(engine.TextureControl.GetTexture("player"))
	Player.SetPosition(3000, 2000*25)
	Player.AttachCollider(0, 0, 30, 50)
	Player.SetGravity(1)

	Player.EnableAnimation()
	Player.SetAnimationSpeed(10)
	Player.AddFrame(engine.TextureControl.GetTexture("player"))
	Player.AddFrame(engine.TextureControl.GetTexture("player3"))

	SkyChild = engine.NewChild2D()
	SkyChild.AttachShader(engine.ShaderControl.GetShader("texture"))
	SkyChild.AttachPrimitive(rapidengine.NewRectangle(2000, 1150, &config))
	SkyChild.AttachTexturePrimitive(engine.TextureControl.GetTexture("back"))

	engine.Instance(&SkyChild)
	engine.Instance(&WorldChild)
	engine.Instance(&Player)

	engine.Initialize()
	engine.StartRenderer()
	<-engine.Done()
	return
}

func render(renderer *rapidengine.Renderer, inputs *input.Input) {
	movePlayer(inputs.Keys)
	renderer.MainCamera.SetPosition(Player.X, Player.Y, 0)
	SkyChild.SetPosition(Player.X-1950/2, Player.Y-1110/2)
}

func movePlayer(keys map[string]bool) {
	if keys["w"] {
		Player.SetVelocityY(10)
	}
	if keys["a"] {
		Player.SetVelocityX(7)
	} else if keys["d"] {
		Player.SetVelocityX(-7)
	} else {
		Player.SetVelocityX(0)
	}
}
