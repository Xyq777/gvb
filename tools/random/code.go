package random

import (
	"math/rand"
	"strconv"
)

func NewCodeSix() string {
	code := rand.Intn(900000) + 100000
	return strconv.Itoa(code)
}
