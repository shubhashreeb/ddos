FROM golang:alpine as builder

WORKDIR /build

COPY . ./

# Build the binary
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ddos 

FROM alpine

EXPOSE 9002

COPY --from=builder /build/ddos /bin/ddos

CMD ["/bin/ddos"]
