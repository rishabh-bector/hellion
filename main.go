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

var Player rapidengine.Child2D
var SkyChild rapidengine.Child2D

var l rapidengine.PointLight

var ScreenWidth = 1440
var ScreenHeight = 900

func init() {
	runtime.LockOSThread()
}

func main() {
	config = rapidengine.NewEngineConfig(ScreenWidth, ScreenHeight, 2)
	engine = rapidengine.NewEngine(config, render)
	engine.Renderer.SetRenderDistance(float32(ScreenWidth/2) + 50)
	engine.Renderer.MainCamera.SetPosition(100, 100, 0)
	engine.Renderer.MainCamera.SetSpeed(0.2)

	//   --------------------------------------------------
	//   Textures
	//   --------------------------------------------------

	engine.TextureControl.NewTexture("./assets/player/player.png", "player")
	engine.TextureControl.NewTexture("./assets/player/playerWalking3.png", "player3")
	engine.TextureControl.NewTexture("./assets/backgrounds/gradient.png", "back")

	//   --------------------------------------------------
	//   Materials
	//   --------------------------------------------------

	playerMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	playerMaterial.BecomeTexture(engine.TextureControl.GetTexture("player"))

	backgroundMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	backgroundMaterial.BecomeTexture(engine.TextureControl.GetTexture("back"))

	//   --------------------------------------------------
	//   Children
	//   --------------------------------------------------

	Player = engine.NewChild2D()
	Player.AttachShader(engine.ShaderControl.GetShader("colorLighting"))
	Player.AttachPrimitive(rapidengine.NewRectangle(30, 50, &config))
	Player.AttachTextureCoordsPrimitive()
	Player.AttachMaterial(&playerMaterial)
	Player.SetPosition(3000, 1000*25)
	Player.AttachCollider(5, 5, 20, 15)
	Player.SetGravity(1)

	SkyChild = engine.NewChild2D()
	SkyChild.AttachShader(engine.ShaderControl.GetShader("colorLighting"))
	SkyChild.AttachPrimitive(rapidengine.NewRectangle(4000, 1100, &config))
	SkyChild.AttachTextureCoordsPrimitive()
	SkyChild.AttachMaterial(&backgroundMaterial)

	light := engine.NewChild3D()
	light.AttachPrimitive(rapidengine.NewCube())
	light.AttachShader(engine.ShaderControl.GetShader("colorLighting"))
	light.AttachMaterial(&playerMaterial)
	light.AttachTextureCoords(rapidengine.CubeTextures)
	light.SetPosition(-1, -1, -10)

	//   --------------------------------------------------
	//   World Gen
	//   --------------------------------------------------

	engine.Config.Logger.Info("Generating world...")
	generateWorld()
	createCopies()

	engine.CollisionControl.CreateGroup("ground")
	engine.CollisionControl.AddChildToGroup(&WorldChild, "ground")

	l = rapidengine.NewPointLight(
		engine.ShaderControl.GetShader("colorLighting"),
		[]float32{0, 0, 0},
		[]float32{0.9, 0.9, 0.9},
		[]float32{0, 0, 0},
		//1.0, -300, 299.8,
		//1.0, -2, 1.9,
		1.0, 0.5, 0.1,
	)
	l.SetPosition([]float32{0, 0, 1})

	engine.InstanceLight(&l)

	engine.Instance(&SkyChild)
	engine.Instance(&CloudChild)
	engine.Instance(&NoCollisionChild)
	engine.Instance(&NatureChild)
	engine.Instance(&WorldChild)
	engine.Instance(&Player)
	//engine.Instance(&light)

	engine.EnableLighting()
	engine.Initialize()
	engine.StartRenderer()
	<-engine.Done()
	return
}

func render(renderer *rapidengine.Renderer, inputs *input.Input) {
	movePlayer(inputs.Keys)
	renderer.MainCamera.SetPosition(Player.X, Player.Y, -10)
	SkyChild.SetPosition(Player.X-1950/2, Player.Y-1110/2)

	x, y := rapidengine.ScaleCoordinates(Player.X, float32(HeightMap[int(Player.X/BlockSize)])*BlockSize, float32(ScreenWidth), float32(ScreenHeight))
	l.SetPosition([]float32{x, y, 1})
}

func movePlayer(keys map[string]bool) {
	if keys["w"] {
		Player.SetVelocityY(20)
	}
	if keys["a"] {
		Player.SetVelocityX(10)
	} else if keys["d"] {
		Player.SetVelocityX(-10)
	} else {
		Player.SetVelocityX(0)
	}
}
