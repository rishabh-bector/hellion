package main

import (
	_ "net/http/pprof"
	"rapidengine/child"
	"rapidengine/cmd"
	"rapidengine/configuration"
	"rapidengine/geometry"
	"rapidengine/input"
	"rapidengine/lighting"
	"rapidengine/material"
	"runtime"
)

var engine *cmd.Engine
var config configuration.EngineConfig

var Player child.Child2D
var SkyChild child.Child2D
var BlockSelect child.Child2D

var l lighting.PointLight

var ScreenWidth = 1920
var ScreenHeight = 1080

func init() {
	runtime.LockOSThread()
}

func main() {
	config = cmd.NewEngineConfig(ScreenWidth, ScreenHeight, 2)
	config.ShowFPS = true
	config.FullScreen = false
	config.VSync = false
	engine = cmd.NewEngine(config, render)
	engine.Renderer.SetRenderDistance(float32(ScreenWidth/2) + 50)
	engine.Renderer.MainCamera.SetPosition(100, 100, 0)
	engine.Renderer.MainCamera.SetSpeed(0.2)

	engine.Renderer.AutomaticRendering = false

	//   --------------------------------------------------
	//   Textures
	//   --------------------------------------------------

	engine.TextureControl.NewTexture("assets/player/OrcBoss.png", "player")
	engine.TextureControl.NewTexture("assets/backgrounds/gradient.png", "back")

	//   --------------------------------------------------
	//   Materials
	//   --------------------------------------------------

	playerMaterial := material.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	playerMaterial.BecomeTexture(engine.TextureControl.GetTexture("player"))

	backgroundMaterial := material.NewMaterial(engine.ShaderControl.GetShader("colorLighting"), &config)
	backgroundMaterial.BecomeTexture(engine.TextureControl.GetTexture("back"))

	//   --------------------------------------------------
	//   Children
	//   --------------------------------------------------

	Player = engine.NewChild2D()
	Player.AttachShader(engine.ShaderControl.GetShader("colorLighting"))
	Player.AttachPrimitive(geometry.NewRectangle(256, 258, &config))
	Player.AttachTextureCoordsPrimitive()
	Player.AttachMaterial(&playerMaterial)
	Player.SetPosition(3000, 1000*BlockSize)
	Player.Gravity = 0

	SkyChild = engine.NewChild2D()
	SkyChild.AttachShader(engine.ShaderControl.GetShader("colorLighting"))
	SkyChild.AttachPrimitive(geometry.NewRectangle(4000, 1100, &config))
	SkyChild.AttachTextureCoordsPrimitive()
	SkyChild.AttachMaterial(&backgroundMaterial)

	BlockSelect = engine.NewChild2D()
	BlockSelect.AttachShader(engine.ShaderControl.GetShader("color"))

	m := material.NewMaterial(engine.ShaderControl.GetShader("color"), &config)
	m.BecomeColor([]float32{0.5, 0.5, 0.5, 0.5})

	BlockSelect.AttachMaterial(&m)
	BlockSelect.AttachPrimitive(geometry.NewRectangle(32, 32, &config))

	//   --------------------------------------------------
	//   World Gen
	//   --------------------------------------------------

	engine.Config.Logger.Info("Generating world...")
	generateWorld()

	l = lighting.NewPointLight(
		engine.ShaderControl.GetShader("colorLighting"),
		[]float32{0, 0, 0},
		[]float32{0.9, 0.9, 0.9},
		[]float32{0, 0, 0},
		//1.0, -300, 299.8,
		1.0, -2, 1.9,
		//1.0, 0.5, 0.1,
	)
	l.SetPosition([]float32{0, 0, 1})

	//   --------------------------------------------------
	//   Instancing
	//   --------------------------------------------------

	engine.InstanceLight(&l)

	engine.Instance(&SkyChild)
	engine.Instance(&CloudChild)
	engine.Instance(&NoCollisionChild)
	engine.Instance(&NatureChild)
	engine.Instance(&WorldChild)
	engine.Instance(&Player)

	engine.Instance(&BlockSelect)

	engine.EnableLighting()
	engine.Initialize()
	engine.StartRenderer()
	<-engine.Done()
	return
}

