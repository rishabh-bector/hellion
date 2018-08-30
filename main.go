package main

import (
	"math/rand"
	"rapidengine"
	"rapidengine/configuration"

	perlin "github.com/aquilax/go-perlin"
)

var engine rapidengine.Engine
var config configuration.EngineConfig

var child1 rapidengine.Child2D
var child2 rapidengine.Child2D

const WorldHeight = 1000
const WorldWidth = 1000
const BlockSize = 25

var p *perlin.Perlin
var world [WorldWidth][WorldHeight]int

var blocks []string

func main() {
	config = configuration.NewEngineConfig(1920, 1080, 2)

	engine = rapidengine.NewEngine(config, render)
	err := engine.Initialize()
	if err != nil {
		panic(err)
	}

	//engine.TextureControl.NewTexture("./sprites/dirt.png", "dirt")
	//engine.TextureControl.NewTexture("./sprites/grass.png", "grass")
	engine.TextureControl.NewTexture("./sprites/dirt.png", "dirt")

	blocks = append(blocks, "dirt")

	p = perlin.NewPerlin(2, 2, 10, int64(rand.Int()))

	/*engine.CollisionControl.CreateGroup("children")
	engine.CollisionControl.AddChildToGroup(&child1, "children")
	engine.CollisionControl.AddChildToGroup(&child2, "children")
	engine.CollisionControl.CreateCollision(&child1, "children", cbk)*/

	engine.Renderer.SetRenderDistance(1000)
	engine.Renderer.MainCamera.SetPosition(17000, 17000)

	child1 = engine.NewChild2D()
	child1.AttachPrimitive(rapidengine.NewRectangle(BlockSize, BlockSize, &config))
	child1.AttachTexturePrimitive(engine.TextureControl.GetTexture("dirt"))
	child1.EnableCopying()

	engine.Config.Logger.Info("Generating world...")
	generateWorld()
	createCopies()

	println(len(child1.GetCopies()))

	engine.Instance(&child1)
	engine.InitializeRenderer()

	engine.StartRenderer()
	<-engine.Done()
}

func render(renderer *rapidengine.Renderer) {
	renderer.RenderChildren()
}

func generateWorld() {
	for x := 0; x < WorldWidth; x += 1 {
		for y := 0; y < WorldHeight; y += 1 {
			world[x][y] = 0
		}
	}
}

func createCopies() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			child1.AddCopy(rapidengine.ChildCopy{float32(x * BlockSize), float32(y * BlockSize), engine.TextureControl.GetTexture(blocks[world[x][y]])})
		}
	}
}
