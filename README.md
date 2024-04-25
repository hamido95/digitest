# Golang Application Dockerfile
in this document, I describe how to create a Dockerfile for our Golang application. As mentioned in the challenge description, we are going to deploy our application in a production environment. so this is very important to use best practices in our dockerfile.

At first, we create a new directory for our project and navigate to it. then we create a Dockerfile and open it in our favorite text editor.
```
mkdir golang-app
cd golang-app
touch Dockerfile
```
## Step by Step Instructions 
1. At the first directive we define our base image. base image is important from two approaches. being lightweight and secure. 
    - Official Golang version 1.21 is considered a secure image. however, we regularly check the vulnerabilities using tools like Trivy, Clair, and Dockerhub Image Scanner.
    - Regarding the lightweight tag, golang:1.21-alpine is more lightweight than the default golang:1.21. because it uses the Alpine Linux distribution.
```
FROM golang:1.21-alpine
```
2. When you create a dockerfile this is best practice to define your working directory before any command. So, we set /app directory as our working directory.
```
WORKDIR /app
```
3. Now, we should copy the "go.mod" and "go.sum" files to our working directory. in the Go programming language, these are two important files that manage dependencies.
```
COPY go.mod ./
COPY go.sum ./
```
We also can combine these two commands:
```
COPY go.mod go.sum ./
```
4. Now, we have go.mod and go.sum files inside our docker image and ready for installing go modules into the directory inside our image.
```
RUN go mod download
```
5. At this point, we need to copy our source code from our host machine to our docker image. We can also use a wildcard "*" to copy all files with .go extension
```
COPY . ./
```

6. Now, we use RUN command to compile our application.

```
RUN CGO_ENABLED=0 GOOS=linux go build -o /risk-ident
```
CGO_ENABLED=0 and GOOS=linux are two environment variables that we use when we want to build our Go application for deploying on linux servers.
-  **CGO_ENABLED=0** is used to disable CGO. It removes the need to include any C libraries. Therefore, it will increase efficeincy. 
- **GOOS=linux** tells the Go compiler to build the application for the linux server.
- **-o** is the output name of our static application file.

we can also set ENV directive in our dockerfile like bellow:
```
ENV CGO_ENABLED=0
ENV GOOS=linux
```
7. in this step we use EXPOSE directive to set the listening port of our application. 
```
EXPOSE 8080
```

8. The last step is to tell the docker what command to run when our image is used to run a container.

```
CMD ["/app/risk-ident"]
```

## Building the Image
At this point we have our Dockerfile and we can easily build our image with the "docker build" command. 

We use -t flag to set tag for our image. If we do not set tag, the docker will set **latest** tag as default.
```
docker build -t risk-ident:v1.0.0 .
```
we can also check our image list to verify image is created
```
docker images
``` 

## Running the Image
according to our application documentation, we need a mongodb instance that our application connects to.
```
docker run -dit --name mongo -p 27017:27017 --network production mongo

docker run -dit --name risk-ident -p 8080:8080 --network production -e GOCRUD_MONGO_URI="mongodb://mongo:27017" risk-ident:v1.0.0
```

## Pushing to Docker Registry
After building our docker image, we can push it to our docker registry to store images with proper tags and then use it in our CI/CD pipelines.
```
docker login dockerhub.io/risk-ident
# enter registry credentials

docker tag risk-ident:v1.0.0 risk-ident/gocrud/risk-ident:v1.0.0

docker push risk-ident/gocrud/risk-ident:v1.0.0
```