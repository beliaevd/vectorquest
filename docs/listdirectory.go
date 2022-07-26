package docs

import (
	"os"
	"path/filepath"
)

func ListDirByWalk(path string) {
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {

		if wPath == path {
			return nil
		}
		if info.IsDir() {

		}
		return nil
	})
}
