FROM golang:alpine as builder
# Install git and certificates
RUN apk --no-cache add tzdata zip ca-certificates git
# Make repository path
RUN mkdir -p /go/src/algodexidx
WORKDIR /go/src/algodexidx
# Install deps
RUN go get -u -v github.com/ahmetb/govvv
#&& \
#	go get -u -v github.com/gorilla/mux
# Copy all project files
ADD . .
# Generate a binary
#RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "$(govvv -flags)" cmd/algodexidxsvr -o app
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/algodexidxsvr

# Second (final) stage, base image is scratch
FROM scratch
EXPOSE 8000

# Copy statically linked binary
COPY --from=builder /go/src/algodexidx/app /app
# Copy SSL certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Notice "CMD", we don't use "Entrypoint" because there is no OS
CMD [ "/app" ]
