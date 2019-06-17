package main

import (
	"math/rand"
	"rapidengine/child"
)

type EnemyManager struct {
	AllEnemies map[int]Enemy
}

func InitializeEnemyManager() *EnemyManager {
	em := EnemyManager{
		AllEnemies: make(map[int]Enemy),
	}

	LoadGoblinTextures()

	return &em
}

func (em *EnemyManager) Update() {
	for i, enemy := range em.AllEnemies {
		if enemy.Activator().IsActive() {
			enemy.Update()
		}
		if enemy.GetCommon().Health <= 0 {
			enemy.GetCommon().Kill()
		}
		if enemy.GetCommon().Dead {
			delete(EM.AllEnemies, i)
		}
	}
}

func (em *EnemyManager) CheckPlayerCollision() Enemy {
	for _, enemy := range em.AllEnemies {
		pdist := Distance(
			Player1.CenterX, Player1.CenterY,
			enemy.GetChild().X, enemy.GetChild().Y,
		)

		if pdist > 200 {
			continue
		}

		if Player1.CheckEnemyCollision(enemy) {
			return enemy
		}
	}

	return nil
}

func (em *EnemyManager) NewGoblin(radius float32) {
	mat := NewGoblinMaterial()

	goblinChild := Engine.ChildControl.NewChild2D()
	goblinChild.AttachMaterial(mat)
	goblinChild.ScaleX = 300
	goblinChild.ScaleY = 300

	screenSide := (rand.Intn(2) * 2) - 1

	goblinChild.X = Player1.PlayerChild.X + float32(screenSide)*((float32(ScreenWidth/2)+100)+radius)
	goblinChild.Y = float32(HeightMap[int(goblinChild.X)/BlockSize]*BlockSize) + 50

	var g = Goblin{
		common: &Common{

			Health:    100,
			MaxHealth: 100,

			MonsterChild:    goblinChild,
			MonsterMaterial: mat,

			VXMult:   (rand.Float32()*2-1.0)*0.2 + 0.8,
			VYMult:   (rand.Float32()*2-1.0)*0.2 + 1.0,
			GravMult: 1,

			NumJumps: 1,

			Hitbox1: NewHitBox(AABB{
				X:      -10,
				Y:      -10,
				Width:  40,
				Height: 120,
			}, 5),

			aHitbox: AABB{
				OffX:   0,
				OffY:   45,
				Width:  55,
				Height: 60,
			},

			State: "normal",
		},

		activator: Activator{},
	}

	g.common.Hitbox1.OffX = -13
	g.common.Hitbox1.OffY = -20

	g.common.HealthBar = Engine.UIControl.NewProgressBar()
	g.common.HealthBar.SetDimensions(50, 10)
	g.common.HealthBar.BackChild.Static = true
	g.common.HealthBar.BarChild.Static = true
	g.common.HOffsetY = -30

	g.Activator().Activate()

	V.AddBox(&g.common.Hitbox1)
	V.AddAABB(&g.common.aHitbox)

	em.AllEnemies[len(em.AllEnemies)-1] = &g
}

type Enemy interface {
	Update()

	Damage(amount float32)

	GetChild() *child.Child2D
	GetCommon() *Common

	Activator() *Activator
}
