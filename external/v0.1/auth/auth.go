package auth

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Auth struct {
	Proxy *httputil.ReverseProxy
}

func NewAuthProxy(urlOrigin string) *Auth {
	url, err := url.Parse(urlOrigin)
	if err != nil {
		panic(err)
	}
	prx := &Auth{Proxy: httputil.NewSingleHostReverseProxy(url)}
	originalDirector := prx.Proxy.Director
	// modify request
	prx.Proxy.Director = func(r *http.Request) {

		originalDirector(r)
		r.Header.Set("Xforwareded-From", "entrnace")
	}
	// Modify response
	prx.Proxy.ModifyResponse = func(r *http.Response) error {
		// Add a response header
		r.Header.Set("Server", "auth")
		return nil
	}
	return prx
}

func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// here i Write things that are for Auth only
	log.Print("login from auth")
	a.Proxy.ServeHTTP(w, r)

}
