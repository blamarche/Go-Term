// Copyright 2010  The "Go-Term" Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
