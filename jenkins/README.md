# Jenkins CI Pipeline for Running Go inside Docker

### Overview

I want a Jenkins pipeline for running a go binary in a docker container.

### Config

Here's a list of the basic setup in a job I'm running:

    SOURCE CODE MANAGEMENT
        
        REPOSITORIES
        https://github.com/smithgeoffrey/go-coop
        */master
        
        ADDITIONAL BEHAVIORS: Checkout to a subdirectory
        $WORKSPACE/src/github.com/smithgeoffrey/go-coop
        
    BUILD ENVIRONMENT
    
        Delete workspace before build starts
        Add timestamps to console output
        Inject environment variables to the build process
            Properties content
                GOROOT=/usr/local/go
                GOPATH=$WORKSPACE
                PATH+=:$GOROOT/bin:$GOPATH/bin
    
    BUILD
        
        EXECUTE SHELL
        # prep a docker buildir having static content for the app
        mkdir $WORKSPACE/docker && \
        cp -a $WORKSPACE/src/github.com/smithgeoffrey/go-coop/ui $WORKSPACE/docker && \
        rm -f $WORKSPACE/docker/ui/*.*

        EXECUTE SHELL
        # build the app binary and put it in the docker buildir
        cd $WORKSPACE/src/github.com/smithgeoffrey/go-coop && \
        go get -u github.com/golang/dep/... && \
        dep init && dep ensure && \
        go build *.go && \
        mv main $WORKSPACE/docker/binary
    
        EXECUTE SHELL
        # create the dockerfile
        cd $WORKSPACE/docker && \
        cat > Dockerfile << EOF
        FROM golang:alpine
        WORKDIR /app
        COPY ./binary /app/
        COPY ./ui /app/
        EXPOSE 8081
        ENTRYPOINT ["./binary"]
        EOF        
    
        EXECUTE DOCKER COMMAND
        Docker command: Create/build image
        Build context folder: $WORKSPACE/docker
        Tag of the resulting docker image: $BUILD_NUMBER
            
    POST-BUILD ACTIONS
        
        SLACK NOTIFICATIONS
        notify failure, success & back to normal