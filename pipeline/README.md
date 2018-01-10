# Pipeline 

Jenkins orchestrates the Go and Docker builds then handles post-build aspects of testing, publishing and deploying. Here's a list of ongoing setup in a Jenkins job I'm running:

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
        # prep a docker buildir
        mkdir $WORKSPACE/docker && \        
        # static ui content for the app
        cp -a $WORKSPACE/src/github.com/smithgeoffrey/go-coop/ui $WORKSPACE/docker && \
        rm -f $WORKSPACE/docker/ui/*.*
        # prometheus node_exporter
        curl -SL https://github.com/prometheus/node_exporter/releases/download/v0.14.0/node_exporter-0.14.0.linux-armv7.tar.gz > $WORKSPACE/node_exporter.tar.gz && \
        tar -xvf $WORKSPACE/node_exporter.tar.gz -C $WORKSPACE/docker/ --strip-components=1

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
        
        WORKDIR /app
        ADD     gobinary .
        ADD     ui/ ./ui/
        ADD     node_exporter /bin/
        
        EXPOSE 8081 9100
        CMD ["/app/gobinary", "nohup /bin/node_exporter &"]
        EOF
    
        EXECUTE DOCKER COMMAND
        Docker command: Create/build image
        Build context folder: $WORKSPACE/docker
        Tag of the resulting docker image: coop
            
    POST-BUILD ACTIONS
        
        SLACK NOTIFICATIONS
        notify failure, success & back to normal