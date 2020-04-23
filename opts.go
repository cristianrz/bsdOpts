// Copyright 2020 Cristian Ariza. All rights reserved.
// License can be found in the LICENSE file.

/*
Package opts implements BSD-style command-line flag parsing.

Usage

Create a set of flags using

	args := make(opts.Opts)

Define each flags default value as

	args['b'] = false
	args['i'] = 3
	args['f'] = file.txt

Parse with

	args.Parse()

*/
package opts

import (
	"errors"
	"fmt"
	"os"
)

// Opts contains a map that includes the values of the opts
type Opts map[rune]interface{}

// Parse parses the command-line options from os.Args[1:].
func (m Opts) Parse() error {
	var (
		lastOpt rune
		expectingString = false
	)

	os.Args = os.Args[1:]
	argv := len(os.Args)

	for i := 0; i < argv; i++ {
		arg := os.Args[0]

		if expectingString {
			m[lastOpt] = arg
			expectingString = false
		} else if arg[0] == '-' {
			arg = arg[1:]

			for j, opt := range arg {
				if _, ok := m[opt]; !ok {
					return errors.New("unknown option -- " + string(opt))
				}

				switch valueType := fmt.Sprintf("%T", m[opt]); valueType {
				case "bool":
					m[opt] = true
				case "string":
					if j != len(arg)-1 {
						m[opt] = ""
					} else {
						expectingString = true
						lastOpt = opt
					}
				}
			}
		} else {
			return nil
		}

		os.Args = os.Args[1:]
	}

	return nil
}
