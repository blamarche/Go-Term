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
