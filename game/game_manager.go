package game

import (
	"fmt"
	"log"

	"caputo.io/scoundrel/components"
)

var deck components.Deck
var player Player
var round int
var gameOver bool
var room []components.Card

func Initialize() {
	fmt.Println("Starting game")

	d, err := components.NewDeck()
	if err != nil {
		fmt.Println(err)
		return
	}
	d = removeUnneededCards(d)
	d.Shuffle()
	deck = d

	player = NewPlayer()
}

func Start() {
	round = 0
	gameOver = false
	GetNewRoom()

	for {
		if !gameOver {
			nextRound()
		} else {
			break
		}
	}
}

func nextRound() {
	round++
	PrintHeader()
	action := getUserAction()
	switch action {
	case 1:
		chosenCardIndex := -1
		fmt.Print("Which card do you wish to equip?\n:>")
		_, err := fmt.Scanf("%d", &chosenCardIndex)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
		player.EquipWeapon(&room[chosenCardIndex-1])
	case 2:
		GetNewRoom()
	case 9:
		fmt.Println("Exiting game")
		gameOver = true
	default:
		fmt.Println("Invalid action")
		round--
	}
}

func PrintHeader() {
	fmt.Println("--------------------------------")
	fmt.Printf("Round %d \\ Health: %d\n", round, player.health)
	fmt.Printf("Room: {1. %s, 2. %s, 3. %s, 4. %s}\n", room[0].Id, room[1].Id, room[2].Id, room[3].Id)
	fmt.Println("--------------------------------")
}

func GetNewRoom() {
	room = []components.Card{}
	for range 4 {
		card, err := deck.Draw()
		if err != nil {
			println("Error drawing card: ", err)
			return
		}
		room = append(room, *card)
	}
}

func getUserAction() int {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Equip a card")
	fmt.Println("2. Reroll room")
	fmt.Println("9. Quit")
	fmt.Print("\n:> ")

	var action int
	_, err := fmt.Scanf("%d", &action)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You chose: ", action)
	return action
}

func equipCard(selectedCard components.Card) {
	player.EquipWeapon(&selectedCard)
}

func removeUnneededCards(deck components.Deck) components.Deck {
	deck.Remove("Jack of Hearts")
	deck.Remove("Jack of Diamonds")
	deck.Remove("Queen of Hearts")
	deck.Remove("Queen of Diamonds")
	deck.Remove("King of Hearts")
	deck.Remove("King of Diamonds")
	deck.Remove("Ace of Hearts")
	deck.Remove("Ace of Diamonds")
	return deck
}
