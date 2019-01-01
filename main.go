package main

import (
	"math"
	_ "net/http/pprof"
	"rapidengine/cmd"
	"rapidengine/input"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

var QUALITY = "LOW" // "MEDIUM" // "LOW"

func main() {
	if runtime.GOOS == "darwin" {
		ScreenWidth = 1440
		ScreenHeight = 900
	}

	//ScreenWidth = 3840
	//ScreenHeight = 2160

	Config = cmd.NewEngineConfig(ScreenWidth, ScreenHeight, 2)

	Config.ShowFPS = true
	Config.FullScreen = false
	Config.GammaCorrection = false
	Config.VSync = false
	Config.Profiling = false

	Engine = cmd.NewEngine(&Config, render)

	Engine.Renderer.SetRenderDistance(float32(ScreenWidth/2) + 50)
	Engine.Renderer.MainCamera.SetSpeed(0.2)

	InitializeTitleScene()
	InitializeChooseScene()
	InitializeLoadingScene()
	InitializeWorldScene()
	InitializeMenuScene()
	InitializeSaveScene()
	InitializeHotbarScene()

	EM = InitializeEnemyManager()

	Engine.SceneControl.InstanceScene(TitleScene)
	Engine.SceneControl.InstanceScene(ChooseScene)
	Engine.SceneControl.InstanceScene(LoadingScene)
	Engine.SceneControl.InstanceScene(WorldScene)
	Engine.SceneControl.InstanceScene(SaveScene)

	WorldScene.InstanceSubscene(MenuScene)
	WorldScene.InstanceSubscene(HotbarScene)

	TitleScene.InstanceSubscene(HotbarScene)

	Engine.SceneControl.SetCurrentScene(TitleScene)
	HotbarScene.Activate()

	if QUALITY == "HIGH" || QUALITY == "MEDIUM" {
		Engine.PostControl.EnablePostProcessing()
		Engine.PostControl.EnableLightScattering(SunChild)

		if QUALITY == "HIGH" {
			Engine.PostControl.EnableBloom(25, 4)
		} else {
			Engine.PostControl.EnableBloom(10, 4)
		}

		Engine.PostControl.BloomIntensity = 1.68
		Engine.PostControl.BloomThreshold = 0.55

		//Engine.PostControl.BloomOffsetX = -12
		Engine.PostControl.BloomOffsetX = -10
		Engine.PostControl.BloomOffsetY = -12
	}

	GamePaused = false

	Engine.Initialize()
	Engine.StartRenderer()
	<-Engine.Done()
	return
}

var JustEnemy = false
var JustKnock = false

func render(renderer *cmd.Renderer, inputs *input.Input) {
	if inputs.Keys["e"] {
		if !JustEnemy {
			EM.NewGoblin(50)
			JustEnemy = true
		}
	} else {
		JustEnemy = false
	}

	if inputs.Keys["q"] {
		if !JustKnock {
			for _, e := range EM.AllEnemies {
				e.Damage(2)
			}
			JustKnock = true
		}
	} else {
		JustKnock = false
	}

	if Engine.SceneControl.GetCurrentScene().ID == "world" {
		renderWorldScene(renderer, inputs)
	}

	if MenuScene.IsActive() {
		renderer.RenderChild(MenuBackChild)
	}

	if WorldScene.IsActive() {
		renderer.RenderChild(ActiveChild)
		for i := 0; i < NumSlots; i++ {
			renderer.RenderChild(BarChildren[i])
		}
	}

	ActiveItem = int(math.Abs(inputs.Scroll*MouseSensitivity)) % NumSlots
	UpdateActiveItem()

	if inputs.Keys["b"] {
		Engine.PostControl.BloomIntensity += 0.01
		println(Engine.PostControl.BloomIntensity)
	}
	if inputs.Keys["v"] {
		Engine.PostControl.BloomIntensity -= 0.01
		println(Engine.PostControl.BloomIntensity)
	}

	if inputs.Keys["c"] {
		Engine.PostControl.BloomThreshold += 0.01
		println(Engine.PostControl.BloomThreshold)
	}
	if inputs.Keys["x"] {
		Engine.PostControl.BloomThreshold -= 0.01
		println(Engine.PostControl.BloomThreshold)
	}

	if inputs.Keys["up"] {
		Engine.PostControl.BloomOffsetY++
		Player1.Gravity++
	}
	if inputs.Keys["down"] {
		Engine.PostControl.BloomOffsetY--
		Player1.Gravity--
	}
	if inputs.Keys["left"] {
		Engine.PostControl.BloomOffsetX++
	}
	if inputs.Keys["right"] {
		Engine.PostControl.BloomOffsetX--
	}
	//println(Engine.PostControl.BloomOffsetX, Engine.PostControl.BloomOffsetY)
}

func renderWorldScene(renderer *cmd.Renderer, inputs *input.Input) {
	// Render Children
	renderer.RenderChild(SkyChild)
	renderer.RenderChildCopies(CloudChild)

	renderer.RenderChild(SunChild)

	renderer.RenderChild(Back4Child)
	renderer.RenderChild(Back3Child)
	renderer.RenderChild(Back2Child)
	renderer.RenderChild(Back1Child)

	renderWorldInBounds(renderer)

	renderer.RenderChild(Player1.PlayerChild)

	// Update and render enemies
	EM.Update()

	renderFrontWorldInBounds(renderer)

	if inputs.Keys["escape"] && !GamePaused {
		GamePaused = true
		MenuScene.Activate()
	}
	if !GamePaused {
		// Update player
		Player1.Update(inputs)

		cx, cy, _ := renderer.MainCamera.GetPosition()
		bx, by := Engine.CollisionControl.ScaleMouseCoords(inputs.MouseX, inputs.MouseY, cx, cy)
		snapx, snapy := int(bx/BlockSize), int(-by/BlockSize)
		BlockSelect.SetPosition(float32(snapx*BlockSize), float32(snapy*BlockSize))

		blockDist := BlockDistance(float32(snapx*BlockSize), float32(snapy*BlockSize), Player1.PlayerChild.X, Player1.PlayerChild.Y)
		if blockDist < 5 {
			renderer.RenderChild(BlockSelect)
		}

		Player1.PlayerChild.Darkness = WorldMap.GetDarkness(
			int(Player1.PlayerChild.X/BlockSize),
			int(Player1.PlayerChild.Y/BlockSize)+1,
		)

		if inputs.LeftMouseButton && blockDist < 5 {
			destroyBlock(snapx, snapy)
		}

		if inputs.RightMouseButton {
			if WorldMap.GetWorldBlockID(snapx, snapy) == "00000" {
				placeBlock(snapx, snapy, HotBarItems[ActiveItem])

				if HotBarItems[ActiveItem] == "torch" {
					CreateLightingLimit(snapx, snapy, 0.72, 18)
				}
			}
		}

		// Camera
		renderer.MainCamera.SetPosition(Player1.PlayerChild.X, Player1.PlayerChild.Y, -10)
		SkyChild.SetPosition(Player1.PlayerChild.X-float32(ScreenWidth/2), Player1.PlayerChild.Y-float32(ScreenHeight/2))

		Back1Child.Y = Player1.PlayerChild.Y - float32(ScreenHeight/2)
		Back2Child.Y = Player1.PlayerChild.Y - float32(ScreenHeight/2)
		Back3Child.Y = Player1.PlayerChild.Y - float32(ScreenHeight/2)
		Back4Child.Y = Player1.PlayerChild.Y - float32(ScreenHeight/2)

		// Parallax: Higher divisor = faster movement = appears closer
		Back1Child.X = (Player1.PlayerChild.X / (WorldWidth * BlockSize / 10000)) / 0.8
		Back2Child.X = (Player1.PlayerChild.X / (WorldWidth * BlockSize / 10000)) / 0.6
		Back3Child.X = (Player1.PlayerChild.X / (WorldWidth * BlockSize / 10000)) / 0.3
		Back4Child.X = (Player1.PlayerChild.X / (WorldWidth * BlockSize / 10000)) / 0.2

		Back1Child.Y = Back1Child.Y + (Player1.PlayerChild.Y/(WorldHeight*BlockSize/10))/0.8
		Back2Child.Y = Back2Child.Y + (Player1.PlayerChild.Y/(WorldHeight*BlockSize/10))/0.6
		Back3Child.Y = Back3Child.Y + (Player1.PlayerChild.Y/(WorldHeight*BlockSize/10))/0.3
		Back4Child.Y = Back4Child.Y + (Player1.PlayerChild.Y/(WorldHeight*BlockSize/10))/0.2
	}
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
			if cpy := WorldMap.GetGrassBlock(int(x/BlockSize), int(y/BlockSize)); cpy.ID != "00000" {
				renderer.RenderCopy(GrassChild, *cpy)
			}
			if cpy := WorldMap.GetLightBlock(int(x/BlockSize), int(y/BlockSize)); cpy.ID != "00000" {
				renderer.RenderCopy(NoCollisionChild, *cpy)
			}
		}
	}
}

func renderFrontWorldInBounds(renderer *cmd.Renderer) {
	for x := int(Player1.PlayerChild.X) - 50 - ScreenWidth/2; x < int(Player1.PlayerChild.X)+50+ScreenWidth/2; x += BlockSize {
		for y := int(Player1.PlayerChild.Y) - 50 - ScreenHeight/2; y < int(Player1.PlayerChild.Y)+50+ScreenHeight/2; y += BlockSize {
			if cpy := WorldMap.GetGrassBlock(int(x/BlockSize), int(y/BlockSize)); cpy.ID != "00000" {
				renderer.RenderCopy(GrassChild, *cpy)
			}
		}
	}
}

func BlockDistance(x1, y1, x2, y2 float32) int {
	dx := ((x1) - (x2)) * ((x1) - (x2))
	dy := ((y1) - (y2)) * ((y1) - (y2))
	return int(math.Sqrt(float64(dx+dy))) / BlockSize
}
