# First stage: build the executable.
ARG GO_VERSION=1.12
FROM golang:${GO_VERSION} AS builder

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./
RUN GOPROXY=https://proxy.golang.org go mod download

# Import the code from the context.
COPY ./ ./

# RUN go build -o ./app ./main.go
# Build the executable to `/app`. Mark the build as statically linked.
RUN CGO_ENABLED=0 GOOS=linux go build \
     -installsuffix 'static' \
     -o /main .

# Final stage: the running container.
FROM scratch AS final

# Import the compiled executable from the first stage.
COPY --from=builder src/config.yml /config.yml
COPY --from=builder /main /main

EXPOSE 8080

# Run the compiled binary.
ENTRYPOINT ["/main"]