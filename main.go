package main

import (
	_ "net/http/pprof"
	"rapidengine/child"
	"rapidengine/cmd"
	"rapidengine/geometry"
	"rapidengine/input"
	"rapidengine/lighting"
	"rapidengine/material"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if runtime.GOOS == "darwin" {
		ScreenWidth = 1440
		ScreenHeight = 900
	}

	Config = cmd.NewEngineConfig(ScreenWidth, ScreenHeight, 2)
	Config.ShowFPS = false
	Config.FullScreen = true
	Config.GammaCorrection = false
	Config.VSync = false
	Config.Profiling = false
	Engine = cmd.NewEngine(&Config, render)
	Engine.Renderer.SetRenderDistance(float32(ScreenWidth/2) + 50)
	Engine.Renderer.MainCamera.SetPosition(100, 100, 0)
	Engine.Renderer.MainCamera.SetSpeed(0.2)

	Engine.Renderer.AutomaticRendering = false

	//   --------------------------------------------------
	//   Textures
	//   --------------------------------------------------

	Engine.TextureControl.NewTexture("assets/player/player.png", "player", "pixel")
	Engine.TextureControl.NewTexture("assets/backgrounds/gradient.png", "back", "pixel")

	//   --------------------------------------------------
	//   Materials
	//   --------------------------------------------------

	playerMaterial := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	playerMaterial.BecomeTexture(Engine.TextureControl.GetTexture("player"))

	backgroundMaterial := material.NewMaterial(Engine.ShaderControl.GetShader("colorLighting"), &Config)
	backgroundMaterial.BecomeTexture(Engine.TextureControl.GetTexture("back"))

	//   --------------------------------------------------
	//   Children
	//   --------------------------------------------------

	Player = Engine.NewChild2D()
	Player.AttachShader(Engine.ShaderControl.GetShader("colorLighting"))
	Player.AttachMesh(geometry.NewRectangle(32, 64, &Config))
	Player.AttachMaterial(&playerMaterial)
	Player.SetPosition(3000, 1000*BlockSize)
	Player.Gravity = 0

	SkyChild = Engine.NewChild2D()
	SkyChild.AttachShader(Engine.ShaderControl.GetShader("colorLighting"))
	SkyChild.AttachMesh(geometry.NewRectangle(4000, 1100, &Config))
	SkyChild.AttachMaterial(&backgroundMaterial)

	BlockSelect = Engine.NewChild2D()
	BlockSelect.AttachShader(Engine.ShaderControl.GetShader("color"))

	m := material.NewMaterial(Engine.ShaderControl.GetShader("color"), &Config)
	m.BecomeColor([3]float32{200, 200, 200})

	BlockSelect.AttachMaterial(&m)
	BlockSelect.AttachMesh(geometry.NewRectangle(32, 32, &Config))

	//   --------------------------------------------------
	//   World Gen
	//   --------------------------------------------------

	Engine.Logger.Info("Generating world...")
	generateWorldTree()
	Engine.Logger.Info("World Complete")

	l = lighting.NewPointLight(
		Engine.ShaderControl.GetShader("colorLighting"),
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

	Engine.InstanceLight(&l)

	Engine.Instance(&SkyChild)
	Engine.Instance(&CloudChild)
	Engine.Instance(&NoCollisionChild)
	Engine.Instance(&NatureChild)
	Engine.Instance(&WorldChild)
	Engine.Instance(&Player)

	Engine.Instance(&BlockSelect)

	Engine.EnableLighting()
	Engine.Initialize()
	Engine.StartRenderer()
	<-Engine.Done()
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
	bx, by := Engine.CollisionControl.ScaleMouseCoords(inputs.MouseX, inputs.MouseY, cx, cy)
	snapx, snapy := int(bx/BlockSize), int(-by/BlockSize)
	BlockSelect.SetPosition(float32(snapx*BlockSize), float32(snapy*BlockSize))

	if inputs.LeftMouseButton {
		destroyBlock(snapx, snapy)
	}

	if inputs.RightMouseButton {
		if WorldMap.GetWorldBlockID(snapx, snapy) == "00000" {
			placeBlock(snapx, snapy, "torch")
			CreateLightingLimit(snapx, snapy, 0.75, 15)
		}
	}

	// Player Logic
	movePlayer(inputs.Keys)
	Player.VY -= 1.05

	top, left, bottom, right := CheckPlayerCollision()
	if bottom && Player.VY < 0 {
		Player.VY = 0
	}
	if left && Player.VX > 1 {
		Player.VX = 0
	}
	if right && Player.VX < 1 {
		Player.VX = 0
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
			if cpy := WorldMap.GetBackBlock(int(x/BlockSize), int(y/BlockSize)); cpy.ID != "00000" {
				renderer.RenderCopy(&NoCollisionChild, *cpy)
			}
			if cpy := WorldMap.GetNatureBlock(int(x/BlockSize), int(y/BlockSize)); cpy.ID != "00000" {
				renderer.RenderCopy(&NatureChild, *cpy)
			}
			if cpy := WorldMap.GetWorldBlock(int(x/BlockSize), int(y/BlockSize)); cpy.ID != "00000" {
				renderer.RenderCopy(&WorldChild, *cpy)
			}
			if cpy := WorldMap.GetLightBlock(int(x/BlockSize), int(y/BlockSize)); cpy.ID != "00000" {
				renderer.RenderCopy(&NoCollisionChild, *cpy)
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

	if WorldMap.GetWorldBlockID(px, py+1) != "00000" {
		top = true
	}
	if WorldMap.GetWorldBlockID(px, py-1) != "00000" {
		bottom = true
	}
	if WorldMap.GetWorldBlockID(px-1, py) != "00000" || WorldMap.GetWorldBlockID(px-1, py+1) != "00000" {
		left = true
	}
	if WorldMap.GetWorldBlockID(px+1, py) != "00000" || WorldMap.GetWorldBlockID(px+1, py+1) != "00000" {
		right = true
	}

	return top, left, bottom, right
}
