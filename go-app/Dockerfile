FROM golang:latest
WORKDIR /app/go
COPY ./ ./
RUN go mod download
RUN go build
ENTRYPOINT [ "./diet-planning-app" ]