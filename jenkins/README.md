# Pipeline for Running Go inside Docker

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
        mv main $WORKSPACE/docker/gobinary
    
        EXECUTE SHELL
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
    
        EXECUTE DOCKER COMMAND
        Docker command: Create/build image
        Build context folder: $WORKSPACE/docker
        Tag of the resulting docker image: coop
            
    POST-BUILD ACTIONS
        
        SLACK NOTIFICATIONS
        notify failure, success & back to normal