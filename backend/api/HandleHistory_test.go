package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"github.com/image-generator/internal/models"
)


type HistoryResponse struct {
    Success bool                                     `json:"success"`
    Message string                                   `json:"message"`
    Data    []models.FileBasedImageGenerationReturn `json:"data"`
}

type FakeHistoryData struct{
	HistoryModel 	[]models.HistoricalData
	Error			error
}

func(h *FakeHistoryData) GetHistoryDataOfUser(userId int)([]models.HistoricalData, error){
	return h.HistoryModel, h.Error
}

func TestHandleHistoryRendering(t *testing.T){
	var modelData []models.HistoricalData

	data := models.HistoricalData{
		UploadedFileId: 1,
		UploadedFileName: "image.raw",
		CreatedAt: time.Now(),
		ImageId: 1,
		ImageName: "image.png",
		MimeType: "image/png",
		Width: 1222,
		Height: 1222,
		FileSize: 1223456,
	}

	modelData = append(modelData, data)
	historicalData := &FakeHistoryData{
		HistoryModel: modelData,
		Error: nil,
	}

	histSrv := &History{
		Data: historicalData,
	}

	req := httptest.NewRequest(http.MethodGet, "/api/history", nil)
	metadata := models.ContextMetaData{
		Email: "bistmilan46@gmail.com",
		UserId: 1,
	}
	ctx := context.WithValue(req.Context(), "metaData", metadata)
	req = req.WithContext(ctx)
	res := httptest.NewRecorder()
	histSrv.HandleHistory(res, req)

	if res.Code != http.StatusOK{
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.Code)
	}


	var responseData HistoryResponse

	err := json.NewDecoder(res.Body).Decode(&responseData)
	if err != nil {
		t.Fatal(err)
	}


	ndata := responseData.Data
	if ndata[0].ActualFile.Filename != "image.raw"{
		t.Fatal("Required image.raw got: ",ndata[0].ActualFile.Filename)
	}

}