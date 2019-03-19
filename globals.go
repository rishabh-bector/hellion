package main

import (
	"rapidengine/child"
	"rapidengine/cmd"
	"rapidengine/configuration"
	"rapidengine/lighting"
	"rapidengine/material"
)

//  --------------------------------------------------
//  Globals.go contains all the global variables in the project.
//  This is not good practice.
//  --------------------------------------------------

// Rapid Engine
var Engine *cmd.Engine
var Config configuration.EngineConfig

// Screen Size
var ScreenWidth = 1920
var ScreenHeight = 1080

// Mouse Settings
var MouseSensitivity = 9.0

//  --------------------------------------------------
//  Children
//  --------------------------------------------------

var BaseGravity = float32(1710.0)
var BaseSpeedX = float32(150.0)
var BaseSpeedY = float32(600.0)

var BlockSelect *child.Child2D

// World
var WorldChild *child.Child2D
var SkyChild *child.Child2D
var NoCollisionChild *child.Child2D
var NatureChild *child.Child2D
var GrassChild *child.Child2D
var CloudChild *child.Child2D
var SunChild *child.Child2D

var Back1Child *child.Child2D
var Back2Child *child.Child2D
var Back3Child *child.Child2D
var Back4Child *child.Child2D
var Back5Child *child.Child2D
var Back6Child *child.Child2D

var l lighting.PointLight

//  --------------------------------------------------
//  General
//  --------------------------------------------------

var GamePaused bool

var CurrentWorld int

//  --------------------------------------------------
//  World Generation
//  --------------------------------------------------

var Seed = int64(0)

// Size
const WorldWidth = 3000
const WorldHeight = 2000
const BlockSize = 32

// Height
const Flatness = 0.25
const GrassMinimum = 1500

// Cave generation
const CaveStartingThreshold = 0.27
const CaveEndingThreshold = 0.42
const CaveThresholdDelta = 0.002

const CaveIterations = 20
const CaveBirthLimit = 4
const CaveDeathLimit = 3

const SecondCaveIterations = 5
const SecondCaveBirthLimit = 3
const SecondCaveDeathLimit = 2

// Stone generation
const StoneFrequencyDelta = 0.001
const StoneStartingFrequency = 0.32
const StoneEndingFrequency = 0.77
const StoneTopDeviation = 10

// Data
var WorldMap WorldTree
var HeightMap [WorldWidth]int
var CaveMap [][]bool

//  --------------------------------------------------
//  Scenes
//  --------------------------------------------------

var TitleScene *cmd.Scene
var ChooseScene *cmd.Scene
var LoadingScene *cmd.Scene
var WorldScene *cmd.Scene
var MenuScene *cmd.Scene
var SaveScene *cmd.Scene
var HotbarScene *cmd.Scene
var InventoryScene *cmd.Scene

var EM *EnemyManager

//  --------------------------------------------------
//  Theme Materials
//  --------------------------------------------------

var ButtonMaterial *material.BasicMaterial

//  --------------------------------------------------
//  Data
//  --------------------------------------------------

var TransparentBlocks = []string{"backdirt", "torch"} //"topGrass1", "topGrass2", "topGrass3", "treeRightRoot", "treeLeftRoot", "treeTrunk", "treeBottomRoot", "treeBranchR1", "treeBranchL1", "flower1", "flower2", "flower3", "pebble"}
var natureBlocks = []string{"leaves", "treeRightRoot", "treeLeftRoot", "treeTrunk", "treeBottomRoot", "treeBranchR1", "treeBranchL1", "topGrass1", "topGrass2", "topGrass3", "flower1", "flower2", "flower3", "pebble"}
var cloudMaterial *material.BasicMaterial
