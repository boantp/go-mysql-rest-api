# Telling to use Docker's golang ready image
FROM golang:onbuild
EXPOSE 8080

RUN apt-get update && apt-get install -y mysql-client
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/boantp/go-mysql-rest-api
# ADD . /home

# WORKDIR /home

# RUN \
#         apk add --no-cache bash git openssh && \
#         go get -u github.com/minio/minio-go
        

# CMD [ "go", "run", "main.go" ]

