package storage

import (
	"log"
	"mime/multipart"
	"os"
	"sync"
)

// Defines the basic needed methods for storage systems
type Storage interface {
	Store(file *multipart.FileHeader, uid string) (string, error)
	Delete(uid string) error
}

var storage Storage
var once sync.Once

// Return an instance of a storage manager
// Later it will take into account config for different storage methods
func GetInstance() Storage {
	once.Do(func() {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		storage = LocalStorage{
			LocalPath:  wd + "/var/images/",
			PublicPath: "http://localhost:8080/img/",
		}
	})
	return storage
}
