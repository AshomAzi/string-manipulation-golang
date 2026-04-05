/* Take a full name like "Adaeze Okonkwo" and return "A.O."*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Input your name: ")
	name := bufio.NewScanner(os.Stdin)
	for name.Scan() {
		val := strings.Fields(name.Text())
		for i := 0; i < len(val); i++ {
			if len(val) > 0 {
				iName := val[i]
				secName := iName[0]
				fmt.Print(strings.ToUpper(string(secName) + "."))
			}
		}
		fmt.Println(" ")
	}
}
