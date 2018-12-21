FROM golang:1.11.3-alpine

# Sets GOPATH environment var and set ports
ENV GOPATH="/project/go"
ENV GUITARIST_BACKEND_PORT=1337

# Install git
RUN apk add --update git make

# Set workdir and copy files
WORKDIR $GOPATH/src/github.com/mauhftw/go-guitarists
COPY . $GOPATH/src/github.com/mauhftw/go-guitarists

# Install dep
RUN go get -u github.com/golang/dep/cmd/dep

# Ensure vendor
RUN $GOPATH/bin/dep ensure -v

# Copile application
RUN make build

# Expose port
EXPOSE $GUITARIST_BACKEND_PORT

# Run application
CMD ["dist/guitarist"]
