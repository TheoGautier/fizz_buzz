#!/bin/bash

function build() {
  docker build -t "fizzbuzz":latest -f ./Dockerfile .
}

function run() {
  docker-compose up fizzbuzz
}

function test() {
  go test -v ./... && \
  go test -v -fuzz FuzzFizzBuzz -fuzztime=5s ./fizzbuzz
}