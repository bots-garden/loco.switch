#!/bin/bash
# -n 10000 -c 1000
#hey -n 10 -c 5 -m POST \
hey -n 10000 -c 1000 -m POST \
-H "Content-Type: text/plain; charset=utf-8" \
-d "John Doe" \
"http://localhost:8080/functions/hello"
