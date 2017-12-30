# Docker

### Build Images, Run Containers

It took me a while to get anything working at all.  I landed on a final Dockerfile that let me hit the app on the pi from my laptop, where Jenkins built the go binary and a docker image serving it, and where I instantiated a docker container from the image via `docker run` on the pi:

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

Rather than running `docker build . -t coop`, I let jenkins docker-build plugin it, setting `coop` as the plugin option called `Tag of the resulting docker image:`. After build, I verified the image was inventoried via `docker images | grep coop`.  

From there, I used an interactive shell to poke around:
 
    docker run --rm -it coop sh

    /app # ls
    gobinary  ui
    /app#

Everything looked ok but the app wouldn't run:

    /app # ./gobinary 
    sh: ./gobinary: not found

It turns out go compiles with glibc but alpine avoids that in favor of muslc, by default.  I changed the Dockerfile to use `FROM golang` instead of `FROM golang:alpine` and it worked!

    docker run --rm -it coop sh

    # pwd
    /app
    # ls
    gobinary  ui
    # ./gobinary
    [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
  
But I couldn't connect to <ip>:8081 from my laptop.  So I added a little port (non-) translation complexity, and then I could connect remotely:
 
    docker run -it -p 8081:8081 coop
     or
    docker run --rm -p 8081:8081 coop

Those were holding on to the shell though, not running in the background.  So I used -d and it looked like I was in business:

    docker run -d -p 8081:8081 coop
    fe89c2315b3343c651912d69f4a5de3005fe67e4a8b8cf4b78fdfd14726c0cc1

I could track it with this:

    docker ps 
    CONTAINER ID        IMAGE               COMMAND             CREATED              STATUS              PORTS                    NAMES
    fe89c2315b33        coop                "/app/gobinary"     About a minute ago   Up About a minute   0.0.0.0:8081->8081/tcp   agitated_euclid

### Where are the Logs?

### How should we monitor?
