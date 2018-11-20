package main

import (
	_ "net/http/pprof"
	"rapidengine/cmd"
	"rapidengine/input"
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
	Config.FullScreen = false
	Config.GammaCorrection = false
	Config.VSync = false
	Config.Profiling = false

	Engine = cmd.NewEngine(&Config, render)

	Engine.Renderer.SetRenderDistance(float32(ScreenWidth/2) + 50)
	Engine.Renderer.MainCamera.SetSpeed(0.2)

	InitializeTitleScene()
	InitializeLoadingScene()
	InitializeWorldScene()

	Engine.ChildControl.SetScene("title")

	Engine.Initialize()
	Engine.StartRenderer()
	<-Engine.Done()
	return
}

func render(renderer *cmd.Renderer, inputs *input.Input) {
	if Engine.ChildControl.GetScene() == "world" {
		renderWorldScene(renderer, inputs)
	}
}

func renderWorldScene(renderer *cmd.Renderer, inputs *input.Input) {
	// Update player
	Player1.Update(inputs)

	// Render Children
	renderer.RenderChild(SkyChild)
	renderer.RenderChildCopies(CloudChild)

	renderWorldInBounds(renderer)

	renderer.RenderChild(Player1.PlayerChild)

	// Block Selector
	renderer.RenderChild(BlockSelect)
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
			CreateLightingLimit(snapx, snapy, 0.65, 15)
		}
	}

	// Camera
	renderer.MainCamera.SetPosition(Player1.PlayerChild.X, Player1.PlayerChild.Y, -10)
	SkyChild.SetPosition(Player1.PlayerChild.X-float32(ScreenWidth/2), Player1.PlayerChild.Y-float32(ScreenHeight/2))
}

func renderWorldInBounds(renderer *cmd.Renderer) {
	for x := int(Player1.PlayerChild.X) - 50 - ScreenWidth/2; x < int(Player1.PlayerChild.X)+50+ScreenWidth/2; x += BlockSize {
		for y := int(Player1.PlayerChild.Y) - 50 - ScreenHeight/2; y < int(Player1.PlayerChild.Y)+50+ScreenHeight/2; y += BlockSize {
			if cpy := WorldMap.GetBackBlock(int(x/BlockSize), int(y/BlockSize)); cpy.ID != "00000" {
				renderer.RenderCopy(NoCollisionChild, *cpy)
			}
			if cpy := WorldMap.GetNatureBlock(int(x/BlockSize), int(y/BlockSize)); cpy.ID != "00000" {
				renderer.RenderCopy(NatureChild, *cpy)
			}
			if cpy := WorldMap.GetWorldBlock(int(x/BlockSize), int(y/BlockSize)); cpy.ID != "00000" {
				renderer.RenderCopy(WorldChild, *cpy)
			}
			if cpy := WorldMap.GetLightBlock(int(x/BlockSize), int(y/BlockSize)); cpy.ID != "00000" {
				renderer.RenderCopy(NoCollisionChild, *cpy)
			}
		}
	}
}
