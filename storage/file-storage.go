package storage

import (
	"encoding/json"
	"errors"
	"os"
)

type Entry struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Storage interface {
	Save([]Entry) error
	Load() ([]Entry, error)
}

type FileStorage struct {
	Path string
}

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{Path: path}
}

// Save password
func (fs *FileStorage) Save(data []Entry) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.Path, bytes, 0644)
}

// Read password from file
func (fs *FileStorage) Load() ([]Entry, error) {
	var entries []Entry
	bytes, err := os.ReadFile(fs.Path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Entry{}, nil
		}
		return nil, err
	}
	err = json.Unmarshal(bytes, &entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}
