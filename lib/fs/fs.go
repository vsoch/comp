package fs

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	fpath "path"
	"path/filepath"
	"strings"
)

// Wrapper to get the PWD
func GetPwd() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot derive the present working directory.")
	}
	return path
}

// Mkdirp creates a nested directory if it doesn't exist
func Mkdirp(path string) bool {

	// If it doesn't exist, create it (and return true)
	if !Exists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatal("Cannot derive the present working directory.")
		}
		return true
	}
	// Not created
	return false
}

// Exists determines if a path exists
func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// Get the realtive path of a file, relative to some root
func RelativePath(root string, path string) string {

	relpath, err := filepath.Rel(root, path)
	if err != nil {
		log.Fatal(err)
	}
	return relpath
}

// Copy a file from a source to a dest (intended for small files_
func CopyFile(source string, dest string) {

	// First open the source file
	src, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	// Now open the destination file
	out, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Do the copy!
	_, err = io.Copy(out, src)
	if err != nil {
		log.Fatal(err)
	}
}

// Which finds the full path to an executable, or returns empty string
func Which(name string) string {
	path, err := exec.LookPath(name)
	if err != nil {
		fmt.Println("Could not find ", name)
	}
	return path
}

// Function to list directory, and either return basesnames of files, folders, or both
func ListDir(path string, includeDir bool, includeFile bool) []string {

	// We will return a list of string results
	results := []string{}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fullpath := fpath.Join(path, file.Name())
		fileInfo, err := os.Stat(fullpath)
		if err != nil {
			log.Fatal(err)
		}

		// If it's a directory and we want directories, add it
		if fileInfo.IsDir() && includeDir {
			results = append(results, file.Name())
		} else if includeFile {
			results = append(results, file.Name())
		}
	}
	return results
}

// ReadFile reads a file from the system
func ReadFile(path string) string {
	filey, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = filey.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	bytes, err := ioutil.ReadAll(filey)
	return string(bytes)
}

// WriteFile and some content to the filesystem
func WriteFile(path string, content string) error {
	filey, err := os.Create(path)
	if err != nil {
		return err
	}
	defer filey.Close()

	_, err = filey.WriteString(content)
	if err != nil {
		return err
	}
	err = filey.Sync()
	return err
}

// Find files in a directory based on a pattern
func RecursiveFind(root string, pattern string, allowPrefix bool) ([]string, error) {

	// Create a list of results to return
	results := []string{}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		match, _ := filepath.Match(pattern, filepath.Base(path))

		// If we don't have a match and the parser allows a prefix
		if !match && allowPrefix {
			fileBasename := filepath.Base(path)
			match = strings.HasPrefix(fileBasename, pattern)
		}

		if match {
			results = append(results, path)
		}
		return nil

	})

	if err != nil {
		log.Fatal("Error running RecursiveFind to find files %s", err)
	}
	return results, err
}
