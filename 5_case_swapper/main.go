/*– Write a function that turns lowercase letters to uppercase and vice versa in the same string.*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
start1:
	fmt.Print("Input your text or exit to cancel: ")
	text := bufio.NewScanner(os.Stdin)

	for text.Scan() {
		line := text.Text()
		var new strings.Builder
		if line == "exit" || line == "EXIT" {
			break
		} else if len(line) > 0 {
			for i := 0; i < len(line); i++ {
				if line[i] >= 'a' && line[i] <= 'z' {
					new.WriteString(strings.ToUpper(string(line[i])))
				} else if line[i] >= 'A' && line[i] <= 'Z' {
					new.WriteString(strings.ToLower(string(line[i])))
				}
			}
			fmt.Println(new.String())
		} else {
			fmt.Println("Invalid Input, try again!")
			goto start1
		}

	}
}
