# ERROR HANDLING

Go's approach to error handling is a little different, mostly centered around:

- no try-catch
- functions return response-error pairs

### Error Returns with Response
 
Here's some suedo code to show the idea:

    import errors
    import fmt
    
    func func1() (string, error) {
        if <error condition> {
            return "", errors.New("My error message")
        }
        return "my response", nil 
    }
    
    res, err := func1(arg1)
    if err != nil {
        return "", fmt.Printf("Error: %s", err)
    }

### Custom Error Types

There's also an understanding of the Error interface in Go:

    type error interface {
        Error() string
    }

Because anything that implements it is an Error, you can create custom error types.  Just create a type and apply a method to it that satisfies the error interface:  

    type ErrorChickenDoorJammed struct {
        message string
    }
    
    func (e *ErrorChickenDoorJammed) Error() string {
        return e.message
    }

One possible use is a way to parse err based on type: [1]  I haven't used this pattern yet but I'll give it a try.

    import fmt
    
    func foo() (string, *ErrorChickenDoorJammed) {
        if <door is jammed> {
            return "", &ErrorChickenDoorJammed{message: "custom error messages"}
        }
        return "success", nil
    }
    
    # now consume it and test for err type
    res, err := foo()
    if err != nil {
        switch err.(type) {
        case *ErrorChickenDoorJammed:
            # now you know why
            fmt.Println(err.Error())
        default:
            fmt.Println("Error: not sure why")
        }
    }        

### Not Returning nil for Response when Error

One usage I've seen in the wild is don't always return nil on the result when there's an error, instead return if possible the empty value of the type expected.  Nil is fine if there isn't one.  This enables users of your library to streamline their use of your lib to do things like this pseudo code:

    ## users of geoff package have the option to ignore 
    ## the error via `_` and work off the result instead 
    
    import geoff
    
    res, _ := geoff.Chickens(arg1)

### Embed The Underlying Call into the Error Conditional

Another thing you'll see in the wild is collapsing the underlying call into the error handling conditional itself.  It leverages Go's if statement having an optional assignment clause before the expression.  Something like:
 
    if res, err := geoff.Chickens(arg1); err != nil {
        fmt.Printf("Error: %s\n", err)
    }
    
Compare to how I had generally thought of it, below.  

    res, err := geoff.Chickens(arg1)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
    }

### References

[1] https://medium.com/@sebdah/go-best-practices-error-handling-2d15e1f0c5ee