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
	if err == os.ErrNotExist || folderExistence == false{
		// create the folder
		err := os.MkdirAll(filePathToSave, 0755)
		if err != nil{
			fmt.Println("Error in creating the folder in destination of .", filePathToSave)
			return "", errors.New("Error in creating file destination")
		}
	}

	var path string
	var err1 error

	switch command {
	case "transformation":
		path, err1 = functionalities.PerformTransformations(specificity, filepath, filePathToSave, mimeType)
		fmt.Println("For the transformation is: ",path, err1)
		if err1 != nil{
			return "", err
		}

	case "resize":
		path, err1 = functionalities.PerformResize(specificity, filepath, filePathToSave, mimeType)
		if err != nil{
			return "", err
		}

	case "filters":
		path, err1 = functionalities.PerformFilters(specificity, filepath, filePathToSave, mimeType)
		if err != nil{
			return "", err
		}
	}

	return path, err1
}