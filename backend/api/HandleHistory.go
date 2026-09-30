package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/image-generator/internal/models"
)

type HistoryData interface{
	GetHistoryDataOfUser(userId int) ([]models.HistoricalData, error)
}

type History struct{
	Data 	HistoryData
}

func(h *History) HandleHistory(w http.ResponseWriter, r *http.Request){
	// get the userId from the context
	value := r.Context().Value("metaData").(models.ContextMetaData)
	userId := value.UserId

	fmt.Println(userId)


	// based on the userId get the history data from the backend server
	historyData, err := h.Data.GetHistoryDataOfUser(userId)
	if err != nil{
		response := models.Response{
			Success: false,
			Message: "Can't retrive the historical data.",
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response)
		return
	}



	var sendingData []models.FileBasedImageGenerationReturn
	var data1 models.FileBasedImageGenerationReturn
	var upload models.BaseFileData
	var image []models.BaseImageMetaData

	for i, value := range historyData{
		// for initial data set the actualFile and data1
		if i == 0{
			upload.Id = int64(value.UploadedFileId)
			upload.Filename = value.UploadedFileName
			upload.CreatedAt = value.CreatedAt
			data1.ActualFile = upload


			var image1 models.BaseImageMetaData
			image1.Id = value.ImageId
			image1.ImageName = value.ImageName
			image1.Mimetype = value.MimeType
			image1.Width = value.Width
			image1.Height = value.Height
			image1.FileSize = value.FileSize
			image = append(image, image1)
		} else{
			if value.UploadedFileId == int(data1.ActualFile.Id){
				// add to the generated image section
				var image1 models.BaseImageMetaData
				image1.Id = value.ImageId
				image1.ImageName = value.ImageName
				image1.Mimetype = value.MimeType
				image1.Width = value.Width
				image1.Height = value.Height
				image1.FileSize = value.FileSize
				image = append(image, image1)
			} else{
				//if not equal then append the whole of the data to sendind data and set it to null
				data1.ActualFile = upload
				data1.GeneratedImage = image
				sendingData = append(sendingData, data1)

				// set all to be null and initiate
				data1 = models.FileBasedImageGenerationReturn{}
				upload = models.BaseFileData{}
				image = []models.BaseImageMetaData{}

				// set all of the upload again
				upload.Id = int64(value.UploadedFileId)
				upload.Filename = value.UploadedFileName
				upload.CreatedAt = value.CreatedAt
				data1.ActualFile = upload

				//set all of the images inside to be null
				var image1 models.BaseImageMetaData
				image1.Id = value.ImageId
				image1.ImageName = value.ImageName
				image1.Mimetype = value.MimeType
				image1.Width = value.Width
				image1.Height = value.Height
				image1.FileSize = value.FileSize
				image = append(image, image1)
			}
		}

	}



	response := models.Response{
		Success: true,
		Message: "Success in retrieving the historical data.",
		Data: sendingData,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}