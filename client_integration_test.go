package main

import (
	"testing"
)

func TestClient(t *testing.T) {
	id := "1"
	client := &Client{
		BaseUrl: "http://localhost:8080",
	}

	user, err := client.FetchUserDetails(id)
	if err != nil {
		t.Errorf("No error is expected while fetching user %s", err.Error())
	}

	if user.Id != id {
		t.Fatalf("id mismatch on fetch")
	}
}
