package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var client *http.Client

func main() {
	router := gin.Default()
	client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 5000,
		},
	}
	router.PUT("/mockserver/status", asyncHandler)
	router.Run(":8080")

}

func syncHandler(c *gin.Context) {
	req, err := http.NewRequest(http.MethodPut, "http://mock-server:1080/mockserver/status", bytes.NewBuffer([]byte{}))
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	reader := resp.Body
	defer reader.Close()

	contentLength := resp.ContentLength
	contentType := resp.Header.Get("Content-Type")
	extraHeaders := map[string]string{}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

type AsyncResponse struct {
	body        []byte
	contentType string
}

func asyncHandler(c *gin.Context) {

	responseChan := make(chan *AsyncResponse, 0)

	go func() {
		req, err := http.NewRequest(http.MethodPut, "http://mock-server:1080/mockserver/status", bytes.NewBuffer([]byte{}))
		if err != nil {
			panic(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		contentType := resp.Header.Get("Content-Type")
		body, _ := ioutil.ReadAll(resp.Body)

		resp.Body.Close()

		responseChan <- &AsyncResponse{
			body:        body,
			contentType: contentType,
		}
	}()

	asyncResponse := <-responseChan

	c.Data(http.StatusOK, asyncResponse.contentType, asyncResponse.body)
}