func render(renderer *cmd.Renderer, inputs *input.Input) {
	// Render Children
	renderer.RenderChild(&SkyChild)
	renderer.RenderChildCopies(&CloudChild)

	renderWorldInBounds(renderer)

	renderer.RenderChild(&Player)

	// Block Selector
	renderer.RenderChild(&BlockSelect)
	cx, cy, _ := renderer.MainCamera.GetPosition()
	bx, by := engine.CollisionControl.ScaleMouseCoords(inputs.MouseX, inputs.MouseY, cx, cy)
	snapx, snapy := int(bx/BlockSize), int(-by/BlockSize)
	BlockSelect.SetPosition(float32(snapx*BlockSize), float32(snapy*BlockSize))

	if inputs.LeftMouseButton {
		destroyBlock(snapx, snapy)
	}

	if inputs.RightMouseButton {
		placeTorch(snapx, snapy)
	}

	// Player Logic
	movePlayer(inputs.Keys)
	Player.VY -= 1.1

	top, left, bottom, right := CheckPlayerCollision()
	if bottom && Player.VY < 0 {
		Player.VY = 0
	}
	if left && Player.VX > 1 {
		Player.VX = -0.1
	}
	if right && Player.VX < 1 {
		Player.VX = 0.1
	}
	if top && Player.VY > 0 {
		Player.VY = 0
	}

	// Camera
	renderer.MainCamera.SetPosition(Player.X, Player.Y, -10)
	SkyChild.SetPosition(Player.X-1950/2, Player.Y-1110/2)

	// Lighting
	//x, y := child.ScaleCoordinates(Player.X, float32(HeightMap[int(Player.X/BlockSize)])*BlockSize, float32(ScreenWidth), float32(ScreenHeight))
	//l.SetPosition([]float32{x, y, 1})
	x, y := child.ScaleCoordinates(Player.X, Player.Y, float32(ScreenWidth), float32(ScreenHeight))
	l.SetPosition([]float32{x, y, 1})
}

func renderWorldInBounds(renderer *cmd.Renderer) {
	for x := int(Player.X) - 50 - ScreenWidth/2; x < int(Player.X)+50+ScreenWidth/2; x += BlockSize {
		for y := int(Player.Y) - 50 - ScreenHeight/2; y < int(Player.Y)+50+ScreenHeight/2; y += BlockSize {
			if cpy := NoCollisionCopies[int(x/BlockSize)][int(y/BlockSize)]; cpy.ID != 0 {
				renderer.RenderCopy(&NoCollisionChild, cpy)
			}
			if cpy := NatureCopies[int(x/BlockSize)][int(y/BlockSize)]; cpy.ID != 0 {
				renderer.RenderCopy(&NatureChild, cpy)
			}
			if cpy := WorldCopies[int(x/BlockSize)][int(y/BlockSize)]; cpy.ID != 0 {
				renderer.RenderCopy(&WorldChild, cpy)
			}
		}
	}
}

func movePlayer(keys map[string]bool) {
	if keys["w"] {
		Player.VY = 20
	}
	if keys["a"] {
		Player.VX = 10
	} else if keys["d"] {
		Player.VX = -10
	} else {
		Player.VX = 0
	}
}

// top, left, bottom, right
func CheckPlayerCollision() (bool, bool, bool, bool) {
	top := false
	left := false
	bottom := false
	right := false

	px := int((Player.X + BlockSize/2) / BlockSize)
	py := int((Player.Y)/BlockSize + 1)

	if WorldCopies[px][py+1].ID != 0 {
		top = true
	}
	if WorldCopies[px][py-1].ID != 0 {
		bottom = true
	}
	if WorldCopies[px-1][py].ID != 0 || WorldCopies[px-1][py+1].ID != 0 {
		left = true
	}
	if WorldCopies[px+1][py].ID != 0 || WorldCopies[px+1][py+1].ID != 0 {
		right = true
	}

	return top, left, bottom, right
}

func placeTorch(x, y int) {
	if WorldMap[x][y].ID != NameMap["sky"] && WorldMap[x][y].ID != NameMap["backdirt"] {
		return
	}

	WorldMap[x][y] = NewBlock("torch")
	WorldMap[x][y].Darkness = 0.8
	createSingleCopy(x, y)
	createNewLightSource(x, y)
}

func destroyTorch(x, y int) {
	destroyBlock(x, y)
	RemoveLightingLimit(x, y, 0.9, 50)
}

func createNewLightSource(x, y int) {
	CreateLightingLimit(x, y, 0.9, 40)
}

func destroyBlock(x, y int) {
	if WorldMap[x][y].ID == NameMap["sky"] || WorldMap[x][y].ID == NameMap["backdirt"] {
		return
	}

	WorldCopies[x][y] = child.ChildCopy{
		ID: 0,
	}
	NoCollisionCopies[x][y] = child.ChildCopy{
		ID: 0,
	}

	if y <= HeightMap[x] {
		WorldMap[x][y] = NewBlock("backdirt")
	} else {
		WorldMap[x][y] = NewBlock("sky")
	}

	FixLightingAt(x, y)

	fixBlock(x, y)

	fixBlock(x+1, y)
	fixBlock(x, y+1)
	fixBlock(x-1, y)
	fixBlock(x, y-1)
}

func fixBlock(x, y int) {
	orientSingleBlock(NameList[WorldMap[x][y].ID], WorldMap[x][y].ID, true, x, y)
	createSingleExtraBackdirt(x, y)
	createSingleCopy(x, y)
}
