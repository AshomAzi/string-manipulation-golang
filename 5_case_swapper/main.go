/*– Write a function that turns lowercase letters to uppercase and vice versa in the same string.*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Input your text: ")
	text := bufio.NewScanner(os.Stdin)

	for text.Scan() {
		line := text.Text()
		var new string
		for i := 0; i < len(line); i++ {
			if line[i] >= 'a' && line[i] <= 'z' {
				new += strings.ToUpper(string(line[i]))
			} else if line[i] >= 'A' && line[i] <= 'Z' {
				new += strings.ToLower(string(line[i]))
			}
		}
		fmt.Println(new)
	}
}
