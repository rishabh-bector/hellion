package main

//   --------------------------------------------------
//   Lighting
//   --------------------------------------------------

func CreateLighting(x, y int, light float32) {
	if !IsValidPosition(x, y) {
		return
	}
	newLight := light - GetLightBlockAmount(x, y)
	if newLight <= WorldMap.GetDarkness(x, y) {
		return
	}

	WorldMap.SetDarkness(x, y, newLight)

	CreateLighting(x+1, y, newLight)
	CreateLighting(x, y+1, newLight)
	CreateLighting(x-1, y, newLight)
	CreateLighting(x, y-1, newLight)
}

func CreateLightingLimit(x, y int, light float32, limit int) {
	if limit < 1 {
		return
	}
	if !IsValidPosition(x, y) {
		return
	}
	newLight := light - GetLightBlockAmount(x, y)
	if newLight <= WorldMap.GetDarkness(x, y) {
		return
	}

	WorldMap.SetDarkness(x, y, newLight)

	CreateLightingLimit(x+1, y, newLight, limit-1)
	CreateLightingLimit(x, y+1, newLight, limit-1)
	CreateLightingLimit(x-1, y, newLight, limit-1)
	CreateLightingLimit(x, y-1, newLight, limit-1)
}

func FixLightingAt(x, y int) {
	maxLight := float32(0)
	if l := WorldMap.GetDarkness(x+1, y); l > maxLight {
		maxLight = l - GetLightBlockAmount(x+1, y)
	}
	if l := WorldMap.GetDarkness(x, y+1); l > maxLight {
		maxLight = l - GetLightBlockAmount(x, y+1)
	}
	if l := WorldMap.GetDarkness(x-1, y); l > maxLight {
		maxLight = l - GetLightBlockAmount(x-1, y)
	}
	if l := WorldMap.GetDarkness(x, y-1); l > maxLight {
		maxLight = l - GetLightBlockAmount(x, y-1)
	}
	WorldMap.SetDarkness(x, y, maxLight-GetLightBlockAmount(x, y))
	WorldMap.UpdateBackBlockMaterial(x, y)
	WorldMap.UpdateWorldBlockMaterial(x, y)
}

func GetLightBlockAmount(x, y int) float32 {
	if WorldMap.GetWorldBlockID(x, y) == "00000" {
		return GetBlock(WorldMap.GetBackBlockName(x, y)).LightBlock
	}
	return GetBlock(WorldMap.GetWorldBlockName(x, y)).LightBlock
}

func IsValidPosition(x, y int) bool {
	if x > 0 && x < WorldWidth {
		if y > 0 && y < WorldHeight {
			if HeightMap[x]+100 < y {
				return false
			}
			return true
		}
	}
	return false
}
