ARIX-search-adapter
===================

[![Build Status](https://travis-ci.org/schul-cloud/arix-search-adapter.svg?branch=master)](https://travis-ci.org/schul-cloud/arix-search-adapter)


Setup
-----

Install Go.
- [Windows][setup-tut]

Testing
- https://github.com/stretchr/testify
```
go get github.com/stretchr/testify/assert
```

This code
```
go get go test github.com/schul-cloud/arix-search-adapter
```

Test this code
```
go test github.com/schul-cloud/arix-search-adapter/arix
```

Run the search engine server
```
go build github.com/schul-cloud/arix-search-adapter/search && search
```


[setup-tut]: http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/
