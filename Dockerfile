# Build. 
# This docker file is for production use.
FROM golang:1.21-bullseye AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
# within the /app directory 
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

# Final stage 
FROM gcr.io/distroless/static-debian11 AS release-stage
COPY --from=build-stage /entrypoint /entrypoint
COPY --from=build-stage /app/web/assets /web/assets
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/opt/go-docker-multistage"]