package storage

import (
	"fmt"

	"githab.com/internal/storage/internal/file"
	"github.com/google/uuid"
)

// чтобы файл стал доступен, нужно создать модуль
// go mod init <название>

type Storage struct {
	files map[uuid.UUID]*file.File
}

// конструктор, возвращаем объект
func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}

}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {

	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	s.files[newFile.ID] = newFile

	return newFile, nil

}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	foundFile, ok := s.files[fileID]
	if !ok {
		return nil, fmt.Errorf("file %v not found", fileID)
	}
	return foundFile, nil
}
