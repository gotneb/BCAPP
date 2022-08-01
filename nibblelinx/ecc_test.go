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
	expected [2]*big.Int
}

var tests []*Test = make([]*Test, 0)
var source []*big.Int = make([]*big.Int, 0)

// Ecc content files
var modp []*big.Int = make([]*big.Int, 0)
var inverse []*big.Int = make([]*big.Int, 0)
var doublep []*big.Int = make([]*big.Int, 0)

// get a file content from a file and storage into dst
func loadFiles(filename string, dst *[]*big.Int) {
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
	Init()
	loadFiles("/home/gabriel/Documents/Development/Go/bcapp/files_tests/source", &source)
	loadFiles("/home/gabriel/Documents/Development/Go/bcapp/files_tests/modp_output", &modp)

	if len(tests) != 0 {
		tests = make([]*Test, 0)
	}

	for i := 0; i < len(source)-1; i += 2 {
		tests = append(tests, &Test{
			source[i],
			source[i+1],
			[2]*big.Int{new(big.Int), new(big.Int)}},
		)
	}
	for i := 0; i < len(modp); i++ {
		tests[i].expected[0] = modp[i]
	}
	for _, test := range tests {
		if out := ModP(test.arg1, test.arg2); out.Cmp(test.expected[0]) != 0 {
			t.Errorf("\nTest Failed:\narg_1: %s\narg_2: %s\nExpected: %s\nReceveid: %s\n", test.arg1, test.arg2, test.expected[0], out)
		}
	}
}

func TestInverse(t *testing.T) {
	Init()
	loadFiles("/home/gabriel/Documents/Development/Go/bcapp/files_tests/source", &source)
	loadFiles("/home/gabriel/Documents/Development/Go/bcapp/files_tests/inverse_output", &inverse)

	if len(tests) != 0 {
		tests = make([]*Test, 0)
	}

	for i := 0; i < len(source)-1; i += 2 {
		tests = append(tests, &Test{
			source[i],
			source[i+1],
			[2]*big.Int{new(big.Int), new(big.Int)}},
		)
	}

	for i := 0; i < len(inverse); i++ {
		tests[i].expected[0] = inverse[i]
	}

	for _, test := range tests {
		if out := Inverse(test.arg1, test.arg2); out.Cmp(test.expected[0]) != 0 {
			t.Errorf("\nTest Failed:\narg_1: %s\narg_2: %s\nExpected: %s\nReceveid: %s\n", test.arg1, test.arg2, test.expected[0], out)
		}
	}
}

// TODO: These 2 functions above they have almost the same code, do only one funciton who does the entire work.

func TestDoubleP(t *testing.T) {
	Init()
	loadFiles("/home/gabriel/Documents/Development/Go/bcapp/files_tests/source", &source)
	loadFiles("/home/gabriel/Documents/Development/Go/bcapp/files_tests/doublep_output", &doublep)

	if len(tests) != 0 {
		tests = make([]*Test, 0)
	}

	for i := 0; i < len(source)-1; i += 2 {
		tests = append(tests, &Test{
			source[i],
			source[i+1],
			[2]*big.Int{doublep[i], doublep[i+1]}},
		)
	}

	for _, test := range tests {
		if out := DoubleP(test.arg1, test.arg2); (out[0].Cmp(test.expected[0]) != 0) || (out[1].Cmp(test.expected[1]) != 0) {
			t.Errorf("\nTest Failed:\narg_1: %s\narg_2: %s\nExpected:\n[0]: %s\n[1]: %s\nReceveid:\n[0]: %s\n[1]: %s\n", test.arg1, test.arg2, test.expected[0], test.expected[1], out[0], out[1])
		}
	}
}
