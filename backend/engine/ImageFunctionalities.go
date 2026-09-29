package engine

import (
	"errors"
	"fmt"
	_ "image/jpeg" // like the init in python  runs its header when decode it called to detect the jpeg files itself
	_ "image/png"
	"os"

	functionalities "github.com/image-generator/engine/functionalites"
)

func PerformAction(command, specificity, filepath, filePathToSave, mimeType string) (string, error){
	// if command is transformations
	folderExistence, err := pathExistence(filePathToSave)
	fmt.Println("Base output path is: ", filePathToSave)
	fmt.Println(folderExistence, err)
	if err == os.ErrNotExist || folderExistence == false{
		// create the folder
		err := os.MkdirAll(filePathToSave, 0755)
		if err != nil{
			fmt.Println("Error in creating the folder in destination of .", filePathToSave)
			return "", errors.New("Error in creating file destination")
		}
	}
	switch command {
	case "transformation":
		transformedPath, err := functionalities.PerformTransformations(specificity, filepath, filePathToSave, mimeType)
		if err != nil{
			return "", err
		}

		return transformedPath, nil

	case "resize":
		functionalities.PerformResize(specificity, filepath)

	case "filters":
		functionalities.PerformFilters(specificity, filepath)
	}

	return "", errors.New("Error in opening the file.")
}