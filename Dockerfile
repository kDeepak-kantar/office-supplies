FROM golang:1.17.6-alpine as build_base 

ARG SERVICE_NAME

RUN apk add --no-cache git

RUN echo ${SERVICE_NAME}

WORKDIR /tmp/${SERVICE_NAME}

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go_bk.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o server ./cmd/${SERVICE_NAME}

# Main
FROM golang:1.17-alpine
RUN apk add ca-certificates
RUN  apk update && \
    apk add --no-cache \
    openssh-keygen

ARG SERVICE_NAME
COPY --from=build_base /tmp/${SERVICE_NAME}/server /app/server
COPY --from=build_base /tmp/${SERVICE_NAME}/public public
# COPY --from=build_base /tmp/${SERVICE_NAME}/data data

# This container exposes port 8080 to the outside world
EXPOSE 8080
CMD ["/app/server"]

ARG GIT_SHA1
ARG EXTERNAL_URL
