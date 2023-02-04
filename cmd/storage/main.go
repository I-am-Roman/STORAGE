package main

import (
	"fmt"
	"log"

	// чтобы получить доступ к файлу storage.go <путь пакета><путь в нашей директории>
	"githab.com/internal/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("text.txt", []byte("hello"))

	if err != nil {
		log.Fatal(err)
	}

	restore, err := st.GetByID(file.ID)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(st.files)

	fmt.Println("Hello, everyone", restore)
}
