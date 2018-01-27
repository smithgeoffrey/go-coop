# OO

Go's approach is a little differentt:

- no inheritance
- composition

I've heard it touched upon in different ways:

- Not `is a __` but `has a __`
- Instead of building large trees of object types, create interfaces that describe desired behavior [1]

From my own experience, Go's answer to OO is that they give you, for any type, methods and interfaces.  Even less, I've been using just methods.  Structs already get me what feels like an object having properties. By adding methods, I feel mostly there.  Any gap should be fillable with interfaces, when I get there. I like the smallness of it in concept.

### Interfaces

Unlike methods, interfaces hadn't clicked with me much.  So I wanted here to dive deeper.  Maybe just using them will help.  Here's a list of breadcrumbs I looked at, trying to get it:

- @icza's answer in https://stackoverflow.com/questions/39092925/why-are-interfaces-needed-in-golang
- https://npf.io/2014/05/intro-to-go-interfaces/
- https://golang.org/doc/effective_go.html#interfaces_and_types

### References

[1] `Go in Practice`, p130 by Butcher & Farina