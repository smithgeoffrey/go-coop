# OO

Go's approach is a little differentt:

- no inheritance
- composition

I've heard it touched upon in different ways:

- Not `is a __` but `has a __`
- "Instead of building large trees of object types, create interfaces that describe desired behavior." [1]

From my own experience, Go's answer to OO is that they give you, for any type, methods and interfaces.  Even less, I've been using just methods.  Structs already get me what feels like an object having properties. By adding methods, I feel mostly there.  Any gap should be fillable with interfaces, when I get there. I like the smallness of it in concept.

### References

[1] `Go in Practice`, p130 by Butcher & Farina
