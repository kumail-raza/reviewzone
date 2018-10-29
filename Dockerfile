FROM golang:1.10.0
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/minhajuddinkhan/reviewzone


COPY . .
RUN dep ensure -v
RUN go install github.com/minhajuddinkhan/reviewzone/api/...
#ENTRYPOINT /go/bin/api

# Document that the service listens on port 8080.
EXPOSE 6000
CMD [ "go", "run", "/go/src/github.com/minhajuddinkhan/reviewzone/api/bin/api/main.go" ]