package helpers

import (
	"sort"
)

/*
1. Laki laki dengan umur > 17 th < 60 th
2. Perempuan dengan umur > 19 th < 60 th
*/

func Soal1(jk string, umur int) bool {
	if jk == "pria" && umur > 17 && umur < 60 {
		return true
	} else if jk == "wanita" && umur > 19 && umur < 60 {
		return true
	}

	return false
}

func Soal2(fighters, powers []int, finalPower int) int {
	var EnemyStatus = map[int]int{}
	for i, powPetarung := range fighters {
		EnemyStatus[powPetarung] = powers[i]
	}
	sort.Ints(fighters)

	for _, fighter := range fighters {
		if finalPower >= fighter {
			finalPower += EnemyStatus[fighter]
		} else {
			break
		}
	}

	return finalPower
}
