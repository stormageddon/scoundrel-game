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
	actionsTaken := 0
	round++

	for actionsTaken < 3 {
		PrintHeader()
		action := getUserAction()
		switch action {
		case 1:
			chosenCardIndex, equipSuccess := TryEquipWeapon()
			fmt.Println("chosenCardIndex: ", chosenCardIndex)
			if equipSuccess {
				actionsTaken++
				if chosenCardIndex == 0 {
					room = room[1:]
				} else {
					room = append(room[:chosenCardIndex], room[chosenCardIndex+1:]...)
				}
				//				room = append(room[:chosenCardIndex-1], room[chosenCardIndex:]...)
			}
		case 2:
			player.TakeDamage(5)
			actionsTaken++
			// if TryFightMonster() {
			// 	actionsTaken++
			// }

		case 3:
			chosenCardIndex := -1
			fmt.Print("Which potion do you wish to drink?\n:>")
			_, err := fmt.Scanf("%d", &chosenCardIndex)
			if err != nil {
				log.Fatal(err)
				continue
			}
			fmt.Println()

			healSuccess := player.Heal(room[chosenCardIndex-1])
			if healSuccess {
				actionsTaken++
				if chosenCardIndex-1 == 0 {
					room = room[1:]
				} else {
					room = append(room[:chosenCardIndex-1], room[chosenCardIndex:]...)
				}
			}
		case 5:
			GetNewRoom()
		case 9:
			fmt.Println("Exiting game")
			gameOver = true
			return
		default:
			fmt.Println("Invalid action")
			round--
		}

	}
}

func PrintHeader() {
	fmt.Println("--------------------------------")
	equippedWeaponName := " -- "
	if player.equippedWeapon != nil {
		fmt.Printf("Player: {health: %d, equippedWeapon: %s}\n", player.health, player.equippedWeapon.Name)
		equippedWeaponName = player.equippedWeapon.Name
	}
	fmt.Printf("Round %d | Health: %d | Weapon: %s\n", round, player.health, equippedWeaponName)
	fmt.Print("Room: {")
	for i := 0; i < len(room); i++ {
		fmt.Printf("%d. %s", i+1, room[i].Id)
		if i < len(room)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("}")
	//	fmt.Printf("Room: {1. %s, 2. %s, 3. %s, 4. %s}\n", room[0].Id, room[1].Id, room[2].Id, room[3].Id)
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
	fmt.Println("1. Equip a weapon (Diamonds only)")
	fmt.Println("2. Fight monster (Spades and Clubs)")
	fmt.Println("3. Heal (Hearts only)")
	fmt.Println("5. Avoid room")
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

func TryEquipWeapon() (int, bool) {
	chosenCardIndex := -1
	fmt.Println("Which card do you wish to equip?")
	printRoom()
	fmt.Print("\n:> ")
	_, err := fmt.Scanf("%d", &chosenCardIndex)
	if err != nil {
		log.Fatal(err)
		return -1, false
	}
	fmt.Printf("====== index: %d card: %v\n", chosenCardIndex-1, room[chosenCardIndex-1])
	equipSuccess := player.EquipWeapon(room[chosenCardIndex-1])
	return chosenCardIndex - 1, equipSuccess
}

func TryFightMonster() bool {
	chosenCardIndex := -1
	fmt.Print("Which monster do you wish to fight?\n:>")
	_, err := fmt.Scanf("%d", &chosenCardIndex)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println()
	return player.FightMonster(&room[chosenCardIndex-1])

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

func printRoom() {
	fmt.Print("Room: {")
	for i := 0; i < len(room); i++ {
		fmt.Printf("%d. %s", i+1, room[i].Id)
		if i < len(room)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("}")
}
