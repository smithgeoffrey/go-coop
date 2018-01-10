# OO

Go's approach to OO sets it apart:

- no inheritance
- composition

I've heard it touched upon in different ways:

- Not `is a __` but `has a __`
- "Instead of building large trees of object types, the Go developer creates interfaces that describe desired behavior." [1]

From my own (noob) experience, it seems like Go's answer to OO is that they give you methods and interfaces.  I've done OK so far by using just methods: structs already get me what feels like an object having properties, and by adding methods to the object, I feel mostly there.  While I can't say I'm right, I like the simplicity or smallness of it.  If I need more, surely interfaces will get me there.

### References

[1] `Go in Practice`, p130 by Butcher & Farina
