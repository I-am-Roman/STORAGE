package main

import (
	"fmt"
	// чтобы получить доступ к файлу storage.go <путь пакета><путь в нашей директории>
	"githab.com/internal/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("Hello, everyone", st)
}
