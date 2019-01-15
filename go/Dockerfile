FROM golang:1.11 as builder
WORKDIR $GOPATH/src/webapp
RUN go get -d -v github.com/gorilla/mux github.com/lib/pq
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/webapp


FROM golang:1.11
WORKDIR /go/bin/webapp
COPY --from=builder /go/bin/webapp /go/bin/webapp
EXPOSE 8080
CMD ["./webapp"]