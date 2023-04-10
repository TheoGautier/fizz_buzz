package server

import (
	"net/http"
	"net/url"
)

type Handler struct {
	Route       string
	HandlerFunc http.HandlerFunc
}

type BadRequestErrorResponse struct {
	Message string `json:"message"`
}

type InternalServerErrorResponse struct {
	Message string `json:"message"`
}

func (s *Server) GetHandlers() []Handler {
	return []Handler{
		{Route: "/health", HandlerFunc: s.GetHealthHandler},
		{Route: "/fizzbuzz", HandlerFunc: s.GetFizzBuzzHandler},
		{Route: "/stats", HandlerFunc: s.GetStatsHandler},
	}
}

func (s *Server) GetHealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Server) StatMiddleware(route string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			s.WriteResponseJson(http.StatusBadRequest, BadRequestErrorResponse{Message: "Could not read request"}, w)
		}
		params := values.Encode() //params are sorted by key
		if s.Stats[route] == nil {
			s.Stats[route] = make(map[string]int)
		}
		s.Stats[route][params] += 1
		next.ServeHTTP(w, r)
	}
}
