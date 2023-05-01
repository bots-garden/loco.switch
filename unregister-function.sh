#!/bin/bash

curl -X DELETE http://localhost:8080/admin/functions/registration \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"hello_js",
    "revision":"default"
}
EOF

echo ""