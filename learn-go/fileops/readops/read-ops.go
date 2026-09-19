package readops

import (
	"errors"
	"os"
	"strconv"
)

func GetBalanceFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 0, errors.New("error reading balance sheet")
	}

	balance, _ := strconv.ParseFloat(string(data), 64)
	return balance, nil
}
