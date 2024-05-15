package main

import (
	"math/rand"
)

type Random_AI struct {
}

func (self Random_AI) Get(player_hand Hand_t) Hand_t {
	return Hand_t(rand.Intn(3))
}
