package storage

import "githab.com/internal/storage/internal/file"

// чтобы файл стал доступен, нужно создать модуль
// go mod init <название>

type Storage struct{}

// конструктор, возвращаем объект
func NewStorage() *Storage {
	return &Storage{}

}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	return file.NewFile(filename, blob)
	// newFile,err := file.NewFile(filename,blob)
	// if err != nil{
	// 	return nil,err
	// }

	// return newFile

}
