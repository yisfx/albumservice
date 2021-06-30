package utils

import (
	"fmt"
	"image/jpeg"
	"os"
)

func CompressJpgResource(orgPath string, targetPath string, quality int) {
	defer ErrorHandler()
	///org
	orgImg, err := os.Open(orgPath)
	if err != nil {
		fmt.Println(err)
	}
	defer orgImg.Close()

	///to
	targetImg, err := os.Create(targetPath)
	if err != nil {
		fmt.Println(err)
	}
	defer targetImg.Close()

	img, err := jpeg.Decode(orgImg)
	if err != nil {
		fmt.Println(err)
	}
	jpeg.Encode(targetImg, img, &jpeg.Options{Quality: quality})
}
