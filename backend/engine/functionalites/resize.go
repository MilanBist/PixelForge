package functionalities

import (
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg" // like the init in python  runs its header when decode it called to detect the jpeg files itself
	_ "image/png"
	"log"
	"os"
	"strconv"
)

// making the image large
func makeLarge(filepath string, times int) {
	// read the file and get it bounds and create a new bound from it which will be of double size perfectly

	// open the file
	file, err1 := os.Open(filepath)

	if err1 != nil {
		log.Println("Error in opening the file for resize. ", err1)
	}

	// decode the file
	img, _, err2 := image.Decode(file)

	if err2 != nil {
		log.Println("Error in decoding the file. ", err2)
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

	// make a newFile called as largeimages
	folderPath := "largeFolder"

	_, err5 := os.Stat(folderPath)
	if err5 != nil {
		// means like the folder doesn't exist so that create
		err6 := os.Mkdir(folderPath, 0755)
		if err6 != nil {
			log.Println("Error in creating a new folder to create a new folder for it")
		}
	}
	// create a new file path and also
	filepath = folderPath + "/" + strconv.Itoa(times) + "times" + "larged" + filepath
	// just print it in the new image
	newImage, err3 := os.Create(filepath)

	if err3 != nil {
		log.Println("Error in creating a new file.", err3)
	}
	err4 := jpeg.Encode(newImage, largedImage, &jpeg.Options{Quality: 100})
	if err4 != nil {
		log.Println("Error in encoding the image in new jpeg file.")
	}

}

// what to do under resizing
func PerformResize(specificity, filepath string) {

	switch specificity {
	case "onex":
		makeLarge(filepath, 1)
	case "twox":
		makeLarge(filepath, 2)
	case "threex":
		makeLarge(filepath, 3)
	}
}