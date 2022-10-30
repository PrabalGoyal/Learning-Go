package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type IDService interface {
	GenerateID() (*ID, error)
	GenerateBulkID(count uint) (*[]ID, error)
}

type (
	ID           string
	UUID4Service struct{}
)

func NewUUID4Service() IDService {
	return &UUID4Service{}
}

func (u UUID4Service) GenerateBulkID(count uint) (*[]ID, error) {
	url := fmt.Sprintf("https://www.uuidtools.com/api/generate/v4/count/%d", count)
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
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var uuids []ID
	err = json.Unmarshal(body, &uuids)
	if err != nil {
		return nil, err
	}
	return &uuids, nil
}

func (u UUID4Service) GenerateID() (*ID, error) {
	uuid4s, err := u.GenerateBulkID(1)
	if err != nil {
		return nil, err
	}
	return &((*uuid4s)[0]), nil
}

func main() {
	uuid4Service := NewUUID4Service()
	uuid4, err := uuid4Service.GenerateID()
	if err != nil {
		panic(err)
	}
	fmt.Println(*uuid4)
}
