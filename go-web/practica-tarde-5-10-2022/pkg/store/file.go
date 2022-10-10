package store


import (
	"encoding/json"
	"os"
)

type Users interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

const (
	FileType  Type = "file"
)

func New(usuario Type, fileName string) Users {
	switch usuario {
	case FileType:
		return &fileUsuario{fileName}
	}

	return nil
}

type fileUsuario struct {
	FilePath string
}

func (fs *fileUsuario) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}

func (fs *fileUsuario) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}