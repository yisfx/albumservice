package utils

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func CompressJpgResource(orgPath string, targetPath string, quality int) {

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
func WaterMark(orgPath string, targetPath string) {

	imgb, _ := os.Open("ying.jpg")
	img, _ := jpeg.Decode(imgb)
	defer imgb.Close()

	wmb, _ := os.Open("fly.png")
	watermark, _ := png.Decode(wmb)
	defer wmb.Close()

	//把水印写到右下角，并向0坐标各偏移10个像素
	offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx()-10, img.Bounds().Dy()-watermark.Bounds().Dy()-10)
	b := img.Bounds()
	m := image.NewNRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	imgw, _ := os.Create("new.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{Quality: 100})

	defer imgw.Close()
}
