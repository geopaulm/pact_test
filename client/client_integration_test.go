package client

import (
	"fmt"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	host := os.Getenv("server_host")
	if host == "" {
		host = "localhost"
	}
	id := "1"
	client := &Client{
		BaseUrl: fmt.Sprintf("http://%s:8080", host),
	}

	user, err := client.FetchUserDetails(id)
	if err != nil {
		t.Errorf("No error is expected while fetching user %s", err.Error())
	}

	if user.Id != id {
		t.Fatalf("id mismatch on fetch")
	}
}
