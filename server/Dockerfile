FROM golang:alpine as development

WORKDIR /app

#Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

# RUN go build -o server .

#CMD [ "go" , "run", "main.go", "process.go"]