package main

type Hitbox struct {
	LAABB AABB
	RAABB AABB
	UAABB AABB
	DAABB AABB

	OffX float32
	OffY float32

	// For viewer
	X float32
	Y float32
}

func NewHitBox(original AABB, width float32) Hitbox {
	s := float32(10)
	return Hitbox{
		LAABB: AABB{
			X:      0,
			Y:      s,
			Width:  width,
			Height: original.Height - s*2,
		},
		RAABB: AABB{
			X:      original.Width - width,
			Y:      s,
			Width:  width,
			Height: original.Height - s*2,
		},
		UAABB: AABB{
			X:      s,
			Y:      original.Height - width,
			Width:  original.Width - s*2,
			Height: width,
		},
		DAABB: AABB{
			X:      s,
			Y:      0,
			Width:  original.Width - s*2,
			Height: width,
		},
	}
}

func (hb *Hitbox) CheckCollisionAABB(other AABB, vx, vy, selfx, selfy float32) (bool, bool, bool, bool) {
	left := hb.LAABB.CheckCollisionTranslated(other, vx, vy, selfx, selfy)
	right := hb.RAABB.CheckCollisionTranslated(other, vx, vy, selfx, selfy)
	up := hb.UAABB.CheckCollisionTranslated(other, vx, vy, selfx, selfy)
	down := hb.DAABB.CheckCollisionTranslated(other, vx, vy, selfx, selfy)
	return left, right, up, down
}

func (hb *Hitbox) CheckCollisionHitbox(other Hitbox, vx, vy float32) (bool, bool, bool, bool) {
	left := hb.LAABB.CheckHitboxCollision(other, vx, vy)
	right := hb.RAABB.CheckHitboxCollision(other, vx, vy)
	up := hb.UAABB.CheckHitboxCollision(other, vx, vy)
	down := hb.DAABB.CheckHitboxCollision(other, vx, vy)
	return left, right, up, down
}

type AABB struct {
	X      float32
	Y      float32
	Width  float32
	Height float32

	OffX float32
	OffY float32
}

func (aabb *AABB) CheckHitboxCollision(other Hitbox, vx, vy float32) bool {
	if col := aabb.CheckCollision(other.LAABB, vx, vy); col {
		return col
	}
	if col := aabb.CheckCollision(other.RAABB, vx, vy); col {
		return col
	}
	if col := aabb.CheckCollision(other.UAABB, vx, vy); col {
		return col
	}
	if col := aabb.CheckCollision(other.DAABB, vx, vy); col {
		return col
	}
	return false
}

func (aabb *AABB) CheckCollisionTranslated(other AABB, vx, vy, selfx, selfy float32) bool {
	ax := aabb.X + (vx * float32(Engine.Renderer.DeltaFrameTime)) + selfx
	ay := aabb.Y + (vy * float32(Engine.Renderer.DeltaFrameTime)) + selfy

	if ax+aabb.Width > other.X &&
		ax < other.X+other.Width &&
		ay+aabb.Height > other.Y &&
		ay < other.Y+other.Height {
		return true
	} else {
		return false
	}
}

func (aabb *AABB) CheckCollision(other AABB, vx, vy float32) bool {
	ax := aabb.X + (vx * float32(Engine.Renderer.DeltaFrameTime))
	ay := aabb.Y + (vy * float32(Engine.Renderer.DeltaFrameTime))

	if ax+aabb.Width > other.X &&
		ax < other.X+other.Width &&
		ay+aabb.Height > other.Y &&
		ay < other.Y+other.Height {
		return true
	} else {
		return false
	}
}

func CheckWorldCollision(hb Hitbox, vx, vy, selfx, selfy float32) (bool, bool, bool, bool, bool, bool) {
	top := false
	left := false
	bottom := false
	right := false

	topleft := false
	topright := false

	px := int((selfx) / BlockSize)
	pex := int((selfx + hb.DAABB.Width) / BlockSize)

	py := int((selfy) / BlockSize)
	pey := int((selfy + hb.LAABB.Height) / BlockSize)

	// Broad phase collision
	for x := px - 3; x < pex+3; x++ {
		for y := py - 3; y < pey+3; y++ {
			if block := WorldMap.GetWorldBlock(x, y); block.ID != "00000" {
				l, r, u, d := hb.CheckCollisionAABB(AABB{block.X, block.Y, BlockSize, BlockSize, 0, 0}, vx, vy, selfx, selfy)
				if l {
					left = true
				}
				if r {
					right = true
				}
				if u {
					top = true
				}
				if d {
					bottom = true
				}
			}
		}
	}

	if block := WorldMap.GetWorldBlock(px-1, py+1); block.ID != "00000" {
		if l, _, _, _ := hb.CheckCollisionAABB(AABB{block.X, block.Y, BlockSize, BlockSize, 0, 0}, vx, vy, selfx, selfy); l {
			topleft = true
		}
	}

	if block := WorldMap.GetWorldBlock(pex+1, py+1); block.ID != "00000" {
		if _, r, _, _ := hb.CheckCollisionAABB(AABB{block.X, block.Y, BlockSize, BlockSize, 0, 0}, vx, vy, selfx, selfy); r {
			topright = true
		}
	}

	return top, left, bottom, right, topleft, topright
}
