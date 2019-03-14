package main

import (
	"rapidengine/child"
	"rapidengine/geometry"
	"rapidengine/material"
)

var V Viewer

type Viewer struct {
	aabbs map[*child.Child2D]*AABB
	hbs   map[*AABB]*Hitbox

	Mat *material.BasicMaterial
}

func InitializeHitboxViewer() {
	V = Viewer{
		aabbs: make(map[*child.Child2D]*AABB),
		hbs:   make(map[*AABB]*Hitbox),
		Mat:   Engine.MaterialControl.NewBasicMaterial(),
	}
}

func (v *Viewer) AddBox(hb *Hitbox) {
	v.AddAABBHitbox(&hb.LAABB, hb)
	v.AddAABBHitbox(&hb.RAABB, hb)
	v.AddAABBHitbox(&hb.DAABB, hb)
	v.AddAABBHitbox(&hb.UAABB, hb)
}

func (v *Viewer) AddAABBHitbox(a *AABB, h *Hitbox) {
	c := Engine.ChildControl.NewChild2D()
	c.AttachMaterial(v.Mat)
	c.AttachMesh(geometry.NewRectangle())

	c.ScaleX = a.Width
	c.ScaleY = a.Height

	v.aabbs[c] = a
	v.hbs[a] = h
}

func (v *Viewer) AddAABB(a *AABB) {
	c := Engine.ChildControl.NewChild2D()
	c.AttachMaterial(v.Mat)
	c.AttachMesh(geometry.NewRectangle())

	c.ScaleX = a.Width
	c.ScaleY = a.Height

	v.aabbs[c] = a
}

func (v *Viewer) Update() {
	for c, ab := range v.aabbs {
		c.X = ab.X
		c.Y = ab.Y

		if h, ok := v.hbs[ab]; ok {
			c.X += h.X
			c.Y += h.Y
		}
	}
}

func (v *Viewer) Render() {
	for c, _ := range v.aabbs {
		Engine.Renderer.RenderChild(c)
	}
}
