package storage

import (
	"log"
	"mime/multipart"
	"os"
	"sync"
)

type Storage interface {
	Store(file *multipart.FileHeader, uid string) (string, error)
	Delete(uid string)
}

var storage Storage
var once sync.Once

func GetInstance() Storage {
	once.Do(func() {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		storage = LocalStorage{
			LocalPath:  wd + "/var/images/",
			PublicPath: "http://localhost:8080/images/",
		}
	})
	return storage
}
