package main

import (
	"math/rand"
	"time"
)

func shafflee(arr []int) {
	rand.Seed(time.Now().Unix())
	for k := range arr {
		random := rand.Intn(len(arr))
		arr[k], arr[random] = arr[random], arr[k]
	}
}
