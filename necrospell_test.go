package main

import (
	"net/http"
	"testing"
)

func TestTheHttp ( t *  testing.T ){

	req, err := http.NewRequest("GET", "/ssh", nil)
	if err != nil{
		t.Error( "Whatever")
	}
	if req.Response.StatusCode != 200 {
		t.Error("Not 200")
	}

}
