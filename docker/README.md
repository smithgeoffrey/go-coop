# Docker

### Build Images, Run Containers

- https://blog.iron.io/an-easier-way-to-create-tiny-golang-docker-images/

With `docker build` create a versioned docker image from a Dockerfile:

    cd <path to Dockerfile>
    docker build . -t coop:1

In jenkins w/ the plugin I was using, I tried using `coop:$BUILD_NUMBER` as the plugin option called `Tag of the resulting docker image:`. Verify the image registered with docker:

    docker images | grep coop
   
It looks like I could create a container from that image using something like:
 
    docker run -it -p 8081:8081 coop 

But it looks like the jenkins plugin already created a container:

    docker container ls --all | grep coop
    docker container start <container id>

Whichever way I try to run it, the container status is `Exited` and throws that it couldn't find the binary. So my Dockerfile wasn't plumbed right I think.

### Troubleshoot Container Startup w/ Logs

