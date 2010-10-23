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


func TestWinsize(t *testing.T) {
	ws, err := GetWinsize()
	if err != nil {
		t.Error(err)
	}

	fmt.Println("\n == Window size\n")
	fmt.Printf(" + Full: rows:%d, columns:%d, X pixels:%d, Y pixels:%d\n",
		ws.Row, ws.Col, ws.Xpixel, ws.Ypixel)

	row, col := GetWinsizeInChar()
	fmt.Printf(" + In characters: rows:%d, columns:%d\n", row, col)
}

