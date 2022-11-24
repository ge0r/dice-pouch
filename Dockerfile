FROM golang:1.19-alpine

WORKDIR /dice-pouch

# TODO only copy go. mod and the directories
COPY . ./

RUN go build -o dp ./cmd/web

EXPOSE 9001

CMD ["./dp"]


