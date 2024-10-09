package main

import (
	c "mosdef/funcs"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 4 || len(os.Args) < 2 {
		fmt.Println("please enter at most one sentence, one color and one substring.")
		return
	}

	var argument string
	var color string
	var subString string
	RGB := [3]int{}

	if len(os.Args) == 4 {
		argument = os.Args[3]
		color = os.Args[1]
		subString = os.Args[2]
		RGB = c.Color(color)
	} else if len(os.Args) == 3 {
		argument = os.Args[2]
		color = os.Args[1]
		RGB = c.Color(color)
	} else if len(os.Args) == 2 {
		argument = os.Args[1]
	} else {
		fmt.Println("invalid input")
	}

	ascii, er1 := os.ReadFile("standard.txt")
	if er1 != nil {
		log.Fatal(er1)
	}

	R := RGB[0]
	G := RGB[1]
	B := RGB[2]
	lines := strings.Split(string(ascii), "\n")
	sentence := strings.Split(argument, "\\n")

	for _, sentence := range sentence {
		starts := c.SubStr(sentence, subString)
		x := []int{}

		for _, i := range sentence {
			x = append(x, int(i))
		}

		// Validate ASCII characters
		for _, ran := range x {
			if ran < 32 || ran > 126 {
				fmt.Println("invalid input")
				return
			}
		}

		// Loop through each line of the ASCII art representation
		for k := 1; k <= 8; k++ {
			for i, ch := range x {
				lineIndex := (ch-32)*9 + k
				line := lines[lineIndex]

				// Check if we need to color the substring or the whole sentence
				if (subString != "" && inRange(i, starts, len(subString))) || subString == "" {
					// Color the entire sentence if subString is empty
					fmt.Printf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", R, G, B, line)
				} else {
					// Print line without color
					fmt.Print(line)
				}

				if i == len(x)-1 {
					fmt.Println()
				}
			}
		}
	}
}

// Function to check if the index is within the range of substring starts
func inRange(i int, starts []int, subLength int) bool {
	for _, start := range starts {
		if i >= start && i < start+subLength {
			return true
		}
	}
	return false
}
