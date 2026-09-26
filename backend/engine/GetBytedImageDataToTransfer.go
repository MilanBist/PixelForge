package engine

import (
	"errors"
	"fmt"
	"os"
)


type SingleImage struct{
	FileTypeRequired string
}

func (s *SingleImage) GetBytedImage(storageKey string) ([]byte, error){
	// read the image based on the provided location
	data, err := os.ReadFile(storageKey)
	if err != nil {
		fmt.Println(err)
		return []byte{}, errors.New("Internal server error.")
	}
	return data, nil
}