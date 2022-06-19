package main

import (
	"fmt"

	"github.com/gotneb/bcapp/nibblelinx"

	"math/big"
)

func main() {
	nibblelinx.Init()

	// Only for very small tests, the real ones will come later
	x1, _ := new(big.Int).SetString("15", 10)
	x2, _ := new(big.Int).SetString("48", 10)

	fmt.Printf("Point[0]: %s\nPoint[1]: %s\n", nibblelinx.DoubleP(x1, x2)[0], nibblelinx.DoubleP(x1, x2)[1])
}
