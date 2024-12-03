package test

import (
	"github.com/CorentinPtrl/evengsdk"
	"os"
	"testing"
)

func TestClient_GetAuth(t *testing.T) {
	client, err := evengsdk.NewBasicAuthClient(os.Getenv("EVE_USER"), os.Getenv("EVE_PASSWORD"), "0", os.Getenv("EVE_HOST"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.GetAuth()
	if err != nil {
		t.Fatal(err)
	}
}
