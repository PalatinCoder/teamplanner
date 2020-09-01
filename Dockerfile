FROM golang:alpine AS build
RUN apk add --no-cache gcc musl-dev

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go test -cover ./...
RUN go build

FROM alpine AS runtime
COPY --from=build /go/src/app/teamplanner /bin/teamplanner
RUN mkdir /data

ENV DBPATH=/data/teamplanner.db
ENV LISTENADDR=":8042"
EXPOSE 8042

ENTRYPOINT [ "/bin/teamplanner" ]

