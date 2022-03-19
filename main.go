package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Fawers/dice-go/dice"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Printf("  %s NUMBER\n", os.Args[0])
	fmt.Println("WHERE:")
	fmt.Println("\tNUMBER is the max value of a Regular 1-based die.")
	fmt.Println("\nThis will roll a die and yield a number between 1 and NUMBER.")
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
		os.Exit(1)
	}

	num, err := strconv.ParseUint(os.Args[1], 10, 64)

	if err != nil {
		printUsage()
		fmt.Printf("\x1b[31merror:\x1b[0m could not convert %q to uint64.\n", os.Args[1])
		os.Exit(2)
	}

	d := dice.New1Based(num)
	fmt.Println(d.Roll())
}
