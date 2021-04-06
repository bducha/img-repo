package storage

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

// As you as you may have guessed, it manages stores files locally
type LocalStorage struct {
	LocalPath  string
	PublicPath string
}

// Stores the file in a local directory defined in the LocalStorage struct.
// It names the the file by it's uid and determines the extension via the original filename
// (Yes it's not ideal but for now it get the job done, i'll improve it later)
func (fileStorage LocalStorage) Store(file *multipart.FileHeader, uid string) (publicPath string, err error) {

	fileReader, err := file.Open()
	if err != nil {
		return publicPath, err
	}

	defer func() {
		err = fileReader.Close()
	}()

	parts := strings.Split(file.Filename, ".")

	if len(parts) < 2 {
		return publicPath, errors.New("wrong filename")
	}

	fileName := uid + "." + parts[len(parts)-1]
	publicPath = fileStorage.PublicPath + fileName

	f, err := os.OpenFile(fileStorage.LocalPath+fileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return publicPath, err
	}

	_, err = io.Copy(f, fileReader)
	return publicPath, err
}

// Deletes a file by it's uid.
// We don't have the extension provided so we search the file by the uid and add a wildcard for the extension
// (Again, not ideal but it works for now :) )
func (fileStorage LocalStorage) Delete(uid string) error {
	files, err := filepath.Glob(fileStorage.LocalPath + uid + ".*")
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return errors.New("file not found")
	}
	for _, file := range files {
		err = os.Remove(file)
	}
	return err
}
