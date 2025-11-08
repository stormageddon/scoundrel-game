package game

import (
	"fmt"

	"caputo.io/scoundrel/components"
)

const maxHealth = 20

type Player struct {
	health         int
	equippedWeapon *components.Card
}

func NewPlayer() Player {
	player := Player{maxHealth, nil}

	return player
}

func (player *Player) TakeDamage(damage int) {
	player.health -= damage
}

func (player *Player) Heal(amount int) {
	player.health = min(maxHealth, player.health+amount)
}

func (player *Player) GetHealth() int {
	return player.health
}

func (player *Player) EquipWeapon(weapon *components.Card) {
	if weapon.Suit != components.Diamond {
		fmt.Println("Cannot equip card that is not a diamond")
		return
	}

	player.equippedWeapon = weapon
	fmt.Println("Equipped weapon: ", weapon.Name)
}

func (player *Player) UnequipWeapon() {
	player.equippedWeapon = nil
}

func (player *Player) GetEquippedWeapon() *components.Card {
	return player.equippedWeapon
}
