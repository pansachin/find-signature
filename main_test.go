package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestGetSignature(t *testing.T) {
	r := &http.Request{}
	// Read a file
	data, err := os.ReadFile("payload.json")
	if err != nil {
		panic(err)
	}
	// Convert byte to reader closer

	dataread := bytes.NewReader(data)
	r.Body = io.NopCloser(dataread)
	r.Header = map[string][]string{
		"X-Hub-Signature": {"sha1=ea546bc4351f4b8ac5fe478f0b327615816a48a0"},
	}

	expectedhmac, err := ForgeSignature(nil, r)
	if err != nil {
		panic(err)
	}
	fmt.Println("Expected HMAC: ", "sha1="+expectedhmac)
}
