package fizzbuzz

import (
	"errors"
	"testing"
)

func TestClassicalFizzBuzz(t *testing.T) {
	expected := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8",
		"fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17",
		"fizz", "19", "buzz"}

	output, err := FizzBuzz(3, 5, 20, "fizz", "buzz")
	if err != nil {
		t.Errorf("Could not execute classical fizz buzz, err: %s", err.Error())
		return
	}
	if len(expected) != len(output) {
		t.Errorf("Output and expected should have the same length, output length is: %d", len(output))
		return
	}
	for index, v := range expected {
		if output[index] != v {
			t.Errorf("Output and expected should be the same, output at index %d is %s, expected is %s", index, output[index], v)
			return
		}
	}
}

func TestFizzBuzzLimitError(t *testing.T) {
	negativeLimit := -20
	output, err := FizzBuzz(3, 5, negativeLimit, "fizz", "buzz")
	if output != nil {
		t.Error("Limit is negative, output should not be available")
		return
	}
	if !errors.Is(err, ErrLimitNegative) {
		t.Errorf("Error is %v, should be %s", err, ErrLimitNegative.Error())
	}
}

func TestFizzBuzzIntegerInputError(t *testing.T) {
	nullInput := 0
	output1, err := FizzBuzz(nullInput, 5, 20, "fizz", "buzz")
	if output1 != nil {
		t.Error("First input is negative, output should not be available")
		return
	}
	if !errors.Is(err, ErrZeroLimit) {
		t.Errorf("Error is %v, should be %s", err, ErrZeroLimit.Error())
	}
	output2, err := FizzBuzz(3, nullInput, 20, "fizz", "buzz")
	if output2 != nil {
		t.Error("Second input is negative, output should not be available")
		return
	}
	if !errors.Is(err, ErrZeroLimit) {
		t.Errorf("Error is %v, should be %s", err, ErrZeroLimit.Error())
	}
}

// FuzzFizzBuzz tests that the output is the same with negative values as
// with positive values
func FuzzFizzBuzz(f *testing.F) {
	limit := 20
	inputs := []int{3, 10, 0, -2}
	for _, input := range inputs {
		f.Add(input)
	}
	f.Fuzz(func(t *testing.T, input int) {
		output1, err := FizzBuzz(input, 3, limit, "fizz", "buzz")
		if input == 0 {
			if !errors.Is(err, ErrZeroLimit) {
				t.Errorf("Error is %v, should be %s", err, ErrZeroLimit.Error())
				return
			}
			return
		}
		if err != nil {
			t.Errorf("Could not execute first fizzbuzz with input %d, err: %s", input, err.Error())
			return
		}
		output2, err := FizzBuzz(-input, 3, limit, "fizz", "buzz")
		if err != nil {
			t.Errorf("Could not execute second fizzbuzz with input %d, err: %s", -input, err.Error())
			return
		}
		if len(output1) != len(output2) {
			t.Errorf("outputs should have the same length, lengths are %d and %d", len(output1), len(output2))
			return
		}
		for index, v1 := range output1 {
			v2 := output2[index]
			if v1 != v2 {
				t.Errorf("outputs should be the same, outputs at index %d are %s and %s", index, v1, v2)
				return
			}
		}
	})
}
