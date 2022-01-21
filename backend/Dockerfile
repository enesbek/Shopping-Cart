FROM golang:1.17

ENV DBNAME="db"

WORKDIR /go/src/backend
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 3000

CMD ["backend"]