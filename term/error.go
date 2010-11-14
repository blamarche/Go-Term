// Copyright 2010  The "go-term" Authors
//
// Use of this source code is governed by the Simplified BSD License
// that can be found in the LICENSE file.
//
// This software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
// OR CONDITIONS OF ANY KIND, either express or implied. See the License
// for more details.

package term


// Represents a failure into a TTY.
type TTYerror string

func (t TTYerror) String() string {
	return "not on a tty: " + string(t)
}

