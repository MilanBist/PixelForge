package api

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"github.com/image-generator/internal/models"
)

type TransformImage interface{
	AddRawFileToDestination(file multipart.File, filename, userId string) (string,int, error)
	GenerateImageTranformation(main, subMain string) error
	GetDimension(location string) (int, int, int64, error)
}

type UploadImages interface{
	AddUploadedImageFiles(uploadedMetaData models.UploadedFilesMetaData)(int, error)
	UploadGeneratedImage() error
}

type ImageTransformation struct{
	UploadGenerate		TransformImage
	Store 				UploadImages		
}


type ImageToBePerformedTasks struct{
	MainTask 		[]string
	SubmainTask  	[]string
	File 			multipart.File
}

func(i *ImageTransformation) HandleSingleImageTransformation(w http.ResponseWriter, r *http.Request){
	// get the image Id just
	file,header,err := r.FormFile("file")
	if err != nil{
        http.Error(w, "failed to get file", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Input file missing.",
		})
        return
	}

	var imgOperations ImageToBePerformedTasks
	mainTask := r.FormValue("mainTask")
	err = json.Unmarshal([]byte(mainTask), &imgOperations.MainTask)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Can't decode the main and submainTask.",
		})
        return
	}
	subMainTask := r.FormValue("subMainTask")
	err = json.Unmarshal([]byte(mainTask), &imgOperations.SubmainTask)
	err = json.Unmarshal([]byte(mainTask), &imgOperations.MainTask)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Can't decode the main and submainTask.",
		})
        return
	}
	imgOperations.File = file

	filename := header.Filename
	extensionName := filepath.Ext(filename)
	// check for the file extension
	if filepath.Ext(filename) != ".png" || filepath.Ext(filename) != ".jpg"{
		fmt.Println("Wrong file name. Should be .png or .jpg file.")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Input .png or .jpg file.",
		})
		return
	}

	// if any of then is unavailable just return false
	if mainTask == "" || subMainTask == ""{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Input both of the main and submain tasks.",
		})
		return
	}

	data := r.Context().Value("metaData")
	metaData := data.(models.ContextMetaData)

	// get the userId
	userId := metaData.UserId


	// based on the userId, imageId first insert them to uploaded files
	// also add its meta data to the database using the interface patterns


	// add the uploaded files
	exactFilePath, statusCode, err := i.UploadGenerate.AddRawFileToDestination(file, filename, strconv.Itoa(userId))

	if err != nil{
		response := models.Response{
			Success: false,
			Message: err.Error(),
		}

		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(response)
		return
	}
	// loop through each of the mian and submain 
	uploadingFile := models.UploadedFilesMetaData{
		UserId: int64(userId),
		Filename: header.Filename,
	}

	
	height, width, size, err := i.UploadGenerate.GetDimension(exactFilePath)
	if err != nil{
		response := models.Response{
			Success: false,
			Message: err.Error(),
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return
	}
	// for the uploadedImages
	var mimetype string
	if extensionName == ".png"{
		mimetype = "image/png"
	}else{
		mimetype = "image/jpg"
	}
	var uploadingMetaData models.UploadedFilesMetaData = models.UploadedFilesMetaData{
		UserId: int64(metaData.UserId),
		Filename: filename,
		StorageKey: exactFilePath,
		FileType: extensionName,
		Mimetype: mimetype,
		FileSize: size,
		Height: height,
		Width: width,
	}

	uploadedId, err := i.Store.AddUploadedImageFiles(uploadingMetaData)
	if err != nil{
		response := models.Response{
			Success: false,
			Message: err.Error(),
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	
	// for i := range imgOperations.MainTask{
	// 	// now get the images data

	// }


	
}
