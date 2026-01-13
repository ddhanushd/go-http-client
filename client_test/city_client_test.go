package client_test

import (
	"net/http"
	"testing"

	"github.com/ddhanushd/go-http-client/client"
	"github.com/therewardstore/httpmatter"
)

func TestGetCity_Success(t *testing.T) {
	h := httpmatter.NewHTTP(t, "basic").
		Add("request_bangalore", "response_bangalore").
		Respond(nil)

	h.Init()
	defer h.Destroy()

	httpClient := &http.Client{}
	cityClient := client.NewCityClient(httpClient, "http://example.com")

	result, err := cityClient.GetCity("Bangalore")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != "City Data" {
		t.Errorf("expected %q, got %q", "City Data", result)
	}
}
