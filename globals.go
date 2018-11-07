package main

import (
	"rapidengine/child"
	"rapidengine/cmd"
	"rapidengine/configuration"
	"rapidengine/lighting"
	"rapidengine/material"

	perlin "github.com/aquilax/go-perlin"
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

//  --------------------------------------------------
//  Children
//  --------------------------------------------------

var Player child.Child2D
var BlockSelect child.Child2D

// World
var WorldChild child.Child2D
var SkyChild child.Child2D
var NoCollisionChild child.Child2D
var NatureChild child.Child2D
var CloudChild child.Child2D

var l lighting.PointLight

var WorldMap WorldTree

//  --------------------------------------------------
//  World Generation Parameters
//  --------------------------------------------------

const WorldHeight = 3000
const WorldWidth = 2000
const BlockSize = 32
const Flatness = 0.1

const GrassMinimum = 700

const CaveNoiseScalar = 30
const CaveNoiseThreshold = 0.75

const StoneNoiseScalar = 30.0
const StoneTop = 600.0
const StoneTopDeviation = 5
const StoneStartingFrequency = -0.3

var Generator *perlin.Perlin
var cloudMaterial material.Material

var HeightMap [WorldWidth]int

//  --------------------------------------------------
//  Data
//  --------------------------------------------------

var TransparentBlocks = []string{"backdirt", "torch"} //"topGrass1", "topGrass2", "topGrass3", "treeRightRoot", "treeLeftRoot", "treeTrunk", "treeBottomRoot", "treeBranchR1", "treeBranchL1", "flower1", "flower2", "flower3", "pebble"}

var natureBlocks = []string{"leaves", "treeRightRoot", "treeLeftRoot", "treeTrunk", "treeBottomRoot", "treeBranchR1", "treeBranchL1", "topGrass1", "topGrass2", "topGrass3", "flower1", "flower2", "flower3", "pebble"}
