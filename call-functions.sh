#!/bin/bash
#curl -X POST -d 'John Doe' http://localhost:8080/functions/hello_go
#echo ""
#curl -X POST -d 'Bob Morane' http://localhost:8080/functions/hello_js/default
#echo ""
#curl -X POST -d 'Jane Doe' http://localhost:8080/functions/hello_js/purple
#echo ""
#curl -X POST -d 'Jane Doe' http://localhost:8080/functions/hello_js/orange
#echo ""
#curl -X POST -d 'BatMan 🦇' http://localhost:8080/functions/hello_js/scale
#echo ""


curl -X POST  http://localhost:8080/functions/hello \
-H 'Content-Type: text/plain; charset=utf-8' \
-d '🤗 Bob Morane'
echo ""

curl -X POST  http://localhost:8080/functions/hey \
-H 'Content-Type: text/plain; charset=utf-8' \
-d '🥰 Jane Doe'
echo ""


