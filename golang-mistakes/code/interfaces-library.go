package lib

import "io"

// instead of
func Do(b []byte) {
	// process data
}

// use interfaces
func Do(r io.Reader) {
	// process data
}
