package main

import (
	"fmt"
	"os"
	"rapidengine/geometry"
	"rapidengine/ui"
)

var b1Text *ui.TextBox
var b2Text *ui.TextBox
var b3Text *ui.TextBox

func InitializeChooseScene() {
	ChooseScene = Engine.SceneControl.NewScene("choose")

	chooseText := Engine.TextControl.NewTextBox("Choose World", "pixel", 720, 600, 1, [3]float32{255, 255, 255})
	ChooseScene.InstanceText(chooseText)

	rect1 := Engine.ChildControl.NewChild2D()
	rect1.AttachMesh(geometry.NewRectangle())
	rect1.ScaleX = 800
	rect1.ScaleY = 75
	rect1.SetPosition(320, 400)

	rect2 := Engine.ChildControl.NewChild2D()
	rect2.AttachMesh(geometry.NewRectangle())
	rect2.ScaleX = 800
	rect2.ScaleY = 75
	rect2.SetPosition(320, 300)

	rect3 := Engine.ChildControl.NewChild2D()
	rect3.AttachMesh(geometry.NewRectangle())
	rect3.ScaleX = 800
	rect3.ScaleY = 75
	rect3.SetPosition(320, 200)

	b1Text = Engine.TextControl.NewTextBox("Play", "pixel", 720, 600, 1, [3]float32{255, 255, 255})
	b2Text = Engine.TextControl.NewTextBox("Play", "pixel", 720, 600, 1, [3]float32{255, 255, 255})
	b3Text = Engine.TextControl.NewTextBox("Play", "pixel", 720, 600, 1, [3]float32{255, 255, 255})

	c1Text := Engine.TextControl.NewTextBox("Save 1", "pixel", 200, 425, 1, [3]float32{255, 255, 255})
	c2Text := Engine.TextControl.NewTextBox("Save 2", "pixel", 200, 325, 1, [3]float32{255, 255, 255})
	c3Text := Engine.TextControl.NewTextBox("Save 3", "pixel", 200, 225, 1, [3]float32{255, 255, 255})

	ChooseScene.InstanceText(c1Text)
	ChooseScene.InstanceText(c2Text)
	ChooseScene.InstanceText(c3Text)

	b1 := Engine.UIControl.NewUIButton(1000, 425, 75, 25)
	b1.SetClickCallback(choose1)
	b1.AttachText(b1Text)

	b2 := Engine.UIControl.NewUIButton(1000, 325, 75, 25)
	b2.SetClickCallback(choose2)
	b2.AttachText(b2Text)

	b3 := Engine.UIControl.NewUIButton(1000, 225, 75, 25)
	b3.SetClickCallback(choose3)
	b3.AttachText(b3Text)

	Engine.UIControl.InstanceElement(b1, ChooseScene)
	Engine.UIControl.InstanceElement(b2, ChooseScene)
	Engine.UIControl.InstanceElement(b3, ChooseScene)

	ChooseScene.InstanceChild(SkyChild)
	ChooseScene.InstanceChild(SunChild)
	ChooseScene.InstanceChild(Back4Child)
	ChooseScene.InstanceChild(Back3Child)
	ChooseScene.InstanceChild(Back2Child)
	ChooseScene.InstanceChild(Back1Child)

	ChooseScene.InstanceChild(rect1)
	ChooseScene.InstanceChild(rect2)
	ChooseScene.InstanceChild(rect3)
}

func updateChooseScene() {
	if doesWorldExist(1) {
		b1Text.Text = "Play"
	} else {
		b1Text.Text = "New"
	}

	if doesWorldExist(2) {
		b2Text.Text = "Play"
	} else {
		b2Text.Text = "New"
	}

	if doesWorldExist(3) {
		b3Text.Text = "Play"
	} else {
		b3Text.Text = "New"
	}
}

func choose1() {
	CurrentWorld = 1
	if b1Text.Text == "Play" {
		loadWorld()
	} else {
		newWorld()
	}
}

func choose2() {
	CurrentWorld = 2
	if b2Text.Text == "Play" {
		loadWorld()
	} else {
		newWorld()
	}
}

func choose3() {
	CurrentWorld = 3
	if b3Text.Text == "Play" {
		loadWorld()
	} else {
		newWorld()
	}
}

func loadWorld() {
	ProgressText.Text = "Loading world..."
	ProgressBar.SetPercentage(0)
	Engine.SceneControl.SetCurrentScene(LoadingScene)
	updateLoadingScreen()

	initializeWorldTree()
	WorldMap.LoadFromFile("./worlds/world" + fmt.Sprint(CurrentWorld) + ".hln")

	Player1.PlayerChild.SetPosition(float32(WorldWidth*BlockSize/2), float32((HeightMap[WorldWidth/2]+25)*BlockSize))
	Engine.SceneControl.SetCurrentScene(WorldScene)
}

func newWorld() {
	ProgressText.Text = "Generating world..."
	Engine.SceneControl.SetCurrentScene(LoadingScene)
	generateWorldTree()
}

func doesWorldExist(world int) bool {
	if _, err := os.Stat("./worlds/world" + fmt.Sprint(world) + ".hln"); os.IsNotExist(err) {
		return false
	}
	return true
}
