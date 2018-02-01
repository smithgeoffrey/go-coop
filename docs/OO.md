# OO

Go's approach is a little different:

- Not inheritance but composition
- Not `is a __` but `has a __`
- Instead of building large trees of object types, create interfaces that describe desired behavior [1]

From my own experience, Go's answer to OO is that they give you, for any type, methods and interfaces.  Even less, I'd been using just methods.  Structs alone got me what felt like an object having properties. By adding methods, I was mostly there.  But my gut kept telling me that any real OO in Go would require my becoming proficient with interfaces.

I had been ignoring interfaces.  They hadn't clicked with me, maybe because I simply wasn't using them.  So I started reading:

- `The Go Programming Language` at Ch. 7 by Donovan and Kernighan
- @icza's answer in https://stackoverflow.com/questions/39092925/why-are-interfaces-needed-in-golang
- https://npf.io/2014/05/intro-to-go-interfaces/
- https://golangbot.com/interfaces-part-1/ & -2/
- https://medium.com/golangspec/type-assertions-in-go-e609759c42e1
- https://newfivefour.com/golang-interface-type-assertions-switch.html

### Assignment

At the end of the day, types get assigned to interfaces.  You can assign if the type has the methods of the interface.  Only those methods will be callable on the interface, even if the concrete type has more.  The interface stores a type and its value.

    var w io.Writer
    w = os.Stdout

There it is said that w is storing a type *os.File having some value, which wasn't obvious to me.  It looks like os.Stdout is just a var Stdout in the os package, where the var is assigned `NewFile(uintptr(syscall.Stdout), "/dev/stdout")`.  See https://golang.org/src/os/file.go.  In turn, NewFile() I found at https://golang.org/src/os/file_unix.go, being a function returning a *File.  All of which is to say, indeed, `w = os.Stdout` does appear to store a type of *os.File.  Phew.

So in this one assignment, the interface is assuming a dynamic type *os.File and a dynamic value being what? I think it's the dereferenced value of *os.File, in this case, /dev/stdout.  So calling the Write() method of the interface will print to stdout.

    w.Write([]byte("cluck cluck"))

And just to be thorough, having to pass a slice of byte also isn't obvious.  So I looked at io.Writer at https://golang.org/src/io/io.go and saw this, which is enough for me for now:

    type Writer interface {
        Write(p []byte) (n int, err error)
    }

Finally, to thoroughly take this to the end, I might have thought instantiating a []byte would be more like `[]byte{foo}` not `[]byte(foo)`.  It looks like my thought is right, but that the () is syntactic sugar to allow entry of a string literal that is human readable, which the compiler converts at runtime to a non-human-unreadable {}.  See, e.g., https://stackoverflow.com/questions/25691879/creating-a-byte-slice-with-a-known-text-string-in-golang.  

### Type Assertions & Type Switches

When working with interfaces, it's handy to check type. You can do so a few ways:

- fmt.Printf("Type: %T; Value: %v", foo, foo)       // %T in string formatting
- reflect.TypeOf(x) or reflect.TypeOf(x).kind()     // reflect package
- someInterface.(type) in switch statements         // only for switch statements against interfaces

Using type checking you can verify the concept of an interface's `concrete` type and value:

    For a given interface Foo and type Bar that implements it, the concrete
    type of the inteface is Bar with a concrete value being Bar's.

Type assertion uses `i.(T)` where T is a type:

    // returns the concrete value but panics if 
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
    
You can call pointer-receiver methods on things (i) that are already a pointer or (ii) whose address can be determined. Yet the Interface Rule says the concrete value of an interface is not addressable, so we are limited to assigning the &t1, meaning only option (i) is available because the interface was implemented using a pointer-recevier.  Compare to the former case where the implementation was a value-receiver, in which case the addressibility is possible, hence both options are available.

### References

[1] `Go in Practice`, p130 by Butcher & Farina
[2] https://golangbot.com/interfaces-part-2/
