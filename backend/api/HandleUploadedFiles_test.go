package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/image-generator/internal/models"
)


type FakeUploadFiles struct{
	FileData 	[]models.BaseFileData
	Error 		error
}


func (f *FakeUploadFiles) GetAllUploadedFiles(userId int)([]models.BaseFileData, error){
	return f.FileData, f.Error
}

type ResponseUpload struct{
    Success bool   `json:"success"`
    Message string `json:"message"`
    Data    []models.BaseFileData    `json:"data,omitempty"`
}


func TestHandleUploadFiles(t *testing.T){
	var data models.BaseFileData
	var fileData []models.BaseFileData

	data.Id = 2
	data.FileType = ".raw"

	fileData = append(fileData, data)
	fuf := &FakeUploadFiles{
		FileData: fileData,
		Error: nil,
	}

	usrv := UploadedFiles{
		Files: fuf,
	}

	metaData := models.ContextMetaData{
		Email: "bistmilan46@gmail.com",
		UserId: 1,
	}

	req := httptest.NewRequest(http.MethodGet, "/uploaded", nil)
	ctx := context.WithValue(req.Context(), "metaData", metaData)
	req = req.WithContext(ctx)

	res := httptest.NewRecorder()
	usrv.HandleUploadedFiles(res, req)


	var response ResponseUpload

	json.NewDecoder(res.Body).Decode(&response)
	

	if res.Code !=  200{
		t.Fatalf("Response should be %d but its %d", 200, res.Code)
	}

	if response.Data[0].Id != 2{
		t.Fatalf("Required %d got %d", 2, response.Data[0].Id)
	}


}