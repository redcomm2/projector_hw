package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Player struct {
	Name      string
	Place     *Location
	Inventory []Item
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
	fmt.Println("Welcome to the New World!!!")
	fmt.Println("Stephen woke up at the entrance to the cave. " +
		"He only remembers his name. Next to him is a backpack in which he finds matches, a flashlight, and a knife.")

	player := Player{
		Name: "Stephen",
	}

	unknownPlace := Location{
		Name:        "Uknown place",
		Description: "Uknown place, where player woked up",
	}

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
			player.Inventory = append(player.Inventory, Item{Name: "Matches"})
			fmt.Println("Matches added to your inventory.")
		case "2":
			player.Inventory = append(player.Inventory, Item{Name: "Knife"})
			fmt.Println("Knife added to your inventory.")
		case "3":
			player.Inventory = append(player.Inventory, Item{Name: "Flashlight"})
			fmt.Println("Flashlight added to your inventory.")
		case "4":
			fmt.Println("No item chosen.")
		default:
			fmt.Println("Invalid choice, please select a valid item number.")
			continue
		}

		break
	}

	fmt.Println("Your inventory items:")
	for _, item := range player.Inventory {
		fmt.Println("-", item.Name)
	}

	desicion1 := Decision{
		Choice: "Go to cave",
	}

	desicion2 := Decision{
		Choice: "Go to forest",
	}

	decisionsInStart := []Decision{desicion1, desicion2}

	for {
		fmt.Println("Choose where you go?")
		for i, item := range decisionsInStart {
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
			unknownPlace.Decision = desicion2

			forest := Location{
				Name:        "Forest",
				Description: "Dark forest",
			}

			fmt.Printf("You find yourself in the %s. You see abandoned . Choose what you want to do with it?", forest.Name)
		default:
			fmt.Println("Invalid choice, please select a valid item number.")
			continue
		}

		break
	}

}
