# build stage
FROM golang:1.22-alpine AS build
WORKDIR /app
COPY main.go .
RUN go build -o main main.go

# run stage
FROM scratch
COPY --from=build /app/main /main
USER 9999
CMD ["/main"]
