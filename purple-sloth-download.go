package main

import (
	"fmt"
	"net/http"
	"os"

	"image"
	"image/color"
	"image/jpeg"
)

// purpleJPEG takes in an image and makes each of pixels more purple by
// incrementing its red and blue hex values by 50.
func purpleJPEG(img image.Image) {
	pixels := image.NewRGBA64(img.Bounds())
	for x := pixels.Bounds().Min.X; x < pixels.Bounds().Max.X; x++ {
		for y := pixels.Bounds().Min.Y; y < pixels.Bounds().Max.Y; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			r = r>>8 + 50
			b = b>>8 + 50
			if r > 256 {
				r = 256
			}
			if b > 256 {
				b = 256
			}
			pixels.Set(x, y, color.RGBA64{R: uint16(r << 8), G: uint16(g), B: uint16(b << 8), A: uint16(a)})
		}
	}
	jpeg.Encode(os.Stdout, pixels, nil)
}

func main() {
	// Make a Client and new HTTP request
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://upload.wikimedia.org/wikipedia/commons/1/18/Bradypus.jpg", nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// Send the request and get back the HTTP response
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	// Decode the response body into an Image and pass it to purpleJPEG
	img, _, err := image.Decode(res.Body)
	if err != nil {
		fmt.Println("Error decoding JPEG image: ", err)
		return
	}
	purpleJPEG(img)
}
