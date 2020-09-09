FROM golang:alpine AS api-build
RUN apk add --no-cache gcc musl-dev

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go test -cover ./...
RUN go build

FROM node:alpine AS spa-build
WORKDIR /home/node
COPY teamplanner-spa .
RUN npm i && npm run build

FROM alpine AS runtime
RUN mkdir /data && mkdir /dist
COPY --from=api-build /go/src/app/teamplanner /bin/teamplanner
COPY --from=spa-build /home/node/dist/ /dist

ENV DBPATH=/data/teamplanner.db
ENV LISTENADDR=":8042"
EXPOSE 8042

ENTRYPOINT [ "/bin/teamplanner" ]

