package main

import (
	"os"
	"rapidengine/ui"
)

var exitButton *ui.Button

func InitializeTitleScene() {
	TitleScene = Engine.SceneControl.NewScene("title")

	Engine.TextControl.LoadFont("./assets/vermin.ttf", "pixel", 32, 15)

	title := Engine.TextControl.NewTextBox("H E L L I O N", "pixel", 720, 700, 3, [3]float32{217, 30, 24})
	TitleScene.InstanceText(title)

	ButtonMaterial = Engine.MaterialControl.NewBasicMaterial()
	ButtonMaterial.Hue = [4]float32{0, 0, 0, 255}

	playText := Engine.TextControl.NewTextBox("Play", "pixel", 100, 100, 1, [3]float32{255, 255, 255})
	settingsText := Engine.TextControl.NewTextBox("Settings", "pixel", 100, 100, 1, [3]float32{255, 255, 255})
	exitText := Engine.TextControl.NewTextBox("Exit", "pixel", 100, 100, 1, [3]float32{255, 255, 255})

	playButton := Engine.UIControl.NewUIButton(100, 500, 200, 50)
	playButton.SetClickCallback(play)
	playButton.AttachText(playText)
	playButton.ButtonChild.AttachMaterial(ButtonMaterial)

	settingsButton := Engine.UIControl.NewUIButton(100, 400, 200, 50)
	settingsButton.SetClickCallback(settings)
	settingsButton.AttachText(settingsText)
	settingsButton.ButtonChild.AttachMaterial(ButtonMaterial)

	exitButton = Engine.UIControl.NewUIButton(100, 300, 200, 50)
	exitButton.SetClickCallback(exit)
	exitButton.AttachText(exitText)
	exitButton.ButtonChild.AttachMaterial(ButtonMaterial)

	Engine.UIControl.AlignCenter(playButton)
	Engine.UIControl.AlignCenter(settingsButton)
	Engine.UIControl.AlignCenter(exitButton)

	Engine.UIControl.InstanceElement(playButton, TitleScene)
	Engine.UIControl.InstanceElement(settingsButton, TitleScene)
	Engine.UIControl.InstanceElement(exitButton, TitleScene)
}

func play() {
	Engine.SceneControl.SetCurrentScene(ChooseScene)
	updateChooseScene()
}

func settings() {

}

func exit() {
	os.Exit(0)
}
