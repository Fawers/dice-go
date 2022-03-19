package dice

import "testing"

// using 16 as seed for a D21 takes
// exactly 36 rolls to roll every single number.
const magicSeed int64 = 16
const magicNumber int = 36

func all(bs []bool) bool {
	for _, b := range bs {
		if !b {
			return false
		}
	}

	return true
}

func TestLoadedDieRollsTheSameNumberEverySingleTime(t *testing.T) {
	d := NewLoaded(42)
	expected := uint64(42)

	for i := 0; i < 100; i++ {
		// roll 42 all 100 iterations
		if roll := d.Roll(); roll != expected {
			t.Errorf("should've rolled %d, rolled %d", expected, roll)
		}
	}
}

func TestLoadedDieMaxAndMinValuesAreEqual(t *testing.T) {
	d := NewLoaded(4)

	if d.MinValue() != d.MaxValue() || d.MaxValue() != d.Roll() {
		t.Fatalf("loaded dice roll, min and max values are bogus")
	}
}

func TestRegular0Based21FacedDieRollsAllNumbersIn36Rolls(t *testing.T) {
	d := New0BasedWithSeed(21, magicSeed)
	var rolls [21]bool

	for i := 0; i < magicNumber; i++ {
		r := d.Roll()
		rolls[r] = true
	}

	if !all(rolls[:]) {
		t.Fatalf("did not roll all numbers of a D21 in %d rolls", magicNumber)
	}
}

func TestRegular1Based6FacedDieMinAndMaxValuesAre1And6(t *testing.T) {
	d := New1Based(6)

	if d.MinValue() != 1 {
		t.Errorf("min value should be 1; is %d", d.MinValue())
	}

	if d.MaxValue() != 6 {
		t.Fatalf("max value should be 6; is %d", d.MaxValue())
	}
}
