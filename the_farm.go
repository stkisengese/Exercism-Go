package exercism

import (
	"errors"
	"fmt"
)

type Calculator struct{}

// FodderAmount(int) method gives the total fodder required.
func (c Calculator) FodderAmount(int) (float64, error) {
	return 50.0, nil
}

// FatteningFactor method returns the multiplier.
func (c Calculator) FatteningFactor() (float64, error) {
	return 1.5, nil
}

// ValidateInputAndDivideFood function to validate number of cows and DividFood
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, cows)
}

// DivideFood divides the total fodder by number of cows and mutiply by fattening factor.
func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	if fatteningFactor, err := fc.FatteningFactor(); err != nil {
		return 0, err
	} else if fodderAmount, err := fc.FodderAmount(cows); err != nil {
		return 0, err
	} else {
		return (fodderAmount * fatteningFactor) / float64(cows), nil
	}
}

// InvalidCowsError for custom message for invalid number of cows.
type InvalidCowsError struct {
	cowsNumber    int
	customMessage string
}

// Error method for InvalidCowsError.
func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cowsNumber, e.customMessage)
}

// ValidateNumberOfCows checks if the number of cows is valid.
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cowsNumber: cows, customMessage: "there are no negative cows"}
	} else if cows == 0 {
		return &InvalidCowsError{cowsNumber: cows, customMessage: "no cows don't need food"}
	}
	return nil
}

// This file contains types used in the exercise and tests and should not be modified.

// FodderCalculator provides helper methods to determine the optimal
// amount of fodder to feed cows.
type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}
