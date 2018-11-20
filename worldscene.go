package main

import (
	"rapidengine/geometry"
	"rapidengine/material"
)

func InitializeWorldScene() {

	Engine.TextureControl.NewTexture("assets/backgrounds/gradient.png", "back", "pixel")
	backgroundMaterial := material.NewMaterial(Engine.ShaderControl.GetShader("texture"), &Config)
	backgroundMaterial.BecomeTexture(Engine.TextureControl.GetTexture("back"))

	InitializePlayer()

	SkyChild = Engine.ChildControl.NewChild2D()
	SkyChild.AttachMaterial(&backgroundMaterial)
	SkyChild.AttachMesh(geometry.NewRectangle())
	SkyChild.ScaleX = float32(ScreenWidth)
	SkyChild.ScaleY = float32(ScreenHeight)

	m := material.NewMaterial(Engine.ShaderControl.GetShader("color"), &Config)
	m.BecomeColor([3]float32{200, 200, 200})
	m.SetTransparency(0.2)

	BlockSelect = Engine.ChildControl.NewChild2D()
	BlockSelect.AttachMaterial(&m)
	BlockSelect.AttachMesh(geometry.NewRectangle())
	BlockSelect.ScaleX = 32
	BlockSelect.ScaleY = 32

	WorldChild = Engine.ChildControl.NewChild2D()
	WorldChild.AttachMesh(geometry.NewRectangle())
	WorldChild.ScaleX = BlockSize
	WorldChild.ScaleY = BlockSize
	WorldChild.EnableCopying()
	WorldChild.AttachCollider(0, 0, BlockSize, BlockSize)

	NoCollisionChild = Engine.ChildControl.NewChild2D()
	NoCollisionChild.AttachMesh(geometry.NewRectangle())
	NoCollisionChild.ScaleX = BlockSize
	NoCollisionChild.ScaleY = BlockSize
	NoCollisionChild.EnableCopying()

	NatureChild = Engine.ChildControl.NewChild2D()
	NatureChild.AttachMesh(geometry.NewRectangle())
	NatureChild.ScaleX = BlockSize
	NatureChild.ScaleY = BlockSize
	NatureChild.EnableCopying()

	Engine.TextureControl.NewTexture("./assets/cloud1.png", "cloud1", "pixel")
	cloudMaterial = material.NewMaterial(Engine.ShaderControl.GetShader("texture"), &Config)
	cloudMaterial.BecomeTexture(Engine.TextureControl.GetTexture("cloud1"))
	CloudChild = Engine.ChildControl.NewChild2D()
	CloudChild.AttachMaterial(&cloudMaterial)
	CloudChild.AttachMesh(geometry.NewRectangle())
	CloudChild.ScaleX = 300
	CloudChild.ScaleY = 145
	CloudChild.EnableCopying()
	CloudChild.SetSpecificRenderDistance(float32(ScreenWidth/2) + 300)

	//   --------------------------------------------------
	//   Instancing
	//   --------------------------------------------------

	Engine.ChildControl.NewScene("world")
	Engine.ChildControl.DisableAutomaticRendering("world")

	Engine.ChildControl.InstanceChild(SkyChild, "world")
	Engine.ChildControl.InstanceChild(CloudChild, "world")
	Engine.ChildControl.InstanceChild(NoCollisionChild, "world")
	Engine.ChildControl.InstanceChild(NatureChild, "world")
	Engine.ChildControl.InstanceChild(WorldChild, "world")
	Engine.ChildControl.InstanceChild(Player1.PlayerChild, "world")
	Engine.ChildControl.InstanceChild(BlockSelect, "world")
}
