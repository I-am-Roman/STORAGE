package storage

// чтобы файл стал доступен, нужно создать модуль
// go mod init <название>

type Storage struct{}

// конструктор, возвращаем объект
func NewStorage() *Storage {
	return &Storage{}

}
