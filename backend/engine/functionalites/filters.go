package functionalities


import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg" // like the init in python  runs its header when decode it called to detect the jpeg files itself
	_ "image/png"
	"os"
)

// making the image look blur
func makeBlur(filepath, filePathToSave, mimetype string) (string, error){
	// open the given image
	file, err1 := os.Open(filepath)

	if err1 != nil {
		fmt.Println("[blurring section]Error in opening the image file. ", err1)
		return "", err1
	}

	// decode the given file
	img, _, err2 := image.Decode(file)

	if err2 != nil {
		fmt.Println("[blurring section]Error in decoding the image file. ", err1)
		return "", err1
	}

	// get the bounds of the given image
	b := img.Bounds()

	// set the new NRGBA for the given image
	blurredImage := image.NewNRGBA(b)

	// get the max X and max Y till where to get the loop
	maxX := b.Max.X
	maxY := b.Max.Y

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			// get the pixel at the given range
			pixel := img.At(x, y)
			r, g, b, a := pixel.RGBA()

			r = r * 0
			g = g * 0
			b = b * 0
			A := uint8(a >> 8)

			// since I got the r,g,b,a value for the given pixel now replace it with the pixel
			// Containing 3x3 matrix including the point itself
			var count uint32
			count = 0
			for i := x - 1; i <= x+1; i++ {
				for j := y - 1; j <= y+1; j++ {
					//  get the image pixel at each of the place
					if j >= 0 && i >= 0 && i < maxX && j < maxY {
						// get the r,g,b and a value at that position
						initialPixel := img.At(i, j)
						r1, g1, b1, _ := initialPixel.RGBA()
						r += r1
						g += g1
						b += b1
						count += 1
					}
				}
			}

			// find the average of all of the given values and then do it
			r = r / count
			g = g / count
			b = b / count

			R := uint8(r >> 8)
			G := uint8(g >> 8)
			B := uint8(b >> 8)

			blurredImage.SetNRGBA(x, y, color.NRGBA{R, G, B, A})
		}
	}
	var newLocation string
	// check for the mimetype and create a new file
	if mimetype == "image/png"{
		newLocation, err1 = blurBasedOnMimeType(mimetype,filePathToSave, ".png", blurredImage)
		fmt.Println("Grey image Portion: ", err1)
		return "", err1
	} else{
		newLocation, err1 = blurBasedOnMimeType(mimetype,filePathToSave, ".jpg", blurredImage)
		if err1 != nil{
			fmt.Println("Grey image Portion: ", err1)
			return "", err1
		}
	}

	return newLocation, nil


}

// making the image look more sharp enough
func makeSharpen(filepath string) {
	// // Sharpening of image via Leplacian filter
	// // basic lap matrix [0 1 0][1-4 1][0 1 0]

	// lapMatrix := [][]int{{0,1,0},{1,-4,1},{0,1,0}}

	// // open the image file
	// file, err1 := os.Open(filepath)

	// if err1 != nil{
	// 	fmt.Println("Error in opening the file with path ", filepath)
	// 	log.Fatal(err1)
	// }

	// // find boundary and get each of the given pixels
	// img, _, err2 := image.Decode(file)

	// if err2 != nil{
	// 	fmt.Println("Error in decoding the file. ")
	// 	log.Fatal(err2)
	// }

	// // find the bounds of teh given files
	// b := img.Bounds()
	// maxY := b.Max.Y
	// maxX := b.Max.X

	// newImg := image.NewNRGBA(b)
	// for y:=0; y<maxY; y++{
	// 	for x:=0; x<maxX; x++{
	// 		// get the pixels at each of the positon
	// 		pixel := img.At(x,y)
	// 		r,g,b,a := pixel.RGBA()

	// 		// now get the matrix of the pixels of the neighbouring pixels from image
	// 		var lapr, lapg, lapb uint32
	// 		if x>0 && x < maxX-1 && y>0 && y<maxX-1{
	// 		for i:=x-1; i<=x+1; i++{
	// 			for j:=y-1; j<= y+1; j++{
	// 				// get the pixels at each of the given position
	// 				currentPix := img.At(i,j)
	// 				r1,g1,b1, _ := currentPix.RGBA()
					
	// 			}
	// 		}
	// 	}
	// 	}
	// }
}

// detecting the edges of the images
func detectEdge(filepath string) {

}

func PerformFilters(specificity, filepath, filePathToSave, mimeType string)(string, error) {
	// blur sharpen and edge detection
	var path string
	var err error
	switch specificity {
	case "blur":
		path, err = makeBlur(filepath, filePathToSave, mimeType)
		if err != nil{
			return "", err
		}
	case "sharpen":
		makeSharpen(filepath)
	case "edge":
		detectEdge(filepath)
	}

	return path, err
}