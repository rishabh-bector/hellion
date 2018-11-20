package main

import (
	"rapidengine/ui"
)

var ProgressBar *ui.ProgressBar
var ProgressText *ui.TextBox

func InitializeLoadingScene() {
	Engine.ChildControl.NewScene("loading")

	ProgressBar = Engine.UIControl.NewProgressBar()
	ProgressBar.SetDimensions(500, 25)
	ProgressBar.SetPosition(0, 500)
	Engine.UIControl.AlignCenter(ProgressBar)

	ProgressText = Engine.TextControl.NewTextBox("Generating world...", "pixel", 720, 600, 1, [3]float32{255, 255, 255})
	Engine.TextControl.AddTextBox(ProgressText, "loading")

	Engine.UIControl.InstanceProgressBar(ProgressBar, "loading")
}

func updateLoadingScreen() {
	ProgressBar.Update(nil)
	Engine.Renderer.ForceUpdate()
}

func create() {
	Engine.ChildControl.SetScene("loading")
	Engine.Logger.Info("Generating world...")

	updateLoadingScreen()

	generateWorldTree()
}
