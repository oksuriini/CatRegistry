FROM golang:1.22.1-bookworm AS build-stage

COPY . /

WORKDIR /src

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /cats

FROM alpine:3.19 AS run-stage

WORKDIR /

COPY --from=build-stage /cats /cats

EXPOSE 8080

ENTRYPOINT [ "./cats" ]
