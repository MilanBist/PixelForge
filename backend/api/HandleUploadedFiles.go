package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/image-generator/internal/models"
)

type UploadedHistory interface{
	GetAllUploadedFiles(userId int) ([]models.BaseFileData, error)
}

type UploadedFiles struct{
	Files 	UploadedHistory
}

func(u *UploadedFiles) HandleUploadedFiles(w http.ResponseWriter, r *http.Request){
	value := r.Context().Value("metaData")
	metaData := value.(models.ContextMetaData)
	fmt.Println(metaData)
	uploadedFiiles, err := u.Files.GetAllUploadedFiles(metaData.UserId)

	if err != nil{
		response := models.Response{
			Success: false,
			Message: "Can't retrieve data.",
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(&response)
		return
	}
	// return the data of the uploaded files now
	response := models.Response{
		Success: true,
		Message: "Success in retrieving the data.",
		Data: uploadedFiiles,
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&response)
}