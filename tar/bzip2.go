package tar

import (
	"archive/tar"
	"github.com/dsnet/compress/bzip2"
	"io"
	"os"
	"path/filepath"
)

func Bzip2(source, target string) error {
	// Create the target file
	targetFile, err := os.Create(source)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	// Create a bzip2 writer
	bw, err := bzip2.NewWriter(targetFile, &bzip2.WriterConfig{Level: bzip2.DefaultCompression})
	if err != nil {
		return err
	}
	defer bw.Close()

	// Create a tar writer
	tw := tar.NewWriter(bw)
	defer tw.Close()

	// Walk through the target directory and add files to the tar writer while keeping parent directories
	return filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Construct the relative file path
		relPath, err := filepath.Rel(target, path)
		if err != nil {
			return err
		}

		// Write directory entries with the parent directories as well
		name := filepath.Join(filepath.Base(target), relPath)
		if info.IsDir() {
			name += string(os.PathSeparator) // Add separator if it's a directory
		}

		// Create a tar header based on the file info
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		// Update the header name to include the parent directories
		header.Name = name

		// Write the header
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}

		// If the file is not a directory, copy its content to the tar writer
		if !info.IsDir() && info.Mode().IsRegular() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			buf := make([]byte, 2048)
			_, err = io.CopyBuffer(tw, file, buf)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func UnBzip2(source, target string) error {
	f, err := os.Open(source)
	if err != nil {
		return err
	}
	defer f.Close()

	bz, err := bzip2.NewReader(f, nil)
	if err != nil {
		return err
	}
	defer bz.Close()

	tr := tar.NewReader(bz)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		filePath := filepath.Join(target, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			err = os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				return err
			}
		case tar.TypeReg:
			f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			_, err = io.Copy(f, tr)

			f.Close()

			if err != nil {
				return err
			}
		}
	}

	return nil
}
