/* Given an email address, extract only the domain name (e.g., "agent@sentinel.org" -> "sentinel.org").*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input your email: ")
	email := bufio.NewScanner(os.Stdin)
	for email.Scan() {
		em := email.Text()
		for i := 0; i < len(em); i++ {
			if string(em[i]) == "@" {
				fmt.Println(em[i:])
			}
		}
	}
}