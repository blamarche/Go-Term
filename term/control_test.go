// Copyright 2010  The "go-term" Authors
//
// Use of this source code is governed by the BSD-2 Clause license
// that can be found in the LICENSE file.
//
// This software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
// OR CONDITIONS OF ANY KIND, either express or implied. See the License
// for more details.

package term

import "testing"


func TestTTY(t *testing.T) {
	if CheckIsatty(stdin) != nil {
		t.Error("should be a terminal")
	}
}

func TestSettings(t *testing.T) {
	term := newTermios(stdin)

	if term.tcgetattr() != nil {
		t.Error("could not retrieve the terminal settings")
	}
}

func TestTTYname(t *testing.T) {
	if _, err := TTYname(stdin); err != nil {
		t.Error("should get the terminal name", err)
	}
}
