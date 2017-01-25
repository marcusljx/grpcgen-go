package functions

import (
	"os"
	"path/filepath"
)

func QualifyFromGopathSrc(path string) string {
	return filepath.Join(os.Getenv("GOPATH"), "src", path)
}
