package main

import (
	"math/rand"
	"rapidengine"
	"rapidengine/configuration"
	"runtime"

	perlin "github.com/aquilax/go-perlin"
)

var engine rapidengine.Engine
var config configuration.EngineConfig

var child1 rapidengine.Child2D
var child2 rapidengine.Child2D

const WorldHeight = 1000
const WorldWidth = 1000
const BlockSize = 25
const Flatness = 0.3

var p *perlin.Perlin
var world [WorldWidth][WorldHeight]int

var blocks []string

func init() {
	runtime.LockOSThread()
}

func main() {
	config = configuration.NewEngineConfig(1920, 1080, 2)

	engine = rapidengine.NewEngine(config, render)
	err := engine.Initialize()
	if err != nil {
		panic(err)
	}

	engine.TextureControl.NewTexture("./sprites/sky.png", "sky")
	engine.TextureControl.NewTexture("./sprites/dirt.png", "dirt")
	engine.TextureControl.NewTexture("./sprites/grass.png", "grass")

	blocks = append(blocks, "sky")
	blocks = append(blocks, "dirt")
	blocks = append(blocks, "grass")

	p = perlin.NewPerlin(2, 2, 10, int64(rand.Int()))

	/*engine.CollisionControl.CreateGroup("children")
	engine.CollisionControl.AddChildToGroup(&child1, "children")
	engine.CollisionControl.AddChildToGroup(&child2, "children")
	engine.CollisionControl.CreateCollision(&child1, "children", cbk)*/

	engine.Renderer.SetRenderDistance(1000)
	engine.Renderer.MainCamera.SetPosition(17000, 17000)
	engine.Renderer.MainCamera.SetSpeed(0.2)

	child1 = engine.NewChild2D()
	child1.AttachPrimitive(rapidengine.NewRectangle(BlockSize, BlockSize, &config))
	child1.AttachTexturePrimitive(engine.TextureControl.GetTexture("grass"))
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
	heights := generateHeights()
	for x := 0; x < WorldWidth; x += 1 {
		world[x][heights[x]] = 2
	}
	fillHeights()
}

func createCopies() {
	for x := 0; x < WorldWidth; x++ {
		for y := 0; y < WorldHeight; y++ {
			child1.AddCopy(rapidengine.ChildCopy{float32(x * BlockSize), float32(y * BlockSize), engine.TextureControl.GetTexture(blocks[world[x][y]])})
		}
	}
}

func noise2D(x, y float64) float64 {
	return (p.Noise2D(x, y) + 0.4) / 0.8
}

func noise1D(x float64) float64 {
	return (p.Noise1D(x) + 0.4) / 0.8
}

func generateHeights() [WorldWidth]int {
	heights := [WorldWidth]int{}
	for x := 0; x < WorldWidth; x++ {
		heights[x] = int(Flatness * noise1D(float64(x)/WorldWidth) * WorldHeight)
	}
	return heights
}

func fillHeights() {
	for x := 0; x < WorldWidth; x += 1 {
		for y := 0; y < WorldHeight; y += 1 {
			world[x][y] = 1
			if world[x][y+1] == 2 {
				break
			}
		}
	}
}
