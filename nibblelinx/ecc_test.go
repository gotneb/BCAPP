package nibblelinx

import (
	"bufio"
	"log"
	"math/big"
	"os"
	"testing"
)

type Test struct {
	arg1     *big.Int
	arg2     *big.Int
	expected *big.Int
}

var source []*big.Int = make([]*big.Int, 0)
var modp []*big.Int = make([]*big.Int, 0)
var tests []*Test = make([]*Test, 0)

func catchLines(filename string, dst *[]*big.Int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, _ := new(big.Int).SetString(scanner.Text(), 10)
		*dst = append(*dst, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func TestModP(t *testing.T) {
	catchLines("/home/gabriel/Documents/Development/Go/bcapp/files_tests/source", &source)
	catchLines("/home/gabriel/Documents/Development/Go/bcapp/files_tests/modp_output", &modp)
	for i := 0; i < len(source)-1; i += 2 {
		tests = append(tests, &Test{
			source[i],
			source[i+1],
			new(big.Int)},
		)
	}
	for i := 0; i < len(modp); i++ {
		tests[i].expected = modp[i]
	}
	for _, test := range tests {
		if out := ModP(test.arg1, test.arg2); out.Cmp(test.expected) != 0 {
			t.Errorf("\nTest Failed:\narg_1: %s\narg_2: %s\nExpected: %s\nReceveid: %s\n", test.arg1, test.arg2, test.expected, out)
		}
	}
}
