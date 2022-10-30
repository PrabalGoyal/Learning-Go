package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	pipeReader, pipeWriter := io.Pipe()
	buf := new(bytes.Buffer)

	go func() {
		defer pipeWriter.Close()
		normalString := strings.NewReader("String reader\n")
		file, err := os.Open("./sample.txt")
		if err != nil {
			panic(err)
		}
		body, err := HTTPRequest()
		if err != nil {
			panic(err)
		}
		// pass data from multiple readers to multiple writers
		_, err = io.Copy(io.MultiWriter(pipeWriter, buf, os.Stdout), io.MultiReader(normalString, file, body))
		if err != nil {
			panic(err)
		}
	}()

	body, err := io.ReadAll(pipeReader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	fmt.Println(buf)
}

func HTTPRequest() (io.ReadCloser, error) {
	url := fmt.Sprintf("https://www.uuidtools.com/api/generate/v4/count/%d", 1)
	method := "GET"

	client := &http.Client{Timeout: 1 * time.Second}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
