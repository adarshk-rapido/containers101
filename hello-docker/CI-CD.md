# CI

### Build
`go build`
`go build -o hello-world` 

### Test
`go test`

### Package
`docker build -t <docker-hub-username>/hello-world:<version> .`

### Publish
`docker login`
`docker push`



# CD

`docker pull <docker-hub-username>/hello-world:<version>`
`docker run --rm -it  -p 8080:8080 <docker-hub-username>/hello-world:<version>`


> Automate the CI part using Github Actions