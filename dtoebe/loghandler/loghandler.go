package loghandler

import (
	"fmt"
	"log"
	"os"

	"github.com/comail/colog"
)

//T int for enum
type T int

const (
	//Ftl Fatal
	Ftl T = 0 + iota
	//Err Error
	Err
	//Wrn Warning
	Wrn
	//Info Standard Message
	Info
)

var t = [...]string{
	"FATAL",
	"ERROR",
	"WARNING",
	"INFO",
}

//LogInit initializes the 3rd party colog package
func LogInit() {
	colog.Register()
}

//Logger is a simple wrapper to use the colog package
//Takes:
// - enum to set log type
// - package string
// = message string
//Also if FATAL then forces exit
func Logger(tp T, pkg, msg string) {
	info := fmt.Sprintf("%s: %s package=%s", t[tp], msg, pkg)
	log.Print(info)
	if tp == 0 {
		os.Exit(1)
	}
}
