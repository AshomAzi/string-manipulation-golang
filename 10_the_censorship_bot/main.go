/* Create a function that replaces a specific "forbidden word" with [REDACTED].*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Input your sentence: ")
	text := bufio.NewScanner(os.Stdin)
	for text.Scan() {
		var newWord string
		line := text.Text()
		NewTest := strings.Fields(line)
		for i := range NewTest {
			if NewTest[i] == "Point" || NewTest[i] == "point" {
				NewTest[i] = "REDACTED"
				newWord += NewTest[i] + " "
			} else {
				newWord += NewTest[i] + " "
			}
		}
		fmt.Println(newWord)

	}

}
