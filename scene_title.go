package main

import (
	"os"
	"rapidengine/ui"
)

var exitButton *ui.Button

var TitleParallax = float32(0)

func InitializeTitleScene() {
	TitleScene = Engine.SceneControl.NewScene("title")

	title := Engine.TextControl.NewTextBox("H E L L I O N", "pixel", (float32(ScreenWidth) / 2), 700, 5, [3]float32{217, 30, 24})
	title.X -= float32(title.GetLength() / 2)

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

	Engine.AudioControl.Load("./assets/music/Theme.wav", "title")
	//Engine.AudioControl.Play("title")

	Engine.UIControl.AlignCenter(playButton)
	Engine.UIControl.AlignCenter(settingsButton)
	Engine.UIControl.AlignCenter(exitButton)

	Engine.UIControl.InstanceElement(playButton, TitleScene)
	Engine.UIControl.InstanceElement(settingsButton, TitleScene)
	Engine.UIControl.InstanceElement(exitButton, TitleScene)

	TitleScene.InstanceChild(SkyChild)
	TitleScene.InstanceChild(SunChild)

	TitleScene.InstanceChild(Back6Child)
	TitleScene.InstanceChild(Back5Child)
	TitleScene.InstanceChild(Back4Child)
	TitleScene.InstanceChild(Back3Child)
	TitleScene.InstanceChild(Back2Child)
	TitleScene.InstanceChild(Back1Child)
}

func updateTitleScreen() {
	Back6Child.X = (TitleParallax / (WorldWidth * BlockSize / 10000)) / 1.4
	Back5Child.X = (TitleParallax / (WorldWidth * BlockSize / 10000)) / 1.2
	Back4Child.X = (TitleParallax / (WorldWidth * BlockSize / 10000)) / 1.0
	Back3Child.X = (TitleParallax / (WorldWidth * BlockSize / 10000)) / 0.6
	Back2Child.X = (TitleParallax / (WorldWidth * BlockSize / 10000)) / 0.3
	Back1Child.X = (TitleParallax / (WorldWidth * BlockSize / 10000)) / 0.2

	TitleParallax -= 5
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
