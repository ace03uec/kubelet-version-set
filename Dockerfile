FROM golang:1.10.2

COPY . /go/src/github.com/rmenn/kubelet-version-set
WORKDIR /go/src/github.com/rmenn/kubelet-version-set

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kubelet-version-set .

FROM alpine:3.6 
COPY --from=0 /go/src/github.com/rmenn/kubelet-version-set/kubelet-version-set /kubelet-version-set
