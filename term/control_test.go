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
	//"fmt"
	"testing"
	//"time"
)


func TestTTY(t *testing.T) {
	if CheckIsatty(stdin) != nil {
		t.Error("should be a terminal")
	}
}

func TestSettings(t *testing.T) {
	term := newTermios()

	if tcgetattr(stdin, term) != nil {
		t.Error("could not retrieve the terminal settings")
	}
}

func TestTTYname(t *testing.T) {
	if _, err := TTYname(stdin); err != nil {
		t.Error("should get the terminal name", err)
	}
}

func TestTerminal(t *testing.T) {
	if !HandleANSI() {
		t.Error("this terminal should be supported")
	}
}
/*
func TestUtility(t *testing.T) {
	fmt.Println("\n == Terminal")

	// === Read single key
	KeyPress()
	in := ReadKey("\n + Mode on single character: ")

	fmt.Printf("\n  pressed: %s", string(in))
	RestoreTermios()

	// === Echo
	fmt.Print("\n + Echo disabled. Try write\n  ")
	Echo(false)
	time.Sleep(5e9)

	fmt.Print(" Echo enabled")
	Echo(true)

	// === Raw mode
	fmt.Printf("\n + Raw mode\n  ")
	MakeRaw(stdin)
	time.Sleep(5e9)
	RestoreTermios()
}
*/
