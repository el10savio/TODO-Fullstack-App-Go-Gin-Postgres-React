FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git 

RUN mkdir /build

WORKDIR /build

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/build


FROM scratch

COPY --from=builder /go/bin/build /go/bin/build
COPY --from=builder /build/wait-for-it.sh ./wait-for-it.sh

CMD ["chmod", "+x", "wait-for-it.sh"]
CMD ["./wait-for-it.sh", "db:5432", "--"]

ENTRYPOINT ["/go/bin/build"]

EXPOSE 8081