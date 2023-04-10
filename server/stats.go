package server

import (
	"fizz_buzz/stats"
	"net/http"
)

type StatsResponse struct {
	Params string `json:"params"`
	Hits   int    `json:"hits"`
}

func (s *Server) GetStatsHandler(w http.ResponseWriter, r *http.Request) {
	params, hits := stats.GetHighestHits(s.Stats["/fizzbuzz"])
	s.WriteResponseJson(http.StatusOK, StatsResponse{Params: params, Hits: hits}, w)
}
