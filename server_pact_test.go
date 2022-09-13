package main

import (
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestServerPact_Verification(t *testing.T) {
	go StartServer()

	// initialize PACT DSL
	pact := dsl.Pact{
		Provider: "example-server",
	}

	// verify Contract on server side
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://127.0.0.1:8080",
		PactURLs:        []string{"pacts/example-client-example-server.json"},
	})

	if err != nil {
		t.Log(err)
	}
}
