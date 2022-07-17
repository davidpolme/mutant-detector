FROM golang:latest AS builder
RUN apt-get-update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 
WORKDIR /go/src/app
COPY . .
RUN go install

FROM scratch 
COPY --from=builder /go/bin/app .
ENTRYPOINT [ "./app" ]


RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]