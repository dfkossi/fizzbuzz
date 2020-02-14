FROM golang:1.13.7-alpine3.11

RUN mkdir /app

ADD . /app
 
WORKDIR /app

RUN go build *.go 

EXPOSE 3000

RUN chmod +x ./app

ENTRYPOINT ["./app"]