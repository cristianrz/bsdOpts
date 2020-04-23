package BSDOpts_test

import (
	"fmt"
	"github.com/cristianrz/BSDOpts"
	"os"
	"path"
	"testing"
)

func TestAll(t *testing.T) {
	os.Args = []string{os.Args[0], "-b", "-d", "-ae", "julian", "assange"}

	programName := os.Args[0]

	fmt.Printf("INITIAL\n\n")
	fmt.Println(os.Args)

	args := make(BSDOpts.Opts)

	args['a'] = false
	args['b'] = false
	args['c'] = false
	args['d'] = false
	args['e'] = "gazette"

	err := args.Parse()
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "%s: %s\n", path.Base(programName), err)
		if err != nil {
			panic(err)
		}

		os.Exit(1)
	}

	fmt.Printf("\nARGS\n\n")
	for k, v := range args {
		fmt.Println(string(k), ":", v)
	}

	fmt.Printf("\nLEFT\n\n")
	fmt.Println(os.Args)
}
