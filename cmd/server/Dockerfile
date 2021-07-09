FROM golang:alpine
RUN apk update && \
    apk add curl \
            git \
            bash \
            make \
            ca-certificates && \
    rm -rf /var/cache/apk/*

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    APP_ENV=local

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container and build
COPY . .
RUN make build

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Prepare builf
RUN cp /build/server . \
&& mkdir config \
&& chmod +x /build/cmd/server/entrypoint.sh \
&& cp /build/cmd/server/entrypoint.sh . \
&& cp /build/config/*.yml ./config/ \
&& mkdir -p /var/log/app
#RUN rm -rf /build

# Export necessary port
EXPOSE 8080

# Command to run when starting the server
ENTRYPOINT ["./entrypoint.sh"]