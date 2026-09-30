package functionalities

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg" // like the init in python  runs its header when decode it called to detect the jpeg files itself
	_ "image/png"
	"log"
	"os"
)

// making the image large
func makeLarge(filepath, filePathToSave, mimetype string, times int) (string, error){
	// read the file and get it bounds and create a new bound from it which will be of double size perfectly

	// open the file
	file, err1 := os.Open(filepath)

	if err1 != nil {
		log.Println("[resizing portion]: Error in opening the file. ", err1)
		return "", err1
	}

	// decode the file
	img, _, err2 := image.Decode(file)

	if err2 != nil {
		log.Println("[resizing portion]: Error in decoding the file. ", err2)
		return "", err2
	}

	// find the bounds of the given image and just make it double
	b := img.Bounds()
	maxX := b.Max.X
	maxY := b.Max.Y

	newMaxX := maxX * times
	newMaxY := maxY * times

	newRectange := image.Rect(0, 0, newMaxX, newMaxY)
	// make a newRGBA color model to store the new Image\
	largedImage := image.NewNRGBA(newRectange)

	// loop over each of the pixels as a grid
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			// get the pixel at the given region
			pixel := img.At(x, y)
			r, g, b, a := pixel.RGBA()

			R := uint8(r >> 8)
			G := uint8(g >> 8)
			B := uint8(b >> 8)
			A := uint8(a >> 8)

			if times == 1 {
				largedImage.SetNRGBA(x, y, color.NRGBA{R, G, B, A})
				continue
			}

			// set at the position 2x,2y then 2x,2y+1 then 2x+1, 2y and at last 2x+1, 2y+1

			if times == 2 {
				largedImage.SetNRGBA(2*x, 2*y, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(2*x+1, 2*y, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(2*x, 2*y+1, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(2*x+1, 2*y+1, color.NRGBA{R, G, B, A})
				continue
			}

			if times == 3 {
				largedImage.SetNRGBA(3*x, 3*y, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(3*x+1, 3*y, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(3*x+2, 3*y, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(3*x, 3*y+1, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(3*x+1, 3*y+1, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(3*x+2, 3*y+1, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(3*x, 3*y+2, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(3*x+1, 3*y+2, color.NRGBA{R, G, B, A})
				largedImage.SetNRGBA(3*x+2, 3*y+2, color.NRGBA{R, G, B, A})
			}

		}
	}

	mapTimes := map[int]string{
		1 : "onex",
		2 : "twox",
		3 : "threex",
	}

	var newLocation string
	// check for the mimetype and create a new file
	if mimetype == "image/png"{
		newLocation, err1 = resizeBasedOnMimeType(mimetype,filePathToSave, ".png", mapTimes[times], largedImage)
		if err1 != nil{
			fmt.Println("Size increasing image Portion: ", err1)
			return "", err1
		}
	} else{
		newLocation, err1 = resizeBasedOnMimeType(mimetype,filePathToSave, ".jpg", mapTimes[times], largedImage)
		if err1 != nil{
			fmt.Println("Grey image Portion: ", err1)
			return "", err1
		}
	}

	return newLocation, nil

}

// what to do under resizing
func PerformResize(specificity, filepath, filePathToSave, mimeType string) (string, error){
	var path string
	var err error
	switch specificity {
	case "onex":
		path, err = makeLarge(filepath, filePathToSave, mimeType, 1)
		if err != nil{
			return "", err
		}
	case "twox":
		path, err = makeLarge(filepath, filePathToSave, mimeType, 2)
		if err != nil{
			return "", err
		}
	case "threex":
		path, err = makeLarge(filepath,filePathToSave, mimeType, 3)
		if err != nil{
			return "", err
		}
	}

	return path, err
}