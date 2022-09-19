package transport

import (
	"container/list"
	"errors"
	"net/http"
	"strings"
)

var (
	ErrEmptyToken   = errors.New("invalid token")
	ErrInvalidToken = errors.New("invalid token")
)

type MiddlewareType func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))

type MiddlewareMux struct {
	http.ServeMux
	middlewares   list.List
	authTransport AuthInterface
}

func NewMiddlewareMux(authTransport AuthInterface) *MiddlewareMux {
	return &MiddlewareMux{authTransport: authTransport}
}

func (mux *MiddlewareMux) AppendMiddleware(middleware func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request))) {
	// Append middleware to the end
	mux.middlewares.PushBack(MiddlewareType(middleware))
}

func (mux *MiddlewareMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	mux.nextMiddleware(mux.middlewares.Front())(w, req)
}

func (mux *MiddlewareMux) nextMiddleware(el *list.Element) func(w http.ResponseWriter, req *http.Request) {
	if el != nil {
		return func(w http.ResponseWriter, req *http.Request) {
			el.Value.(MiddlewareType)(w, req, mux.nextMiddleware(el.Next()))
		}
	}
	return mux.ServeMux.ServeHTTP
}

func (mux *MiddlewareMux) authMiddleware(w http.ResponseWriter, req *http.Request, next func(http.ResponseWriter, *http.Request)) {
	if strings.Contains(req.RequestURI, "/v1/auth/") {
		next(w, req)
		return
	}
	token, err := parseToken(req)
	if err != nil {
		NewResponse(w, http.StatusUnauthorized, err.Error())
		return
	}
	if !mux.authTransport.CheckToken(token) {
		NewResponse(w, http.StatusUnauthorized, ErrInvalidToken.Error())
		return
	}
	next(w, req)
}

func parseToken(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", ErrEmptyToken
	}
	headerAuth := strings.Split(header, " ")
	if len(headerAuth) != 2 || headerAuth[0] != "Bearer" {
		return "", ErrInvalidToken
	}
	return headerAuth[1], nil
}
