package main

import (
	"rapidengine/child"
	"rapidengine/geometry"
	"rapidengine/material"
)

const NumSlots = 5
const SlotSize = 50
const SlotSpacing = 25

var ActiveItem int
var HotBarItems [NumSlots]string

var BarChildren [NumSlots]*child.Child2D
var BarMats [NumSlots]*material.BasicMaterial

var ActiveChild *child.Child2D

func InitializeHotbarScene() {
	HotbarScene = Engine.SceneControl.NewScene("hotbar")

	Engine.TextureControl.NewTexture("./assets/selector.png", "selector", "pixel")
	activeMat := Engine.MaterialControl.NewBasicMaterial()
	activeMat.DiffuseLevel = 1
	activeMat.DiffuseMap = Engine.TextureControl.GetTexture("selector")

	ActiveChild = Engine.ChildControl.NewChild2D()
	ActiveChild.AttachMesh(geometry.NewRectangle())
	ActiveChild.AttachMaterial(activeMat)
	ActiveChild.ScaleX = SlotSize + 10
	ActiveChild.ScaleY = SlotSize + 10
	ActiveChild.Static = true

	HotbarScene.InstanceChild(ActiveChild)

	for i := 0; i < NumSlots; i++ {
		BarMats[i] = Engine.MaterialControl.NewBasicMaterial()
		BarMats[i].DiffuseLevel = 1
		BarMats[i].DiffuseMap = Engine.TextureControl.GetTexture("cloud1")
		BarMats[i].DiffuseMapScale = 1

		BarChildren[i] = Engine.ChildControl.NewChild2D()
		BarChildren[i].AttachMesh(geometry.NewRectangle())
		BarChildren[i].AttachMaterial(BarMats[i])
		BarChildren[i].ScaleX = SlotSize
		BarChildren[i].ScaleY = SlotSize
		BarChildren[i].Static = true
		BarChildren[i].SetPosition(float32(ScreenWidth)-120, 500-float32(i*(SlotSize+SlotSpacing)))

		HotbarScene.InstanceChild(BarChildren[i])
	}

	HotBarItems = [NumSlots]string{
		"dirt",
		"stone",
		"torch",
		"stoneBrick",
		"dirt",
	}
	ActiveItem = 0

	Player1Health = Engine.UIControl.NewProgressBar()
	Player1Health.BackChild.Static = true
	Player1Health.BarChild.Static = true
	Player1Health.SetDimensions(400, 25)
	Player1Health.SetPosition(50, 0.9*float32(Engine.Config.ScreenHeight))
	Engine.UIControl.InstanceElement(Player1Health, HotbarScene)

	UpdateHotBar()
	HotbarScene.Deactivate()
}

func UpdateHotBar() {
	for i := 0; i < NumSlots; i++ {
		if HotBarItems[i] != "" {
			BarMats[i].DiffuseMap = GetBlock(HotBarItems[i]).GetMaterial("NN").DiffuseMap
		}
	}
}

func UpdateActiveItem() {
	ActiveChild.SetPosition(1800-5, 500-5-float32(ActiveItem*(SlotSize+SlotSpacing)))

	// UI Updates
	Player1Health.SetPercentage(Player1.Health / Player1.MaxHealth * 100)
	Player1Health.SetPosition(50, 0.9*float32(Engine.Config.ScreenHeight))
}
