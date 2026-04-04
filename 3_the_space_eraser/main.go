/* Create a tool that removes all spaces from a sentence.*/

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
		fmt.Println(strings.Join(strings.Fields(text.Text()), ""))
	}
}
