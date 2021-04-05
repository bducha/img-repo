package storage

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

type LocalStorage struct {
	LocalPath  string
	PublicPath string
}

func (fileStorage LocalStorage) Store(file *multipart.FileHeader, uid string) (path string, err error) {

	fileReader, err := file.Open()
	if err != nil {
		return path, err
	}

	defer func() {
		err = fileReader.Close()
	}()

	parts := strings.Split(file.Filename, ".")

	if len(parts) < 2 {
		return path, errors.New("wrong filename")
	}

	fileName := uid + "." + parts[len(parts)-1]
	path = fileStorage.PublicPath + fileName

	f, err := os.OpenFile(fileStorage.LocalPath+fileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return path, err
	}

	_, err = io.Copy(f, fileReader)
	return path, err
}

func (fileStorage LocalStorage) Delete(uid string) {

}
