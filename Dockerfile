# Define builder stage
FROM golang:1.11.3-alpine as builder

# Sets GOPATH environment var and set ports
ENV GOPATH="/project/go"
ENV GUITARIST_BACKEND_PORT=1337

# Install git
RUN apk add --update git && \
    rm -rf /var/cache/apk/*
WORKDIR $GOPATH/src/github.com/mauhftw/go-guitarists/

# Install dep, cache dependencies and ensure vendor
RUN go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.* $GOPATH/src/github.com/mauhftw/go-guitarists/
RUN $GOPATH/bin/dep ensure --vendor-only -v
COPY . $GOPATH/src/github.com/mauhftw/go-guitarists

# Copile application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o dist/guitarist


# Define release stage
FROM scratch as release

# Expose port and copy release binary
EXPOSE $GUITARIST_BACKEND_PORT
COPY --from=builder /project/go/src/github.com/mauhftw/go-guitarists/dist/guitarist /go/bin/guitarist

# Run application
ENTRYPOINT ["/go/bin/guitarist"]
