package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func getBody(t *testing.T) *bytes.Buffer {
	file, err := os.Open("./input.png")
	if err != nil {
		t.Error("error in opening file")
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		t.Error("error in copy")
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Error("error in copy")
	}
	err = writer.Close()
	if err != nil {
		t.Error("error in close writer")
	}
	return body
}
func TestAPI(t *testing.T) {
	ts := httptest.NewServer(getHandlers())
	defer ts.Close()

	newreq := func(method, url string, body io.Reader) *http.Request {
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		return r
	}

	tests := []struct {
		name   string
		r      *http.Request
		status int
	}{
		{name: "1: testing get", r: newreq("GET", ts.URL+"/", nil), status: 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tt.r)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.status {
				t.Error("error in debug api")
			}
		})
	}
}

func TestAPI1(t *testing.T) {
	ts := httptest.NewServer(getHandlers())
	defer ts.Close()

	newreq := func(method, url string) *http.Request {
		file, err := os.Open("./input.png")
		if err != nil {
			t.Error("error in opening file")
		}
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", file.Name())
		if err != nil {
			t.Error("error in copy")
		}
		_, err = io.Copy(part, file)
		if err != nil {
			t.Error("error in copy")
		}
		err = writer.Close()
		if err != nil {
			t.Error("error in close writer")
		}
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		r.Header.Set("Content-Type", writer.FormDataContentType())
		return r
	}

	tests := []struct {
		name   string
		r      *http.Request
		status int
	}{
		{name: "1: testing get", r: newreq("POST", ts.URL+"/upload"), status: 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tt.r)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.status {
				t.Error("error in debug api")
			}
		})
	}
}

func TestAPI2(t *testing.T) {
	ts := httptest.NewServer(getHandlers())
	defer ts.Close()

	newreq := func(method, url string, body io.Reader) *http.Request {
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		return r
	}

	tests := []struct {
		name   string
		r      *http.Request
		status int
	}{
		{name: "1: testing get", r: newreq("GET", ts.URL+"/modify/input.png?mode=2", nil), status: 200},
		{name: "3: testing get", r: newreq("GET", ts.URL+"/modify/input.png?mode=2&number=100", nil), status: 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tt.r)
			if err != nil {
				t.Error("error in response")
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.status {
				t.Error("error in debug api")
			}
		})
	}
}
