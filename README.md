# dice-go

Roll dice with Go.

## Basic usage
```go
import (
	"fmt"

	"github.com/Fawers/dice-go/dice"
)

func main() {
	die := dice.New1Based(6)
	roll := die.Roll()
	fmt.Printf("Rolled a %d\n", roll)
}
```

## Terminology
| Term | Description |
|:-:|:-:
| 0-based die | A die whose rollings yield any number between 0 and its max value minus 1. A 6 faced 0 based die will yield numbers between 0 and 5. |
| 1-based die | A die whose rollings yield any number between 1 and its max value. A 6 faced 1 based die will yield numbers between 1 and 6. |
| loaded die | A loaded/weighted die that always yields the same number. |

# Inspiration
I was teaching Go to some of my students. While teaching types, structs and
interfaces, I had to come up with an example why one would use interfaces and
how one would use different structs to implement their methods. Dice came to
mind. Maybe someone out there may want to use this lib to program a Yahtzee
game in Go?

# CLI Usage
```
$ go build main.go args.go 
$ ./main 
Usage:
        ./main FACES
        ./main DIE_KIND FACES
        ./main DIE_KIND FACES NUM_ROLLS
WHERE:
        FACES is the max value of a die.
        DIE_KIND is the kind of a die. Valid kinds are:
          1 - 1-based die
          z - 0-based die
          L - loaded die
        NUM_ROLLS is how many times the die will be rolled.

Examples:
        ./main 6
        ./main 1 6
        ./main 1 6 1
        These will roll a die once and yield a number between 1 and 6.
```
