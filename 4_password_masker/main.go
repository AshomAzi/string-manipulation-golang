/* – Take a string and return it replaced entirely by asterisks (e.g., "secret" becomes ******).*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Input a string: ")
	text := bufio.NewScanner(os.Stdin)
	for text.Scan() {
		line := text.Text()
		new := strings.Repeat("*", len(line))
		fmt.Println(new)
	}
}
