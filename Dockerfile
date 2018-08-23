FROM golang:1.10-alpine AS builder

WORKDIR /go/src/github.com/alileza/sample-app

COPY . ./

RUN go build -o /sample-app main.go

# ---

FROM alpine

COPY --from=builder /sample-app /bin/sample-app

EXPOSE 9000
ENTRYPOINT [ "/bin/sample-app" ]
