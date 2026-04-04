/* Create a function that returns true only if the entire string is in ALL CAPS.*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Input a text: ")
	text := bufio.NewScanner(os.Stdin)
	for text.Scan() {
		line := text.Text()
		if strings.ToUpper(line) == line {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
		if len(line) == 0 {
			fmt.Println("Empty string is not allowed!")
		}
	}
}
