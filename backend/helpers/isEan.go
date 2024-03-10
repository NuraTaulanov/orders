package helpers

import (
	"strconv"
)

func IsEan(str string) bool {
	if IsNumber(str) && len(str) == 13 {
		return true
	} else {
		return false
	}

}

func IsNumber(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	} else {
		return false
	}
}
