FROM golang:1.22.4

RUN useradd golang
USER root

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go mod download

RUN go build stori .

USER golang

EXPOSE 8080
CMD [ "/app/stori"]
