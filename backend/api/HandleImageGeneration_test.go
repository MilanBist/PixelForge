package api

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/image-generator/engine"
	"github.com/image-generator/internal/models"
)

type FakeMetaDataUpload struct{
	ExactPath string
	StatusCode int
	Error error
	AllFiles engine.AllFiles 
}

type UploadGeneration struct{
	UploadedId 	int
	ImageId  int64
	Error 	error
}

type FakeDimension struct{
	Height int
	Width int
	Size int64
	Error error
}

func(f *FakeMetaDataUpload) AddRawFileToDestination(file multipart.File, fileName, userId string)(string, int, error){
	return f.ExactPath, f.StatusCode, f.Error
}

func (f *FakeMetaDataUpload) GenerateImage(exactFilePath, userId string) (engine.AllFiles, int, error){
	return f.AllFiles, f.StatusCode, f.Error
}

func (u *UploadGeneration) AddUploadedFiles(uploadedMetaData models.UploadedFilesMetaData)(int, error){
	return u.UploadedId, u.Error
}

func (u *UploadGeneration) AddGeneratedFiles(generatedFilesMetaData models.GeneratedImageMetaData)(int64, error){
	return u.ImageId, u.Error
}
func (u *UploadGeneration) AddToHistoryOfUser(historyData models.History) (error){
	return nil
}

func (fd *FakeDimension) GetDimension(location string) (int, int, int64, error){
	return fd.Height, fd.Width, fd.Size, fd.Error
}

// create the testing function here
func HandleImageGeneration_CheckingActualResponse(t *testing.T){
	allFilesData := engine.AllFiles{
		JpgFiles: []string{"./images/generated/user/1/image2.jpg"},
		PngFiles: []string{"./images/generated/user/1/image2.png"},
	}
	forFileUpload := &FakeMetaDataUpload{
		ExactPath: "./images/uploaded/user/1/card.raw",
		StatusCode: http.StatusOK,
		Error: nil,
		AllFiles: allFilesData,
	}
	store := &UploadGeneration{
		UploadedId: 1,
		ImageId: 1,
	}

	dimension := &FakeDimension{
		Height: 1024,
		Width: 1024,
		Size: 123456,
		Error: nil,
	}


	ih := &ImageGeneratorHandler{
		Savator: forFileUpload,
		Store: store,
		Dimension: dimension,
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "card.raw")
	if err != nil {
		t.Fatal(err)
	}

	fakeRawData := []byte("fake raw file contents")

	_, err = part.Write(fakeRawData)
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Close()
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/getImages",
		body,
	)

	req.Header.Set("Content-Type", writer.FormDataContentType())

	res := httptest.NewRecorder()

	ih.HandleImageGeneration(res, req)
	

	if res.Code != http.StatusOK {
    	t.Errorf("expected status %d, got %d", http.StatusOK, res.Code)
	}


	var response map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	// each method is being called so the handler works perfectly
}