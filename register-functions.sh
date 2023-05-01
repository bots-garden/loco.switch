#!/bin/bash

curl -X POST http://localhost:8080/admin/functions/registration \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"hello_go",
    "revision":"default",
    "httpPort":8888,
    "status":0,
    "https":false,
    "domain":"localhost"
}
EOF

curl -X POST http://localhost:8080/admin/functions/registration \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"hello_js",
    "revision":"default",
    "httpPort":6666,
    "status":0,
    "https":false,
    "domain":"localhost"
}
EOF

echo ""

curl -X POST http://localhost:8080/admin/functions/registration \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"hello_js",
    "revision":"purple",
    "httpPort":5555,
    "status":0,
    "https":false,
    "domain":"localhost"
}
EOF


echo ""


curl -X POST http://localhost:8080/admin/functions/registration \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"hello_js",
    "revision":"orange",
    "httpPort":4444,
    "status":0,
    "https":false,
    "domain":"localhost"
}
EOF


echo ""

curl -X POST http://localhost:8080/admin/functions/registration \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"hello_js",
    "revision":"scale",
    "httpPort":1111,
    "status":0,
    "https":false,
    "domain":"localhost"
}
EOF

# Add endpoint
curl -X POST http://localhost:8080/admin/functions/endpoint \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"hello_js",
    "revision":"scale",
    "httpPort":2222,
    "status":0,
    "https":false,
    "domain":"localhost"
}
EOF

# Add endpoint
curl -X POST http://localhost:8080/admin/functions/endpoint \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"hello_js",
    "revision":"scale",
    "httpPort":3333,
    "status":0,
    "https":false,
    "domain":"localhost"
}
EOF


echo ""