package main

import (
	"math"
	_ "net/http/pprof"
	"rapidengine/child"
	"rapidengine/cmd"
	"rapidengine/geometry"
	"rapidengine/input"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

var QUALITY = "LOW" // "EPIC" // "HIGH" // "MEDIUM" // "LOW"

var colChild *child.Child2D

func main() {
	if runtime.GOOS == "darwin" {
		ScreenWidth = 1440
		ScreenHeight = 900
	}
	if QUALITY == "EPIC" {
		ScreenWidth = 3840
		ScreenHeight = 2160
	}

	//ScreenWidth = 3840
	//ScreenHeight = 2160

	Config = cmd.NewEngineConfig(ScreenWidth, ScreenHeight, 2)

	Config.ShowFPS = true
	Config.FullScreen = true
	Config.GammaCorrection = false
	Config.AntiAliasing = true
	Config.VSync = false
	Config.Profiling = false
	Config.Blending = false

	Engine = cmd.NewEngine(&Config, render)

	Engine.Renderer.SetRenderDistance(float32(ScreenWidth/2) + 50)
	Engine.Renderer.MainCamera.SetSpeed(0.2)
	Engine.Renderer.MainCamera.SetSmoothSpeed(0.05)

	Engine.TextControl.LoadFont("./assets/vermin.ttf", "pixel", 32, 15)

	InitializeHitboxViewer()
	V.Mat.Hue = [4]float32{200, 100, 0, 255}
	V.Mat.DiffuseLevel = 0

	InitializeLoadingScene()
	InitializeWorldScene()
	InitializeMenuScene()
	InitializeSaveScene()
	InitializeHotbarScene()
	InitializeChooseScene()
	InitializeTitleScene()
	InitializeRespawnScene()

	EM = InitializeEnemyManager()

	Engine.SceneControl.InstanceScene(TitleScene)
	Engine.SceneControl.InstanceScene(ChooseScene)
	Engine.SceneControl.InstanceScene(LoadingScene)
	Engine.SceneControl.InstanceScene(WorldScene)
	Engine.SceneControl.InstanceScene(SaveScene)

	WorldScene.InstanceSubscene(MenuScene)
	WorldScene.InstanceSubscene(HotbarScene)
	WorldScene.InstanceSubscene(RespawnScene)

	Engine.SceneControl.SetCurrentScene(TitleScene)
	HotbarScene.Activate()

	if QUALITY == "HIGH" || QUALITY == "MEDIUM" || QUALITY == "EPIC" {
		Engine.PostControl.EnablePostProcessing()

		if QUALITY == "HIGH" || QUALITY == "EPIC" {
			Engine.PostControl.EnableLightScattering(SunChild)
			Engine.PostControl.EnableBloom(50, 8)

			Engine.PostControl.BloomThreshold = -0.35
			Engine.PostControl.BloomIntensity = 0.73
		} else {
			Engine.PostControl.EnableBloom(10, 4)
			Engine.PostControl.BloomThreshold = 0.64
			Engine.PostControl.BloomIntensity = 0.57
		}

		Engine.PostControl.BloomIntensity = 1.5
		Engine.PostControl.BloomThreshold = -2.6
	}

	GamePaused = false

	colChild = Engine.ChildControl.NewChild2D()
	colChild.AttachMesh(geometry.NewRectangle())
	colChild.AttachMaterial(Engine.MaterialControl.NewBasicMaterial())

	Engine.Initialize()
	Engine.StartRenderer()
	<-Engine.Done()
	return
}

var JustEnemy = false
var JustKnock = false

var ViewerEnabled = false

func render(renderer *cmd.Renderer, inputs *input.Input) {
	if inputs.Keys["e"] {
		if !JustEnemy {
			EM.NewGoblin(50)
			JustEnemy = true
		}
	} else {
		JustEnemy = false
	}

	if inputs.Keys["l"] {
		renderer.MainCamera.Shake(0.3, 0.01)
	}

	//println(Engine.PostControl.BloomOffsetX, Engine.PostControl.BloomOffsetY)

	if inputs.Keys["h"] {
		ChangeParallax()
	} else {
		JustParallax = false
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

	if Engine.SceneControl.GetCurrentScene().ID == "title" || Engine.SceneControl.GetCurrentScene().ID == "choose" {
		updateTitleScreen()
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
	//renderer.RenderChildCopies(CloudChild)

	renderer.RenderChild(SunChild)

	renderer.RenderChild(Back6Child)
	renderer.RenderChild(Back5Child)
	renderer.RenderChild(Back4Child)
	renderer.RenderChild(Back3Child)
	renderer.RenderChild(Back2Child)
	renderer.RenderChild(Back1Child)

	renderWorldInBounds(renderer)

	//renderer.RenderChild(colChild)
	renderer.RenderChild(Player1.PlayerChild)

	// Update and render enemies
	EM.Update()

	renderFrontWorldInBounds(renderer)

	if inputs.Keys["escape"] && !GamePaused {
		GamePaused = true
		MenuScene.Activate()
	}

	if Player1.Dead {
		RespawnScene.Activate()
	}

	if !GamePaused {
		// Update player
		Player1.Update(inputs)

		cx, cy, _ := renderer.MainCamera.GetPosition()
		bx, by := Engine.CollisionControl.ScaleMouseCoords(inputs.MouseX, inputs.MouseY, cx, cy)
		snapx, snapy := int(bx/BlockSize), int(-by/BlockSize)
		BlockSelect.SetPosition(float32(snapx*BlockSize), float32(snapy*BlockSize))

		blockDist := BlockDistance(float32(snapx*BlockSize), float32(snapy*BlockSize), Player1.CenterX, Player1.CenterY)
		if blockDist < 5 {
			renderer.RenderChild(BlockSelect)
		}

		Player1.PlayerChild.Darkness = WorldMap.GetDarkness(
			int(Player1.CenterX/BlockSize),
			int(Player1.CenterY/BlockSize)+1,
		)

		if Player1.CurrentMiningTimer <= 0 {
			Player1.Lastsnapx, Player1.Lastsnapy = snapx, snapy
		}

		if inputs.LeftMouseButton && blockDist < 5 {
			if Player1.Lastsnapx == snapx && Player1.Lastsnapy == snapy {
				Player1.CurrentMiningTimer += renderer.DeltaFrameTime
			}
			if Player1.CurrentMiningTimer > 1 { //GetBlock(WorldMap.GetBackBlockName(snapx, snapy)).Durability {
				destroyBlock(snapx, snapy)
				Player1.CurrentMiningTimer = 0
				Player1.Lastsnapx, Player1.Lastsnapy = snapx, snapy
			}
		} else {
			Player1.CurrentMiningTimer = 0
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
		renderer.MainCamera.SetPosition(Player1.CenterX, Player1.CenterY, -10)
		SkyChild.SetPosition(cx-float32(ScreenWidth/2), cy-float32(ScreenHeight/2))

		Back1Child.Y = cy - float32(ScreenHeight/2)
		Back2Child.Y = cy - float32(ScreenHeight/2)
		Back3Child.Y = cy - float32(ScreenHeight/2)
		Back4Child.Y = cy - float32(ScreenHeight/2)
		Back5Child.Y = cy - float32(ScreenHeight/2)
		Back6Child.Y = cy - float32(ScreenHeight/2)

		// Parallax: Higher divisor = faster movement = appears closer
		Back1Child.X = (cx / (WorldWidth * BlockSize / 10000)) / 0.8
		Back2Child.X = (cx / (WorldWidth * BlockSize / 10000)) / 0.6
		Back3Child.X = (cx / (WorldWidth * BlockSize / 10000)) / 0.3
		Back4Child.X = (cx / (WorldWidth * BlockSize / 10000)) / 0.2
		Back5Child.X = (cx / (WorldWidth * BlockSize / 10000)) / 0.1
		Back6Child.X = (cx / (WorldWidth * BlockSize / 10000)) / 0.05
	}

	if ViewerEnabled {
		V.Update()
		V.Render()
	}

	if inputs.Keys["v"] {
		if ViewerEnabled {
			ViewerEnabled = false
		} else {
			ViewerEnabled = true
		}
	}
}

func renderWorldInBounds(renderer *cmd.Renderer) {
	for x := int(Player1.CenterX) - 50 - ScreenWidth/2; x < int(Player1.CenterX)+50+ScreenWidth/2; x += BlockSize {
		for y := int(Player1.CenterY) - 50 - ScreenHeight/2; y < int(Player1.CenterY)+50+ScreenHeight/2; y += BlockSize {
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
	for x := int(Player1.CenterX) - 50 - ScreenWidth/2; x < int(Player1.CenterX)+50+ScreenWidth/2; x += BlockSize {
		for y := int(Player1.CenterY) - 50 - ScreenHeight/2; y < int(Player1.CenterY)+50+ScreenHeight/2; y += BlockSize {
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
