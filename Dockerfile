FROM golang:1.7-alpine

ADD . /go/src/github.com/TobiEiss/freshboard

WORKDIR /go/src/github.com/TobiEiss/freshboard

RUN apk --update add git nodejs

RUN cd freeboard/ && npm install && npm install -g grunt-cli && grunt && cd ..

RUN go get -d -v; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o freshboard

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/src/github.com/TobiEiss/freshboard/freshboard

# Document that the service listens on port 8080.
EXPOSE 8080