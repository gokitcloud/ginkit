package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	go main()
}

func TestIsOK(t *testing.T) {

	tests := [][]string{
		{
			"a", "12345678",
		},
		{
			"b", "abcdefgh",
		},
	}

	for _, args := range tests {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/%s", args[0]), nil)
		if err != nil {
			t.Errorf(err.Error())
			t.FailNow()
		}

		req.Header.Set("X-token", args[1])

		respObj, res, err := doRequest(http.DefaultClient, req)
		if err != nil {
			t.Errorf(err.Error())
			t.FailNow()
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("status code should be %d, but is %d", http.StatusOK, res.StatusCode)
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

}

func TestUnauth(t *testing.T) {

	tests := [][]string{
		{
			"a", "abcdefgh",
		},
		{
			"b", "12345678",
		},
	}

	for _, args := range tests {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/%s", args[0]), nil)
		if err != nil {
			t.Errorf(err.Error())
			t.FailNow()
		}
		req.Header.Set("X-token", args[1])
		_, res, err := doRequest(http.DefaultClient, req)
		if err != nil {
			t.Errorf(err.Error())
			t.FailNow()
		}

		if res.StatusCode != http.StatusForbidden {
			t.Errorf("status code should be %d, but is %d", http.StatusForbidden, res.StatusCode)
			t.FailNow()
		}
	}
}

func doRequest(client *http.Client, req *http.Request) (respObj gin.H, res *http.Response, err error) {
	waitCounter := 0
	for waitCounter < 10 {
		res, err = client.Do(req.Clone(context.Background()))

		if err != nil {
			waitCounter += 1
			time.Sleep(100 * time.Millisecond)
		} else {
			waitCounter = 11
		}
	}

	if err != nil {
		return nil, res, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}

	err = json.Unmarshal(b, &respObj)
	if err != nil {
		return nil, res, err
	}

	return respObj, res, nil
}
