package imguri

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"reflect"
	"strconv"

	"../fontrender"
	log "../loghandler"
)

type empty struct{}

var pkg = reflect.TypeOf(empty{}).PkgPath()

//newImg takes the length of the twxt string and creates an
// image.RGBA
func newImg(txtLen int) *image.RGBA {
	dpiPadding := fontrender.GetDPI() + 3
	return image.NewRGBA(image.Rect(0, 0, dpiPadding*txtLen, 200))
}

//CreatImgData takes the count int and the path to the wanted font
// and outputs the basic datauri and an error
func CreatImgData(count int, fontPath string) (dataURI string, err error) {
	log.LogInit()
	text := strconv.Itoa(count)
	rgba := newImg(len(text))
	rgba, err = fontrender.DrawFont(rgba, text, fontPath)
	if err != nil {
		return "", err
	}
	data := new(bytes.Buffer)
	if err := png.Encode(data, rgba); err != nil {
		log.Logger(log.Err, pkg, "Unable to encode PNG: "+err.Error())
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data.Bytes()), nil
}
