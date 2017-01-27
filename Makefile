
build: main.go rest.go
  go get \
  go build

run: restmb
  ./restmb

image: Dockerfile
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o restmb . \
  docker build -t paoloo/restmb .

docker: Dockerfile
	docker run -p 8080:8080 paoloo/restmb
