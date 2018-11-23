package main

import (
	"rapidengine/ui"
)

var ProgressBar *ui.ProgressBar
var ProgressText *ui.TextBox

func InitializeLoadingScene() {
	LoadingScene = Engine.SceneControl.NewScene("loading")

	ProgressBar = Engine.UIControl.NewProgressBar()
	ProgressBar.SetDimensions(500, 25)
	ProgressBar.SetPosition(0, 500)
	Engine.UIControl.AlignCenter(ProgressBar)

	ProgressText = Engine.TextControl.NewTextBox("Generating world...", "pixel", 720, 600, 1, [3]float32{255, 255, 255})
	LoadingScene.InstanceText(ProgressText)

	Engine.UIControl.InstanceElement(ProgressBar, LoadingScene)
}

func updateLoadingScreen() {
	ProgressBar.Update(nil)
	Engine.Renderer.ForceUpdate()
}

func create() {
	Engine.SceneControl.SetCurrentScene(LoadingScene)
	Engine.Logger.Info("Generating world...")

	updateLoadingScreen()

	generateWorldTree()
}
