# Docker

### Build Images, Run Containers

- https://blog.iron.io/an-easier-way-to-create-tiny-golang-docker-images/

It took me a while to get anything working at all, but I landed on a Dockerfile that let me hit the app:

    # create the dockerfile
    cd $WORKSPACE/docker && \
    cat > Dockerfile << EOF
    FROM golang
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

Everything looks ok but the app wouldn't run in the interactive shell:

    /app # ./gobinary 
    sh: ./gobinary: not found

It turns out go compiles with glibc but alpine avoids that in favor of muslc, by default.  So I changed the Dockerfile to use `FROM golang` instead of `FROM golang:alpine` and it worked!

    docker run --rm -it coop sh

    # pwd
    /app
    # ls
    gobinary  ui
    # ./gobinary
    [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
  
But I couldn't connect to <ip>:8081 from my laptop.  I run same but added some port (non-) translation complexity, and now I could connect remotely while running it iteractively:
 
    docker run -it -p 8081:8081 coop 

I tried this and it still worked:

    docker run -rm -p 8081:8081 coop 

All of these though we're holding on to the shell, not running it in the background.  I added -d and it looked like I was in business:

    docker run -d -p 8081:8081 coop
    fe89c2315b3343c651912d69f4a5de3005fe67e4a8b8cf4b78fdfd14726c0cc1

I could track it with this:

    docker ps 
    CONTAINER ID        IMAGE               COMMAND             CREATED              STATUS              PORTS                    NAMES
    fe89c2315b33        coop                "/app/gobinary"     About a minute ago   Up About a minute   0.0.0.0:8081->8081/tcp   agitated_euclid

### Where are the Logs?

### How should we monitor?
