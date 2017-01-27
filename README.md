# GoRESTmb
A RESTful modbus to http connector API to make testing industrial devices easier.
It's based on github.com/paoloo/modbuscli, a pure **go** modbus driver.

```
+--------+                  +----------+           |======================= = = =
| device < TCP/502 --------<  GoRESTmb < 8080/TCP -+ GET, POST, PUT, DELETE
+--------+                  +----------+           |======================= = = =
```

## Build

First, change modbus device endpoint address and port:
```go
mb.EndPoint = "127.0.0.1:502"
```

Locally:
  ```bash
  go get
  go build
  ```
Docker image:
  ```bash
  go get
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o restmb .
  docker build -t paoloo/restmb .
  ```

Additionally, a Make file was provided where:

- `make build` to build locally

- `make mage` to create docker image

- `make run` to run local build

- `make docker` to run docker image

## Usage

Run it with `./gorestmb` or `docker run -p 8080:8080 paoloo/restmb` then:

- To write an array of numbers ( holding register ), do:

 - `$ curl -v -X POST -H "Content-Type: application/json" -d '{"address": 2, "values":[5,8]}' http://localhost:8080/modbus/`

 - whose result is: `[2,2]`

- To read an array of 16 bits data( holding register )

 - `$ curl -v -X GET -H "Content-Type: application/json" http://localhost:8080/modbus/1/10`

 - whose result is: `[0,5,8,0,0,0,0,0,0,0]`

- To update a register(could be done with POST)

 - `url -v -X PUT -H "Content-Type: application/json" -d '{"values":[3,7]}' http://localhost:8080/modbus/2`

 - whose result is: `[2,2]`

- To delete(return to zero any register:

 - `curl -v -X DELETE -H "Content-Type: application/json" http://localhost:8080/modbus/2`

 - whose result is: `[2,0]`

## Enjoy
