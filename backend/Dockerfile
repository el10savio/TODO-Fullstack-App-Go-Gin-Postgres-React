FROM golang:alpine 

RUN apk add --update git

ADD . $GOPATH/src/build/

WORKDIR $GOPATH/src/build/

RUN go get -u "github.com/gin-gonic/contrib/cors"
RUN go get -u "github.com/lib/pq"

EXPOSE 8081

CMD ["wget", "https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh"]
CMD ["chmod", "+x", "wait-for-it.sh"]
CMD ["./wait-for-it.sh", "db:5432", "--"]

CMD ["go", "run", "main.go"]
