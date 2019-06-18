FROM golang as builder
COPY . /go/src/github.com/flaccid/kunsul
WORKDIR /go/src/github.com/flaccid/kunsul
RUN go get ./... && \
    CGO_ENABLED=0 GOOS=linux go build -o /kunsul cmd/kunsul/kunsul.go

FROM gcr.io/distroless/static
COPY --from=builder /kunsul /kunsul
COPY template.html /usr/share/kunsul/template.html
ENTRYPOINT ["/kunsul"]
