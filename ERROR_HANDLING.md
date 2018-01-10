# ERROR HANDLING

Go's approach to error handling is a little different, mostly centered around there being no try-catch and functions returning response/error pairs.

### Error Returns with Response
 
Here's some suedo code to show the idea:

    import errors
    import fmt
    
    func func1() (string, error) {
        if <error condition> {
            return "", error.New("My error message")
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

Because anything that implements it is an Error, you can create custom error types:  

    type ErrorChickenDoorJammed struct {
        message string
    }
    
    func (e *ErrorChickenDoorJammed) Error() string {
        return e.message
    }

We just created a type and applied a method to it that satisfies the interface.  Hence the arbitrary type has become a custom error type.  Having one isn't much fun if you don't use it:

    func HelperFuncToReturnTheCustomError(message string) *ErrorChickenDoorJammed {
        # instantiate the error object then return it
        err := &ErrorChickenDoorJammed{message: message}
        return err
    }

### Not Reterning nil for Response when Error

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