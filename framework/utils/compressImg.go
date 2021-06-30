package utils

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	resize "github.com/nfnt/resize"
)

func CompressImg(source string, wide uint, newName string) error {
	defer ErrorHandler()
	var err error
	var file *os.File
	reg, _ := regexp.Compile(`^.*\.((png)|(jpg)|(JPG)|(PNG))$`)
	if !reg.MatchString(source) {
		err = errors.New("%s is not a .png or .jpg file")
		return err
	}
	if file, err = os.Open(source); err != nil {
		return err
	}
	defer file.Close()
	name := file.Name()
	var img image.Image
	switch {
	case strings.HasSuffix(name, ".PNG"):
		if img, err = png.Decode(file); err != nil {
			return err
		}
	case strings.HasSuffix(name, ".png"):
		if img, err = png.Decode(file); err != nil {
			return err
		}
	case strings.HasSuffix(name, ".JPG"):
		if img, err = jpeg.Decode(file); err != nil {
			return err
		}
	case strings.HasSuffix(name, ".jpg"):
		if img, err = jpeg.Decode(file); err != nil {
			return err
		}
	default:
		err = fmt.Errorf("Images %s name not right!", name)
		return err
	}
	resizeImg := resize.Resize(wide, 0, img, resize.Lanczos3)
	if outFile, err := os.Create(newName); err != nil {
		return err
	} else {
		defer outFile.Close()
		err = jpeg.Encode(outFile, resizeImg, nil)
		if err != nil {
			return err
		}
	}
	filepath.Abs(newName)
	return nil
}

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
