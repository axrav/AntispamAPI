FROM golang:latest

# Set the Current Working Directory inside the container

WORKDIR /app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container

COPY . .



# This container exposes port 8080 to the outside world

EXPOSE 8080

# Build the Go app

RUN go build -o app ./cmd

# Download the dataset

ADD https://huggingface.co/datasets/thehamkercat/telegram-spam-ham/raw/main/dataset.csv dataset.csv

# Command to run the executable

CMD ["./app"]