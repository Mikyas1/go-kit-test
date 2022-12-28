FROM golang

#RUN mkdir -p /go/src//home/mike/Documents/projects/go/go-kit-1
WORKDIR /app

#ADD . /go/src//home/mike/Documents/projects/go/go-kit-1
COPY ./bugs/* ./

RUN go mod download

#RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run ./cmd  -watch ./
