FROM golang

ADD . /go/src/github.com/koki/kubernetes-viewer/server

RUN go install github.com/koki/kubernetes-viewer/server

ENTRYPOINT ["/go/bin/server"]
CMD []
