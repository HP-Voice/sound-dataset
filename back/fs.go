package main

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/pgtype"
	"io"
	"os"
)

func initFs() error {
	return os.MkdirAll(config.Fs.Path, os.ModePerm)
}

func filenameOf(sampleId pgtype.UUID) string {
	return config.Fs.Path + fmt.Sprintf("%x", sampleId.Bytes) + config.Fs.Extension
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
