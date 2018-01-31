# OO

Go's approach is a little different:

- no inheritance
- composition

I've heard it touched upon in different ways:

- Not `is a __` but `has a __`
- Instead of building large trees of object types, create interfaces that describe desired behavior [1]

From my own experience, Go's answer to OO is that they give you, for any type, methods and interfaces.  Even less, I've been using just methods.  Structs already get me what feels like an object having properties. By adding methods, I feel mostly there.  Any gap should be fillable with interfaces, when I get there. I like the smallness of it in concept.

### Interfaces

Unlike methods, interfaces hadn't clicked with me as well.  So I wanted here to dive deeper.  Maybe just using them will help.  Here's a list of breadcrumbs I looked at, trying better to get it:

- @icza's answer in https://stackoverflow.com/questions/39092925/why-are-interfaces-needed-in-golang
- https://npf.io/2014/05/intro-to-go-interfaces/
- https://golangbot.com/interfaces-part-1/ & -2/

### Type Assertions & Type Switches

- https://golangbot.com/interfaces-part-1/
- https://medium.com/golangspec/type-assertions-in-go-e609759c42e1
- https://newfivefour.com/golang-interface-type-assertions-switch.html

When working with interfaces, it's handy to check type. You can do so a few ways:

- fmt.Printf("Type: %T; Value: %v", foo, foo)       // %T in string formatting
- reflect.TypeOf(x) or reflect.TypeOf(x).kind()     // reflect package
- someInterface.(type) in switch statements         // only for switch statements against interfaces

Using type checking you can verify the concept of an interface's `concrete` type and value:

    For a given interface Foo and type Bar that implements it, the concrete
    type of the inteface is Bar with a concrete value being Bar's. There's

Type assertion uses `i.(T)` where T is a type:

    // returns the concrete value if of that type, else panic
    someInterface.(int|string|float64)
    
    // avoid panic, return the concrete value and ok=true if of that type, else
    // ok=false and zero value of that type 
    v, ok := someInterface.(int|string|float64)

Type switch is similar and uses `i.(type)`:

    switch someInterface.(type) {
    case string:
        fmt.Printf("string value: %s\n", i.(string))
    case int:
        fmt.Printf("int value: %d\n", i.(int))
    default:
        fmt.Printf("Type unknown\n")
    }

There we compare the interface's concrete type against known ones and use type assertion once matched.  A similar pattern is an empty interface argument to a function: you can compare a type with the interface it implements then call a method within the interface's signature:

    
    func foo(i interface{}) {
        switch t := i.(type) {
        case someInteface:
            t.Method1()
        default:
            fmt.Printf("Type unknown\n")
        }
    }

### Pointer Receivers & Value Receivers

A type can can implement interfaces with methods having receivers that are pointer or value based:

    // interface calling for method1
    type MyInterface interface {
        method1(arg1) (string, error)
    }
    
    type MyType1 struct {
        foo int
        bar string
    }
        
    type MyType2 struct {
        a int
        b string
    }
    
    // value receiver
    func (r MyType1) method1(arg1) (string, error) {
        <insert>
    }
    
    // pointer receiver 
    func (r *MyType2) method1(arg1) (string, error) {
        <insert>
    }

An Interface Rule is that the concrete value stored in an interface is not addressable.  That complicates the usage of interfaces when implemented using a pointer receiver:

    // value receivers mean interface can be assigned the type or the &type
    // and the complier can render the method for either
    var i1 MyInterface
    t1 := MyType1{"32", "green"}
    t2 := MyType1{"23", "eggs"}
    i1 = t1
    i1.method1(arg1)
    i1 = &t2
    i1.method1(arg1)
        
    // pointer receivers mean interface can be assigned only the &type not the type
    var i2 MyInterface
    t1 := MyType2{"23", "eggs"}
    //i2 = t1  // fails
    i2 = &t1   // works
    i2.method1(arg1)
    
You can call pointer-receiver methods on things (i) that are already a pointer or (ii) whose address can be determined. Yet the Interface Rule says the concrete value of an interface is not addressable, so we are limited to assigning the &t1, meaning only option (i) is available.

### References

[1] `Go in Practice`, p130 by Butcher & Farina
[2] https://golangbot.com/interfaces-part-2/
