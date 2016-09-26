package settingsloader

import (
	"encoding/json"
	"os"
	"reflect"
)

type empty struct{}

var pkg = reflect.TypeOf(empty{}).PkgPath()

//Settings takes all the JSON data types
type Settings struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	ImgDir   string `json:"image-dir"`
	FontPath string `json:"font-path"`
	RootURL  string `json:"root-url"`
	URLs     []Urls `json:"urls"`
}

//Urls is reference to the URLs allowed
type Urls struct {
	URL      string `json:"url"`
	DataType string `json:"data-type"`
}

//NewSettings takes a path string and loads decodes the JSON to the settings struct
func NewSettings(p string) *Settings {
	file, err := os.Open(p)
	if err != nil {

	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	set := new(Settings)
	decoder.Decode(set)

	return set
}
