package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Simple client fetching user details
type Client struct {
	BaseUrl string
}

func (client Client) FetchUserDetails(id string) (*User, error) {
	url := fmt.Sprintf("%s/users/%s", client.BaseUrl, id)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
