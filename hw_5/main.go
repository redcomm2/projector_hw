package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	lines, err := readTextFromFile()
	if err != nil {
		fmt.Println("Failed to read text from file:", err)
		return
	}

	word, err := getUserWord()
	if err != nil {
		fmt.Println("Failed to get user word:", err)
		return
	}

	searchMap, err := createSearchHashMap(lines)

	// Old search like in HW_4_1
	start := time.Now()
	oldSearchResult := oldSearch(lines, word)
	fmt.Printf("Old result number: %d", len(oldSearchResult))
	oldSearchDuration := time.Since(start)

	fmt.Printf("oldSearch took %v\n", oldSearchDuration)

	fmt.Println("-------------------")
	// End of Old search

	// New searchby hash
	start = time.Now()
	searchByHashMapResult, err := searchByMap(searchMap, word)
	fmt.Printf("Hash result number: %d", len(searchByHashMapResult))
	if err != nil {
		fmt.Println("Error with searchByMap:", err)
		return
	}
	searchByHashMapDuration := time.Since(start)

	fmt.Printf("searchByMap took %v\n", searchByHashMapDuration)
	//End of new search
}

func searchByMap(searchMap map[string][]string, userWord string) ([]string, error) {
	return searchMap[userWord], nil
}

func oldSearch(lines []string, userWord string) []string {

	var containingLines []string

	userWordLower := strings.ToLower(userWord)

	for _, line := range lines {
		words := strings.Fields(line)

		for _, word := range words {
			if strings.ToLower(word) == userWordLower {
				containingLines = append(containingLines, line)
				break
			}
		}
	}

	return containingLines
}

func getUserWord() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", fmt.Errorf("error reading user input: %w", err)
	}

	input = strings.TrimSpace(input)

	return strings.ToLower(input), nil
}

func createSearchHashMap(lines []string) (map[string][]string, error) {

	wordMap := make(map[string][]string)

	for _, line := range lines {
		lineWords := strings.Split(line, " ")

		for _, word := range lineWords {
			if !sliceContains(wordMap[word], line) {
				wordMap[word] = append(wordMap[word], line)
			}
		}
	}

	return wordMap, nil
}

func sliceContains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

func readTextFromFile() ([]string, error) {
	myfile, err := os.Open("text.txt")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer myfile.Close()

	scanner := bufio.NewScanner(myfile)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	return lines, nil
}
