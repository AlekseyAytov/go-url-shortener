package filestorage

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/AlekseyAytov/go-url-shortener/internal/urlobject"
)

type FileStorage struct {
	fileName string
}

func NewFileStorage(fileName string) *FileStorage {
	return &FileStorage{fileName: fileName}
}

func (s *FileStorage) SaveObject(u urlobject.URLObject) error {
	if s.fileName == "" {
		return nil
	}

	file, err := os.OpenFile(s.fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	data, err := json.Marshal(&u)
	if err != nil {
		return err
	}
	data = append(data, '\n')
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *FileStorage) ReadObjects() ([]urlobject.URLObject, error) {
	if s.fileName == "" {
		return []urlobject.URLObject{}, nil
	}

	file, err := os.OpenFile(s.fileName, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := make([]urlobject.URLObject, 10)

	for scanner.Scan() {
		uo := &urlobject.URLObject{}
		err = json.Unmarshal(scanner.Bytes(), uo)
		if err != nil {
			return nil, err
		}
		result = append(result, *uo)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
