package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func GetTestHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test entered test handler, this should not happen")
	}
	return http.HandlerFunc(fn)
}

func TestRecoveryMW(t *testing.T) {
	handler := http.HandlerFunc(panicHandler)
	executeRequest("Get", "/panic", recoveryMw(handler))
}

func executeRequest(method string, url string, handler http.Handler) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	rr := httptest.NewRecorder()
	rr.Result()
	handler.ServeHTTP(rr, req)
	return rr, err
}
func TestDebugAPI(t *testing.T) {
	// Create server using the a router initialized elsewhere. The router
	// can be a Gorilla mux as in the question, a net/http ServeMux,
	// http.DefaultServeMux or any value that statisfies the net/http
	// Handler interface.
	ts := httptest.NewServer(gethandles())
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
		{name: "1: testing get", r: newreq("GET", ts.URL+"/debug?path=/home/mohit/go/src/github.com/MohitArora1/gophercises/recover/main.go&line=64", nil), status: 200},
		{name: "2: testing get", r: newreq("GET", ts.URL+"/debug?path=/home/mohit/go/src/github.com/MohitArora1/gophercises/recoer/main.go&line=64", nil), status: 500},
		{name: "2: testing get", r: newreq("GET", ts.URL+"/debug?path=/home/mohit/go/src/github.com/MohitArora1/gophercises/recover/main.go&line=et", nil), status: 500},
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

func TestMakeLinks(t *testing.T) {
	str := `
	goroutine 11 [running]:
runtime/debug.Stack(0xc42005db48, 0x1, 0x1)
	/usr/local/go/src/runtime/debug/stack.go:24 +0xa7
main.recoveryMw.func1.1(0x9ffc40, 0xc4200c0000)
	/home/mohit/go/src/github.com/MohitArora1/gophercises/recover/main.go:64 +0xac
panic(0x82fc80, 0x9f6f60)
	/usr/local/go/src/runtime/panic.go:502 +0x229
main.funcThatPanic()
	/home/mohit/go/src/github.com/MohitArora1/gophercises/recover/main.go:56 +0x39
main.panicHandler(0x9ffc40, 0xc4200c0000, 0xc420157a00)
	/home/mohit/go/src/github.com/MohitArora1/gophercises/recover/main.go:52 +0x20
net/http.HandlerFunc.ServeHTTP(0x92be60, 0x9ffc40, 0xc4200c0000, 0xc420157a00)
	/usr/local/go/src/net/http/server.go:1947 +0x44
net/http.(*ServeMux).ServeHTTP(0xc4202101e0, 0x9ffc40, 0xc4200c0000, 0xc420157a00)
	/usr/local/go/src/net/http/server.go:2337 +0x130
main.recoveryMw.func1(0x9ffc40, 0xc4200c0000, 0xc420157a00)
	/home/mohit/go/src/github.com/MohitArora1/gophercises/recover/main.go:70 +0x95
net/http.HandlerFunc.ServeHTTP(0xc420421520, 0x9ffc40, 0xc4200c0000, 0xc420157a00)
	/usr/local/go/src/net/http/server.go:1947 +0x44
net/http.serverHandler.ServeHTTP(0xc4203c5ee0, 0x9ffc40, 0xc4200c0000, 0xc420157a00)
	/usr/local/go/src/net/http/server.go:2694 +0xbc
net/http.(*conn).serve(0xc4200a4140, 0x9fffc0, 0xc420200000)
	/usr/local/go/src/net/http/server.go:1830 +0x651
created by net/http.(*Server).Serve
	/usr/local/go/src/net/http/server.go:2795 +0x27b`

	makeLinks(str)
}

func TestPanic(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8000/panic", nil)
	if err != nil {
		t.Fatalf("not able to request %v", err)
	}
	rec := httptest.NewRecorder()
	defer func() {
		if err := recover(); err != nil {

		}
	}()
	panicHandler(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("not expected error in panic %v", res.StatusCode)
	}
}

func TestMainFunc(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Main panicked ??")
		}
	}()

	go main()
	time.Sleep(1 * time.Second)
}
