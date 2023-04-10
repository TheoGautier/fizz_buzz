package server

import (
	"encoding/json"
	"net/http"
	"time"
)

type Server struct {
	Port  string
	Stats map[string]map[string]int
}

func New(port string) *Server {
	s := Server{Port: port}
	s.Stats = make(map[string]map[string]int)
	return &s
}

func (s *Server) Start() {
	for _, handler := range s.GetHandlers() {
		s.Stats[handler.Route] = make(map[string]int)
		http.Handle(handler.Route, http.TimeoutHandler(s.StatMiddleware(handler.Route, handler.HandlerFunc), time.Second, ""))
	}
	if err := http.ListenAndServe(":"+s.Port, nil); err != nil {
		panic(err)
	}
}

func (s *Server) WriteResponseJson(statusCode int, v any, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(statusCode)
	data, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
