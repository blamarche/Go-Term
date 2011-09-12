// goheader -w -p term -s linux /usr/include/asm-generic/ioctls.h
// MACHINE GENERATED; DO NOT EDIT


package term

//!!! #ifndef __ASM_GENERIC_IOCTLS_H
//!!! #define __ASM_GENERIC_IOCTLS_H

//!!! #include <linux/ioctl.h>

// These are the most common definitions for tty ioctl numbers.
// Most of them do not use the recommended _IOC(), but there is
// probably some source code out there hardcoding the number,
// so we might as well use them for all new platforms.
//
// The architectures that use different values here typically
// try to be compatible with some Unix variants for the same
// architecture.

// 0x54 is just a magic number to make these relatively unique ('T')

const (
	TCGETS       = 0x5401
	TCSETS       = 0x5402
	TCSETSW      = 0x5403
	TCSETSF      = 0x5404
	TCGETA       = 0x5405
	TCSETA       = 0x5406
	TCSETAW      = 0x5407
	TCSETAF      = 0x5408
	TCSBRK       = 0x5409
	TCXONC       = 0x540A
	TCFLSH       = 0x540B
	TIOCEXCL     = 0x540C
	TIOCNXCL     = 0x540D
	TIOCSCTTY    = 0x540E
	TIOCGPGRP    = 0x540F
	TIOCSPGRP    = 0x5410
	TIOCOUTQ     = 0x5411
	TIOCSTI      = 0x5412
	TIOCGWINSZ   = 0x5413
	TIOCSWINSZ   = 0x5414
	TIOCMGET     = 0x5415
	TIOCMBIS     = 0x5416
	TIOCMBIC     = 0x5417
	TIOCMSET     = 0x5418
	TIOCGSOFTCAR = 0x5419
	TIOCSSOFTCAR = 0x541A
	FIONREAD     = 0x541B
	TIOCINQ      = FIONREAD
	TIOCLINUX    = 0x541C
	TIOCCONS     = 0x541D
	TIOCGSERIAL  = 0x541E
	TIOCSSERIAL  = 0x541F
	TIOCPKT      = 0x5420
	FIONBIO      = 0x5421
	TIOCNOTTY    = 0x5422
	TIOCSETD     = 0x5423
	TIOCGETD     = 0x5424
	TCSBRKP      = 0x5425 // Needed for POSIX tcsendbreak()
	TIOCSBRK     = 0x5427 // BSD compatibility
	TIOCCBRK     = 0x5428 // BSD compatibility
	TIOCGSID     = 0x5429 // Return the session ID of FD
	//!!! TCGETS2 = _IOR('T', 0x2A, struct termios2)
	//!!! TCSETS2 = _IOW('T', 0x2B, struct termios2)
	//!!! TCSETSW2 = _IOW('T', 0x2C, struct termios2)
	//!!! TCSETSF2 = _IOW('T', 0x2D, struct termios2)
	TIOCGRS485 = 0x542E
	TIOCSRS485 = 0x542F
	//!!! TIOCGPTN = _IOR('T', 0x30, unsigned int) // Get Pty Number (of pty-mux device)
	//!!! TIOCSPTLCK = _IOW('T', 0x31, int)  // Lock/unlock Pty
	TCGETX  = 0x5432 // SYS5 TCGETX compatibility
	TCSETX  = 0x5433
	TCSETXF = 0x5434
	TCSETXW = 0x5435
)

const (
	FIONCLEX        = 0x5450
	FIOCLEX         = 0x5451
	FIOASYNC        = 0x5452
	TIOCSERCONFIG   = 0x5453
	TIOCSERGWILD    = 0x5454
	TIOCSERSWILD    = 0x5455
	TIOCGLCKTRMIOS  = 0x5456
	TIOCSLCKTRMIOS  = 0x5457
	TIOCSERGSTRUCT  = 0x5458 // For debugging only
	TIOCSERGETLSR   = 0x5459 // Get line status register
	TIOCSERGETMULTI = 0x545A // Get multiport config
	TIOCSERSETMULTI = 0x545B // Set multiport config
)

const (
	TIOCMIWAIT  = 0x545C // wait for a change on serial input line(s)
	TIOCGICOUNT = 0x545D // read serial port __inline__ interrupt counts
)

// some architectures define FIOQSIZE as 0x545E, which is used for
// TIOCGHAYESESP on others
//!!! #ifndef FIOQSIZE
const (
	TIOCGHAYESESP = 0x545E // Get Hayes ESP configuration
	TIOCSHAYESESP = 0x545F // Set Hayes ESP configuration
	FIOQSIZE      = 0x5460
	//!!! #endif
)

// Used for packet mode
const (
	TIOCPKT_DATA       = 0
	TIOCPKT_FLUSHREAD  = 1
	TIOCPKT_FLUSHWRITE = 2
	TIOCPKT_STOP       = 4
	TIOCPKT_START      = 8
	TIOCPKT_NOSTOP     = 16
	TIOCPKT_DOSTOP     = 32
)

const TIOCSER_TEMT = 0x01 // Transmitter physically empty

//!!! #endif // __ASM_GENERIC_IOCTLS_H
