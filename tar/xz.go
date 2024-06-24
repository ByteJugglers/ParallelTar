package tar

import (
	"archive/tar"
	"github.com/ulikunitz/xz"
	"io"
	"os"
	"path/filepath"
)

func Xz(source, target string) error {
	// Create the target file
	targetFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	// Create an xz writer
	xw, err := xz.NewWriter(targetFile)
	if err != nil {
		return err
	}
	defer xw.Close()

	// Create a tar writer on top of xz writer
	tw := tar.NewWriter(xw)
	defer tw.Close()

	// Walk through the source directory and add files to the tar writer
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Construct the relative file path
		relPath, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}

		// Write directory entries with the parent directories as well
		name := filepath.ToSlash(relPath)
		if info.IsDir() {
			name += "/" // Add separator if it's a directory
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

			_, err = io.Copy(tw, file)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func UnXz(source, target string) error {
	// Open the source xz file
	xzFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer xzFile.Close()

	// Create an xz reader
	xr, err := xz.NewReader(xzFile)
	if err != nil {
		return err
	}

	// Create a tar reader on top of xz reader
	tr := tar.NewReader(xr)

	// Walk through each tar entry and extract files
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Construct the file path for the extracted file
		filePath := filepath.Join(target, header.Name)

		// Ensure the parent directories exist
		if header.FileInfo().IsDir() {
			err := os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			// Create the file
			file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, header.FileInfo().Mode())
			if err != nil {
				return err
			}
			defer file.Close()

			// Write the file content
			_, err = io.Copy(file, tr)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
