package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	myfile, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer myfile.Close()

	scanner := bufio.NewScanner(myfile)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var containingLines []string
	for _, element := range lines {

		if strings.Contains(element, input) {
			containingLines = append(containingLines, element)
		}
	}

	for _, line := range containingLines {
		fmt.Println(line)
	}
}
