package functionalities

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"github.com/google/uuid"
)


func transformBasedOnMimeType(mimeType, orignalLocation, fileWithExtension string, img *image.Gray)(string, error){
	id := uuid.New()
	imageName := "grey_"+id.String()+fileWithExtension
	newFilePath := filepath.Join(orignalLocation, imageName)

	newFile, err := os.Create(newFilePath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()

	if mimeType == "image/png"{
		err = png.Encode(newFile, img)
		if err != nil {
			return "", err
		}
	} else{
		err = jpeg.Encode(newFile, img, &jpeg.Options{Quality: 100})
		if err != nil {
			return "", err
		}
	}

	return newFilePath, nil
}

func brightnessBasedOnMimeType(mimeType, orignalLocation, fileWithExtension string, img *image.RGBA, typo string)(string, error){
	id := uuid.New()
	imageName := typo+"brightness_"+id.String()+fileWithExtension
	newFilePath := filepath.Join(orignalLocation, imageName)

	newFile, err := os.Create(newFilePath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()

	if mimeType == "image/png"{
		err = png.Encode(newFile, img)
		if err != nil {
			return "", err
		}
	} else{
		err = jpeg.Encode(newFile, img, &jpeg.Options{Quality: 100})
		if err != nil {
			return "", err
		}
	}
	

	return newFilePath, nil
}

func resizeBasedOnMimeType(mimeType, orignalLocation, fileWithExtension, specificity string, img *image.NRGBA)(string, error){
	id := uuid.New()
	imageName := specificity+"_"+id.String()+fileWithExtension
	newFilePath := filepath.Join(orignalLocation, imageName)

	newFile, err := os.Create(newFilePath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()

	if mimeType == "image/png"{
		err = png.Encode(newFile, img)
		if err != nil {
			return "", err
		}
	} else{
		err = jpeg.Encode(newFile, img, &jpeg.Options{Quality: 100})
		if err != nil {
			return "", err
		}
	}
	return newFilePath, nil
}


// save the blur image in the certain place and return the newFilepath of the blurred image
func blurBasedOnMimeType(mimeType, orignalLocation, fileWithExtension string, img *image.NRGBA)(string, error){
	id := uuid.New()
	imageName := "blur_"+id.String()+fileWithExtension
	newFilePath := filepath.Join(orignalLocation, imageName)

	newFile, err := os.Create(newFilePath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()

	if mimeType == "image/png"{
		err = png.Encode(newFile, img)
		if err != nil {
			return "", err
		}
	} else{
		err = jpeg.Encode(newFile, img, &jpeg.Options{Quality: 100})
		if err != nil {
			return "", err
		}
	}
	return newFilePath, nil
}