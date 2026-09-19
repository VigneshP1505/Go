package writeops

import (
	"fmt"
	"os"
)

func WriteBalanceToFile(balance float64, fileName string) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}
