package dice

import (
	"math/rand"
	"time"
)

// Die is the interface that provides the methods Roll, MaxValue,
// and MinValue. This package provides (private) implementors that
// are returned through the New* functions.
type Die interface {
	// Roll rolls a new number.
	Roll() uint64

	// MaxValue returns the max number this die will ever roll,
	// inclusive.
	MaxValue() uint64

	// MinValue returns the min number this die will ever roll.
	MinValue() uint64
}

// New0Based returns a new die that will rolls numbers
// between 0 and `maxValue` exclusive. I.e., 0 <= die.Roll() < `maxValue`.
//
// New0Based calls New0BasedWithSeed passing time.Now().UnixNano()
// as the seed.
func New0Based(maxValue uint64) Die {
	return New0BasedWithSeed(maxValue, time.Now().UnixNano())
}

// New0BasedWithSeed returns a new die that will roll numbers
// between 0 and `maxValue` exclusive. I.e., 0 <= die.Roll() < `maxValue`.
//
// The internal rand.Rand object will use `seed` as its seed.
func New0BasedWithSeed(maxValue uint64, seed int64) Die {
	d := new(regular0Die)
	d.max = maxValue
	d.rand = rand.New(rand.NewSource(seed))
	return d
}

// New1Based returns a new die that will rolls numbers
// between 1 and `maxValue` inclusive. I.e., 1 <= die.Roll() <= `maxValue`.
//
// New1Based calls New1BasedWithSeed passing time.Now().UnixNano()
// as the seed.
func New1Based(maxValue uint64) Die {
	return New1BasedWithSeed(maxValue, time.Now().UnixNano())
}

// New1BasedWithSeed returns a new die that will roll numbers
// between 1 and `maxValue` inclusive. I.e., 1 <= die.Roll() <= `maxValue`.
//
// The internal rand.Rand object will use `seed` as its seed.
func New1BasedWithSeed(maxValue uint64, seed int64) Die {
	d := new(regular1Die)
	d.r0d = regular0Die{
		max:  maxValue,
		rand: rand.New(rand.NewSource(seed)),
	}
	return d
}

// NewLoaded returns a loaded die that always returns the same
// number, no matter how many times you roll it. Can be used to
// test code that uses dice, or to cheat in dice games.
func NewLoaded(constantValue uint64) Die {
	d := new(loadedDie)
	d.constant = constantValue
	return d
}

type regular0Die struct {
	max  uint64
	rand *rand.Rand
}

func (d *regular0Die) Roll() uint64 {
	num := d.rand.Uint64()
	roll := num % d.max
	return roll
}

func (d *regular0Die) MaxValue() uint64 {
	return d.max - 1
}

func (d *regular0Die) MinValue() uint64 {
	return 0
}

type regular1Die struct {
	r0d regular0Die
}

func (d *regular1Die) Roll() uint64 {
	return d.r0d.Roll() + 1
}

func (d *regular1Die) MaxValue() uint64 {
	return d.r0d.max
}

func (d *regular1Die) MinValue() uint64 {
	return 1
}

type loadedDie struct {
	constant uint64
}

func (d *loadedDie) Roll() uint64 {
	return d.constant
}

func (d *loadedDie) MaxValue() uint64 {
	return d.constant
}

func (d *loadedDie) MinValue() uint64 {
	return d.constant
}
