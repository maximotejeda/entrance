package proxy

import (
	"fmt"
	"net/http"
	"regexp"
)

func NewMultiplexer() *Multiplexer {
	return &Multiplexer{
		handler: make(map[string]http.Handler),
		cache:   make(map[string]*regexp.Regexp),
	}
}

type Multiplexer struct {
	handler map[string]http.Handler
	cache   map[string]*regexp.Regexp
}

func (m *Multiplexer) Add(regex string, handler http.Handler) {
	m.handler[regex] = handler
	cache := regexp.MustCompile(regex)
	m.cache[regex] = cache
}

func (m *Multiplexer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	check := fmt.Sprintf("%s %s", r.Method, r.URL.Path)
	for pattern, handler := range m.handler {
		if m.cache[pattern].MatchString(check) {
			handler.ServeHTTP(w, r)
			return
		}

	}

	http.NotFound(w, r)
}
