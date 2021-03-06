FROM golang:alpine AS build-env

# Set up dependencies
# bash for debugging
# git, make for installation
# libc-dev, gcc, linux-headers, eudev-dev are used for cgo and ledger installation (possibly)
RUN apk add bash git make libc-dev gcc linux-headers eudev-dev jq


# Set working directory for the build
WORKDIR /root/straightedge
# default home directory is /root

COPY go.mod .
COPY go.sum .

RUN go mod download

# Add source files
COPY . .

# Install strd, strcli
#ENV LEDGER_ENABLED False
RUN make install

# Run strd by default, omit entrypoint to ease using container with strcli
CMD ["strd"]