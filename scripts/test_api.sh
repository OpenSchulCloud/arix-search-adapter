#!/bin/bash
set -e

cd "`dirname \"$0\"`"

go build github.com/schul-cloud/arix-search-adapter
./arix-search-adapter &
PID="$!"

python -m schul_cloud_search_tests.search \
       --url=http://localhost:8080/v1/search

kill "$PID"
