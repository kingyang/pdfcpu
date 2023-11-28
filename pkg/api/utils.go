package api

import (
	"io"
	"os"
)

func Rename(src, dst string) error {
	// Open the source file for reading
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file for writing
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// Close the destination file
	err = dstFile.Close()
	if err != nil {
		return err
	}

	// Remove the source file
	err = os.Remove(src)
	if err != nil {
		return err
	}

	// Return nil if everything succeeds
	return nil
}
