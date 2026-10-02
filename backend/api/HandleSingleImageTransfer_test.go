package api

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/image-generator/internal/models"
)

type FakeSingleImageProperty struct{}

func(f *FakeSingleImageProperty) GetBytedImage(storageKey string)([]byte,error){ 
	return []byte("fake image data"),nil 
}

func(f *FakeSingleImageProperty) GetImageStorageKey(imageId,userId int)(string,error){ 
	return "fake/storage/image.jpg",nil 
}

func TestHandleSingleImageProperty(t *testing.T){
	fake := &FakeSingleImageProperty{}
	handler := &SingleImageProperty{
		Transfer:fake,
		Store:fake,
	}

	req := httptest.NewRequest(http.MethodGet,"/getSingleImage?id=10&mimetype=image/jpeg",nil)

	metaData := models.ContextMetaData{
		UserId:1,
		Email: "bistmilan46@gmail.com",
	}
	ctx := context.WithValue(req.Context(),"metaData",metaData)
	req = req.WithContext(ctx)

	res := httptest.NewRecorder()

	handler.HandleSingleImageProperty(res,req)



	if res.Code !=  200{
		t.Fatalf("Response should be %d but its %d", 200, res.Code)
	}

	task := res.Body
	bytedTask := task.Bytes()

	toBeCheckedOne := []byte("fake image data")
	if !bytes.Equal(bytedTask, toBeCheckedOne){
		t.Fatal("Failed to get the correct byted response.")
	}
}