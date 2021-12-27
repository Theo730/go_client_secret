FROM golang:1.16 as builder

LABEL maintainer="go_client_secret"

WORKDIR /app

COPY go.mod go.sum ./
#COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

#WORKDIR /root/

ADD ./cert.pem /etc/ssl/certs/
ADD ./key.pem /etc/ssl/certs/
#ADD ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main .

EXPOSE 5959

CMD ["/main", "-Cert", "/etc/ssl/certs/cert.pem", "-Key", "/etc/ssl/certs/key.pem"]

