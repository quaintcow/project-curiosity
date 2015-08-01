package gotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/simplq/handlers"
)

//const addr = "http://127.0.0.1:2/"

var addr = "" //"http://::1/"

func TestIndex(t *testing.T) {
	ts, h := helpPrepServer()
	defer ts.Close()
	helpCreateRequest(t, h, "GET", "")

}
func TestGetLogin(t *testing.T) {
	ts, h := helpPrepServer()
	defer ts.Close()
	helpCreateRequest(t, h, "GET", "login")
}
func TestPostLogin(t *testing.T) {
	ts, h := helpPrepServer()
	defer ts.Close()
	helpPostLogin(t, h, "login")
}

// ----- Helper Functions ------ //
func helpPostLogin(t *testing.T, h *handlers.Mux, path string) {
	res, err := http.PostForm(addr+path,
		url.Values{"email": {"zlaw777@gmail.com"}, "password": {"EEUE12345"}})
	if err != nil {
		log.Fatal(err)
	}
	bs, _ := ioutil.ReadAll(res.Body)
	tr := string(bs)
	fmt.Println("Response:", tr)
	//fmt.Println(req)
}
func helpCreateRequest(t *testing.T, h *handlers.Mux, method string, path string) {
	req, err := http.NewRequest(method, addr+path, nil)
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Fatal("Code is", w.Code)
	}
	fmt.Println(w.Body.String())
}
func helpPrepServer() (*httptest.Server, *handlers.Mux) {
	h := handlers.GetMux()
	server := httptest.NewServer(h)
	addr = server.URL + "/" // the server URL is generated by NewServer
	return server, h
}
