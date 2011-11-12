// Copyright 2010  The "Go-Term" Authors
//
// Use of this source code is governed by the BSD 2-Clause License
// that can be found in the LICENSE file.
//
// This software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
// OR CONDITIONS OF ANY KIND, either express or implied. See the License
// for more details.

package term

import "os"

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
