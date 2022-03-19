package main

import (
	"fmt"
	"strconv"

	"github.com/Fawers/dice-go/dice"
)

type args struct {
	dieMaker func(uint64) dice.Die
	maxValue uint64
	numRolls uint
}

func parseArgs(argv []string) (*args, error) {
	a := args{dice.New1Based, 6, 1}

	switch len(argv) {
	case 3:
		rolls := argv[2]
		numRolls, err := strconv.ParseUint(rolls, 10, 0)

		if err != nil {
			return nil, fmt.Errorf(
				"failed to understand NUM_ROLLS: %q is not a uint", rolls)
		}
		a.numRolls = uint(numRolls)
		fallthrough

	case 2:
		dieKind := argv[0]
		switch dieKind {
		case "1":
			break

		case "z":
			a.dieMaker = dice.New0Based

		case "L":
			a.dieMaker = dice.NewLoaded

		default:
			return nil, fmt.Errorf(
				"%q is not a valid die kind", dieKind)
		}

		argv[0] = argv[1]
		fallthrough

	case 1:
		faces := argv[0]
		var err error
		a.maxValue, err = strconv.ParseUint(faces, 10, 64)

		if err != nil {
			return nil, fmt.Errorf(
				"failed to understand FACES: %q is not a uint", faces)
		}
	}

	return &a, nil
}
