package gupdown

import (
	"io"
	"mime/multipart"
	"os"
)

func Upload(file multipart.File, handler *multipart.FileHeader, err error, target string) (int64, error) {
	defer file.Close()

	// fmt.Fprintf(w, "%v", handler.Header)
	newFile, err := os.OpenFile(target+"/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return 0, err
	}
	defer newFile.Close()

	size, err := io.Copy(newFile, file)
	if err != nil {
		return 0, err
	}

	return size, nil
}
