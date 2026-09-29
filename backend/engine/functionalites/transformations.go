package functionalities

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg" // like the init in python  runs its header when decode it called to detect the jpeg files itself
	_ "image/png"
	"log"
	"os"
)

// make the image grey
func makeGreyScale(path, filePathToSave, mimetype string) (string, error){
	// open the given file
	file, err1 := os.Open(path)
	if err1 != nil {
		log.Println("Error in opening the filepath. ", err1)
	}

	// now find the bound of the given file
	// decode the given file
	img, _, err2 := image.Decode(file)

	if err2 != nil {
		fmt.Println("Error in decoding the file. ", err2)
		return "", err2
	}

	// now I have the access to the bounds of the given file
	b := img.Bounds()

	// setup for the newGray which will store the gray pix of each of pixels that I got
	grayImage := image.NewGray(b)

	// process each point row by row
	for y := 0; y < b.Max.Y; y++ {
		for x := 0; x < b.Max.X; x++ {
			// get the pixel at each of the point here
			pixel := img.At(x, y)
			r, g, b, _ := pixel.RGBA()
			R := uint8(r >> 8)
			G := uint8(g >> 8)
			B := uint8(b >> 8)

			grayPix := 0.299*float64(R) + 0.587*float64(G) + 0.114*float64(B)

			// set the gray Pix to the given grayImage
			grayImage.SetGray(x, y, color.Gray{Y: uint8(grayPix)})
		}
	}

	var newLocation string
	// check for the mimetype and create a new file
	if mimetype == "image/png"{
		newLocation, err1 = transformBasedOnMimeType(mimetype,filePathToSave, ".png", grayImage)
		fmt.Println("Grey image Portion: ", err1)
		return "", err1
	} else{
		newLocation, err1 = transformBasedOnMimeType(mimetype,filePathToSave, ".jpg", grayImage)
		if err1 != nil{
			fmt.Println("Grey image Portion: ", err1)
			return "", err1
		}
	}

	return newLocation, nil
}


// increase the brightness of the image
func increaseBrightness(path, filePathToSave, mimetype string) (string,error){
	// add certain value to the r,g and b value in order to make it more brightening

	// open the file
	file, err1 := os.Open(path)
	defer file.Close()

	if err1 != nil {
		log.Println("Error in opening the file")
	}
	// decode the file
	img, _, err2 := image.Decode(file)
	if err2 != nil {
		log.Println("Error in decoding the file", err2)
	}
	// get each of the pixel grid
	b := img.Bounds()
	maxX := b.Max.X
	maxY := b.Max.Y

	// make a new grid with the rgba
	brightedImage := image.NewRGBA(b)

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			pix := img.At(x, y)
			r, g, b, a := pix.RGBA()

			// convert each of them in the uint8 format
			R := int(uint8(r >> 8))
			G := int(uint8(g >> 8))
			B := int(uint8(b >> 8))
			A := uint8(a >> 8)

			// check for the overflow
			R = R + 20
			if R > 255 {
				R = 255
			}
			G = G + 20
			if G > 255 {
				G = 255
			}
			B = B + 20
			if B > 255 {
				B = 255
			}

			r1 := uint8(R)
			g1 := uint8(G)
			b1 := uint8(B)

			// now store them all in the brightedImge place
			brightedImage.SetRGBA(x, y, color.RGBA{r1, g1, b1, A})
		}
	}
	// add the images data to the cerating
	var newLocation string
	// check for the mimetype and create a new file
	if mimetype == "image/png"{
		newLocation, err1 = brightnessBasedOnMimeType(mimetype,filePathToSave, ".png", brightedImage, "increase")
		fmt.Println("Increase Brigntness image Portion: ", err1)
		return "", err1
	} else{
		newLocation, err1 = brightnessBasedOnMimeType(mimetype,filePathToSave, ".jpg", brightedImage, "increase")
		if err1 != nil{
			fmt.Println("Increase Brightness image Portion: ", err1)
			return "", err1
		}
	}

	return newLocation, nil

}

// decrease the brightness of the image
func decreaseBrightness(path, filePathToSave, mimetype string) (string,error) {
	// subtract the given value with the delta value to make it lose it being more brightening
	// open the file
	file, err1 := os.Open(path)
	if err1 != nil {
		log.Println("Error in opening the file. ", err1)
	}
	defer file.Close()

	// decode the contents and make a new RGB there
	img, _, err2 := image.Decode(file)
	if err2 != nil {
		log.Println("Error in decoding the file. ", err2)
	}

	// find the bounds
	b := img.Bounds()
	maxX := b.Max.X
	maxY := b.Max.Y

	darkenedImage := image.NewRGBA(b)
	// get each of the pixel and make a new whole grid from it with the RGBA
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			pixel := img.At(x, y)
			r, g, b, a := pixel.RGBA()

			R := int(uint8(r >> 8))
			G := int(uint8(g >> 8))
			B := int(uint8(b >> 8))
			A := uint8(a >> 8)

			gamma := 40
			// now subtract certain
			R = R - gamma
			if R < 0 {
				R = 0
			}
			G = G - gamma
			if G < 0 {
				G = 0
			}
			B = B - gamma
			if B < 0 {
				B = 0
			}

			// again convert each of the following in the same type
			r1 := uint8(R)
			g1 := uint8(G)
			b1 := uint8(B)

			// set at the position of the given one
			darkenedImage.SetRGBA(x, y, color.RGBA{r1, g1, b1, A})
		}
	}

	// add the images data to the cerating
	var newLocation string
	// check for the mimetype and create a new file
	if mimetype == "image/png"{
		newLocation, err1 = brightnessBasedOnMimeType(mimetype,filePathToSave, ".png",darkenedImage, "decrease")
		fmt.Println("Decrease Brigntness image Portion: ", err1)
		return "", err1
	} else{
		newLocation, err1 = brightnessBasedOnMimeType(mimetype,filePathToSave, ".jpg", darkenedImage, "decrease")
		if err1 != nil{
			fmt.Println("Decrease Brightness image Portion: ", err1)
			return "", err1
		}
	}

	return newLocation, nil
}

func PerformTransformations(specificity, filepath, filePathToSave, mimetype string) (string, error){
	var path string
	var err error
	switch specificity {
	case "greyscale":
		path, err = makeGreyScale(filepath, filePathToSave, mimetype)
		if err != nil{
			fmt.Println("From grey scale place. Error: ", err)
			return "", errors.New("Error from making grey scale part.")
		}

	case "increaseBrightness":
		path, err = increaseBrightness(filepath, filePathToSave, mimetype)
		if err != nil{
			fmt.Println("From increase brightness place. Error: ", err)
			return "", errors.New("Error in increasing brightness part.")
		}

	case "decreaseBrightness":
		path, err = decreaseBrightness(filepath, filePathToSave,mimetype)
		if err != nil{
			fmt.Println("From decrease brightness place. Error: ", err)
			return "", errors.New("Error from decreaseBrightness part.")
		}

	}

	return path, nil
}