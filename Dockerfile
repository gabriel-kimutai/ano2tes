FROM golang:1.21

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .
RUN go mod tidy
# RUN go build -v -o /usr/local/bin/app ./...

# CMD ["app"]
