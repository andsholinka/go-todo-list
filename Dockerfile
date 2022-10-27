FROM golang:1.19-alpine

# Create app directory
WORKDIR /app

COPY . .

RUN go build -o go-todo-list

EXPOSE 3030

# CMD [ "go", "run", "/app/main.go" ]
CMD ./go-todo-list