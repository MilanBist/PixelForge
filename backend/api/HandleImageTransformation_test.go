package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/image-generator/internal/models"
)

type FakeImageTransformation struct {
	// Values returned by the fake methods
	UploadedID  int
	GeneratedID int64
	UploadedPath  string
	GeneratedPath string
	Width  int
	Height int
	FileSize int64
}

type SingleImageTransformationResponse struct{
    Success bool                                     	`json:"success"`
    Message string                                   	`json:"message"`
    Data    models.FileBasedImageGenerationReturn 	`json:"data"`
}

func (f *FakeImageTransformation) AddRawFileToDestination(file multipart.File, filename string, userId string,) (string, int, error) {
	return f.UploadedPath, http.StatusOK, nil
}

func (f *FakeImageTransformation) GenerateImageTransformations(mainTask string,subMainTask string,filePath string,userId string,mimetype string,) (string, error) {
	return f.GeneratedPath, nil
}

func (f *FakeImageTransformation) AddUploadedImageFiles(uploadedMetaData models.UploadedFilesMetaData,) (int, error) {
	return f.UploadedID, nil
}

func (f *FakeImageTransformation) AddGeneratedFiles(generatedFilesMetaData models.GeneratedImageMetaData,) (int64, error) {
	return f.GeneratedID, nil
}


func (f *FakeImageTransformation) GetDimension(location string,) (int, int, int64, error) {
	return f.Width, f.Height, f.FileSize, nil
}

func TestHandleSingleImageTransformation(t *testing.T) {

	// Fake data
	fake := &FakeImageTransformation{
		UploadedID:    10,
		GeneratedID:   25,
		UploadedPath:  "/tmp/uploaded/image.jpg",
		GeneratedPath: "/tmp/generated/image.jpg",
		Width:         512,
		Height:        512,
		FileSize:      12345,
	}

	// Create the handler
	handler := &ImageTransformation{
		UploadGenerate: fake,
		Store:          fake,
		Dimension:      fake,
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// Add image file
	fileWriter, err := writer.CreateFormFile("file", "test.jpg")
	if err != nil {
		t.Fatal(err)
	}

	// Fake image contents
	_, err = fileWriter.Write([]byte("fake image data"))
	if err != nil {
		t.Fatal(err)
	}

	// Add mainTask
	err = writer.WriteField("mainTask", "resize")
	if err != nil {
		t.Fatal(err)
	}

	// Add subMainTask
	err = writer.WriteField("subMainTask", "2x")
	if err != nil {
		t.Fatal(err)
	}

	writer.Close()


	req := httptest.NewRequest(http.MethodPost,"/transformImage",&body,)

	req.Header.Set("Content-Type",writer.FormDataContentType(),)

	// Add user metadata to context
	metaData := models.ContextMetaData{
		UserId: 1,
		Email: "bistmilan46@gmail.com",
	}

	ctx := context.WithValue(
		req.Context(),
		"metaData",
		metaData,
	)

	req = req.WithContext(ctx)

	res := httptest.NewRecorder()

	handler.HandleSingleImageTransformation( res, req )


	if res.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d",http.StatusOK,res.Code)
	}

	var response SingleImageTransformationResponse

	json.NewDecoder(res.Body).Decode(&response)

	var generatedId int64
	if floatVal, ok := response.Data.GeneratedImage[0].Id.(float64); ok {
		generatedId = int64(floatVal)
	} else {
		t.Fatalf("Error: The ID is neither an int64 nor a float64.")
	}
	fmt.Printf("%T",int(generatedId))
	if generatedId != 25{
		t.Fatalf("Expected %d and got %d",fake.GeneratedID, generatedId)
	}
}