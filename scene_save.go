package main

import "rapidengine/ui"

var SaveProgressBar *ui.ProgressBar

func InitializeSaveScene() {
	SaveScene = Engine.SceneControl.NewScene("save")

	SaveProgressBar = Engine.UIControl.NewProgressBar()
	SaveProgressBar.SetDimensions(200, 25)
	SaveProgressBar.SetPosition(0, 500)
	SaveProgressBar.SetPercentage(0)
	Engine.UIControl.AlignCenter(SaveProgressBar)

	saveText := Engine.TextControl.NewTextBox("Saving World...", "pixel", 720, 600, 1, [3]float32{255, 255, 255})
	SaveScene.InstanceText(saveText)

	Engine.UIControl.InstanceElement(SaveProgressBar, SaveScene)

	SaveScene.Deactivate()
}

func updateSaveScreen() {
	SaveProgressBar.Update(nil)
	Engine.Renderer.ForceUpdate()
}
