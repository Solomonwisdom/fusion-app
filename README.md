### build in docker
* `docker run -v ~/workspace/go:/go -e GO111MODULE=on -e GOPROXY=https://goproxy.io -it golang:1.11 bash`
* `cd src/github.com/lxs137/fusion-app/`
* `go build -v github.com/lxs137/fusion-app/cmd/manager` or `go test ...`
