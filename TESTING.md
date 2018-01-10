# Testing

### Overview

There's a few things about Go's testing that sets it apart:
  
- no heavy focus on assertion tools
- test code and main code live side by side, not in a separate dir or package

A supposed benefit of the latter includes making it possible to test unexported code as well as the public API.  I'll start with this approach here:

    foo/source.go
    foo/source_test.go

Any go source file ending in `_test.go` is treated as a test file by `go test`.  In a _test.go file you create functions starting in `Test` that take a param `*testing.T`:

    package foo
    
    import "testing"
    
    func TestBar(t *testing.T) {
        <insert>
    }

### Breadcrumbs

Eventually, I'd like to read these:

- https://github.com/golang/lint
- https://github.com/tebeka/go2xunit
- https://nathanleclaire.com/blog/2015/03/09/youre-not-using-this-enough-part-one-go-interfaces/
- https://nathanleclaire.com/blog/2015/10/10/interfaces-and-composition-for-effective-unit-testing-in-golang/
- https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742
- https://elithrar.github.io/article/testing-http-handlers-go/
- https://medium.com/@zarkopafilis/building-a-ci-system-for-go-with-jenkins-4ab04d4bacd0
