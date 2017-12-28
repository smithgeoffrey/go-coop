# Jenkins CI Pipeline for Running Go inside Docker

### Overview

We want a Jenkins pipeline for running a go binary in a docker container.  Possible use case is for deploying microservices in a small, secure way.

### Config

Eventually I'll land on some steady state and can include the job's .xml file.  For now, here's a list of the basic setup in a job I'm running:

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
        cd $WORKSPACE && go get -u github.com/golang/dep/...
        
        EXECUTE SHELL
        cd $WORKSPACE/src/github.com/smithgeoffrey/go-coop && dep init && dep ensure
        
        EXECUTE SHELL
        cd $WORKSPACE && go build src/github.com/smithgeoffrey/go-coop/*.go && mv main app
        
    POST-BUILD ACTIONS
        
        SLACK NOTIFICATIONS
        notify failure, success & back to normal