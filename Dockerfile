FROM golang:alpine AS build
WORKDIR /go/src/veritran
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/veritran cmd/api/main.go

FROM scratch
COPY --from=build /go/bin/veritran /go/bin/veritran
EXPOSE 3000
ENTRYPOINT ["/go/bin/veritran"]
