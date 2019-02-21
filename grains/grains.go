package grains

import "errors"

// square returns grains of rice on the chessboard

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("number must be between 1 and 64")
	}

	return (1 << uint(number-1)), nil
}

// Total returns the number of grains of rice on the entire mythical
// chessboard.
func Total() uint64 {
	return 1<<64 - 1

}
