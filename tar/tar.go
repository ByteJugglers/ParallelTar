package tar

import (
	"fmt"
	"path/filepath"
	"strings"
)

func Tar(target, packType string) error {
	switch packType {
	case "gzip":
		fileName := target + ".tar.gz"
		return Gzip(fileName, target)
	case "bzip2":
		fileName := target + ".tar.bz2"
		return Bzip2(fileName, target)
	case "xz":
		fileName := target + ".tar.xz"
		return Xz(fileName, target)
	default:
		return fmt.Errorf("unknown compress type: %s", packType)
	}
}

func UnTar(source string) error {
	target := filepath.Dir(source)

	exts := strings.Split(source, ".")
	ext := strings.Join(exts[len(exts)-2:], ".")

	switch ext {
	case "tar.gz":
		return UnGzip(source, target)
	case "tar.bz2":
		return UnBzip2(source, target)
	case "tar.xz":
		return UnXz(source, target)
	default:
		return fmt.Errorf("unsupport compress type: %s", ext)
	}
}
