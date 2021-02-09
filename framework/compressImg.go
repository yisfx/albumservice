package framework

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

	"github.com/nfnt/resize"
)

func CompressImg(source string, hight uint) error {
	var err error
	var file *os.File
	reg, _ := regexp.Compile(`^.*\.((png)|(jpg))$`)
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
	case strings.HasSuffix(name, ".png"):
		if img, err = png.Decode(file); err != nil {
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
	resizeImg := resize.Resize(hight, 0, img, resize.Lanczos3)
	newName := newName(source, int(hight))
	if outFile, err := os.Create(newName); err != nil {
		return err
	} else {
		defer outFile.Close()
		err = jpeg.Encode(outFile, resizeImg, nil)
		if err != nil {
			return err
		}
	}
	abspath, _ := filepath.Abs(newName)
	return nil
}

//create a file name for the iamges that after resize
func newName(name string, size int) string {
	dir, file := filepath.Split(name)
	return fmt.Sprintf("%s_%d%s", dir, size, file)
}
