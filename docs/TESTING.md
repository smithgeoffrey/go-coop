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

For assertions use https://github.com/stretchr/testify.

    import (
      "errors"
      "testing"
      "github.com/stretchr/testify/assert"
    )
    
    func TestSomething(t *testing.T) {
      res, err := fun1()
      assert.IsType(t, error, err)
      assert.Equal(t, <expected res>, res, "should be equal")
      assert.NotEqual(t, <expected not res>, res, "should be not equal")
      assert.Nil(t, object)
      if assert.NotNil(t, object) {
        // further assertions without causing any nil errors
        assert.Equal(t, "Something", object.Value)
      }
    }    

Some like to load tests into a structure of some kind then iterate over it for testing different inputs or mock responses:

    import "testify"
        
    type Tests struct {
        result string
        err MyError
    }
    
    func TestGetDoor(t *testing.T) {
        result, err := dosomething(arg1)
        tests := []Tests{
            {result: "foo", err: error},
            {result: "bar", err: error},
        }
        for _, test := range tests {
            <asserts>
        }
    }

### Mocking Interfaces

Interface mocking is a typical component of testing.  To avoid boilerplate coding required to do it, try https://github.com/vektra/mockery.




### Breadcrumbs

Eventually, I'd like to read these:

- https://github.com/golang/lint
- https://github.com/tebeka/go2xunit
- https://nathanleclaire.com/blog/2015/03/09/youre-not-using-this-enough-part-one-go-interfaces/
- https://nathanleclaire.com/blog/2015/10/10/interfaces-and-composition-for-effective-unit-testing-in-golang/
- https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742
- https://elithrar.github.io/article/testing-http-handlers-go/
- https://medium.com/@zarkopafilis/building-a-ci-system-for-go-with-jenkins-4ab04d4bacd0