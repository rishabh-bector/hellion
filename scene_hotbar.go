package main

var ActiveItem int
var HotBarItems []string

func InitializeHotbarScene() {
	HotbarScene = Engine.SceneControl.NewScene("hotbar")

	HotBarItems = []string{
		"dirt",
		"stone",
	}
	ActiveItem = 0

	HotbarScene.Deactivate()
}
