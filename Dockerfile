FROM golang:1.16-buster as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/portainer-templates cmd/main.go

FROM gcr.io/distroless/base-debian10

COPY --from=build /go/bin/portainer-templates /

ENV PORT 4345
ENV HOST 0.0.0.0
ENV DB_FILE /data/portainer-templates.db
ENV TEMPLATE_FILE /data/templates.json
ENV LOG_LEVEL info

VOLUME /data

CMD ["/portainer-templates"]