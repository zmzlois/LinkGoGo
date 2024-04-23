# Build. 
# This docker file is for production use.
FROM golang:1.21-bullseye AS build-stage
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
# within the /app directory 
RUN CGO_ENABLED=0 GOOS=linux go build -o /opt/go-docker-multistage 

# Final stage 
FROM alpine:3.14 
COPY --from=build-stage /opt/go-docker-multistage /opt/go-docker-multistage 
EXPOSE 8080
ENTRYPOINT ["/opt/go-docker-multistage"]