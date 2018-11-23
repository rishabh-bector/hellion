package main

import (
	"fmt"
	"rapidengine/child"
)

var MenuBackChild *child.Child2D

func InitializeMenuScene() {
	MenuScene = Engine.SceneControl.NewScene("menu")

	MenuBackChild = Engine.ChildControl.NewChild2D()
	MenuBackChild.ScaleX = 200
	MenuBackChild.ScaleY = 300
	MenuBackChild.Static = true
	MenuBackChild.SetPosition(620, 250)

	resumeText := Engine.TextControl.NewTextBox("Resume", "pixel", 100, 500, 1, [3]float32{255, 255, 255})
	saveText := Engine.TextControl.NewTextBox("Save", "pixel", 100, 400, 1, [3]float32{255, 255, 255})
	exitText := Engine.TextControl.NewTextBox("Exit", "pixel", 100, 300, 1, [3]float32{255, 255, 255})

	resumeButton := Engine.UIControl.NewUIButton(100, 500, 200, 50)
	resumeButton.SetClickCallback(resume)
	resumeButton.AttachText(resumeText)
	resumeButton.ButtonChild.AttachMaterial(ButtonMaterial)

	saveButton := Engine.UIControl.NewUIButton(100, 400, 200, 50)
	saveButton.SetClickCallback(save)
	saveButton.AttachText(saveText)
	saveButton.ButtonChild.AttachMaterial(ButtonMaterial)

	exitButton := Engine.UIControl.NewUIButton(100, 300, 200, 50)
	exitButton.SetClickCallback(exitToTitle)
	exitButton.AttachText(exitText)
	exitButton.ButtonChild.AttachMaterial(ButtonMaterial)

	Engine.UIControl.AlignCenter(resumeButton)
	Engine.UIControl.AlignCenter(saveButton)
	Engine.UIControl.AlignCenter(exitButton)

	Engine.UIControl.InstanceElement(resumeButton, MenuScene)
	Engine.UIControl.InstanceElement(saveButton, MenuScene)
	Engine.UIControl.InstanceElement(exitButton, MenuScene)
	MenuScene.InstanceChild(MenuBackChild)

	MenuScene.Deactivate()
}

func resume() {
	GamePaused = false
	MenuScene.Deactivate()
}

func save() {
	Engine.SceneControl.SetCurrentScene(SaveScene)
	WorldMap.WriteToFile("./worlds/world" + fmt.Sprint(CurrentWorld) + ".hln")
	Engine.SceneControl.SetCurrentScene(WorldScene)
	MenuScene.Activate()
}

func exitToTitle() {
	exitButton.Block()
	Engine.SceneControl.SetCurrentScene(TitleScene)
}
