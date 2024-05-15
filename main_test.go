package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudge_Game(t *testing.T) {
	assert.Equal(
		t,
		Judge_Game(Hand_chyoki, Hand_gu),
		Game_Result_lose,
	)
	assert.Equal(
		t,
		Judge_Game(Hand_chyoki, Hand_pa),
		Game_Result_win,
	)

	assert.NotEqual(
		t,
		Judge_Game(Hand_chyoki, Hand_gu),
		Game_Result_draw,
	)
}
