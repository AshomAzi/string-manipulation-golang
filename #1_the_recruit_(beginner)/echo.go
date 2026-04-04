/*1.	Operation: Echo Chamber –
Create a program that repeats a user's string back to them 5 times,
but each time it adds an extra exclamation mark.*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Write you test: ")
	text := bufio.NewScanner(os.Stdin)

	for text.Scan(){
		line := text.Text()
		for i:= 1; i <= 5; i++{
			fmt.Println(line+strings.Repeat("!", i))
		}
	}
}

