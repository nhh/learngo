package main

import (
	"strings"
	"testing"
)

const request = `GET / HTTP/1.0
Accept: text/plain
Accept-Language: en, de
`

const expectedResponse = `HTTP/1.0 200 OK
`
const contentLength = len("Hallo Welt")

func Test(t *testing.T) {

	resp := Handle(request)

	if resp != expectedResponse {
		t.Errorf("Expected %q, \n got %q", expectedResponse, resp)
	}

	if !strings.Contains(resp, "200") {
		t.Errorf("Expected %q, \n got %q", expectedResponse, resp)
	}

}
