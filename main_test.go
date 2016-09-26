package main

import (
	"net/http"
	"testing"

	. "github.com/dtoebe/gophx-img-api/testerror"
)

func TestCheckAccept(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/something", nil)
	if err != nil {
		ErrorOut("CheckAccept", "", err.Error(), t)
	}
	req.Header.Add("Accept", "application/json")
	if !checkAccept(req) {
		ErrorOut("checkAccept", "true", "false", t)
	}
}
