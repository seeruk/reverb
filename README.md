# Reverb

Reverb is a simple HTTP server for testing / debugging / logging requests from some service. It 
stores a request details in memory, and provides an interface to retrieve them later to inspect 
them.

## Usage

Reverb is go-gettable, and super easy to run and use.

```
$ go get -u github.com/SeerUK/reverb/...
$ reverb
$ # Or if you want to configure it a bit more:
$ reverb -addr 0.0.0.0 -port 8080
```

@TODO

## License

MIT
