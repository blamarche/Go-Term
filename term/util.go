// Copyright 2010  The "go-term" Authors
//
// Use of this source code is governed by the Simplified BSD License
// that can be found in the LICENSE file.
//
// This software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
// OR CONDITIONS OF ANY KIND, either express or implied. See the License
// for more details.

package term

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var unsupportedTerm = []string{"dumb", "cons25"}


// Checks if the terminal supports ANSI terminal escape controls.
func HandleANSI() bool {
	term := os.Getenv("TERM")
	if term == "" {
		return false
	}

	for _, value := range unsupportedTerm {
		if value == term {
			return false
		}
	}
	return true
}

// Reads the key pressed. The argument `prompt` is written to standard output,
// if any.
func ReadKey(prompt string) (rune int, err os.Error) {
	in, err := bufio.NewReaderSize(os.Stdin, 1)
	if err != nil {
		return 0, err
	}

	if prompt != "" {
		fmt.Print(prompt)
	}

	rune, _, err = in.ReadRune()
	if err != nil {
		return 0, err
	}

	return rune, nil
}

// Reads a line from input until Return is pressed (stripping a trailing
// newline), and returns that. The argument `prompt` is written to standard
// output, if any.
func ReadString(prompt string) (line string, err os.Error) {
	in := bufio.NewReader(os.Stdin)

	if prompt != "" {
		fmt.Print(prompt)
	}

	line, err = in.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimRight(line, "\n"), nil
}

