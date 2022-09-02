package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// parse errors
var (
	ErrEventNotSpecifiedToParse  = errors.New("no Event specified to parse")
	ErrInvalidHTTPMethod         = errors.New("invalid HTTP Method")
	ErrMissingGithubEventHeader  = errors.New("missing X-GitHub-Event Header")
	ErrMissingHubSignatureHeader = errors.New("missing X-Hub-Signature Header")
	ErrEventNotFound             = errors.New("event not defined to be parsed")
	ErrParsingPayload            = errors.New("error parsing payload")
	ErrHMACVerificationFailed    = errors.New("HMAC verification failed")
)

func main() {
	http.HandleFunc("/", GetSignature)

	fmt.Println("Server starting on port 9090....")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}

func GetSignature(w http.ResponseWriter, r *http.Request) {
	// Parse payload
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		log.Printf("error occured: %v", ErrParsingPayload)
		return
	}

	// set webhook secret
	secret := "somevalue"
	signature := r.Header.Get("X-Hub-Signature")
	if len(signature) == 0 {
		log.Printf("error occured: %v", ErrMissingHubSignatureHeader)
		return
	}
	log.Printf("Actual signature: %s", signature)

	mac := hmac.New(sha1.New, []byte(secret))
	_, _ = mac.Write(payload)
	expectedMAC := "sha1=" + hex.EncodeToString(mac.Sum(nil))
	log.Printf("Expected signature: %s", expectedMAC)

	// Verify secret
	if !hmac.Equal([]byte(signature), []byte(expectedMAC)) {
		log.Printf("error occured: %v", ErrHMACVerificationFailed)
		return
	}

	log.Println("👏🎉Congratulations!!!!Signature matched")
}
