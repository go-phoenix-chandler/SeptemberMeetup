package fontrender

import (
	"image"
	"reflect"
	"testing"

	"github.com/dtoebe/gophx-img-api/settingsloader"
	. "github.com/dtoebe/gophx-img-api/testerror"
)

var testRGBA = image.NewRGBA(image.Rect(0, 0, 10, 10))

func TestGetDPI(t *testing.T) {
	res := GetDPI()
	if res != dpi {
		t.Errorf("GetDPI: Expected: %d; Got: %d\n", dpi, res)
	}
}

func TestDrawFont(t *testing.T) {
	s := settingsloader.NewSettings("../server-config.json")
	exp := "*image.RGBA"
	df, err := DrawFont(testRGBA, "999", s.FontPath)
	res := reflect.TypeOf(df).String()
	if err != nil {
		ErrorOut("DrawFont", exp, res, t)
	}
	ErrorOut("DrawFont", exp, res, t)
}

func TestGetFontFace(t *testing.T) {
	s := settingsloader.NewSettings("../server-config.json")
	exp := "*truetype.Font"
	f, err := getFontFace(s.FontPath)
	res := reflect.TypeOf(f).String()
	ErrorOut("getFontFace", exp, res, t)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetContext(t *testing.T) {
	s := settingsloader.NewSettings("../server-config.json")
	exp := "*freetype.Context"
	f, _ := getFontFace(s.FontPath)
	res := reflect.TypeOf(getContext(f, testRGBA)).String()
	ErrorOut("getContext", exp, res, t)
}
