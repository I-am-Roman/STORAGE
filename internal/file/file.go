package file

import "github.com/google/uuid"

type File struct {
	Name string
	Data []byte
	ID   uuid.UUID
}

// конструктор файла
func NewFile(filename string, blob []byte) (*File, error) {

	fileID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &File{
		Name: filename,
		Data: blob,
		ID:   fileID,
	}, nil
}
