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

var blocks []rapidengine.Material

var l rapidengine.PointLight

func init() {
	runtime.LockOSThread()
}

func main() {
	config = rapidengine.NewEngineConfig(1920, 1080, 2)
	engine = rapidengine.NewEngine(config, render)
	engine.Renderer.SetRenderDistance(1000)
	engine.Renderer.MainCamera.SetPosition(100, 100, 0)
	engine.Renderer.MainCamera.SetSpeed(0.2)

	//   --------------------------------------------------
	//   Textures
	//   --------------------------------------------------

	engine.TextureControl.NewTexture("./assets/player/player.png", "player")
	engine.TextureControl.NewTexture("./assets/player/playerWalking3.png", "player3")
	engine.TextureControl.NewTexture("./assets/blocks/dirt.png", "dirt")
	engine.TextureControl.NewTexture("./assets/testBlocks/topDirt.png", "topDirt")
	engine.TextureControl.NewTexture("./assets/testBlocks/topLeftDirt.png", "topLeftDirt")
	engine.TextureControl.NewTexture("./assets/testBlocks/topRightdirt.png", "topRightDirt")
	engine.TextureControl.NewTexture("./assets/testBlocks/bottomDirt.png", "bottomDirt")
	engine.TextureControl.NewTexture("./assets/testBlocks/bottomRightDirt.png", "bottomRightDirt")
	engine.TextureControl.NewTexture("./assets/testBlocks/bottomLeftDirt.png", "bottomLeftDirt")
	engine.TextureControl.NewTexture("./assets/testBlocks/leftDirt.png", "leftDirt")
	engine.TextureControl.NewTexture("./assets/testBlocks/rightDirt.png", "rightDirt")
	engine.TextureControl.NewTexture("./assets/blocks/grass.png", "grass")
	engine.TextureControl.NewTexture("./assets/blocks/stone.png", "stone")
	engine.TextureControl.NewTexture("./assets/back.png", "back")

	//   --------------------------------------------------
	//   Materials
	//   --------------------------------------------------

	grassMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"))
	grassMaterial.BecomeTexture(engine.TextureControl.GetTexture("grass"))

	playerMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"))
	playerMaterial.BecomeTexture(engine.TextureControl.GetTexture("player"))

	dirtMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"))
	dirtMaterial.BecomeTexture(engine.TextureControl.GetTexture("dirt"))

	topDirtMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"))
	topDirtMaterial.BecomeTexture(engine.TextureControl.GetTexture("topDirt"))

	topRightDirtMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"))
	topRightDirtMaterial.BecomeTexture(engine.TextureControl.GetTexture("topRightDirt"))

	topLeftDirtMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"))
	topLeftDirtMaterial.BecomeTexture(engine.TextureControl.GetTexture("topLeftDirt"))

	bottomRightDirtMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"))
	bottomRightDirtMaterial.BecomeTexture(engine.TextureControl.GetTexture("bottomRightDirt"))

	bottomLeftDirtMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"))
	bottomLeftDirtMaterial.BecomeTexture(engine.TextureControl.GetTexture("bottomLeftDirt"))

	rightDirtMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"))
	rightDirtMaterial.BecomeTexture(engine.TextureControl.GetTexture("rightDirt"))

	leftDirtMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"))
	leftDirtMaterial.BecomeTexture(engine.TextureControl.GetTexture("leftDirt"))

	stoneMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"))
	stoneMaterial.BecomeTexture(engine.TextureControl.GetTexture("stone"))

	backgroundMaterial := rapidengine.NewMaterial(engine.ShaderControl.GetShader("colorLighting"))
	backgroundMaterial.BecomeTexture(engine.TextureControl.GetTexture("back"))

	blocks = append(blocks, playerMaterial)          //0
	blocks = append(blocks, dirtMaterial)            //1
	blocks = append(blocks, grassMaterial)           //2
	blocks = append(blocks, stoneMaterial)           //3
	blocks = append(blocks, topDirtMaterial)         //4
	blocks = append(blocks, topLeftDirtMaterial)     //5
	blocks = append(blocks, topRightDirtMaterial)    //6
	blocks = append(blocks, bottomRightDirtMaterial) //7
	blocks = append(blocks, bottomLeftDirtMaterial)  //8
	blocks = append(blocks, rightDirtMaterial)       //9
	blocks = append(blocks, leftDirtMaterial)        //10

	//   --------------------------------------------------
	//   Children
	//   --------------------------------------------------

	WorldChild = engine.NewChild2D()
	WorldChild.AttachShader(engine.ShaderControl.GetShader("colorLighting"))
	WorldChild.AttachPrimitive(rapidengine.NewRectangle(BlockSize, BlockSize, &config))
	WorldChild.AttachTextureCoordsPrimitive()
	WorldChild.EnableCopying()
	WorldChild.AttachCollider(0, 0, BlockSize, BlockSize)

	Player = engine.NewChild2D()
	Player.AttachShader(engine.ShaderControl.GetShader("colorLighting"))
	Player.AttachPrimitive(rapidengine.NewRectangle(30, 50, &config))
	Player.AttachTextureCoordsPrimitive()
	Player.AttachMaterial(&playerMaterial)
	Player.SetPosition(3000, 1000*25)
	Player.AttachCollider(0, 0, 30, 50)
	Player.SetGravity(1)

	SkyChild = engine.NewChild2D()
	SkyChild.AttachShader(engine.ShaderControl.GetShader("colorLighting"))
	SkyChild.AttachPrimitive(rapidengine.NewRectangle(2000, 1500, &config))
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
		[]float32{0.9, 0.9, 0.7},
		[]float32{0, 0, 0},
		//1.0, -300, 299.8,
		1.0, 0.5, 0.1,
	)
	l.SetPosition([]float32{0, 0, 1})

	engine.InstanceLight(&l)

	engine.Instance(&SkyChild)
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
	renderer.MainCamera.SetPosition(Player.X, Player.Y, 0)
	SkyChild.SetPosition(Player.X-1950/2, Player.Y-1110/2)

	x, y := rapidengine.ScaleCoordinates(Player.X, Player.Y+30, 1920, 1080)
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
