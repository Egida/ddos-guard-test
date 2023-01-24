package tgbot

import (
	"strconv"
	"strings"
)

func validateMathResult(number float64) string {
	return strconv.FormatFloat(number, 'f', -1, 64)
}

func validateCalculate(text string) string {
	return strings.Replace(text, "/calculate ", "", -1)
}
