package main

import (
	"bufio"
	"fmt"
	"os"

	calc "github.com/KoyamaSohei/calculator"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		v, err := calc.Eval(s)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Answer: %d\n", v)
	}
	if scanner.Err() != nil {
		os.Exit(1)
	}
}
