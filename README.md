### Snake over Telnet in Go

Try it:

```
telnet pliutau.com 8080
```


### Development

```
go get github.com/plutov/go-snake-telnet
go-snake-telnet --host localhost --port 8080
```

### Tests

```
go test ./... -bench=.
```
