FROM golang:1.22-alpine AS build

WORKDIR /app

COPY main.go .
#RUN go build main.go
RUN go build -o main main.go

FROM scratch

COPY --from=build /app/main /main

CMD ["/main"]
