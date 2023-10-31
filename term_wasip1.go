// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1
// +build wasip1

package readline

type State struct {
	mode uint32
}

// IsTerminal returns true if the given file descriptor is a terminal.
func IsTerminal(fd int) bool {
	return false
}

// GetSize returns the dimensions of the given terminal.
func GetSize(fd int) (int, int, error) {
	return 0, 0, nil
}

// Restore restores the terminal connected to the given file descriptor to a
// previous state.
func restoreTerm(fd int, state *State) error {
	return nil
}

// MakeRaw put the terminal connected to the given file descriptor into raw
// mode and returns the previous state of the terminal so that it can be
// restored.
func MakeRaw(fd int) (*State, error) {
	return nil, nil
}
