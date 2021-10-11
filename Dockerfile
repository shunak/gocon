FROM golang:1.9
WORKDIR /echo
COPY . .
CMD ["go", "run", "/echo/main.go"]
