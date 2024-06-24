package main

import (
	"os"
	"path/filepath"
	"regexp"
)

func listSubDirectories(path string) ([]string, error) {
	// Get the list of files and subdirectories in the provided path
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var subDirectories []string

	// Iterate through the entries to find subdirectories
	for _, entry := range entries {
		if entry.IsDir() && entry.Name()[0] != '.' {
			subDirectories = append(subDirectories, filepath.Join(path, entry.Name()))
		}
	}

	return subDirectories, nil
}

func listSubPackedFiles(path string) ([]string, error) {
	// Get the list of files and subdirectories in the provided path
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var subDirectories []string

	reg := regexp.MustCompile(`.tar.gz$|.tar.bz2$`)
	// Iterate through the entries to find subdirectories
	for _, entry := range entries {
		if !entry.IsDir() && reg.MatchString(entry.Name()) {
			subDirectories = append(subDirectories, filepath.Join(path, entry.Name()))
		}
	}

	return subDirectories, nil
}
