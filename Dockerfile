FROM golang:1.11-alpine
RUN apk add --update --no-cache git
WORKDIR /go/src/git.coinninja.net/ssm
COPY . .

RUN go get -d -v
RUN go install -v ./...
EXPOSE 8080

CMD ["ssm"]