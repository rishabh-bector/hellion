package main

//   --------------------------------------------------
//   Lighting
//   --------------------------------------------------

func CreateLighting(x, y int, light float32) {
	if !IsValidPosition(x, y) {
		return
	}
	newLight := light - GetLightBlockAmount(x, y)
	if newLight <= GetLightAt(x, y) {
		return
	}
	WorldMap[x][y].Darkness = newLight
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
	if newLight <= GetLightAt(x, y) {
		return
	}
	WorldMap[x][y].Darkness = newLight
	createSingleCopy(x, y)
	CreateLightingLimit(x+1, y, newLight, limit-1)
	CreateLightingLimit(x, y+1, newLight, limit-1)
	CreateLightingLimit(x-1, y, newLight, limit-1)
	CreateLightingLimit(x, y-1, newLight, limit-1)
}

func RemoveLightingLimit(x, y int, light float32, limit int) {
	if limit < 1 {
		return
	}
	if !IsValidPosition(x, y) {
		return
	}
	newLight := light + GetLightBlockAmount(x, y)
	if newLight >= GetLightAt(x, y) {
		return
	}
	WorldMap[x][y].Darkness = newLight
	createSingleCopy(x, y)
	RemoveLightingLimit(x+1, y, newLight, limit-1)
	RemoveLightingLimit(x, y+1, newLight, limit-1)
	RemoveLightingLimit(x-1, y, newLight, limit-1)
	RemoveLightingLimit(x, y-1, newLight, limit-1)
}

func FixLightingAt(x, y int) {
	maxLight := float32(0)
	if l := GetLightAt(x+1, y); l > maxLight {
		maxLight = l
	}
	if l := GetLightAt(x, y+1); l > maxLight {
		maxLight = l
	}
	if l := GetLightAt(x-1, y); l > maxLight {
		maxLight = l
	}
	if l := GetLightAt(x, y-1); l > maxLight {
		maxLight = l
	}
	WorldMap[x][y].Darkness = maxLight - GetLightBlockAmount(x, y)
}

func GetLightAt(x, y int) float32 {
	return WorldMap[x][y].Darkness
}

func GetLightBlockAmount(x, y int) float32 {
	return BlockMap[NameList[WorldMap[x][y].ID]].LightBlock
}

func IsValidPosition(x, y int) bool {
	if x > 0 && x < WorldWidth {
		if y > 0 && y < WorldHeight {
			return true
		}
	}
	return false
}
