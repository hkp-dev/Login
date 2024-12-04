FROM golang:latest
EXPOSE 3000
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o app cmd/main.go
CMD ["./app"]

#loads image từ file tar 
#syntax: docker save <image_name> -o <file_name>
#step1: docker save app -o myapp_image.tar
#step2: docker load -i myapp_image.tar
#step3: docker run -p 3000:3000 app

#build image: docker build -t <image name> .
#run container: docker run -p 3000:3000 -name <container name> <image name> 
# -name: name of the container
# -p: port mapping
# --restart always: restart the container if it fails
# -d: detached mode
# -it : interactive mode
# --restart always: restart the container if it fails