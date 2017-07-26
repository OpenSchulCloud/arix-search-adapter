ARIX-search-adapter
===================

[![Build Status](https://travis-ci.org/schul-cloud/arix-search-adapter.svg?branch=master)](https://travis-ci.org/schul-cloud/arix-search-adapter)
[![Dockerhub Automated Build Status](https://img.shields.io/docker/build/schulcloud/arix-search-adapter.svg)](https://hub.docker.com/r/schulcloud/arix-search-adapter/builds/)

This is a search adapter for the content search of Antares.

Setup
-----

Install Go.
- [Windows][setup-windows]
- [Ubuntu][setup-ubuntu]

Testing
- https://github.com/stretchr/testify
```
go get github.com/stretchr/testify/assert
```

This code
```
go get github.com/schul-cloud/arix-search-adapter
```

Test this code
```
go test github.com/schul-cloud/arix-search-adapter/arix
```

Run the search engine server
```
go build github.com/schul-cloud/arix-search-adapter/search && ./search
```

Try it out
----------

After Setup, you can run this command to request a search:

```
curl -i 'http://localhost:8080/v1/search?Q=Einstein'
```

Configuration
-------------

You can configure the server by setting environment variables.
Here you see the environment variables with their explanation:

- `ARIX_SEARCH_SERVER` defaults to `http://arix.datenbank-bildungsmedien.net/`  
  This is the URL the server uses for requests. It is expected to find an ARIX compatible endpoint there.
- `ARIX_SEARCH_CONTEXT` defaults to `HH`  
  This is the search context. This influences wich search results can be found.
- `ARIX_SEARCH_SECRET` defaults to `` (not set)  
  This is the secret to verify that the user has the license for the material which is requested.
- `ARIX_SEARCH_PORT` defaults to `8080`  
  This is the port the search server listens on.
- `ARIX_SEARCH_LIMIT` defaults to `10`  
  This is the maximum number of resources to request.
- `ARIX_SEARCH_SERVER_ID` defaults to `ARIX`  
  This is the id of the arix server. This can be used if you use mutiple servers to identify the resources.
- `ARIX_SEARCH_LINK_TYPE` defaults to `direct`  
  This is the type of the resource link. These options are available:
  - `direct` - Links show a media player or the resource in the users browser.
  - `download` - Links make the browser download the content.

[setup-windows]: http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/
[setup-ubuntu]: https://wiki.ubuntu.com/Go
