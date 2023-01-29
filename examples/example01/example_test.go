package main

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestMain(t *testing.T) {
	go main()

	waitCounter := 0

	var resp *http.Response
	var err error

	for waitCounter < 10 {
		resp, err = http.Get("http://localhost:8080/ping")
		if err != nil {
			waitCounter += 1
			time.Sleep(100 * time.Millisecond)
		} else {
			waitCounter = 11
		}
	}

	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}

	var respObj gin.H

	err = json.Unmarshal(b, &respObj)
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}

	if msg, ok := respObj["message"]; ok {
		if msg != "pong" {
			t.Errorf("message should be pong, but is %s", msg)
			t.FailNow()
		}
	} else {
		t.Errorf("missing message")
		t.FailNow()
	}

}
