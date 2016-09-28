package imgbin

import (
	"bufio"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"reflect"
	"strconv"

	"../fontrender"
	log "../loghandler"
)

type empty struct{}

var pkg = reflect.TypeOf(empty{}).PkgPath()

func newImg(txtLen int) *image.RGBA {
	dpiPadding := fontrender.GetDPI() + 3
	return image.NewRGBA(image.Rect(0, 0, dpiPadding*txtLen, 200))
}

func getWd() (path string, err error) {
	return os.Getwd()
}

func createImgPath(text, imgPath string) (path, name string, err error) {
	imgName := "img_" + text + ".png"
	return filepath.Join(imgPath, imgName), imgName, nil
}

//CreateImgBin something
func CreateImgBin(count int, fontPath, imgPath string) (path string, err error) {
	log.LogInit()
	text := strconv.Itoa(count)
	imgPath, imgName, err := createImgPath(text, imgPath)
	if err != nil {
		log.Logger(log.Err, pkg, "Error setting image path:"+err.Error())
		return "", err
	}
	imgFile, err := os.Create(imgPath)
	if err != nil {
		log.Logger(log.Err, pkg, "Error creating image file")
		return "", err
	}
	defer imgFile.Close()

	rgba := newImg(len(text))
	rgba, err = fontrender.DrawFont(rgba, text, fontPath)
	if err != nil {
		return "", err
	}
	buff := bufio.NewWriter(imgFile)
	if err := png.Encode(buff, rgba); err != nil {
		log.Logger(log.Err, pkg, "Error encoding png for file: "+err.Error())
		return "", err
	}
	if err := buff.Flush(); err != nil {
		log.Logger(log.Err, pkg, "Error flushing buffer to file: "+err.Error())
		return "", err
	}

	return imgName, nil
}
