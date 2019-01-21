package main

type AABB struct {
	X      float32
	Y      float32
	Width  float32
	Height float32

	// To create the second hitbox, for direction detection
	DirectionOffset float32

	// Minimum distance for collision
	MinimumXDist float32
	MinimumYDist float32
}

func (aabb *AABB) CheckCollision(other AABB, vx, vy float32) int {
	ax := aabb.X + (vx * float32(Engine.Renderer.DeltaFrameTime))
	ay := aabb.Y // + (vy * float32(Engine.Renderer.DeltaFrameTime))

	if ax+aabb.Width > other.X &&
		ax < other.X+other.Width &&
		ay+aabb.Height > other.Y &&
		ay < other.Y+other.Height {
	} else {
		return 0
	}

	b_collision := ay - (other.Y + other.Height)
	t_collision := (ay + aabb.Height) - other.Y

	l_collision := ax - (other.X + other.Width)
	r_collision := (ax + aabb.Width) - other.X

	b_collision *= b_collision
	t_collision *= t_collision
	l_collision *= l_collision
	r_collision *= r_collision

	l_collision += aabb.DirectionOffset
	r_collision += aabb.DirectionOffset

	if l_collision < r_collision && l_collision < t_collision && l_collision < b_collision {
		if l_collision >= aabb.MinimumXDist {
			return 3
		}
	}

	if r_collision < l_collision && r_collision < t_collision && r_collision < b_collision {
		if r_collision >= aabb.MinimumXDist {
			return 1
		}
	}

	if t_collision < b_collision && t_collision < l_collision && t_collision < r_collision {
		if t_collision >= aabb.MinimumYDist {
			return 2
		}
	}

	if b_collision < t_collision && b_collision < l_collision && b_collision < r_collision {
		if b_collision >= aabb.MinimumYDist && vy <= 0 {
			return 4
		}
	}

	return 0
}

func CheckWorldCollision(hb AABB, vx, vy float32) (bool, bool, bool, bool, bool, bool) {
	top := false
	left := false
	bottom := false
	right := false

	topleft := false
	topright := false

	px := int((hb.X) / BlockSize)
	pex := int((hb.X + hb.Width) / BlockSize)

	py := int((hb.Y) / BlockSize)
	pey := int((hb.Y + hb.Height) / BlockSize)

	// Broad phase collision
	for x := px - 3; x < pex+3; x++ {
		for y := py - 3; y < pey+3; y++ {
			if block := WorldMap.GetWorldBlock(x, y); block.ID != "00000" {
				if cols := hb.CheckCollision(AABB{block.X, block.Y, BlockSize, BlockSize, 0, 0, 0}, vx, vy); cols != 0 {
					if cols == 1 {
						right = true
					} else if cols == 2 {
						top = true
					} else if cols == 3 {
						left = true
					} else if cols == 4 {
						bottom = true
					}
				}
			}
		}
	}

	if block := WorldMap.GetWorldBlock(px-1, py+1); block.ID != "00000" {
		if cols := hb.CheckCollision(AABB{block.X, block.Y, BlockSize, BlockSize, 0, 0, 0}, vx, vy); cols != 0 {
			if cols == 3 {
				topleft = true
			}
		}
	}

	if block := WorldMap.GetWorldBlock(pex+1, py+1); block.ID != "00000" {
		if cols := hb.CheckCollision(AABB{block.X, block.Y, BlockSize, BlockSize, 0, 0, 0}, vx, vy); cols != 0 {
			if cols == 1 {
				topright = true
			}
		}
	}

	return top, left, bottom, right, topleft, topright
}
