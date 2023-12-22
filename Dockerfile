FROM golang:alpine3.19

WORKDIR /app

COPY . .

RUN go mod download 

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /captain-ai-bot

CMD ["/captain-ai-bot"]