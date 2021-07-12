FROM golang:alpine as builder
# Install git and certificates
RUN apk --no-cache add tzdata zip ca-certificates git
# Make repository path
RUN mkdir -p /go/src/algodexidx
WORKDIR /go/src/algodexidx
# Install deps
RUN go get -u -v github.com/ahmetb/govvv

# Copy all project files
ADD . .
# Generate a binary
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -ldflags "$(govvv -flags)" -o app ./cmd/algodexidxsvr

# Use centos as new base as we want an algorand node install for the goal cli (for the debugging inspection endpoint)
FROM centos:latest
EXPOSE 8000

RUN yum install wget -y
RUN mkdir -p /node && cd /node && wget https://raw.githubusercontent.com/algorand/go-algorand-doc/master/downloads/installers/update.sh && chmod 544 ./update.sh
RUN cd /node && ./update.sh -i -c stable -p /node -d /node/data -n

COPY --from=builder /go/src/algodexidx/app /app

ENTRYPOINT [ "/app" ]
