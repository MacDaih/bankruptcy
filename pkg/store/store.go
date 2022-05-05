package store

import "os"

func createStore(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		return err
	}
	return nil
}

func GetFile(filePath string, file string) (*os.File, error) {
	if err := createStore(filePath); err != nil {
		return nil, err
	}
	return os.OpenFile(filePath+file, os.O_RDWR|os.O_CREATE, 0755)
}
