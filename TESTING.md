# Testing in Go

### Overview

I hadn't done a lot of testing yet, much less in go.  

I wanted to keep testing as a top level package, until I read that may be a no no in go, instead prefering to have testing live side by side with the main code:

    package/foo.go
    package/foo_test.go

So I'll start with that approach here.  Eventually, I'd like to read these breadcrumbs.

- https://github.com/golang/lint
- https://github.com/tebeka/go2xunit
- https://medium.com/@zarkopafilis/building-a-ci-system-for-go-with-jenkins-4ab04d4bacd0
- https://nathanleclaire.com/blog/2015/03/09/youre-not-using-this-enough-part-one-go-interfaces/
- https://nathanleclaire.com/blog/2015/10/10/interfaces-and-composition-for-effective-unit-testing-in-golang/
- https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742
- https://elithrar.github.io/article/testing-http-handlers-go/
