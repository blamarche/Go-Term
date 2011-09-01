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

func TestTerminal(t *testing.T) {
	if !HandleANSI() {
		t.Error("this terminal should be supported")
	}
}
