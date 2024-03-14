# Choose whatever you want, version >= 1.16
FROM golang:1.22-alpine

# Install required packages for Tailwind CSS compilation
RUN apk update && \
    apk add --no-cache npm

# Install tailwindcss/cli globally
RUN npm install -g tailwindcss

# Install sass compiler for SCSS support
RUN npm install -g sass

# Install Air for Go live-reloading
RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

# Copy your Tailwind CSS configuration file
COPY tailwind.config.js ./

# Start watchers for Tailwind CSS and SCSS, and run Go application with Air
# CMD ["sh", "-c", "npx tailwindcss build -i ./web/static/styles/tailwind.css -o ./web/static/styles/tailwind.min.css --watch & go run ./cmd/app/main.go"]
CMD ["sh", "-c", "npx tailwindcss --minify build -i web/static/styles/tailwind.css -o web/static/styles/tailwind.min.css & sass --watch --style=compressed web/static/styles/main.scss:web/static/styles/style.min.css & air -c .air.toml -d"]

