/* Check if a string ends with a question mark. If not, add one.*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input a string: ")
	text := bufio.NewScanner(os.Stdin)
	for text.Scan() {
		line := text.Text()
		if string(line[len(line) -1 ]) == "?" {
			fmt.Println(line)
		} else if line == "exit" {
			break
		} else {
			fmt.Println(line+"?")
		}
	}
}
