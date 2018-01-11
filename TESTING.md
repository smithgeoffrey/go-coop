# Testing

### Not In A Separate Dir Or Package

- test code lives with main code

A supposed benefit includes making it possible to test unexported code as well as the public API. I'll try it in this project:

    foo/source.go
    foo/source_test.go

### _test.go, go test & *testing.T

Any go source file ending in `_test.go` is treated as a test file by `go test`.  In a _test.go file you create functions starting in `Test` that take an argument `*testing.T`:

    package foo
    
    import "testing"
    
    func TestBar(t *testing.T) {
        <insert>
    }

### Assertions

If you want common assertions and mocks, see https://github.com/stretchr/testify.

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

  // assert equality
  assert.Equal(t, 123, 123, "they should be equal")

  // assert inequality
  assert.NotEqual(t, 123, 456, "they should not be equal")

  // assert for nil (good for errors)
  assert.Nil(t, object)

  // assert for not nil (good when you expect something)
  if assert.NotNil(t, object) {

    // now we know that object isn't nil, we are safe to make
    // further assertions without causing any errors
    assert.Equal(t, "Something", object.Value)

  }

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
