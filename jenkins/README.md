# Jenkins CI Pipeline for Running Go inside Docker

### Overview

We want a Jenkins pipeline for go/docker.

### Breadcrumbs

- go get https://github.com/golang/dep
- go get <packages>
- cd $GOPATH/src/project && dep ensure