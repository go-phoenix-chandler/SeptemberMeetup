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

//count int is the amount of hits made on the api
var count = 0

//empty and pkg are at the top of the packages in this application
//to make the package name a string for the logger
type empty struct{}

var pkg = reflect.TypeOf(empty{}).PkgPath()

//settings simple struct that holds loaded setings from the settingsload pkg
type settings struct {
	Data *settingsloader.Settings
}

//response is the structure of every API response
type response struct {
	DataType  string `json:"data-type"`
	Data      string `json:"data"`
	Error     string `json:"error"`
	TimeStamp string `json:"time-stamp"`
}

//checkAccept takes the request and outputs a bool based on whether or not `Accept`
//is set to `application/json` in the requesters sent header
func checkAccept(r *http.Request) bool {
	if strings.TrimSpace(r.Header.Get("Accept")) == "application/json" {
		return true
	}
	return false
}

//checkFilePath takes a path to a filr as a string and returns a bool.
//True if it exists, false if it does not
func checkFilePath(p string) bool {
	if _, err := os.Stat(p); err == nil {
		return true
	}
	return false
}

//createImageDir takes a path to a directory as a string and checks is the directory exists.
//If the directory does not exist it creates it.
func createImageDir(p string) {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		// I know.. I know 777 is bad!
		if err := os.Mkdir(p, 0777); err != nil {
			log.Logger(log.Ftl, pkg, "Unable to create the image directory: "+err.Error())
		}
	}
}

//urlHandler is a method of the settings struct.  this is an http handler func takes a dynamic url
//based off of the RootURL, and []URLs from the settings
func (s *settings) urlHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//create the empty response
	res := new(response)
	log.Logger(log.Info, pkg, "Received connection from "+r.RemoteAddr)
	log.Logger(log.Info, pkg, "Requested URL: "+r.URL.String())
	//Check if header contains the Accept: application/json.
	//if not add error to response.  If so then build the data
	if !checkAccept(r) {
		res.Error = "must use Accept: application/json"
		log.Logger(log.Err, pkg, "Accept: application/json not set. No data sent")
	} else {
		found := false
		var pathType settingsloader.Urls
		//Here we try to match the url fomr the slice of URLs in the settings file
		for _, url := range s.Data.URLs {
			if url.URL == ps.ByName("path") {
				found = true
				pathType = url
			}
		}
		count++
		//If URL is found in the settings file
		if found {
			//Switch on the URL DataType so multiple urls can go to the same DataType if needed
			switch pathType.DataType {
			case "bin/png":
				//Here we create the image in binary format and save it the setings ImgDir
				// there we have a static fileserver watching to load the images
				dataBin, err := imgbin.CreateImgBin(count, s.Data.FontPath, s.Data.ImgDir)
				if err != nil {
					//If failed to create the binary in any way send an error though the API instead
					res.Error = "Unable to create image binary"
				}
				url := "http://" + s.Data.Host
				if len(s.Data.Port) > 0 {
					url += ":" + s.Data.Port
				}
				url += "/images/" + dataBin
				res.Data = url
				res.DataType = pathType.DataType
				log.Logger(log.Info, pkg, "Sent "+pathType.DataType)
			case "data/png":
				//Here we create the dataURI string.  this string does not get saved and is up to the
				//client to retain the data
				dataURI, err := imguri.CreatImgData(count, s.Data.FontPath)
				if err != nil {
					//If failed to create the data string the send error message over API
					res.Error = "Unable to generate data URI"
				}
				res.Data = "data:image/png;base64," + dataURI

				res.DataType = pathType.DataType
				log.Logger(log.Info, pkg, "Sent "+pathType.DataType)
			default:
				//If for some reason the DataType is not listed in a case send an error through the api
				res.Error = "404: Requested URL not found"
				log.Logger(log.Info, pkg, "Requested URL not found: "+ps.ByName("path"))
			}

		} else {
			//If the URL is not found send an error message throught the API
			res.Error = "404: Requested URL not found"
			log.Logger(log.Info, pkg, "Requested URL not found: "+ps.ByName("path"))
		}

	}
	res.TimeStamp = time.Now().String()
	//Encode response to json and send it
	if err := json.NewEncoder(w).Encode(res); err != nil {
		//If sending for some reason fails send a status
		w.WriteHeader(http.StatusInternalServerError)
		log.Logger(log.Err, pkg, "Unable to send JSON: "+err.Error())
		return
	}
	log.Logger(log.Info, pkg, "Sent")
}

func main() {
	//Initialize the logger wrapper
	log.LogInit()
	log.Logger(log.Info, pkg, "Loading Settings...")
	s := new(settings)
	if len(os.Args) < 2 {
		log.Logger(log.Ftl, pkg, "Please set absolute path to server-settings.json file")
	}
	s.Data = settingsloader.NewSettings(os.Args[1])
	//Check if Font File exists if not fatal out
	if !checkFilePath(s.Data.FontPath) {
		log.Logger(log.Ftl, pkg, "Font file does not exist: "+s.Data.FontPath)
	}
	log.Logger(log.Info, pkg, "Checking for image directory")
	//Check and/or create the Image Path
	checkFilePath(s.Data.ImgDir)
	log.Logger(log.Info, pkg, "Starting server on "+s.Data.Host+":"+s.Data.Port)
	router := httprouter.New()
	//Serve Image Directory as static files
	router.ServeFiles("/images/*filepath", http.Dir(s.Data.ImgDir))
	//serve the API
	router.GET(s.Data.RootURL+"/:path", s.urlHandler)
	if err := http.ListenAndServe(s.Data.Host+":"+s.Data.Port, router); err != nil {
		log.Logger(log.Ftl, pkg, "Unable to start server. Error: "+err.Error())
	}
}
