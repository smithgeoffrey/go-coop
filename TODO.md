# TODOs

### Panics

Panics unlike errors should be used sparingly and with intention.  Here I'll start by ignoring panics in favor of error handling.  After I've laid down the error handling and have a sense of finality with it, I'll loop back and consider injecting any panic's that may make sense.

### Errors

I'm learning that Go unlike other languages doesn't place much importance on the overhead of custom error types: most errors have no special attributes that would be better conveyed by a special error type. This jives with some of my past experiences, and until I stumble on cause to expand, I'll just use the std lib's base error type.

So far I'm ignoring errors largely. Try modifying everything to return multiple values: result, err.  

One guideline I've seen and want to try is don't return nil on the result when there's an error, instead return if possible the empty value of the type expected.  That enables users of your library to streamline their use of your lib to do things like this pseudo code:

    ## geoff package that returns and error
    import errors
    
    func Chickens(arg1) (string, error) {
        ...
        
        // error case
        return "", errors.New("snafu happened")
    }

    ## users of geoff package have the option to ignore 
    ## the error via `_` and work off the result instead 
    ## as desired
    import geoff.Chickens
    
    res, _ := geoff.Chickens(arg1)

Another guideline I've seen and want to try is, when handling errors, to collapse the underlying call that is to be error handled into the error handling conditional itself.  It leverages Go's if statement having an optional assignment clause before the expression.  Something like:
 
    if res, err := geoff.Chickens(arg1); err != nil {
        fmt.Printf("Error: %s\n", err)
    }
    
Compare to how I had generally thought of it, below.  

    res, err := geoff.Chickens(arg1)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
    }

### Testing

See ~/test/README.md, a major todo.