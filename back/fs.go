package main

import (
	"errors"
	"io"
	"os"
)

func initFs() error {
	return os.MkdirAll(config.Fs.Path, os.ModePerm)
}

func filenameOf(sampleId UUID) string {
	return config.Fs.Path + sampleId.String() + config.Fs.Extension
}

func writeFile(filename string, data io.Reader) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, config.Fs.BlockSize)
	more := true
	for more {
		read, err := data.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				more = false
			} else {
				return err
			}
		}
		if read == 0 {
			break
		}
		_, err = file.Write(buffer[:read])
		if err != nil {
			return err
		}
	}

	return nil
}
