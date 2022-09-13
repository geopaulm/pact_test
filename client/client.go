package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/geo/pacttest/models"
)

// Simple client fetching user details
type Client struct {
	BaseUrl string
}

func (client Client) FetchUserDetails(id string) (*models.User, error) {
	url := fmt.Sprintf("%s/users/%s", client.BaseUrl, id)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
