package main

type AI_t interface {
	Get(player_hand Hand_t) Hand_t
}
