# Docker

### Build Image, Run Container

I landed on a Dockerfile that let me hit the app on the pi:

    # create the dockerfile
    cd $WORKSPACE/docker && \
    cat > Dockerfile << EOF
    FROM golang
    MAINTAINER Geoff Smith "smithgeoffrey123@gmail.com"
    EXPOSE 8081
    
    WORKDIR /app
    ADD gobinary .
    ADD ui/ ./ui/
    
    ENV PORT=8081
    CMD ["/app/gobinary"]
    EOF

Rather than running `docker build . -t coop`, I let the jenkins docker-build plugin do it, setting `coop` as the plugin option called `Tag of the resulting docker image:`. Once built, I verified via `docker images | grep coop` and `docker inspect coop`.  I used an interactive shell of a container running the image:
 
    docker run -it coop sh

    /app # ls
    gobinary  ui
    /app#

Everything looked ok except that the app wouldn't run:

    /app # ./gobinary 
    sh: ./gobinary: not found

It turns out go compiles with glibc but alpine avoids that in favor of muslc, by default.  I changed the Dockerfile to use `FROM golang` instead of `FROM golang:alpine` and that issue went away:

    docker run -it coop sh

    # pwd
    /app
    # ls
    gobinary  ui
    # ./gobinary
    [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
  
But I couldn't connect to <pi ip>:8081 from my laptop, until I added port translation between the host and container:
 
    docker run -p 8081:8081 coop

Finally, it was holding on to the shell until I specified the detachment option:

    docker run -d -p 8081:8081 coop
    fe89c2315b3343c651912d69f4a5de3005fe67e4a8b8cf4b78fdfd14726c0cc1

I could track it with this:

    docker ps 
    CONTAINER ID  IMAGE  COMMAND          CREATED              STATUS              PORTS                    NAMES
    fe89c2315b33  coop   "/app/gobinary"  About a minute ago   Up About a minute   0.0.0.0:8081->8081/tcp   agitated_euclid

That shows running containers.  To see them all regarless of status:

    docker container ls --all

I noticed the strange names being assigned to containers, like agitated_euclid.  It seems the code for that is at https://github.com/moby/moby/blob/master/pkg/namesgenerator/names-generator.go.

### Where are the Logs?

### Now have Jenkins deploy the container instead of me running it manually after build

### How should we test?

### How should we monitor?
