package test

import (
	"testing"

	"caputo.io/scoundrel/components"
	"caputo.io/scoundrel/game"
)

func TestCreatePlayer(t *testing.T) {
	player := game.NewPlayer()

	playerHealth := player.GetHealth()

	if playerHealth != 20 {
		t.Errorf("Player health is not %d", playerHealth)

	}
}

func TestPlayerCanTakeDamage(t *testing.T) {
	player := game.NewPlayer()

	player.TakeDamage(5)

	playerHealth := player.GetHealth()

	if playerHealth != 15 {
		t.Errorf("Player health should be %d but was %d", 15, playerHealth)
	}
}

func TestPlayerCanHeal(t *testing.T) {
	player := game.NewPlayer()

	player.TakeDamage(10)
	player.Heal(components.Card{Value: 1, Suit: components.Heart})
	playerHealth := player.GetHealth()

	if playerHealth != 11 {
		t.Errorf("Player health should be %d but was %d", 11, playerHealth)
	}
}

func TestPlayerCannotHealOverMax(t *testing.T) {
	player := game.NewPlayer()

	player.Heal(components.Card{Value: 10, Suit: components.Heart})
	playerHealth := player.GetHealth()

	if playerHealth > 20 {
		t.Error("Player should not be able to heal over max health")
	}
}
