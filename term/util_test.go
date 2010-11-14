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
	"fmt"
	"testing"
)


func TestTerminal(t *testing.T) {
	if !HandleANSI() {
		t.Error("this terminal should be supported")
	}
}

func TestReading(t *testing.T) {
	fmt.Println("\n == Terminal")

	// === Read single key
	KeyPress()
	rune, _ := ReadKey("\n + Mode on single character: ")

	fmt.Printf("\n  pressed: %q", string(rune))
	RestoreTerm()

	// === Echo
	Echo(false)
	fmt.Print("\n + Echo disabled. Write and press Enter: ")
	line, _ := ReadString("")

	fmt.Printf("\n  pressed: %q\n", line)
	fmt.Println(" + Echo enabled")
	Echo(true)
}

func TestReadSequence(t *testing.T) {
	fmt.Println("\n == Read sequence")
	fmt.Println(" Press Enter to exit\n")

	for {
		line, _ := ReadString(" + enter: ")
		if line == "" {
			break
		}
		fmt.Printf("   code: %q\n\n", line)
	}
}

