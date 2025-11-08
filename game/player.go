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

func (player *Player) FightMonster(card *components.Card) bool {
	if card.Suit != components.Spade && card.Suit != components.Club {
		fmt.Println("Cannot fight with card that is not a spade or club")
		return false
	}
	fmt.Println("Fighting monster: ", card.Name)
	player.TakeDamage(card.Value)
	return true
}

func (player *Player) TakeDamage(damage int) {
	player.health -= damage
}

func (player *Player) Heal(card components.Card) bool {
	if card.Suit != components.Heart {
		fmt.Println("Cannot heal with card that is not a heart")
		return false
	}
	player.health = min(maxHealth, player.health+card.Value)
	return true

}

func (player *Player) GetHealth() int {
	return player.health
}

func (player *Player) EquipWeapon(weapon components.Card) bool {
	if weapon.Suit != components.Diamond {
		fmt.Println("Cannot equip card that is not a diamond")
		return false
	}

	player.equippedWeapon = &weapon
	fmt.Println("Equipped weapon: ", weapon.Name)
	return true
}

func (player *Player) UnequipWeapon() {
	player.equippedWeapon = nil
}

func (player *Player) GetEquippedWeapon() *components.Card {
	return player.equippedWeapon
}
