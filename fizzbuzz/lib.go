package fizzbuzz

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrLimitNegative = errors.New("wrong value error, limit should be at least 1")
var ErrZeroLimit = errors.New("wrong value error, integer input cannot be 0")

func FizzBuzz(int1 int, int2 int, limit int, str1 string, str2 string) ([]string, error) {
	if limit < 1 {
		return nil, ErrLimitNegative
	}
	if int1 == 0 {
		return nil, fmt.Errorf("int1 is wrong: err: %w", ErrZeroLimit)
	}
	if int2 == 0 {
		return nil, fmt.Errorf("int2 is wrong: err: %w", ErrZeroLimit)
	}

	list := make([]string, limit)
	intProduct := int1 * int2
	strProduct := fmt.Sprintf("%s%s", str1, str2)
	for i := 0; i < limit; i++ {
		value := i + 1
		if value%intProduct == 0 {
			list[i] = strProduct
		} else if value%int1 == 0 {
			list[i] = str1
		} else if value%int2 == 0 {
			list[i] = str2
		} else {
			list[i] = strconv.Itoa(value)
		}
	}
	return list, nil
}
