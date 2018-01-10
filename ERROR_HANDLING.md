# ERROR HANDLING

There's a few things about Go's error handling that sets it apart:
  
- no try-catch
- no heavy focus on custom error types

### Ideas

One guideline I've seen and want to try is don't return nil on the result when there's an error, instead return if possible the empty value of the type expected.  That enables users of your library to streamline their use of your lib to do things like this pseudo code:

    ## geoff package that returns an error
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