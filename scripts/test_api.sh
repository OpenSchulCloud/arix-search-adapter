#!/bin/bash
set -e

go build github.com/schul-cloud/arix-search-adapter/search
./search &
PID="$!"

python3 -m schul_cloud_search_tests.search \
        --url=http://localhost:8080/v1/search

kill "$PID"
