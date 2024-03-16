package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Player struct {
	Name      string
	Location  *Location
	Inventory *Item
}

type Item struct {
	Name string
}

type Location struct {
	Name        string
	Description string
	Decision    Decision
}

type Decision struct {
	Choice string
}

func main() {
	runGame()
}

func runGame() {
	fmt.Println("Welcome to the New World!!!")
	fmt.Println("Stephen woke up at the entrance to the cave. " +
		"He only remembers his name. Next to him is a backpack in which he finds matches, a flashlight, and a knife.")

	player := Player{
		Name: "Stephen",
	}

	unknownPlace := Location{
		Name:        "Uknown place",
		Description: "Uknown place, where player woke up",
	}

	player.Location = &unknownPlace

	for {
		fmt.Println("Choose an item to add to your inventory (or choose not to take anything):")
		items := []string{"Matches", "Knife", "Flashlight", "None"}

		for i, item := range items {
			fmt.Printf("[%d] %s\n", i+1, item)
		}

		fmt.Print("Enter the number of your choice: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			player.Inventory = &Item{Name: "Matches"}
			fmt.Println("Matches added to your inventory.")
		case "2":
			player.Inventory = &Item{Name: "Knife"}
			fmt.Println("Knife added to your inventory.")
		case "3":
			player.Inventory = &Item{Name: "Flashlight"}
			fmt.Println("Flashlight added to your inventory.")
		case "4":
			fmt.Println("No item chosen.")
		default:
			fmt.Println("Invalid choice, please select a valid item number.")
			continue
		}

		break
	}

	desicionInUnknownPlace1 := Decision{
		Choice: "Go to cave",
	}

	desicionInUnknownPlace2 := Decision{
		Choice: "Go to forest",
	}

	decisionsInUnknownPlace := []Decision{desicionInUnknownPlace1, desicionInUnknownPlace2}

	forest := Location{
		Name:        "Forest",
		Description: "Dark forest",
	}

	for {
		fmt.Println("Choose where you go?")
		for i, item := range decisionsInUnknownPlace {
			fmt.Printf("[%d] %s\n", i+1, item.Choice)
		}

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("There is beast in cave. You died.")
			fmt.Println("Game over.")
			os.Exit(1)
		case "2":
			unknownPlace.Decision = desicionInUnknownPlace2

			fmt.Printf("You find yourself in the %s.", forest.Name)
		default:
			fmt.Println("Invalid choice, please select a valid item number.")
			continue
		}

		break
	}

	fmt.Println("You see motionless animal. Choose what you want to do with it?")

	desicionInForest1 := Decision{
		Choice: "Do nothing",
	}

	desicionInForest2 := Decision{
		Choice: "Go to animal",
	}

	decisionsInForest := []Decision{desicionInForest1, desicionInForest2}

	for {
		fmt.Println("Choose where you go?")
		for i, item := range decisionsInForest {
			fmt.Printf("[%d] %s\n", i+1, item.Choice)
		}

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			forest.Decision = desicionInForest1
		case "2":
			fmt.Println("You are approaching an animal. It turned out to be a wild bear. He wakes up and attacks you.")

			if player.Inventory.Name == "Knife" {
				fmt.Println("There was a knife in your pocket. With its help you kill the bear. " +
					"Next to him you find a bloody piece of paper with the inscription: safe code 43.")

			} else {
				fmt.Println("You didn’t have anything to protect yourself from the wild beast. You are dead. Game over.")
				os.Exit(1)
			}
			forest.Decision = desicionInForest2
		default:
			fmt.Println("Invalid choice, please select a valid item number.")
			continue
		}

		break
	}

	campInForest := Location{
		Name:        "Camp in Forest",
		Description: "Adonded camp in forest",
	}

	fmt.Printf("You see the %s. You decided to go there", campInForest.Name)

	fmt.Println("In the camp, in one of the tents you see a safe. You approach it and try to open it, but it is closed. " +
		"It has a panel for entering a password. You are trying to retrieve the password:")

	attempts := 0
	for {
		if attempts >= 3 {
			fmt.Println("You still haven't guessed the password. You feel weak and lose consciousness.")
			runGame()
		}

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "43" {
			fmt.Println("In the safe you find a map. You understand how to get out of the forest and reach people. " +
				"Congratulations, you've won!")
			os.Exit(1)
		} else {
			fmt.Println("Wrong password...")
		}

		attempts++
	}
}
