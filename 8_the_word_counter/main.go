/* Count how many words are in a string (handle multiple spaces between words!).*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input a text: ")
	text := bufio.NewScanner(os.Stdin)
	for text.Scan() {
		line := text.Text()
		var counter string

		for i := 0; i < len(line); i++ {
			if string(line[i]) != " " && len(line) > 0 {
				counter += string(line[i])
			} else if line == "exit" {
				break
			}
		}
		fmt.Println(len(counter))
	}
}
