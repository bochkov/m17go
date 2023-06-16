FROM golang:alpine AS builder
WORKDIR /build
COPY . /build/
RUN go mod tidy && go build -o m17 -ldflags "-s -w"

FROM scratch
COPY --from=builder /build/m17 /
EXPOSE 5000
ENTRYPOINT ["/m17"]