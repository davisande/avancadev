FROM golang:1.15

WORKDIR /go/src/fullcycle

COPY . .

RUN go get github.com/hashicorp/go-retryablehttp

RUN GOOS=linux go build

EXPOSE 9090

ENTRYPOINT [ "./fullcycle" ]