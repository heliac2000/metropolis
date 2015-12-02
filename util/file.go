package util

import "os"

// File exists ?
//
func FileExists(fname string) bool {
	_, err := os.Stat(fname)
	return err == nil
}
