FROM golang:1.16-alpine

RUN apt-get update && apt-get install -y \
  libpam0g-dev \
  && rm -rf /var/lib/apt/lists/*

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go buil -o ./out/app ./src/readout/*.go


# This container exposes port 1080 to the outside world
EXPOSE 1080

# Run the binary program produced by `go install`
CMD ["./out/app"]