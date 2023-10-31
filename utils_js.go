//go:build js
// +build js

package readline

import (
	"io"
	"sync"
	"syscall"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// SuspendMe use to send suspend signal to myself, when we in the raw mode.
// For OSX it need to send to parent's pid
// For Linux it need to send to myself
func SuspendMe() {

}

// get width of the terminal
func getWidth(stdoutFd int) int {
	cols, _, err := GetSize(stdoutFd)
	if err != nil {
		return -1
	}
	return cols
}

func GetScreenWidth() int {
	w := getWidth(syscall.Stdout)
	if w < 0 {
		w = getWidth(syscall.Stderr)
	}
	return w
}

// ClearScreen clears the console screen
func ClearScreen(w io.Writer) (int, error) {
	return w.Write([]byte("\033[H"))
}

func DefaultIsTerminal() bool {
	return IsTerminal(syscall.Stdin) && (IsTerminal(syscall.Stdout) || IsTerminal(syscall.Stderr))
}

func GetStdin() int {
	return syscall.Stdin
}

// -----------------------------------------------------------------------------

var (
	widthChange         sync.Once
	widthChangeCallback func()
)

func DefaultOnWidthChanged(f func()) {

}
