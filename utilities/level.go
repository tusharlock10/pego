package utilities

import "math"

func GetChampionLevelFromXP(totalXP uint) uint {
	const maxLevel = 999
	const thresholdXP = 25500000

	level := (float64(totalXP-thresholdXP) / 1000000) + 50
	if totalXP < thresholdXP {
		level = (-100 + math.Sqrt(10000+float64(4*totalXP))) / 200
	}
	if level > maxLevel {
		return maxLevel
	}
	return uint(math.Floor(level))
}

func GetPlayerLevelFromXP(totalXP uint) uint {
	const maxLevel = 999
	const thresholdXP = 25500000

	totalXP += 20000
	level := (float64(totalXP-thresholdXP) / 1000000) + 50
	if totalXP < thresholdXP {
		level = (-100 + math.Sqrt(10000+float64(4*totalXP))) / 200
	}
	if level > maxLevel {
		return maxLevel
	}
	return uint(math.Floor(level))
}
