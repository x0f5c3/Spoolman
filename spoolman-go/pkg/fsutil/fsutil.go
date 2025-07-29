package fsutil

import "path/filepath"

func Abs(p string) string {
	np, err := filepath.Abs(p)
	if err != nil {
		return p
	} else {
		return np
	}
}
