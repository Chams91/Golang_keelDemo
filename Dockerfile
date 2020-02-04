FROM golang:1.11-alpine AS build
WORKDIR /
COPY main.go go.* /
RUN CGO_ENABLED=0 go build -o /bin/demo
FROM scratch 
COPY --from=build /bin/demo /bin/demo
ENTRYPOINT ["/bin/demo"] 