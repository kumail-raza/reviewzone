FROM golang:1.10.0

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/minhajuddinkhan/reviewzone/dump
WORKDIR /go/src/github.com/minhajuddinkhan/reviewzone/dump
# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get
RUN go install github.com/minhajuddinkhan/reviewzone/dump/bin/dump

# Run the outyet command by default when the container starts.
#ENTRYPOINT /go/bin/dump

# Document that the service listens on port 8080.
EXPOSE 4000


CMD [ "go", "run", "/go/src/github.com/minhajuddinkhan/reviewzone/dump/bin/dump/main.go" ]