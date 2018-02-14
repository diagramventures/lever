Lever API library for Go
=========================

[![GoDoc](https://godoc.org/github.com/diagramventures/lever?status.svg)](https://godoc.org/github.com/diagramventures/lever)

This library aims to provide simple access to data structures and API
calls to [Lever](https://www.lever.co).

Basic usage
-----------

```go
api := lever.New("apiKey...")

candidates, _ := api.ListCandidates()
```

Contributing
------------

Any contributions are welcome, use your standard GitHub-fu to pitch in and improve.


License
-------

MIT