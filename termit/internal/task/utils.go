package task

import (
	"math/rand"
	"time"
)

func GetRandomNonNegInt(limit int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(limit)
}