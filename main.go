package main

import (
	"fmt"
	"os"
)

func printUsage() {
	p := os.Args[0]
	fmt.Println("Usage:")
	fmt.Printf("\t%s FACES\n", p)
	fmt.Printf("\t%s DIE_KIND FACES\n", p)
	fmt.Printf("\t%s DIE_KIND FACES NUM_ROLLS\n", p)
	fmt.Println("WHERE:")
	fmt.Println("\tFACES is the max value of a die.")
	fmt.Println("\tDIE_KIND is the kind of a die. Valid kinds are:")
	fmt.Println("\t  1 - 1-based die")
	fmt.Println("\t  z - 0-based die")
	fmt.Println("\t  L - loaded die")
	fmt.Println("\tNUM_ROLLS is how many times the die will be rolled.")
	fmt.Println("\nExamples:")
	fmt.Printf("\t%s 6\n", p)
	fmt.Printf("\t%s 1 6\n", p)
	fmt.Printf("\t%s 1 6 1\n", p)
	fmt.Println("\tThese will roll a die once and yield a number between 1 and 6.")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	args, err := parseArgs(os.Args[1:])

	if err != nil {
		printUsage()
		fmt.Printf("\x1b[31merror:\x1b[0m %s.\n", err)
		os.Exit(2)
	}

	d := args.dieMaker(args.maxValue)

	for i := uint(0); i < args.numRolls; i++ {
		fmt.Println(d.Roll())
	}
}
