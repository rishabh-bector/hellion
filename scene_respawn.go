package main

func InitializeRespawnScene() {
	RespawnScene = Engine.SceneControl.NewScene("respawn")

	heading := Engine.TextControl.NewTextBox("YOU DIED", "pixel", (float32(ScreenWidth) / 2), 700, 5, [3]float32{217, 30, 24})
	RespawnScene.InstanceText(heading)
	respawnText := Engine.TextControl.NewTextBox("Respawn", "pixel", 100, 100, 1, [3]float32{255, 255, 255})
	respawnButton := Engine.UIControl.NewUIButton(100, 500, 200, 50)
	respawnButton.SetClickCallback(Player1.Respawn)
	respawnButton.AttachText(respawnText)
	respawnButton.ButtonChild.AttachMaterial(ButtonMaterial)
	Engine.UIControl.AlignCenter(respawnButton)
	Engine.UIControl.InstanceElement(respawnButton, RespawnScene)

	RespawnScene.Deactivate()
}
