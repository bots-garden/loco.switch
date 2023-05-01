# loco.switch

**LocoSwitch** is a kind of tiny reverse proxy.

## Usage

First, start LocoSwitch:

### Start
```bash
./loco-switch
```
> the default HTTP port of LocoSwitch is `8080`

### Register

For example, you started several HTTP services on various port:

```bash
http://localhost:8888
http://localhost:6666
```

To register the two services, use this:
```bash
curl -X POST http://localhost:8080/admin/functions/register \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"hello_go",
    "revision":"default",
    "httpPort":8888,
    "status":0,
    "https":false,
    "domain":"localhost",
    "endpoint":""
}
EOF

curl -X POST http://localhost:8080/admin/functions/register \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"hello_js",
    "revision":"default",
    "httpPort":6666,
    "status":0,
    "https":false,
    "domain":"localhost",
    "endpoint":""
}
EOF
```

### Call the services with the new endpoints

```bash
curl -X POST -d 'John Doe' http://localhost:8080/functions/hello_go
curl -X POST -d 'Bob Morane' http://localhost:8080/functions/hello_js
```

## Revisions

> 🚧 this is a work in progress

## Load balancing

> 🚧 this is a work in progress

