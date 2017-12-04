package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	canvas := Canvas{0, 0, []string{}}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("enter command: ")
		text, _ := reader.ReadString('\n')
		text = strings.Trim(strings.Trim(text, "\n"), "\r")
		cmd := Parse(text)
		cmd.Execute(&canvas)
		for i := 0; i < canvas.height; i++ {
			for j := 0; j < canvas.width; j++ {
				el, _ := canvas.Element(j, i)
				fmt.Print(el)
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}
