package server

import (
	"errors"
	"fizz_buzz/fizzbuzz"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func (s *Server) GetFizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		s.WriteResponseJson(http.StatusBadRequest, BadRequestErrorResponse{Message: "Could not read request"}, w)
		return
	}
	int1, err := strconv.Atoi(values.Get("int1"))
	if err != nil {
		s.WriteResponseJson(http.StatusBadRequest, BadRequestErrorResponse{Message: fmt.Sprintf("Could not read request int1, given is %s", values.Get("int1"))}, w)
		return
	}
	int2, err := strconv.Atoi(values.Get("int2"))
	if err != nil {
		s.WriteResponseJson(http.StatusBadRequest, BadRequestErrorResponse{Message: fmt.Sprintf("Could not read request int2, given is %s", values.Get("int2"))}, w)
		return
	}
	limit, err := strconv.Atoi(values.Get("limit"))
	if err != nil {
		s.WriteResponseJson(http.StatusBadRequest, BadRequestErrorResponse{Message: fmt.Sprintf("Could not read request limit, given is %s", values.Get("limit"))}, w)
		return
	}
	str1 := values.Get("str1")
	str2 := values.Get("str2")

	output, err := fizzbuzz.FizzBuzz(int1, int2, limit, str1, str2)
	if errors.Is(err, fizzbuzz.ErrZeroLimit) || errors.Is(err, fizzbuzz.ErrLimitNegative) {
		s.WriteResponseJson(http.StatusBadRequest, BadRequestErrorResponse{Message: fmt.Sprintf("Could not execute fizzbuzz, err: %s", err.Error())}, w)
		return
	} else if err != nil {
		s.WriteResponseJson(http.StatusInternalServerError, InternalServerErrorResponse{Message: fmt.Sprintf("Could not read request limit, given is %s", values.Get("limit"))}, w)
		return
	}
	s.WriteResponseJson(http.StatusOK, output, w)
}
