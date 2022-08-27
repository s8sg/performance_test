package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.PUT("/mockserver/status", syncHandler)
	router.Run(":8080")
}

func syncHandler(c *gin.Context) {
	client := &http.Client{}

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

func asyncHandler(c *gin.Context) {
	client := &http.Client{}

	responseChan := make(chan *http.Response, 0)

	go func() {
		req, err := http.NewRequest(http.MethodPut, "http://mock-server:1080/mockserver/status", bytes.NewBuffer([]byte{}))
		if err != nil {
			panic(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		responseChan <- resp
	}()

	resp := <-responseChan
	reader := resp.Body
	defer reader.Close()
	contentLength := resp.ContentLength
	contentType := resp.Header.Get("Content-Type")
	extraHeaders := map[string]string{}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}
