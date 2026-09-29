package storage

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"github.com/image-generator/engine"
)


type StoreFile struct{
	BasePath string
}

// check for path existence
func pathExistence(filePath string) (bool, error){
	fmt.Println("Filepath: ", filePath)
	_, err := os.Stat(filePath)

	if err == nil{
		return true, nil
	} else if errors.Is(err, os.ErrNotExist){
		return false, err
	}

	return false, err
}

// add file to the given location
func addToLocation(file multipart.File, filepath string) error{
	fmt.Println(filepath)
	inputFile, err := os.Create(filepath)
	if err != nil {
		// Error in creating the file
		fmt.Println("Error in creating the file. Actual error is: ", err)
		return errors.New("Error in creating the file.")
	}

	defer inputFile.Close()

	//copy each item to the given inputFile
	_, err =io.Copy(inputFile, file)
	if err != nil{
		fmt.Println("Error in copying the file item. Actual error: ", err)
		return errors.New("Error in copying the file item.")
	}
	fmt.Println("Successfully added the file to the location.")
	return nil
}



// add the file to the certain location
func (f *StoreFile) AddRawFileToDestination(file multipart.File, fileName, userId string) (string, int, error){
	newBasePath := filepath.Join(f.BasePath,userId,"uploaded")
	folderExistence, err := pathExistence(newBasePath)
	if err == os.ErrNotExist || folderExistence == false{
		// create the folder
		err := os.MkdirAll(newBasePath, 0755)
		if err != nil{
			fmt.Println("Error in creating the folder in destination of .", f.BasePath + fileName)
			return "", http.StatusInternalServerError, errors.New("Error in creating file destination")
		}
		fmt.Println("Created Successfully.")
	}


	fmt.Println("File name is: ", fileName)
	err = addToLocation(file, filepath.Join(newBasePath, fileName))
	if err != nil {
		return "", http.StatusInternalServerError, errors.New("Error in creating the file.")
	}
	return filepath.Join(newBasePath, fileName), http.StatusAccepted, nil
}

// Generate image from the raw file
func (f *StoreFile) GenerateImage(exactFilePath string, userId string) (engine.AllFiles,int, error){
	//generate image from the raw file
	filePathToSave := filepath.Join(f.BasePath, userId, "generated")
	allImageFiles, _, err := engine.GenerateImageFromRaw(exactFilePath, filePathToSave)
	
	if err != nil{
		fmt.Println("Internal Server error: ",err)
		return  engine.AllFiles{}, -1, err
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("All images files are: ", allImageFiles)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	return allImageFiles, http.StatusAccepted, nil
}

// Generate image transformation based on the main and submain task and filepath
func(f *StoreFile) GenerateImageTransformations(mainTask, subMainTask, filePath, userId, mimetype string) (string, error){
	filePathToSave := filepath.Join(f.BasePath, userId, "generated")
	path, err := engine.PerformAction(mainTask, subMainTask, filePath, filePathToSave,mimetype)
	if err != nil{
		fmt.Println(err)
		return "", err
	}
	return path, err
}