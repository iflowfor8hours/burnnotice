# Our base image is Go 1.7 on Alpine Linux.
FROM golang:1.7-alpine

# Establish a working directory and copy our application files into it.
WORKDIR /opt/hello-go
COPY . .

# Build your application.
RUN go build ./...

# Run the application.
ENTRYPOINT ["/opt/hello-go/hello-go"]

# You can test this Docker image locally by running:
#
#    $ docker build -t hello-go .
#    $ docker run --rm -it --expose 8081 -p 8081:8081 -e PORT=8081 hello-go
#
# and then visiting http://localhost:8081/ in your browser.
