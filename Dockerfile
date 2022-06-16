FROM golang:1.18-alpine AS base
WORKDIR /app

FROM base AS build
ADD . /app
RUN go build -o bin/pingpong

FROM alpine:3.16 AS runtime
COPY --from=build /app/bin/pingpong /usr/local/bin/
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/pingpong"]
CMD ["server"]