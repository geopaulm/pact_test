package main

import (
	//
	// some imports
	//
	"errors"
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

func TestClientPact(t *testing.T) {
	pact := dsl.Pact{
		Consumer: "example-client",
		Provider: "example-server",
	}

	pact.Setup(true)

	t.Run("get user by id", func(t *testing.T) {
		id := "1"
		pact.AddInteraction().
			Given("User Geo exists.").
			UponReceiving("User Geo is requested").
			WithRequest(dsl.Request{
				Method: "GET",
				Path:   dsl.Term("users/1", "users/[0-9]+"),
			}).
			WillRespondWith(dsl.Response{
				Status: 200,

				Body: dsl.Like(User{
					Id:        id,
					FirstName: "Geo",
					LastName:  "Paul",
				}),
			})

		err := pact.Verify(func() error {
			client := &Client{
				BaseUrl: fmt.Sprintf("http://%s:%d", pact.Host, pact.Server.Port),
			}

			user, err := client.FetchUserDetails(id)
			if err != nil {
				return fmt.Errorf("No error is expected while fetching user %s", err.Error())
			}

			if user.Id != id {
				return errors.New("User id mismatch")
			}
			return nil
		})

		if err != nil {
			t.Fatalf(err.Error())
		}

		// write Contract into file
		if err := pact.WritePact(); err != nil {
			t.Fatal(err)
		}

		// stop PACT mock server
		pact.Teardown()
	})
}
