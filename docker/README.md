# Docker

### Build Images

- https://blog.iron.io/an-easier-way-to-create-tiny-golang-docker-images/

With `docker build` create a versioned docker image from a Dockerfile:

    cd <path to Dockerfile>
    docker build -t "smithgeoffrey/go-coop:v1" .

In jenkins w/ the plugin I was using, I tried using `smithgeoffrey/coop:$BUILD_NUMBER` as the plugin option called `Tag of the resulting docker image:`.

Verify the image registered with docker:

    docker images | grep smithgeoffrey
   
Create a container from the image, referencing the image by it's image id as seen in the prior verification step:

    docker run -d --name geoff1 e42d47b5e84d            

Verify the container registered with docker:

    docker container ls --all | grep geoff
