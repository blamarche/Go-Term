// Copyright 2010  The "go-term" Authors
//
// Use of this source code is governed by the Simplified BSD License
// that can be found in the LICENSE file.
//
// This software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
// OR CONDITIONS OF ANY KIND, either express or implied. See the License
// for more details.

package term

/* Terminal Information and Control. */

// #include <termios.h>
// #include <unistd.h>
import "C"

import (
	"os"
)


var stdin = 0

var (
	IsRawMode       = false        // To check if restore is needed.
	origTermios     = newTermios() // In order to restore the original settings.
	unsupportedTerm = []string{"dumb", "cons25"}
)


// === Init
// ===

func init() {
	// Store the actual terminal settings.
	if err := tcgetattr(stdin, origTermios); err != nil {
		panic("terminal settings could not be got: " + err.String())
	}
}


// === Wrappers around types.
// ===

// Allows to manipulate the terminal.
type termios struct {
	wrap *_Ctype_struct_termios
}

// Deep copy for pointer fields.
func (self *termios) copyto(to *termios) {
	*to.wrap = *self.wrap
}

func newTermios() *termios {
	return &termios{new(_Ctype_struct_termios)}
}


// === Wrappers around functions.
// ===

// Determines if the device is a terminal.
//
// int isatty(int fd);
func Isatty(fd int) bool {
	exitCode, _ := C.isatty(C.int(fd))

	if exitCode == 1 {
		return true
	}
	return false
}

// Determines if the device is a terminal. Return an error, if any.
//
// int isatty(int fd);
func CheckIsatty(fd int) os.Error {
	exitCode, errno := C.isatty(C.int(fd))

	if exitCode == 1 {
		return nil
	}
	return errno
}

// Gets terminal state.
//
// int tcgetattr(int fd, struct termios *termios_p);
func tcgetattr(fd int, termios_p *termios) os.Error {
	exitCode, errno := C.tcgetattr(C.int(fd), termios_p.wrap)

	if exitCode == 0 {
		return nil
	}
	return errno
}

// Sets terminal state.
//
// int tcsetattr(int fd, int optional_actions, const struct termios *termios_p);
func tcsetattr(fd int, optional_actions int, termios_p *termios) os.Error {
	exitCode, errno := C.tcsetattr(C.int(fd), C.int(optional_actions), termios_p.wrap)

	if exitCode == 0 {
		return nil
	}
	return errno
}

// Gets the name of a terminal.
//
// char *ttyname(int fd)
func TTYname(fd int) (string, os.Error) {
	name, errno := C.ttyname(C.int(fd))
	if errno != nil {
		return "", errno
	}
	return C.GoString(name), nil
}


// === Utility
// ===

// Turns the echo mode.
func Echo(echo bool) {
	if !echo {
		origTermios.wrap.c_lflag &^= (C.ECHO)
	} else {
		origTermios.wrap.c_lflag |= (C.ECHO)
	}

	if err := tcsetattr(stdin, C.TCSANOW, origTermios); err != nil {
		panic("echo mode")
	}
}

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

// Sets the terminal to single-character mode.
func KeyPress() {
	newSettings := newTermios()
	origTermios.copyto(newSettings)

	// Disable canonical mode, and set buffer size to 1 byte.
	newSettings.wrap.c_lflag &^= (C.ICANON)
	newSettings.wrap.c_cc[C.VTIME] = 0
	newSettings.wrap.c_cc[C.VMIN] = 1

	if err := tcsetattr(stdin, C.TCSANOW, newSettings); err != nil {
		panic("single-character mode")
	}
}

// Sets the terminal to something like the "raw" mode. Input is available
// character by character, echoing is disabled, and all special processing of
// terminal input and output characters is disabled.
//
// Based in C call: void cfmakeraw(struct termios *termios_p)
func MakeRaw(fd int) os.Error {
	raw := newTermios()
	origTermios.copyto(raw)
	IsRawMode = true

	// Input modes - no break, no CR to NL, no NL to CR, no carriage return,
	// no strip char, no start/stop output control, no parity check.
	raw.wrap.c_iflag &^= (C.BRKINT | C.IGNBRK | C.ICRNL | C.INLCR | C.IGNCR |
		C.ISTRIP | C.IXON | C.PARMRK)

	// Output modes - disable post processing.
	raw.wrap.c_oflag &^= (C.OPOST)

	// Local modes - echoing off, canonical off, no extended functions,
	// no signal chars (^Z,^C).
	raw.wrap.c_lflag &^= (C.ECHO | C.ECHONL | C.ICANON | C.IEXTEN | C.ISIG)

	// Control modes - set 8 bit chars.
	raw.wrap.c_cflag &^= (C.CSIZE | C.PARENB)
	raw.wrap.c_cflag |= (C.CS8)

	// Control chars - set return condition: min number of bytes and timer.
	// We want read to return every single byte, without timeout.
	raw.wrap.c_cc[C.VMIN] = 1 // Read returns when one char is available.
	raw.wrap.c_cc[C.VTIME] = 0

	// Put terminal in raw mode after flushing
	if err := tcsetattr(fd, C.TCSAFLUSH, raw); err != nil {
		return err
	}

	return nil
}

// Recovers the original settings for this terminal.
func RestoreTermios() {
	if err := tcsetattr(stdin, C.TCSANOW, origTermios); err != nil {
		panic("restoring the terminal")
	}
}

