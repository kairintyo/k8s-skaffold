FROM golang:1.14.3-alpine3.11 as builder
COPY main.go .
RUN CGO_ENABLED=0 go build -o /hello main.go

FROM scratch
COPY --from=builder /hello /opt/app/bin/
COPY assets/ /opt/app/assets/
CMD ["/opt/app/bin/hello"]
