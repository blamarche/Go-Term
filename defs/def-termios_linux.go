// goheader -w -p term -s linux /usr/include/asm-generic/termios.h
// MACHINE GENERATED; DO NOT EDIT


package term

//!!! #ifndef _ASM_GENERIC_TERMIOS_H
//!!! #define _ASM_GENERIC_TERMIOS_H
// Most architectures have straight copies of the x86 code, with
// varying levels of bug fixes on top. Usually it's a good idea
// to use this generic version instead, but be careful to avoid
// ABI changes.
// New architectures should not provide their own version.

//!!! #include <asm/termbits.h>
//!!! #include <asm/ioctls.h>

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

const NCC = 8

type Termio struct {
	Iflag uint16     // input mode flags
	Oflag uint16     // output mode flags
	Cflag uint16     // control mode flags
	Lflag uint16     // local mode flags
	Line  uint8      // line discipline
	Cc    [NCC]uint8 // control characters
}

// modem lines
const (
	TIOCM_LE   = 0x001
	TIOCM_DTR  = 0x002
	TIOCM_RTS  = 0x004
	TIOCM_ST   = 0x008
	TIOCM_SR   = 0x010
	TIOCM_CTS  = 0x020
	TIOCM_CAR  = 0x040
	TIOCM_RNG  = 0x080
	TIOCM_DSR  = 0x100
	TIOCM_CD   = TIOCM_CAR
	TIOCM_RI   = TIOCM_RNG
	TIOCM_OUT1 = 0x2000
	TIOCM_OUT2 = 0x4000
	TIOCM_LOOP = 0x8000
)

// ioctl (fd, TIOCSERGETLSR, &result) where result may be as below


//!!! #endif // _ASM_GENERIC_TERMIOS_H
