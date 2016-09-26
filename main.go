package main

import (
	"encoding/json"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/dtoebe/gophx-img-api/imgbin"
	"github.com/dtoebe/gophx-img-api/imguri"
	log "github.com/dtoebe/gophx-img-api/loghandler"
	"github.com/dtoebe/gophx-img-api/settingsloader"
	"github.com/julienschmidt/httprouter"
)

var count = 0

type empty struct{}

var pkg = reflect.TypeOf(empty{}).PkgPath()

type settings struct {
	Data *settingsloader.Settings
}

type response struct {
	DataType  string `json:"data-type"`
	Data      string `json:"data"`
	Error     string `json:"error"`
	TimeStamp string `json:"time-stamp"`
}

func checkAccept(r *http.Request) bool {
	if strings.TrimSpace(r.Header.Get("Accept")) == "application/json" {
		return true
	}
	return false
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := new(response)
	if !checkAccept(r) {
		res.Error = "must use Accept: application/json"
	} else {
		res.Data = "Hello World"
	}

	res.TimeStamp = time.Now().String()

	if err := json.NewEncoder(w).Encode(res); err != nil {
		//error handler
	}
}

func checkFilePath(p string) bool {
	if _, err := os.Stat(p); err == nil {
		return true
	}
	return false
}

func createImageDir(p string) {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		// I know.. I know 777 is bad!
		if err := os.Mkdir(p, 0777); err != nil {
			log.Logger(log.Ftl, pkg, "Unable to create the image directory: "+err.Error())
		}
	}
}

func (s *settings) urlHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res := new(response)
	log.Logger(log.Info, pkg, "Received connection from "+r.RemoteAddr)
	log.Logger(log.Info, pkg, "Requested URL: "+r.URL.String())
	if !checkAccept(r) {
		res.Error = "must use Accept: application/json"
		log.Logger(log.Err, pkg, "Accept: application/json not set. No data sent")
	} else {
		found := false
		var pathType settingsloader.Urls
		for _, url := range s.Data.URLs {
			if url.URL == ps.ByName("path") {
				found = true
				pathType = url
			}
		}
		count++
		if found {
			switch pathType.DataType {
			case "bin/png":
				dataBin, err := imgbin.CreateImgBin(count, s.Data.FontPath, s.Data.ImgDir)
				if err != nil {
					res.Error = "Unable to create image binary"
				}
				res.Data = dataBin
				res.DataType = pathType.DataType
				log.Logger(log.Info, pkg, "Sent "+pathType.DataType)
			case "data/png":
				dataURI, err := imguri.CreatImgData(count, s.Data.FontPath)
				if err != nil {
					res.Error = "Unable to generate data URI"
				}
				res.Data = "data:image/png;base64," + dataURI

				res.DataType = pathType.DataType
				log.Logger(log.Info, pkg, "Sent "+pathType.DataType)
			default:
				res.Error = "404: Requested URL not found"
				log.Logger(log.Info, pkg, "Requested URL not found")
			}

		}
	}
	res.TimeStamp = time.Now().String()
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Logger(log.Err, pkg, "Unable to send JSON: "+err.Error())
		return
	}
	log.Logger(log.Info, pkg, "Sent")
}

func main() {
	log.LogInit()
	log.Logger(log.Info, pkg, "Loading Settings...")
	s := new(settings)
	s.Data = settingsloader.NewSettings(os.Args[1])
	if !checkFilePath(s.Data.FontPath) {
		log.Logger(log.Ftl, pkg, "Font file does not exist: "+s.Data.FontPath)
	}
	log.Logger(log.Info, pkg, "Checking for image directory")
	checkFilePath(s.Data.ImgDir)
	log.Logger(log.Info, pkg, "Starting server on "+s.Data.Host+":"+s.Data.Port)
	router := httprouter.New()
	router.GET("/", indexHandler)
	router.ServeFiles("/images/*filepath", http.Dir(s.Data.ImgDir))
	router.GET(s.Data.RootURL+"/:path", s.urlHandler)
	if err := http.ListenAndServe(s.Data.Host+":"+s.Data.Port, router); err != nil {
		log.Logger(log.Ftl, pkg, "Unable to start server. Error: "+err.Error())
	}
}
