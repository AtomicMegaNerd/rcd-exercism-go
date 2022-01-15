package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

type SillyNephewError struct {
	numCows int
}

func (s SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.numCows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}
	if cows < 0 {
		return 0.0, SillyNephewError{cows}
	}

	fodder, err := weightFodder.FodderAmount()
	if err == ErrScaleMalfunction {
		fodder *= 2
	} else if err != nil {
		return 0.0, err
	}
	if fodder < 0 {
		return 0.0, errors.New("negative fodder")
	}

	return fodder / float64(cows), nil
}
