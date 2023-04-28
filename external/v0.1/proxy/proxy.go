package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type SimpleProxy struct {
	Proxy *httputil.ReverseProxy
}

// NewProxy: return a new single host proxy
// https://www.codedodle.com/go-reverse-proxy-example.html
func NewProxy(rawURL string) (*SimpleProxy, error) {
	url, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	prx := &SimpleProxy{Proxy: httputil.NewSingleHostReverseProxy(url)}
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
	return prx, nil
}

func (s *SimpleProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Do anything you want here
	// e.g. blacklisting IP, log time, modify headers, etc
	//log.Printf("Proxy receives request.")
	//log.Printf("Proxy forwards request to origin.")
	//	s.Proxy.ServeHTTP(w, r)
	//log.Printf("Origin server completes request.")
}
