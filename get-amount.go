package rupiah

import (
	"regexp"
	"strconv"
	"strings"
)

// GetAmount for decode Rp 1.000 to 1000
func GetAmount(rupiah string) int64 {
	re := regexp.MustCompile("[0-9]+")
	getNumber := re.FindAllString(rupiah, -1)
	amountString := strings.Join(getNumber, "")
	amount, err := strconv.ParseInt(amountString, 10, 64)
	if err != nil {
		return 0
	}
	return amount
}
