# Telling to use Docker's golang ready image
FROM golang:1.8

MAINTAINER boantp

EXPOSE 3000 443

RUN apt-get update -y && \
	apt-get install -y \
	vim

# Install and build binaries revel for this server
RUN go get github.com/boantp/go-mysql-rest-api
RUN go get github.com/go-sql-driver/mysql

# Set pwd to the go folder
RUN mkdir /app
ADD . /app/
WORKDIR  /app

RUN go build -o server
RUN go run server.go
CMD ["/app/server"]

# Copy the entrypoint file
ADD /docker/web/infinite_loop.sh /infinite_loop.sh
RUN chmod +x /infinite_loop.sh

# Start process
ENTRYPOINT ["/infinite_loop.sh"]