package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	Name      string
	Inventory []string
}

type Decision struct {
	Description string
	NextSceneID string
	ActionID    string
}

type Scene struct {
	ID          string
	Description string
	Decisions   []Decision
}

var scenes map[string]Scene

func initGame() {
	scenes = map[string]Scene{
		"start": {
			ID:          "start",
			Description: "Stephen woke up at the entrance to the cave with a backpack containing matches, a flashlight, and a knife. Which do you take?",
			Decisions: []Decision{
				{Description: "Take the matches.", NextSceneID: "choosePath"},
				{Description: "Take the flashlight.", NextSceneID: "choosePath"},
				{Description: "Take the knife.", NextSceneID: "choosePath", ActionID: "takeKnife"},
				{Description: "Don't take anything.", NextSceneID: "choosePath"},
			},
		},
		"choosePath": {
			ID:          "choosePath",
			Description: "There are two paths ahead: one leading to a cave and another to a forest. Where do you go?",
			Decisions: []Decision{
				{Description: "Enter the cave.", NextSceneID: "caveDeath"},
				{Description: "Go to the forest.", NextSceneID: "forest"},
			},
		},
		"caveDeath": {
			ID:          "caveDeath",
			Description: "You enter the cave and encounter a beast. Unfortunately, you are not prepared to fight it. Game Over.",
		},
		"forest": {
			ID:          "forest",
			Description: "You are in a dark forest. There's a motionless figure ahead. It's a bear.",
			Decisions: []Decision{
				{Description: "Approach the bear.", NextSceneID: "bearEncounter"},
				{Description: "Avoid the bear and look for a way out.", NextSceneID: "campScene"},
			},
		},
		"bearEncounter": {
			ID:          "bearEncounter",
			Description: "As you approach, the bear wakes up and attacks you.",
			Decisions: []Decision{
				{Description: "Fight back.", NextSceneID: "campScene", ActionID: "fightBear"},
				{Description: "Run away.", NextSceneID: "campScene"},
			},
		},
		"campScene": {
			ID:          "campScene",
			Description: "You find an abandoned camp with a locked safe.",
			Decisions: []Decision{
				{Description: "Try to open the safe.", NextSceneID: "startScene", ActionID: "openSafe"},
			},
		},
	}
}

func main() {
	player := &Player{Name: "Stephen"}
	initGame()
	enterScene("start", player)
}

func enterScene(sceneID string, player *Player) {
	scene, exists := scenes[sceneID]
	if !exists {
		fmt.Println("Scene does not exist.")
		return
	}

	fmt.Println(scene.Description)
	for index, decision := range scene.Decisions {
		fmt.Printf("[%d] %s\n", index+1, decision.Description)
	}

	makeDecision(scene, player)
}

func makeDecision(scene Scene, player *Player) {
	fmt.Print("Choose an option: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)

	if err != nil || choice < 1 || choice > len(scene.Decisions) {
		fmt.Println("Invalid choice, try again.")
		makeDecision(scene, player)
	} else {
		decision := scene.Decisions[choice-1]
		handleDecision(decision, player)
	}
}

func handleDecision(decision Decision, player *Player) {
	fmt.Println(decision.ActionID)
	switch decision.ActionID {
	case "takeKnife":
		player.Inventory = append(player.Inventory, "Knife")
		fmt.Println("Knife added to your inventory.")
	case "fightBear":
		if contains(player.Inventory, "Knife") {
			fmt.Println("You successfully defend yourself with the knife. You find a clue that says '43'.")
		} else {
			fmt.Println("Without a weapon, you have no chance. Game Over.")
			os.Exit(1)
			return
		}
	case "openSafe":
		fmt.Println("Enter the code for the safe:")
		reader := bufio.NewReader(os.Stdin)
		code, _ := reader.ReadString('\n')
		code = strings.TrimSpace(code)
		if code == "43" {
			fmt.Println("The safe opens, revealing a way out of the forest.")
		} else {
			fmt.Println("Wrong code. The safe remains locked.")
		}
		return
	}

	if decision.NextSceneID != "" {
		enterScene(decision.NextSceneID, player)
	}
}

func contains(slice []string, item string) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}
