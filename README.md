### Snake over Telnet in Go [![Build Status](https://travis-ci.org/plutov/go-snake-telnet.svg?branch=master)](https://travis-ci.org/plutov/go-snake-telnet)


### Run it locally

```
go get github.com/plutov/go-snake-telnet
go-snake-telnet --host localhost --port 8080

# docker
docker build -t go-snake-telnet .
docker run -d -p $PORT:$PORT go-snake-telnet --host $HOST --port $PORT
```

```
telnet localhost 8080
```

### Tests

```
go test ./... -bench=.
```
