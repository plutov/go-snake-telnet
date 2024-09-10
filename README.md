### Snake Game over Telnet in Go

### Run it with go

```bash
go get github.com/plutov/go-snake-telnet
go-snake-telnet
```

## Run with Docker

```bash
docker build -t snake-telnet .
docker run -d -p 8080:8080 snake-telnet
```

## Play!

Make sure to install telnet first:

```bash
brew install telnet
```

Then connect to the game:
```bash
telnet 127.0.0.1 8080
```

### Tests

```
go test ./... -bench=.
```
