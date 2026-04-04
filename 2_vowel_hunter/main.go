/*Vowel Hunter – Write a function that counts only the vowels in a string.*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

var vowels []string = []string{"a", "e", "i", "o", "u"}

func main() {
	fmt.Print("Input a string: ")
	text := bufio.NewScanner(os.Stdin)

	for text.Scan() {
		var counter int
		line := text.Text()
		for i := 0; i < len(line); i++ {
			for j := 0; j < len(vowels); j++ {
				if string(line[i]) == vowels[j] {
					counter += 1
				}
			}
		}
		fmt.Println(counter)
	}

}
