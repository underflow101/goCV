package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	fmt.Println("Starting system...")
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
