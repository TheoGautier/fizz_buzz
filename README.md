# fizz-buzz REST server

## Overview
Fizz_buzz is a Rest server implementing an API interface to the fizzbuzz package - an expansion of the classical
fizzbuzz to wider limits and outputs.

**Version used:** go 1.20.3.  
This project uses generics, and fuzz testing, so it requires **go > 1.18**. 

The project exposes three routes on the port 8080:
- `/health`: Writes 200 response - used to check if server is alive
- `/fizzbuzz`: Accepts five parameters as query parameters :
  - `int1, int2`: Integers replacing `3` and `5` in the classical fizzbuzz
  - `limit`: Integer to count display to (limit included).
  - `str1, str2`: Strings replacing `"fizz"` and `"buzz"` in the classical fizzbuzz
  Example query: `/fizzbuzz?int1=3&int2=5&limit=10000&str1=fizz&str2=buzz`
- `/stats`: Expose the highest used parameters on the `/fizzbuzz` route.

Note that for technical performance, routes timeout after 1min.

## Start project
You can source `build.sh` file to use testing/building/running helpers.
```shell
source build.sh
```

### Testing
Test project using `test` or the following:
```shell
go test -v ./... && \
go test -v -fuzz FuzzFizzBuzz -fuzztime=5s ./fizzbuzz
```

### Building
To build project in docker, use `build` or
```shell
docker build -t "fizzbuzz":latest -f ./Dockerfile .
```

### Running
To run project in a docker container, using docker-compose, use `run` or 
```shell
docker-compose up fizzbuzz
```
The project runs on PORT 8080. 