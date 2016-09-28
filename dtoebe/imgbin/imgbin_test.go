package imgbin

import (
	"os"
	"reflect"
	"testing"

	"github.com/dtoebe/gophx-img-api/settingsloader"
	. "github.com/dtoebe/gophx-img-api/testerror"
)

func TestNewImg(t *testing.T) {
	exp := "*image.RGBA"
	res := reflect.TypeOf(newImg(100)).String()
	ErrorOut("newImg", exp, res, t)
}

func TestGetWD(t *testing.T) {
	exp, _ := os.Getwd()
	res, err := getWd()
	if exp != res {
		ErrorOut("getWd", exp, res, t)
	}
	if err != nil {
		ErrorOut("getWd", "", err.Error(), t)
	}
}

func TestCreateImgPath(t *testing.T) {
	s := settingsloader.NewSettings("../server-config.json")
	txt := "9999"
	path, name, err := createImgPath(txt, s.ImgDir)
	expName := "img_" + txt + ".png"
	if name != expName {
		ErrorOut("createImgPath: name", expName, name, t)
	}
	if path != s.ImgDir+"/"+expName {
		ErrorOut("createImgPath: path", s.ImgDir+"/"+expName, path, t)
	}
	if err != nil {
		ErrorOut("createImgPath", "", err.Error(), t)
	}
}

func TestCreateImgBin(t *testing.T) {
	s := settingsloader.NewSettings("../server-config.json")
	_, exp1, _ := createImgPath("1000", s.ImgDir)
	exp2 := "string"
	res1, err := CreateImgBin(1000, s.FontPath, s.ImgDir)
	res2 := reflect.TypeOf(res1).String()
	ErrorOut("CreateImgBin", exp1, res1, t)
	ErrorOut("CreateImgBin", exp2, res2, t)
	if err != nil {
		ErrorOut("CreateImgBin: error", "", err.Error(), t)
	}
}
