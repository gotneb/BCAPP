package main

import (
	"fmt"

	"github.com/gotneb/bcapp/nibblelinx"

	"math/big"
)

func main() {
	nibblelinx.Init()

	x1, _ := new(big.Int).SetString("30", 10)
	x2, _ := new(big.Int).SetString("22", 10)

	fmt.Println("Point: ", nibblelinx.DoubleP(x1, x2)[0])
}
