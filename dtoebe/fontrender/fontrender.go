package fontrender

//FontRender package:
//Takes thirdparty packages:
// - freetype
// - freetype/truetype
// - image/font
//Purpose:
// - Takes a created image.RGBA and the path to the fontFile
//   and the text wanted to write to the image.RGBA writes the
//   text and returns it.

import (
	"image"
	"io/ioutil"
	"reflect"

	log "../loghandler"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type empty struct{}

var pkg = reflect.TypeOf(empty{}).PkgPath()

const (
	fontSize = 125.0
	dpi      = 72
)

//GetDPI I don't like mking package const exportable
func GetDPI() int {
	return dpi
}

//DrawFont takes in the image what test to be printed and the path to
//the font file draws the text to the image and returns the image
func DrawFont(rgba *image.RGBA, text, fontPath string) (*image.RGBA, error) {
	log.LogInit()
	fontParse, err := getFontFace(fontPath)
	if err != nil {
		return new(image.RGBA), err
	}

	ctx := getContext(fontParse, rgba)

	opts := truetype.Options{}
	opts.Size = fontSize
	face := truetype.NewFace(fontParse, &opts)

	for i, x := range text {
		aWidth, ok := face.GlyphAdvance(rune(x))
		if ok != true {
			log.Logger(log.Err, pkg, "Failed to place letter on image")
			return new(image.RGBA), nil
		}

		iWidthf := int(float64(aWidth) / 20)
		pt := freetype.Pt(i*dpi+(int(fontSize-iWidthf/2)), 128)
		ctx.DrawString(string(x), pt)
	}

	return rgba, nil
}

//getFontFace takes the path to the font file and returns a buffer of
// the font
func getFontFace(fontPath string) (*truetype.Font, error) {
	fontFile, err := ioutil.ReadFile(fontPath)
	if err != nil {
		log.Logger(log.Err, pkg, "Unable to read from font file: "+err.Error())
		return new(truetype.Font), err
	}

	return truetype.Parse(fontFile)
}

//getContext takes the font buffer and the images to set the font
// settings and returns them
func getContext(f *truetype.Font, rgba *image.RGBA) (ctx *freetype.Context) {
	ctx = freetype.NewContext()
	ctx.SetDPI(dpi)
	ctx.SetFont(f)
	ctx.SetFontSize(fontSize)
	ctx.SetClip(rgba.Bounds())
	ctx.SetDst(rgba)
	ctx.SetSrc(image.Black)
	ctx.SetHinting(font.HintingNone)

	return ctx
}
