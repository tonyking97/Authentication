FROM golang

ADD . /go/src/auth_server
WORKDIR /go/src/auth_server
#COPY . .

RUN go get -d -v ./...
#RUN go install -v ./...
EXPOSE 50051

CMD ["go", "run","server.go"]

