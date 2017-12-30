# Docker

### Build Images, Run Containers

- https://blog.iron.io/an-easier-way-to-create-tiny-golang-docker-images/

It took me a while to get anything working at all, but I landed on a Dockerfile that let me hit the app:

    # create the dockerfile
    cd $WORKSPACE/docker && \
    cat > Dockerfile << EOF
    FROM golang:alpine
    MAINTAINER Geoff Smith "smithgeoffrey123@gmail.com"
    EXPOSE 8081
    
    WORKDIR /app
    COPY gobinary .
    ADD ui/ ./ui/
    
    ENV PORT=8081
    CMD ["/app/gobinary"]
    EOF

Build the image manually with `docker build . -t coop` or let jenkins do same by setting `coop` as the (docker-build) plugin option called `Tag of the resulting docker image:`. After build, verify the image is inventoried via `docker images | grep coop`.  I liked to verify on a container spawned from the image, in an interactive shell; note that it dumps me to the `workdir` immediately:
 
    docker run --rm -it coop sh

    /app # ls
    gobinary  ui
    /app#

Everything looks ok but the app wouldn't run:

    /app # ./gobinary 
    sh: ./gobinary: not found

It turns out go compiles with glibc but alpine with muslc, by default.  So I have a little homework how to patch the two so a go binary runs like usual on alpine.  Maybe I'll just avoid alpine for the moment.
  
Run the image not interactively:
 
    docker run -it -p 8081:8081 coop 

Also, try inspecting the image:

    docker inspect 10d765f63e2c


### Troubleshoot Container Startup w/ Logs

