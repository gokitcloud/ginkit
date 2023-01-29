package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestMain(t *testing.T) {
	go main()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/ping", nil)
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}

	respObj, err := doRequest(http.DefaultClient, req)
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

func doRequest(client *http.Client, req *http.Request) (respObj gin.H, err error) {
	waitCounter := 0
	var resp *http.Response
	for waitCounter < 10 {
		resp, err = client.Do(req.Clone(context.Background()))

		if err != nil {
			waitCounter += 1
			time.Sleep(100 * time.Millisecond)
		} else {
			waitCounter = 11
		}
	}

	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}
