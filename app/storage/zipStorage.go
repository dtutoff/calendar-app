package storage

import (
	"archive/zip"
	"io"
	"os"

	"github.com/SamiRemi/project/app/logger"
	"github.com/SamiRemi/project/app/validation"
)

type ZipStorage struct {
	*Storage
}

func NewZipStorage(filename string) *ZipStorage {
	logger.Info("Запуск фукции NewZipStorage")
	return &ZipStorage{
		&Storage{
			filename: filename,
		},
	}
}
func (z *ZipStorage) Save(data []byte) error {
	logger.Info("Запуск фукции Save")
	f, err := os.Create(z.GetFilename())
	if err != nil {
		return err
	}
	defer f.Close()
	zw := zip.NewWriter(f)
	defer zw.Close()
	w, err := zw.Create("data")
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (z *ZipStorage) Load() ([]byte, error) {
	logger.Info("Запуск функции Load")
	r, err := zip.OpenReader(z.GetFilename())
	if err != nil {
		return nil, err
	}
	defer r.Close()

	if len(r.File) == 0 {
		return nil, validation.ArchiveEmptyError
	}
	file := r.File[0]
	rc, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return io.ReadAll(rc)
}
