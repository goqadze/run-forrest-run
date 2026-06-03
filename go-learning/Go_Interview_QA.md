# Go / Golang Interview Questions & Answers

---

## 1. Go Basics

**1. What is Go?**
Go (also called Golang) is a statically typed, compiled programming language designed at Google. It focuses on simplicity, performance, and built-in concurrency support.

**2. Why was Go created?**
Google built Go to address slow build times, dependency management issues, and lack of good concurrency in C++ and Java. The goal was a language that is simple, fast to compile, and easy to write concurrent code in.

**3. What are the main advantages of Go?**
Fast compilation, simple syntax, built-in concurrency (goroutines + channels), strong standard library, automatic garbage collection, and static typing with type inference.

**4. What type of language is Go: compiled or interpreted?**
Compiled. Go code is compiled directly to native machine code, which makes it fast at runtime.

**5. Is Go object-oriented?**
Not in the traditional sense. Go has no classes or inheritance. Instead it uses structs with methods and interfaces, which achieves similar goals without the complexity of classical OOP.

**6. What is the difference between Go and Java?**
Go compiles to native code (no JVM), has no classes/inheritance, uses goroutines instead of threads, has simpler generics, and lacks exceptions. Go is generally simpler and faster to compile.

**7. What is the difference between Go and Python?**
Go is statically typed and compiled — it catches type errors at compile time and is much faster at runtime. Python is dynamically typed and interpreted. Go has explicit error handling; Python uses exceptions.

**8. What is the difference between Go and Node.js?**
Go is compiled and statically typed. Node.js is JavaScript running on V8. Go handles concurrency via goroutines; Node.js uses an event loop. Go is generally better for CPU-bound and high-concurrency backend work.

**9. What is the difference between `var`, `:=`, and `const`?**
- `var` declares a variable with explicit type or zero value: `var x int`
- `:=` is shorthand declaration+assignment inside functions: `x := 5`
- `const` declares a compile-time constant that cannot be changed: `const Pi = 3.14`

**10. Can you use `:=` outside a function?**
No. `:=` is only valid inside functions. At package level, use `var`.

**11. What are zero values in Go?**
When a variable is declared without initialization, Go sets it to its zero value automatically. This prevents uninitialized memory bugs.

**12. What is the zero value of `int`, `string`, `bool`, pointer, slice, map, channel?**
- `int` → `0`
- `string` → `""`
- `bool` → `false`
- pointer → `nil`
- slice → `nil`
- map → `nil`
- channel → `nil`

**13. What is `package main`?**
It marks the package as an executable program. Every Go program that produces a binary must have `package main` with a `main()` function.

**14. What is the purpose of the `main()` function?**
It is the entry point of the program. Execution starts here when the binary runs.

**15. What is the difference between `fmt.Print`, `fmt.Println`, and `fmt.Printf`?**
- `Print` writes values with no formatting and no newline
- `Println` adds spaces between values and a newline at the end
- `Printf` uses format verbs like `%s`, `%d`, `%v` — no automatic newline

**16. What is `go fmt`?**
A tool that automatically formats Go code to the standard style. All Go code is expected to be `go fmt`-compliant.

**17. What is `go mod`?**
The module system for managing dependencies. `go mod init` creates a module, `go mod tidy` cleans up unused dependencies.

**18. What is the difference between `go run`, `go build`, and `go install`?**
- `go run` compiles and runs in one step (no output binary saved)
- `go build` compiles and saves the binary to the current directory
- `go install` compiles and installs the binary to `$GOPATH/bin`

**19. What is the purpose of `init()`?**
A special function that runs automatically before `main()`. Used for package-level setup like registering drivers or initializing globals.

**20. Can a package have multiple `init()` functions?**
Yes. A single file can even have multiple `init()` functions. They all run in the order they appear.

---

## 2. Variables, Types, and Constants

**1. What are basic data types in Go?**
`bool`, `string`, integer types (`int`, `int8`, `int16`, `int32`, `int64`, `uint`, etc.), float types (`float32`, `float64`), `complex64`, `complex128`, `byte` (alias for `uint8`), and `rune` (alias for `int32`).

**2. What is the difference between `int`, `int32`, and `int64`?**
`int` is platform-dependent — 32-bit on 32-bit systems, 64-bit on 64-bit systems. `int32` is always 32 bits, `int64` is always 64 bits. Use `int` for general use unless you need a specific size.

**3. What is `rune`?**
An alias for `int32`. It represents a Unicode code point. Use `rune` when working with individual characters in a string.

**4. What is `byte`?**
An alias for `uint8`. Represents a single byte of data. Strings in Go are sequences of bytes.

**5. What is the difference between `byte` and `rune`?**
`byte` is one byte (ASCII range). `rune` is 4 bytes and can represent any Unicode character. `'A'` fits in a byte, `'გ'` (Georgian) requires a rune.

**6. What is the difference between single quotes and double quotes in Go?**
Double quotes `"hello"` create a `string`. Single quotes `'a'` create a `rune` (character literal). There are no single-quoted strings in Go.

**7. What is a raw string literal?**
A string enclosed in backticks: `` `hello\nworld` ``. Backslashes are not interpreted, and the string can span multiple lines.

**8. What is type inference?**
Go can infer the type from the assigned value: `x := 42` — Go knows `x` is an `int` without you specifying it.

**9. What is type conversion?**
Explicitly converting one type to another: `float64(x)` or `int(f)`. Go never does this implicitly.

**10. Does Go support implicit type conversion?**
No. You must always convert explicitly. `var x int = 3.14` is a compile error.

**11. What is `iota`?**
A predeclared identifier used in `const` blocks. It starts at 0 and increments by 1 for each constant in the block. Used to define enumerations.

**12. How do you define enums in Go?**
Go has no `enum` keyword. Use `const` blocks with `iota`:
```go
type Direction int
const (
    North Direction = iota // 0
    South                  // 1
    East                   // 2
    West                   // 3
)
```

**13. What are typed and untyped constants?**
Typed: `const x int = 5` — fixed to `int`. Untyped: `const x = 5` — flexible, can be used wherever a compatible numeric type is needed.

**14. What happens when an integer overflows?**
For fixed-size types like `int32`, it wraps around silently (no panic). This can cause subtle bugs. Go does not check for overflow at runtime.

**15. What is the difference between `float32` and `float64`?**
`float64` has more precision (15-17 significant digits vs 6-9 for `float32`). Always prefer `float64` unless memory is a real concern.

**16. How does Go handle strings internally?**
A string is an immutable sequence of bytes (not characters). Internally it's a struct with a pointer to byte data and a length.

**17. Are strings mutable in Go?**
No. Strings are immutable. To manipulate them, convert to `[]byte` or `[]rune`, modify, then convert back.

**18. How do you convert a string to `[]byte`?**
```go
b := []byte("hello")
```

**19. How do you convert a string to `[]rune`?**
```go
r := []rune("გამარჯობა")
```

**20. Why is `len("hello")` different from `len("გამარჯობა")`?**
`len()` returns the number of **bytes**, not characters. `"hello"` is 5 bytes. `"გამარჯობა"` has 9 Georgian characters but each takes 3 bytes in UTF-8, so `len` returns 27. Use `len([]rune(s))` to count characters.

---

## 3. Control Flow

**1. How does `if` work in Go?**
Like most languages, but no parentheses around the condition:
```go
if x > 0 {
    fmt.Println("positive")
}
```

**2. Can you declare a variable inside an `if` statement?**
Yes. The variable is scoped to the `if` block:
```go
if err := doSomething(); err != nil {
    return err
}
```

**3. How does `switch` work in Go?**
Evaluates cases top to bottom and stops at the first match. No need for `break`.
```go
switch x {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
}
```

**4. Does Go require `break` in `switch`?**
No. Cases do NOT fall through by default — `break` is implicit. This is the opposite of C/Java.

**5. What is `fallthrough`?**
Explicitly continues to the next case even after a match:
```go
case 1:
    fmt.Println("one")
    fallthrough
case 2:
    fmt.Println("two") // also runs if case 1 matched
```

**6. What is the difference between `for` and `while` in Go?**
Go has no `while` keyword. `for` serves all looping needs.

**7. Does Go have a `while` loop?**
No, but you write one using `for`:
```go
for condition {
    // same as while in other languages
}
```

**8. How do you create an infinite loop?**
```go
for {
    // runs forever
}
```

**9. What is `range`?**
A keyword for iterating over arrays, slices, maps, strings, and channels. Returns index/key and value.

**10. What does `range` return for arrays and slices?**
Index and value: `for i, v := range slice { ... }`

**11. What does `range` return for maps?**
Key and value: `for k, v := range m { ... }`

**12. What does `range` return for strings?**
Byte index and `rune` (Unicode code point), not byte. This handles multi-byte characters correctly.

**13. What does `range` return for channels?**
Only values, one at a time, until the channel is closed: `for v := range ch { ... }`

**14. What is the difference between `break` and `continue`?**
`break` exits the loop entirely. `continue` skips the rest of the current iteration and goes to the next.

**15. What are labeled breaks?**
Labels let you `break` or `continue` an outer loop from inside a nested loop:
```go
outer:
for i := range rows {
    for j := range cols {
        if done {
            break outer
        }
    }
}
```

**16. When would you use a labeled `break`?**
When you need to exit multiple nested loops at once from a condition deep inside them.

**17. Can you use `goto` in Go?**
Yes, Go has `goto`. It jumps to a labeled statement within the same function.

**18. Should you use `goto` in production code?**
Almost never. `goto` makes code hard to follow. Use structured control flow (`break`, `return`, functions) instead.

---

## 4. Arrays and Slices

**1. What is an array in Go?**
A fixed-size sequence of elements of the same type. The size is part of the type: `[3]int` and `[5]int` are different types.

**2. What is a slice?**
A dynamic, flexible view over an underlying array. Most Go code uses slices, not arrays directly.

**3. What is the difference between an array and a slice?**
Arrays have fixed size and are value types (copied when assigned). Slices are references to an underlying array, have dynamic length, and are much more common.

**4. Are arrays passed by value or by reference?**
By value. The entire array is copied. Changes to the copy don't affect the original.

**5. Are slices passed by value or by reference?**
The slice header (pointer, length, capacity) is copied by value, but the underlying array is shared. So modifications to elements affect the original, but `append` beyond capacity won't.

**6. What does a slice contain internally?**
Three fields: a pointer to the underlying array, a length (number of elements), and a capacity (total space in the array from the pointer onward).

**7. What are `len` and `cap`?**
`len` returns the number of elements in the slice. `cap` returns the total capacity from the slice's start to the end of the underlying array.

**8. What is the difference between length and capacity?**
Length is how many elements are currently in the slice. Capacity is how many can be added before a new array is allocated.

**9. How does `append` work?**
Adds elements to a slice. If there's enough capacity, it adds in-place. If not, it allocates a new, larger array, copies data, and returns the new slice.

**10. Can `append` modify the original array?**
If there's remaining capacity, `append` writes into the underlying array, which can affect other slices sharing it. If a new array is allocated, the original is untouched.

**11. When does `append` allocate a new backing array?**
When `len == cap` — there's no room left. Go typically doubles capacity for small slices.

**12. What happens when slice capacity is exceeded?**
`append` allocates a new, larger backing array, copies all elements over, and returns a new slice pointing to it.

**13. How do you copy a slice?**
```go
dst := make([]int, len(src))
copy(dst, src)
```

**14. What is the difference between `copy` and `append`?**
`copy` copies into an existing slice (won't grow it). `append` adds to a slice and can allocate a new one.

**15. How do you remove an element from a slice?**
```go
s = append(s[:i], s[i+1:]...)
```
Note: this modifies the original slice's backing array.

**16. How do you insert an element into a slice?**
```go
s = append(s[:i], append([]T{v}, s[i:]...)...)
```

**17. How do you reverse a slice?**
```go
for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
}
```

**18. How do you avoid memory leaks with large slices?**
When sub-slicing a large slice, the backing array is still held in memory. Copy needed elements to a new slice instead of keeping a sub-slice reference.

**19. What is a nil slice?**
A slice declared but not initialized: `var s []int`. Its pointer, length, and capacity are all zero.

**20. What is an empty slice?**
A slice with length 0 but a non-nil pointer: `s := []int{}` or `s := make([]int, 0)`.

**21. What is the difference between `nil` slice and empty slice?**
A nil slice has no backing array at all. An empty slice has a backing array but zero elements. Both have `len == 0` and both work with `append`.

**22. How does JSON encode nil slice vs empty slice?**
- nil slice → `null`
- empty slice → `[]`
This matters in APIs.

**23. Why can modifying a sub-slice affect the original slice?**
Sub-slices share the same backing array. Writing to `b[0]` where `b = a[1:3]` writes to `a[1]`.

**24. What is slice aliasing?**
Two slices pointing to the same backing array. Changes via one slice are visible through the other.

**25. How do you clone a slice safely?**
```go
clone := append([]T(nil), original...)
// or
clone := make([]T, len(original))
copy(clone, original)
```

**Tricky question output:**
```go
a := []int{1, 2, 3}
b := a[:2]
b = append(b, 10)
```
Output: `a = [1 2 10]`, `b = [1 2 10]`
Because `b` still has capacity in `a`'s array, `append` writes `10` into position 2 of the same backing array, overwriting `a[2]`.

---

## 5. Maps

**1. What is a map in Go?**
An unordered collection of key-value pairs. Keys must be comparable types. Implemented as a hash table.

**2. How do you create a map?**
```go
m := make(map[string]int)
// or with a literal:
m := map[string]int{"a": 1}
```

**3. What is the zero value of a map?**
`nil`. A nil map cannot be written to.

**4. Can you write to a nil map?**
No. It panics: `assignment to entry in nil map`. Always initialize with `make`.

**5. Can you read from a nil map?**
Yes. Reading from a nil map returns the zero value for the value type — it does not panic.

**6. How do you check if a key exists?**
Use the two-value form:
```go
value, ok := m["key"]
if ok {
    // key exists
}
```

**7. What does `value, ok := m["key"]` mean?**
`value` is the value for that key (or zero value if absent). `ok` is `true` if the key exists, `false` if not.

**8. Are maps ordered?**
No. Maps are unordered. Iteration order is intentionally randomized in Go to prevent code from accidentally depending on it.

**9. Is map iteration order guaranteed?**
No. Every iteration may produce a different order. Sort keys explicitly if you need order.

**10. Are maps safe for concurrent use?**
No. Concurrent reads are safe, but concurrent read+write or write+write causes a race condition and will panic in Go 1.6+ (fatal error, not a regular panic).

**11. What happens if multiple goroutines write to the same map?**
It causes a data race. The Go runtime detects concurrent map writes and panics with "concurrent map writes".

**12. How do you delete from a map?**
```go
delete(m, "key")
```
Deleting a non-existent key is a no-op.

**13. Can a map key be a slice?**
No. Slices are not comparable, so they can't be map keys.

**14. What types can be map keys?**
Any comparable type: `bool`, `int`, `float`, `string`, pointer, array, struct (if all fields are comparable). Not slices, maps, or functions.

**15. Why must map keys be comparable?**
The hash table implementation needs to check equality between keys. Types that can't be compared with `==` can't be hashed reliably.

**16. What is the difference between `map[string]int{}` and `make(map[string]int)`?**
Both create an empty, initialized map ready to use. They are functionally equivalent. `make` allows specifying an initial capacity hint: `make(map[string]int, 100)`.

**17. How do you initialize a nested map?**
```go
m := map[string]map[string]int{}
m["outer"] = map[string]int{"inner": 1}
```
Each nested map must also be initialized before writing.

**18. How do you copy a map?**
There's no built-in copy. Iterate and copy manually:
```go
copy := make(map[string]int)
for k, v := range original {
    copy[k] = v
}
```

**19. What is `sync.Map`?**
A concurrency-safe map from the `sync` package. It allows safe concurrent reads and writes without external locking.

**20. When should you use `sync.Map`?**
When the map is read-heavy with infrequent writes, or when keys are only written once and read many times. For most cases, a regular map with a `sync.Mutex` or `sync.RWMutex` is simpler and clearer.
---

## 6. Structs

**1. What is a struct?**
A composite type that groups named fields together. It's the main way to define data structures in Go.

**2. How do you define a struct?**
```go
type User struct {
    Name string
    Age  int
}
```

**3. How do you initialize a struct?**
```go
u := User{Name: "Alice", Age: 30} // named fields (preferred)
u := User{"Alice", 30}            // positional (fragile)
u := User{}                       // zero values
var u User                        // zero values
```

**4. What is an anonymous struct?**
A struct without a named type, defined and used inline:
```go
p := struct{ X, Y int }{X: 1, Y: 2}
```
Useful for one-off data grouping, like in tests.

**5. What is struct embedding?**
Including one struct type inside another without a field name. The outer struct "inherits" the embedded struct's fields and methods:
```go
type Animal struct{ Name string }
type Dog struct{ Animal; Breed string }
d := Dog{}
d.Name = "Rex" // promoted from Animal
```

**6. Is embedding inheritance?**
No. It's composition. There is no polymorphism or subtyping. The embedded struct's methods are promoted, but you can't pass a `Dog` where an `Animal` is expected.

**7. What are struct tags?**
Metadata attached to struct fields using backtick strings:
```go
type User struct {
    Name string `json:"name" db:"user_name"`
}
```
Read at runtime via reflection.

**8. How are struct tags used in JSON?**
The `encoding/json` package reads `json:"..."` tags to control marshaling. The tag specifies the JSON key name.

**9. What is the difference between exported and unexported fields?**
Fields starting with uppercase are exported (visible outside the package). Lowercase fields are unexported (package-private).

**10. What happens if a JSON field is unexported?**
It is ignored during JSON marshaling and unmarshaling.

**11. How do you omit empty fields in JSON?**
Add `omitempty` to the tag:
```go
Name string `json:"name,omitempty"`
```
The field is omitted if it holds its zero value.

**12. What does `json:"-"` mean?**
The field is always excluded from JSON output and input, regardless of its value.

**13. How do you compare structs?**
With `==` if all fields are comparable:
```go
u1 == u2 // true if all fields are equal
```

**14. When are structs comparable?**
When all their fields are comparable types (no slices, maps, or functions).

**15. Can a struct contain methods?**
Methods are defined on types, not inside structs. But you define methods "on" a struct type:
```go
func (u User) Greet() string { return "Hello, " + u.Name }
```

**16. Should you pass a struct by value or pointer?**
By pointer (`*User`) when the struct is large, when you need to modify it, or for consistency. By value for small, simple structs that shouldn't be modified.

**17. What is the difference between value receiver and pointer receiver?**
Value receiver gets a copy: `func (u User) Greet()` — changes don't affect the original. Pointer receiver gets the actual struct: `func (u *User) SetName()` — changes persist.

**18. Can embedded structs have method promotion?**
Yes. Methods of the embedded struct are promoted to the outer struct and can be called directly.

**19. What happens if embedded structs have conflicting method names?**
The outer struct's method takes priority. If two embedded structs have the same method name, it causes a compile error when you try to call it — you must qualify it: `d.Animal.Speak()`.

**20. How do you validate struct fields?**
Manually in code, or using a validation library like `github.com/go-playground/validator` that reads struct tags.

---

## 7. Pointers

**1. What is a pointer?**
A variable that stores the memory address of another variable.

**2. What is the zero value of a pointer?**
`nil` — it points to nothing.

**3. What is `nil`?**
The zero value for pointers, interfaces, maps, slices, channels, and functions. It means "no value" or "uninitialized".

**4. What does `&` do?**
Returns the address of a variable: `p := &x` means `p` holds the address of `x`.

**5. What does `*` do?**
Dereferences a pointer — accesses the value at the address: `*p = 10` sets the value `p` points to.

**6. Can Go do pointer arithmetic?**
No. Go does not allow pointer arithmetic like C/C++.

**7. Why does Go not allow pointer arithmetic?**
For memory safety. Without pointer arithmetic, Go can guarantee programs don't access arbitrary memory locations, preventing a large class of bugs.

**8. What is the difference between passing by value and passing by pointer?**
By value: a copy is made, the original is unaffected. By pointer: you pass the address, so the function can modify the original.

**9. When should you use pointers?**
When you need to modify the original, when the struct is large (avoid expensive copies), or to express "optional" values (a nil pointer means absent).

**10. When should you avoid pointers?**
For small, simple values (`int`, `bool`, small structs) where a copy is cheap. Unnecessary pointers add GC pressure and indirection.

**11. Can a pointer point to another pointer?**
Yes: `var pp **int` is a pointer to a pointer to an int. Rarely needed in practice.

**12. What is a nil pointer dereference?**
Accessing a value through a nil pointer: `var p *int; _ = *p` causes a panic at runtime.

**13. How do you avoid nil pointer panic?**
Always check if a pointer is nil before dereferencing:
```go
if p != nil {
    fmt.Println(*p)
}
```

**14. What is escape analysis?**
A compile-time process where Go decides whether a variable should be allocated on the stack or the heap. If a variable's address is returned or stored outside the function, it "escapes" to the heap.

**15. What does it mean when a variable escapes to the heap?**
It's allocated on the heap instead of the stack, which means the GC is responsible for cleaning it up. Heap allocations are slower than stack allocations.

**16. Are local variables always allocated on the stack?**
No. If they escape (e.g., their address is returned), they go to the heap. Escape analysis determines this.

**17. How can pointers affect garbage collection?**
Every pointer to heap data keeps that data alive. Holding pointers longer than necessary prevents GC from freeing memory.

**18. What is the difference between pointer receiver and value receiver?**
Pointer receiver: `func (u *User)` — operates on the actual struct, can modify it. Value receiver: `func (u User)` — operates on a copy.

**19. Can you take the address of a map element?**
No. `&m["key"]` is a compile error. Map values are not addressable because the map may reallocate internally.

**20. Why can't you take the address of a map value directly?**
The map's internal hash table can be reorganized at any time, which would invalidate the address. Go prevents this at compile time.

---

## 8. Functions

**1. How do you define a function in Go?**
```go
func add(a, b int) int {
    return a + b
}
```

**2. Can functions return multiple values?**
Yes. This is idiomatic in Go, especially for returning a result and an error:
```go
func divide(a, b float64) (float64, error) { ... }
```

**3. What are named return values?**
Return values can be named in the signature:
```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // "naked return"
}
```

**4. Should you use named return values?**
Sparingly. Useful for short functions and for `defer`-modified returns. In longer functions they reduce clarity.

**5. What is a variadic function?**
A function that accepts a variable number of arguments:
```go
func sum(nums ...int) int { ... }
sum(1, 2, 3)
```

**6. What does `...` mean?**
In a function parameter: accepts zero or more arguments of that type. When calling: spreads a slice into individual arguments: `sum(nums...)`.

**7. What is a closure?**
A function that captures and uses variables from its enclosing scope:
```go
func counter() func() int {
    n := 0
    return func() int { n++; return n }
}
```

**8. What is an anonymous function?**
A function without a name, defined inline:
```go
f := func(x int) int { return x * 2 }
```

**9. Can functions be passed as arguments?**
Yes. Functions are first-class values in Go:
```go
func apply(f func(int) int, x int) int { return f(x) }
```

**10. Can functions be returned from other functions?**
Yes. This enables closures, middleware, and factory patterns.

**11. What is a higher-order function?**
A function that takes a function as an argument or returns a function.

**12. What is recursion?**
A function that calls itself. Must have a base case to stop.

**13. Does Go optimize tail recursion?**
No. Go does not perform tail call optimization. Deep recursion can overflow the goroutine stack.

**14. What is `defer`?**
Schedules a function call to run when the surrounding function returns:
```go
defer f.Close() // runs when the function exits, regardless of how
```

**15. When is `defer` executed?**
Just before the surrounding function returns — whether it returns normally, early, or due to a panic.

**16. Are deferred calls executed in FIFO or LIFO order?**
LIFO (last in, first out). The last `defer` registered runs first.

**17. What is the cost of `defer`?**
There is a small overhead. In extremely tight loops this matters, but for most code it's negligible.

**18. Can deferred functions modify named return values?**
Yes. A deferred function can change named return values because it runs after the return expression is evaluated but before the function actually returns.

**19. What happens if you call `defer` inside a loop?**
Each iteration registers a new deferred call. They all run when the function returns, not at the end of each iteration. For many iterations this can be a problem.

**20. How do you recover from panic?**
Use `recover()` inside a deferred function:
```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("recovered:", r)
    }
}()
```

---

## 9. Methods and Interfaces

**1. What is a method in Go?**
A function with a receiver — it's associated with a specific type:
```go
func (u User) Greet() string { return "Hello, " + u.Name }
```

**2. What is a receiver?**
The type a method is defined on. It appears between `func` and the method name: `(u User)` or `(u *User)`.

**3. What is the difference between value receiver and pointer receiver?**
Value receiver operates on a copy — changes don't persist. Pointer receiver operates on the original — changes persist. Use pointer receiver when you need to mutate the struct or when the struct is large.

**4. Can you define methods on built-in types?**
No. You can only define methods on types declared in the same package. You must create a new named type first: `type MyInt int`.

**5. Can you define methods on types from another package?**
No. You can't add methods to types you don't own. Use a wrapper type or a standalone function instead.

**6. What is an interface?**
A set of method signatures. Any type that implements all methods of the interface satisfies it — no explicit declaration needed.
```go
type Greeter interface {
    Greet() string
}
```

**7. How are interfaces implemented in Go?**
Implicitly. If a type has all the methods of an interface, it implements that interface automatically.

**8. Does Go have explicit `implements` keyword?**
No. Interface satisfaction is structural (duck typing), not nominal.

**9. What is implicit interface implementation?**
A type satisfies an interface simply by having the required methods — no `implements` declaration needed. This allows interfaces to be defined after the types that satisfy them.

**10. What is the empty interface?**
`interface{}` — an interface with no methods. Every type satisfies it. Used to hold any value.

**11. What is `any`?**
A type alias for `interface{}` introduced in Go 1.18. `any` and `interface{}` are identical.

**12. What is the difference between `interface{}` and `any`?**
None — `any` is just a more readable alias for `interface{}`.

**13. What is type assertion?**
Extracting the underlying concrete type from an interface:
```go
s, ok := i.(string) // ok is false if i doesn't hold a string
```

**14. What is type switch?**
A switch statement that checks the dynamic type of an interface value:
```go
switch v := i.(type) {
case string:  fmt.Println("string:", v)
case int:     fmt.Println("int:", v)
}
```

**15. What happens if type assertion fails?**
Without `ok`: panics. With two-value form (`v, ok := i.(T)`): `ok` is `false` and `v` is zero value — no panic.

**16. What is interface composition?**
Combining multiple interfaces into one:
```go
type ReadWriter interface {
    io.Reader
    io.Writer
}
```

**17. What is the difference between concrete type and interface type?**
A concrete type has a specific in-memory representation. An interface type is an abstraction — it holds a value of any type that implements its methods.

**18. Can an interface value be nil?**
Yes, but carefully. An interface is nil only when both its type and value are nil.

**19. What is the nil interface problem?**
An interface can be non-nil even if the underlying value is nil. The interface has a type but a nil pointer value, so `i == nil` is `false`.

**20. Why can this be confusing?**
```go
var p *User = nil
var i interface{} = p
fmt.Println(i == nil) // false!
```
`i` has type `*User` (non-nil type) even though the pointer is nil. The interface is not nil because its type field is set.

**21. What is duck typing?**
"If it walks like a duck and quacks like a duck, it's a duck." In Go: if a type has the right methods, it satisfies the interface — no explicit declaration.

**22. How do you design small interfaces?**
Define interfaces where they are used (at the consumer), not where types are defined. Keep them to 1-3 methods.

**23. Why does Go prefer small interfaces?**
Small interfaces are easier to implement, mock, and compose. They allow more types to satisfy the interface, improving flexibility.

**24. What is the `io.Reader` interface?**
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```
Anything that can read bytes. Files, network connections, buffers all implement it.

**25. What is the `io.Writer` interface?**
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
Anything that can accept bytes: files, HTTP responses, buffers.

**26. Why are `io.Reader` and `io.Writer` powerful abstractions?**
Code written against `io.Reader`/`io.Writer` works with any data source or sink — files, network, in-memory buffers, etc. — without changes.

**27. What is dependency injection using interfaces?**
Passing dependencies as interface types instead of concrete types. The function or struct doesn't know the concrete implementation — it only cares about the behavior.

**28. How do you mock dependencies in Go?**
Create a test struct that implements the same interface and returns controlled responses. No mocking library required.

**29. What is interface pollution?**
Defining too many unnecessary interfaces, especially before you know you need the abstraction. Go prefers interfaces to emerge from usage, not be pre-emptively defined.

**30. When should you return a concrete type instead of an interface?**
When there's only one implementation and you don't expect variability. Return the concrete type; the caller can always accept it as an interface.

---

## 10. Error Handling

**1. How does Go handle errors?**
Errors are plain values returned from functions. You check them explicitly: `if err != nil { ... }`.

**2. Why does Go not use exceptions?**
Exceptions create hidden control flow. Go's explicit error returns make error paths visible in the code, easier to handle, and harder to accidentally ignore.

**3. What is the `error` interface?**
```go
type error interface {
    Error() string
}
```
Any type with an `Error() string` method is an error.

**4. How do you create an error?**
```go
errors.New("something went wrong")
fmt.Errorf("user %d not found", id)
```

**5. What is `errors.New`?**
Creates a simple error with a static message. Two errors created with `errors.New` with the same message are not equal.

**6. What is `fmt.Errorf`?**
Creates a formatted error message. Use `%w` to wrap another error for inspection later.

**7. What is error wrapping?**
Adding context to an existing error while preserving the original:
```go
fmt.Errorf("fetching user: %w", err)
```

**8. What does `%w` do?**
Wraps an error so it can be unwrapped with `errors.Is` and `errors.As`. Without `%w`, the original error is lost.

**9. What is `errors.Is`?**
Checks if an error (or any error in its chain) matches a target error:
```go
if errors.Is(err, ErrNotFound) { ... }
```

**10. What is `errors.As`?**
Extracts a specific error type from the chain:
```go
var notFound *NotFoundError
if errors.As(err, &notFound) { ... }
```

**11. What is `errors.Unwrap`?**
Returns the next error in the chain. Usually you use `errors.Is`/`errors.As` instead of calling `Unwrap` directly.

**12. What is a sentinel error?**
A package-level error variable used for comparison:
```go
var ErrNotFound = errors.New("not found")
if err == ErrNotFound { ... } // or errors.Is(err, ErrNotFound)
```

**13. What is a custom error type?**
A struct that implements the `error` interface, allowing it to carry extra data:
```go
type ValidationError struct{ Field string; Message string }
func (e *ValidationError) Error() string { return e.Field + ": " + e.Message }
```

**14. When should you use custom errors?**
When the caller needs structured data from the error (field name, status code, etc.) beyond just a message string.

**15. When should you use sentinel errors?**
For simple, well-known conditions callers need to check (`ErrNotFound`, `io.EOF`). Keep them package-level and document them.

**16. What is the difference between panic and error?**
Errors are expected, handled conditions. Panic is for truly unexpected, unrecoverable situations (programmer bugs, not user input errors).

**17. When should you use panic?**
Only for unrecoverable programmer errors — invalid arguments in a library that "can't happen", or initialization failures. Not for normal control flow.

**18. When should you recover from panic?**
At the top-level of goroutines (e.g., HTTP handlers) to prevent one panic from crashing the entire server. Always convert recovered panics back to errors.

**19. Should business logic use panic?**
No. Business logic should return errors. Panic is not a substitute for error handling.

**20. How do you add context to errors?**
```go
return fmt.Errorf("creating user: %w", err)
```
This builds an error chain with context at each layer.

**21. How do you handle errors in HTTP handlers?**
Check the error, log it if needed, and write an appropriate HTTP status code and message back to the client. Don't expose internal error details to clients.

**22. How do you avoid ignoring errors?**
Always assign errors to a variable and check them. Enable a linter like `errcheck` to catch ignored errors in CI.

**23. Why is `_ = err` usually bad?**
It silently discards the error. Problems are hidden and debugging becomes very hard later.

**24. How do you handle multiple errors?**
Collect errors in a slice and use `errors.Join` (Go 1.20+) or a library like `multierr`.

**25. What is `errors.Join`?**
Combines multiple errors into one. `errors.Is` and `errors.As` work through joined errors.
```go
err = errors.Join(err1, err2)
```
---

## 11. Goroutines

**1. What is a goroutine?**
A lightweight, concurrently executing function managed by the Go runtime. Start one with `go fn()`.

**2. How is a goroutine different from a thread?**
Goroutines are much cheaper — they start at ~2KB of stack (vs ~1MB for OS threads), are multiplexed onto a small number of OS threads by the Go scheduler, and context-switching between them is faster.

**3. How do you start a goroutine?**
```go
go myFunction()
go func() { fmt.Println("concurrent") }()
```

**4. Are goroutines cheap?**
Yes. You can easily run thousands or millions of goroutines in a single program.

**5. How much memory does a goroutine need initially?**
About 2–8 KB for the initial stack. The stack grows dynamically as needed.

**6. Who schedules goroutines?**
The Go runtime scheduler, not the OS. It maps goroutines onto OS threads.

**7. What is the Go scheduler?**
A cooperative/preemptive scheduler built into the Go runtime that multiplexes goroutines onto OS threads (since Go 1.14, it can preempt goroutines at safe points).

**8. What is the GMP model?**
The model the Go scheduler is based on: G (goroutines), M (OS threads/machines), P (processors/schedulers). Each P runs one goroutine at a time on an M.

**9. What are G, M, and P?**
- **G**: A goroutine — the unit of work
- **M**: An OS thread that executes Go code
- **P**: A logical processor that holds a run queue of goroutines. `GOMAXPROCS` controls the number of Ps.

**10. What is cooperative scheduling?**
Goroutines yield control voluntarily at scheduling points (function calls, channel ops, etc.). Early Go used this; since Go 1.14, preemption is also supported.

**11. What is preemption in Go?**
Since Go 1.14, the scheduler can interrupt a goroutine at safe points even if it doesn't yield. This prevents goroutines from hogging the scheduler.

**12. Can a goroutine leak?**
Yes. A goroutine that is blocked forever (waiting on a channel that never receives, etc.) and never exits is a leaked goroutine.

**13. What is a goroutine leak?**
A goroutine that was started but never terminates, holding memory and resources forever. Common cause of memory issues in long-running services.

**14. How do you prevent goroutine leaks?**
Always give goroutines a way to exit: use `context.Context` for cancellation, or a done channel. Ensure goroutines don't block forever on channels.

**15. How do you wait for goroutines to finish?**
Use `sync.WaitGroup`:
```go
var wg sync.WaitGroup
wg.Add(1)
go func() { defer wg.Done(); doWork() }()
wg.Wait()
```

**16. What is `sync.WaitGroup`?**
A counter that blocks until it reaches zero. Call `Add(n)` before launching goroutines, `Done()` when each finishes, and `Wait()` to block until all are done.

**17. What happens if you call `Done()` too many times?**
The counter goes negative and panics: "sync: negative WaitGroup counter".

**18. What happens if you forget `Done()`?**
`Wait()` blocks forever — a deadlock.

**19. Can you copy a `WaitGroup`?**
No. Never copy a `WaitGroup` after first use. Always pass by pointer.

**20. How do you limit the number of goroutines?**
Use a buffered channel as a semaphore:
```go
sem := make(chan struct{}, 10) // max 10 concurrent
sem <- struct{}{}              // acquire
go func() { defer func() { <-sem }(); doWork() }()
```

**21. How do you implement a worker pool?**
Create N worker goroutines that read from a shared jobs channel. The main goroutine sends work to the channel; workers process and send results to an output channel.

**22. What is fan-out/fan-in?**
Fan-out: one goroutine distributes work to many goroutines. Fan-in: many goroutines send results back to one channel.

**23. What is a pipeline pattern?**
A series of stages connected by channels. Each stage reads from an input channel, processes data, and writes to an output channel.

**24. How do you stop a goroutine gracefully?**
Signal it via a done channel or `context.WithCancel`. The goroutine checks for the signal and exits cleanly.

**25. How do you handle errors from goroutines?**
Send errors back via an error channel, use `errgroup` from `golang.org/x/sync`, or collect errors in a shared variable protected by a mutex.

---

## 12. Channels

**1. What is a channel?**
A typed conduit for communicating between goroutines. Sends and receives synchronize goroutines.

**2. How do you create a channel?**
```go
ch := make(chan int)          // unbuffered
ch := make(chan int, 10)      // buffered, capacity 10
```

**3. What is an unbuffered channel?**
A channel with no buffer. A send blocks until a receiver is ready, and a receive blocks until a sender is ready. Both sides must be ready simultaneously.

**4. What is a buffered channel?**
A channel with a queue. Sends don't block until the buffer is full. Receives don't block while there are items in the buffer.

**5. What is the difference between buffered and unbuffered channels?**
Unbuffered: synchronizes sender and receiver (rendezvous). Buffered: decouples sender and receiver up to the buffer size.

**6. Does sending to an unbuffered channel block?**
Yes, until a receiver is ready.

**7. Does receiving from an unbuffered channel block?**
Yes, until a sender sends.

**8. When does a buffered channel block?**
Sending blocks when the buffer is full. Receiving blocks when the buffer is empty.

**9. What happens when you send to a closed channel?**
Panic: "send on closed channel". Never send to a closed channel.

**10. What happens when you receive from a closed channel?**
Returns the zero value immediately and the `ok` flag is `false`:
```go
v, ok := <-ch // ok is false if ch is closed and empty
```

**11. What happens when you close a nil channel?**
Panic: "close of nil channel".

**12. What happens when you send to a nil channel?**
Blocks forever.

**13. What happens when you receive from a nil channel?**
Blocks forever.

**14. Who should close a channel?**
Only the sender. Closing is a signal "no more values". Never close from the receiver side.

**15. Should receivers close channels?**
No. Only the sender should close. If multiple senders exist, coordinate with a WaitGroup to close after all senders are done.

**16. How do you range over a channel?**
```go
for v := range ch {
    // runs until ch is closed
}
```

**17. What does `close(ch)` do?**
Marks the channel as closed. Any pending and future receives return immediately with the zero value. Any sends panic.

**18. How do you check if a channel is closed?**
Use the two-value receive: `v, ok := <-ch`. If `ok` is `false`, the channel is closed.

**19. What is a directional channel?**
A channel restricted to only send or only receive. Used in function signatures to express intent.

**20. What is `chan<- int`?**
Send-only channel. You can only send to it: `ch <- 1`. Cannot receive from it.

**21. What is `<-chan int`?**
Receive-only channel. You can only receive from it: `v := <-ch`. Cannot send to it.

**22. What is `select`?**
A control structure that waits on multiple channel operations and executes the first one that's ready:
```go
select {
case v := <-ch1:
    // ch1 had a value
case ch2 <- x:
    // sent to ch2
}
```

**23. What happens if multiple select cases are ready?**
Go picks one at random. No priority ordering.

**24. What is `default` in select?**
If no other case is ready, `default` executes immediately. Makes `select` non-blocking.

**25. How do you implement timeout with select?**
```go
select {
case v := <-ch:
    // got value
case <-time.After(5 * time.Second):
    // timed out
}
```

**26. How do you implement cancellation with channels?**
Use a done channel that gets closed to signal all goroutines to stop:
```go
select {
case work := <-jobs:
    process(work)
case <-done:
    return
}
```

**27. What is a done channel?**
A channel (usually `chan struct{}`) that is closed to broadcast a cancellation signal to multiple goroutines simultaneously.

**28. What is channel ownership?**
The goroutine that creates and closes a channel "owns" it. Only the owner should close the channel.

**29. What is backpressure?**
When a producer slows down because the consumer can't keep up. Buffered channels provide natural backpressure — the producer blocks when the buffer is full.

**30. When should you use channels instead of mutexes?**
Use channels for passing data and ownership between goroutines. Use mutexes for protecting shared state that multiple goroutines need to read/write concurrently.

---

## 13. Context

**1. What is `context.Context`?**
An interface for propagating cancellation signals, deadlines, and request-scoped values across API boundaries and goroutines.

**2. Why is context important?**
It lets you cancel in-flight operations cleanly — stop a database query or HTTP call when the client disconnects or a timeout fires.

**3. What is `context.Background()`?**
The root context. Never cancelled, has no deadline. Use as the top-level context in `main`, tests, or server startup.

**4. What is `context.TODO()`?**
A placeholder context for when you're not sure which context to use yet. Functionally identical to `Background()` but signals intent to add proper context later.

**5. What is `context.WithCancel()`?**
Creates a child context with a cancel function:
```go
ctx, cancel := context.WithCancel(parent)
defer cancel()
```
Calling `cancel()` signals all work using `ctx` to stop.

**6. What is `context.WithTimeout()`?**
Creates a context that cancels automatically after a duration:
```go
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel()
```

**7. What is `context.WithDeadline()`?**
Like `WithTimeout` but takes a specific time instead of a duration:
```go
ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
```

**8. What is `context.WithValue()`?**
Attaches a key-value pair to the context. Use for request-scoped values (request ID, user ID) — not for optional function parameters.

**9. When should you avoid `context.WithValue()`?**
When the value is a function parameter, not truly request-scoped. Don't use it to pass config, databases, or loggers — use explicit parameters instead.

**10. How do you cancel an HTTP request?**
```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()
req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
```

**11. How do you pass context to database queries?**
Most database drivers accept context: `db.QueryContext(ctx, query)`, `db.ExecContext(ctx, ...)`.

**12. Why should context be the first function parameter?**
It's a Go convention for consistency and readability. Makes it easy to identify functions that support cancellation.

**13. Should context be stored in a struct?**
Generally no. The Go documentation says: "Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it."

**14. How do you prevent context leaks?**
Always call the cancel function returned by `WithCancel`, `WithTimeout`, and `WithDeadline`. Using `defer cancel()` is the standard pattern.

**15. Why should you call `cancel()`?**
Not calling it leaks resources. The context and its children stay in memory until the parent is cancelled or the deadline fires.

**16. What happens if you forget to call cancel?**
The context and its goroutines may stay alive longer than expected, causing resource leaks.

**17. How do you propagate cancellation across goroutines?**
Pass the same context to all goroutines. When the context is cancelled, all goroutines that check `ctx.Done()` can exit.

**18. How do you use context in worker pools?**
Pass the context to each worker. Each worker selects between work and `ctx.Done()`:
```go
select {
case job := <-jobs: process(job)
case <-ctx.Done(): return
}
```

**19. What is request-scoped data?**
Data that belongs to a single request's lifecycle — like a request ID or authenticated user. Context values are appropriate for these.

**20. How does context help with graceful shutdown?**
A root cancel context is passed to all goroutines. When shutdown is triggered, `cancel()` is called, all goroutines detect `ctx.Done()` and stop cleanly.

---

## 14. Synchronization

**1. What is a race condition?**
When two or more goroutines access shared data concurrently and at least one writes, without synchronization. The outcome depends on timing.

**2. What is a data race?**
An unsynchronized concurrent access to the same memory location where at least one access is a write. Causes undefined behavior.

**3. How do you detect race conditions in Go?**
Run your program with the race detector: `go run -race`, `go test -race`, `go build -race`.

**4. What does `go test -race` do?**
Compiles and runs tests with the race detector enabled. It instruments memory accesses and reports if any data races are detected.

**5. What is `sync.Mutex`?**
A mutual exclusion lock. Only one goroutine can hold it at a time.
```go
mu.Lock()
defer mu.Unlock()
// critical section
```

**6. What is `sync.RWMutex`?**
A reader/writer lock. Multiple goroutines can hold a read lock simultaneously. Only one can hold a write lock, and no readers can hold the lock while a writer does.

**7. What is the difference between `Mutex` and `RWMutex`?**
`Mutex` blocks all concurrent access. `RWMutex` allows concurrent reads while still protecting against concurrent writes.

**8. When should you use `RWMutex`?**
When reads vastly outnumber writes. Concurrent reads improve throughput. If writes are common, the overhead of `RWMutex` may not be worth it.

**9. What is `sync.Once`?**
Ensures a function is executed exactly once, even if called from multiple goroutines:
```go
var once sync.Once
once.Do(func() { /* initialization */ })
```
Used for lazy initialization of singletons.

**10. What is `sync.Cond`?**
A condition variable. Allows goroutines to wait for a condition to become true, and be signaled when it does.

**11. What is `sync.Pool`?**
A temporary object pool for reusing allocations. Reduces GC pressure for frequently allocated/freed objects.

**12. When should you use `sync.Pool`?**
For expensive-to-allocate objects that are used briefly and returned often — like buffers in HTTP handlers. Note: pooled objects can be GC'd at any time.

**13. What is `atomic`?**
The `sync/atomic` package provides low-level atomic operations (add, load, store, compare-and-swap) on primitive types without a mutex.

**14. What is the difference between mutex and atomic operations?**
Atomics are faster and work on individual integers/pointers. Mutexes protect arbitrary code sections. Use atomics for simple counters; mutexes for complex shared state.

**15. What is a deadlock?**
Two or more goroutines are each waiting for the other to release a lock — they block forever. Go detects all-goroutine deadlocks and panics.

**16. How do you avoid deadlocks?**
Always acquire locks in a consistent order. Minimize the scope of locks. Use `defer mu.Unlock()` to ensure unlocking. Avoid holding a lock while waiting on a channel.

**17. What is lock contention?**
When multiple goroutines compete for the same lock, causing them to wait. High contention reduces performance.

**18. What is lock granularity?**
How much data a single lock protects. Fine-grained locks protect small pieces of data (less contention). Coarse-grained locks protect more data (simpler but more contention).

**19. Can you copy a mutex?**
No. Copying a `sync.Mutex` copies its internal state, which can cause deadlocks. Always use a pointer to a mutex or embed it without copying.

**20. Why should you avoid copying structs containing mutexes?**
The mutex state (locked/unlocked) is copied too. The copy and the original become independent locks, defeating the purpose. The `go vet` tool warns about this.

---

## 15. Memory Management and Garbage Collection

**1. Does Go have garbage collection?**
Yes. Go has a concurrent, tri-color mark-and-sweep garbage collector that runs alongside your program.

**2. How does Go garbage collection work at a high level?**
The GC traces all reachable objects from root references (globals, stack variables). Unreachable objects are swept (freed). It runs concurrently with your program to minimize pauses.

**3. What is the heap?**
Memory managed by the GC, used for objects that need to outlive the function that created them or are too large for the stack.

**4. What is the stack?**
Per-goroutine memory for local variables and function call frames. Allocation and deallocation are instant (just move a pointer). Goroutine stacks start small and grow dynamically.

**5. What is escape analysis?**
A compile-time analysis that determines whether a variable can be allocated on the stack or must go to the heap. Stack is faster; heap adds GC pressure.

**6. How do you check escape analysis?**
```bash
go build -gcflags="-m" ./...
```
Shows which variables escape to the heap.

**7. What does `go build -gcflags="-m"` do?**
Prints the compiler's escape analysis decisions — what escapes to heap and why. Useful for understanding and reducing allocations.

**8. What causes allocations?**
Taking the address of a local variable (if it escapes), interface boxing, closures capturing variables, `make`, `new`, `append` beyond capacity, string conversions to/from `[]byte`.

**9. How do you reduce allocations?**
Pre-allocate slices with the right capacity, use `sync.Pool` for reusable buffers, avoid unnecessary interface boxing, use `strings.Builder` instead of `+`, pass values instead of pointers for small types.

**10. What is pointer pressure?**
Having many pointers the GC must scan. More pointers = more work for the GC. Reducing pointer-heavy data structures improves GC performance.

**11. What is GC pressure?**
The amount of work your GC has to do. High allocation rates create high GC pressure — the GC runs more frequently, reducing throughput.

**12. What is `GOGC`?**
An environment variable controlling GC aggressiveness. Default is 100, meaning GC runs when heap doubles. Lower values = more frequent GC (less memory, more CPU). Higher = less frequent (more memory, less CPU).

**13. How does `sync.Pool` help GC?**
Pooled objects are reused instead of freed. This reduces allocation rate, which directly reduces GC work.

**14. Can Go have memory leaks?**
Yes — not through typical use of GC, but through goroutine leaks, lingering references, large slice/string sub-slices, or incorrect use of finalizers.

**15. What are common causes of memory leaks in Go?**
Goroutine leaks, holding references in global maps that are never cleaned up, sub-slicing large arrays keeping the entire backing array alive, unclosed HTTP response bodies.

**16. How can goroutines cause memory leaks?**
A leaked goroutine (one that never exits) holds its stack and anything it references in memory indefinitely.

**17. How can slices cause memory leaks?**
A sub-slice keeps the entire backing array alive. If you take a small slice of a huge array, the huge array cannot be GC'd.

**18. What is stack growth?**
Goroutine stacks start small (~2KB) and grow automatically when more stack space is needed. Go uses a copying or segmented stack strategy.

**19. Are goroutine stacks fixed size?**
No. They start small and grow dynamically up to a default maximum of 1GB (configurable).

**20. How do you profile memory usage?**
Use `pprof`:
```go
import _ "net/http/pprof"
// then: go tool pprof http://localhost:6060/debug/pprof/heap
```
Or write a heap profile to a file with `runtime/pprof`.
---

## 16. Generics

**1. What are generics?**
A way to write functions and types that work with any type, specified at compile time. Introduced in Go 1.18.

**2. Why were generics added to Go?**
To eliminate repetitive code — writing the same logic for `int`, `string`, `float64` separately. Generics allow reusable, type-safe algorithms.

**3. How do you define a generic function?**
```go
func Min[T constraints.Ordered](a, b T) T {
    if a < b { return a }
    return b
}
```
The `[T constraints.Ordered]` is the type parameter.

**4. What is a type parameter?**
A placeholder for a type, declared inside `[]` after the function or type name. Replaced with a concrete type at compile time.

**5. What is a type constraint?**
An interface that limits what types can be used as a type parameter:
```go
func Sum[T int | float64](nums []T) T { ... }
```

**6. What is `any`?**
As a constraint, `any` means no restrictions — any type is allowed. Same as `interface{}`.

**7. What is `comparable`?**
A built-in constraint meaning the type supports `==` and `!=`. Required for map keys or when you need to compare values.

**8. When do you use `comparable`?**
When writing generic functions that need to compare values with `==`, like a generic `Contains` or `IndexOf`.

**9. How do you define a generic struct?**
```go
type Stack[T any] struct {
    items []T
}
func (s *Stack[T]) Push(item T) { s.items = append(s.items, item) }
```

**10. Can methods have type parameters?**
No. Methods cannot add new type parameters. Type parameters must be declared on the type itself.

**11. Can types have type parameters?**
Yes: `type Pair[K, V any] struct { Key K; Value V }`.

**12. What is a type set?**
The set of types that satisfy a constraint interface. A type is in the set if it implements all methods and matches any type terms.

**13. What does `~int` mean in a constraint?**
`~int` means `int` and any type whose underlying type is `int`:
```go
type MyInt int // ~int includes MyInt
```

**14. What is the difference between `int` and `~int` in constraints?**
`int` only matches exactly `int`. `~int` also matches named types with underlying type `int`, like `type UserID int`.

**15. How do you constrain numeric types?**
Use `constraints.Integer`, `constraints.Float`, or `constraints.Ordered` from `golang.org/x/exp/constraints`, or define your own:
```go
type Number interface { ~int | ~float64 }
```

**16. When should you avoid generics?**
When the code is already simple without them, when you only have one concrete type, or when generics would make the code harder to read. Don't use generics just to be clever.

**17. Should you use generics instead of interfaces?**
Use interfaces when behavior (methods) matters. Use generics when the type itself matters and you need to work with the actual value (not just call methods on it).

**18. What are the limitations of Go generics?**
No generic methods (only generic types/functions), no operator overloading, some constraints require external packages, and type inference doesn't always work.

**19. How do generics affect readability?**
Simple generic functions are readable. Complex nested type constraints can be hard to parse. Prefer readable code over clever generic abstractions.

**20. Give an example of a generic `Min` function.**
```go
func Min[T constraints.Ordered](a, b T) T {
    if a < b { return a }
    return b
}
// Usage:
Min(3, 5)        // int
Min(3.14, 2.71)  // float64
Min("a", "b")    // string
```

---

## 17. HTTP and Web Development

**1. How do you create an HTTP server in Go?**
```go
http.HandleFunc("/", handler)
http.ListenAndServe(":8080", nil)
```

**2. What is `net/http`?**
The standard library package for HTTP clients and servers. No third-party framework required for most use cases.

**3. What is `http.Handler`?**
An interface with one method: `ServeHTTP(ResponseWriter, *Request)`. Anything implementing it can handle HTTP requests.

**4. What is `http.HandlerFunc`?**
An adapter that lets a plain function with the right signature be used as an `http.Handler`. Avoids needing to create a new type.

**5. What is middleware?**
A function that wraps an `http.Handler` to add behavior before/after the actual handler — like logging, authentication, or panic recovery.

**6. How do you create middleware in Go?**
```go
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
```

**7. How do you read query parameters?**
```go
r.URL.Query().Get("name")
```

**8. How do you read path parameters?**
With the standard library (Go 1.22+): `r.PathValue("id")`. With older versions, use a router library like `gorilla/mux` or `chi`.

**9. How do you read JSON request body?**
```go
var payload MyStruct
json.NewDecoder(r.Body).Decode(&payload)
```

**10. How do you write JSON response?**
```go
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(data)
```

**11. How do you handle HTTP errors?**
```go
http.Error(w, "not found", http.StatusNotFound)
```

**12. How do you set headers?**
```go
w.Header().Set("X-Custom-Header", "value")
```
Must set headers before writing the status code or body.

**13. How do you handle CORS?**
Set the appropriate response headers (`Access-Control-Allow-Origin`, etc.) in middleware, or use a library like `rs/cors`.

**14. How do you handle file uploads?**
```go
r.ParseMultipartForm(10 << 20) // 10 MB limit
file, header, err := r.FormFile("upload")
```

**15. How do you serve static files?**
```go
http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
```

**16. What is graceful shutdown?**
Stopping the server without abruptly cutting active connections. New connections are rejected, existing requests are allowed to complete.

**17. How do you implement graceful shutdown?**
```go
srv := &http.Server{Addr: ":8080"}
go srv.ListenAndServe()
// On shutdown signal:
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
srv.Shutdown(ctx)
```

**18. What is request timeout?**
The total time allowed for a client request to complete. Set via context or `http.Server` timeout fields.

**19. What is read timeout?**
The maximum time allowed to read the complete request, including body. Set with `http.Server{ReadTimeout: 5*time.Second}`.

**20. What is write timeout?**
The maximum time allowed to write the response. Set with `http.Server{WriteTimeout: 10*time.Second}`.

**21. Why should you configure timeouts on HTTP servers?**
Without timeouts, slow or malicious clients can hold connections open forever, exhausting server resources.

**22. How do you create an HTTP client?**
```go
client := &http.Client{Timeout: 10 * time.Second}
resp, err := client.Get(url)
```

**23. Why should you reuse `http.Client`?**
`http.Client` manages a connection pool. Reusing it allows connection reuse (HTTP keep-alive), which is faster and uses fewer resources than creating a new client per request.

**24. What is connection pooling?**
Reusing existing TCP connections for multiple HTTP requests instead of creating a new connection each time.

**25. What is `Transport`?**
`http.Transport` is the low-level mechanism managing connections, TLS, proxies, and connection pooling. It's shared across an `http.Client`.

**26. How do you set client timeout?**
```go
client := &http.Client{Timeout: 10 * time.Second}
```
This covers the full round trip (connect + send + receive).

**27. How do you retry failed requests?**
Implement a loop with backoff, check if the error is retryable (network errors, 5xx), and limit retries:
```go
for attempt := 0; attempt < maxRetries; attempt++ {
    resp, err = client.Do(req)
    if err == nil && resp.StatusCode < 500 { break }
    time.Sleep(backoff(attempt))
}
```

**28. Should you retry all failed HTTP requests?**
No. Only retry idempotent, safe operations (GET, PUT). Retrying POST can cause duplicate actions.

**29. What is idempotency?**
An operation is idempotent if performing it multiple times has the same effect as once. GET, PUT, DELETE are idempotent; POST is generally not.

**30. How do you handle context cancellation in HTTP requests?**
Use `http.NewRequestWithContext(ctx, ...)`. If the context is cancelled, the HTTP client aborts the request.

---

## 18. Databases

**1. How do you connect Go to SQL databases?**
Use `database/sql` with a database driver (e.g., `lib/pq` for PostgreSQL, `go-sql-driver/mysql`):
```go
db, err := sql.Open("postgres", connStr)
```

**2. What is `database/sql`?**
The standard library package providing a generic SQL interface. Database-specific behavior is handled by drivers.

**3. What is a database driver?**
A package that implements the `database/sql/driver` interface for a specific database (e.g., `pgx`, `lib/pq` for PostgreSQL).

**4. What is `sql.DB`?**
A handle representing a pool of database connections. It's safe for concurrent use.

**5. Is `sql.DB` a single connection?**
No. It's a connection pool. It manages multiple connections automatically.

**6. What is connection pooling?**
`sql.DB` maintains a pool of open connections. Requests reuse idle connections instead of opening a new one each time.

**7. How do you configure max open connections?**
```go
db.SetMaxOpenConns(25)
```

**8. How do you configure max idle connections?**
```go
db.SetMaxIdleConns(10)
```

**9. How do you configure connection lifetime?**
```go
db.SetConnMaxLifetime(5 * time.Minute)
```
Prevents stale connections from being reused.

**10. What is `Query`?**
Executes a SELECT returning multiple rows:
```go
rows, err := db.QueryContext(ctx, "SELECT id, name FROM users")
defer rows.Close()
```

**11. What is `QueryRow`?**
Executes a SELECT expected to return one row:
```go
var name string
db.QueryRowContext(ctx, "SELECT name FROM users WHERE id=$1", id).Scan(&name)
```

**12. What is `Exec`?**
Executes an INSERT, UPDATE, or DELETE (no rows returned):
```go
result, err := db.ExecContext(ctx, "DELETE FROM users WHERE id=$1", id)
```

**13. What is prepared statement?**
A pre-compiled SQL statement. Faster for repeated execution and protects against SQL injection:
```go
stmt, err := db.PrepareContext(ctx, "SELECT * FROM users WHERE id=$1")
```

**14. How do you prevent SQL injection?**
Always use parameterized queries (`$1`, `?`). Never build SQL strings by concatenating user input.

**15. How do you handle transactions?**
```go
tx, err := db.BeginTx(ctx, nil)
if err != nil { return err }
defer tx.Rollback() // no-op if committed
// ... db operations using tx ...
return tx.Commit()
```

**16. How do you commit a transaction?**
`tx.Commit()` — must be called explicitly after all operations succeed.

**17. How do you rollback a transaction?**
`tx.Rollback()` — call if any operation fails. A deferred `Rollback()` is a no-op after `Commit()`.

**18. Why should rollback often be deferred?**
Calling `defer tx.Rollback()` ensures the transaction is always rolled back on error, even if you forget or a panic occurs. After a successful `Commit()`, the `Rollback()` does nothing.

**19. How do you pass context to DB queries?**
Use the `Context` variants: `QueryContext`, `ExecContext`, `QueryRowContext`. This enables timeout and cancellation.

**20. What happens if a context timeout occurs during a query?**
The query is cancelled. The database driver sends a cancel signal if possible. The query returns a context deadline exceeded error.

**21. How do you scan nullable database fields?**
Use `sql.NullString`, `sql.NullInt64`, etc.:
```go
var name sql.NullString
row.Scan(&name)
if name.Valid { fmt.Println(name.String) }
```

**22. What is `sql.NullString`?**
A struct representing a string that can be NULL in the database: `{ String string; Valid bool }`. Valid is false when the DB value is NULL.

**23. What are common database mistakes in Go?**
Not closing `rows`, not using context, ignoring `rows.Err()`, SQL injection from string concatenation, not configuring connection pool limits.

**24. What is N+1 query problem?**
Loading a list of N records, then making N additional queries (one per record) to fetch related data. Solve with JOINs or batch queries.

**25. How do you test database code?**
Use an in-memory database (SQLite), a test database with transactions rolled back after each test, or mock the database layer using an interface.

---

## 19. Testing

**1. How do you write tests in Go?**
Create a `_test.go` file, import `testing`, and write functions named `TestXxx(t *testing.T)`:
```go
func TestAdd(t *testing.T) {
    if add(1, 2) != 3 { t.Error("expected 3") }
}
```

**2. What is the `testing` package?**
The standard library package that provides testing infrastructure, benchmarking, fuzzing, and test helpers.

**3. How do you name test files?**
End with `_test.go`. Go only compiles these during testing.

**4. How do you name test functions?**
`Test` + anything starting with a capital letter: `TestAdd`, `TestUserCreate`.

**5. What is table-driven testing?**
Defining test cases as a slice of structs and ranging over them:
```go
tests := []struct{ input, want int }{{1, 2}, {3, 4}}
for _, tc := range tests {
    if got := fn(tc.input); got != tc.want { t.Errorf(...) }
}
```

**6. Why is table-driven testing common in Go?**
It keeps test cases compact, makes it easy to add new cases, and reduces code duplication.

**7. What is `t.Run()`?**
Creates a subtest with its own name, useful with table-driven tests:
```go
t.Run("case name", func(t *testing.T) { ... })
```

**8. What is `t.Helper()`?**
Marks the calling function as a test helper. Error messages will show the caller's line number, not the helper's.

**9. What is `t.Fatal`?**
Logs the error and immediately stops the current test function. Use when continuing doesn't make sense.

**10. What is `t.Error`?**
Marks the test as failed and logs the message, but continues running the test function.

**11. What is the difference between `Fatal` and `Error`?**
`Fatal` stops the test immediately. `Error` marks it as failed but lets the test continue so you can see all failures at once.

**12. What is a benchmark?**
A function named `BenchmarkXxx(b *testing.B)` that measures performance:
```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ { add(1, 2) }
}
```

**13. How do you write benchmark tests?**
Use `b.N` as the loop count. The testing framework adjusts `N` until stable results are obtained. Run with `go test -bench=.`.

**14. What is `b.N`?**
The number of iterations the benchmark loop should run. The framework sets this automatically to get statistically stable results.

**15. What is fuzz testing?**
Automated testing where the framework generates random inputs to find crashes or panics:
```go
func FuzzAdd(f *testing.F) {
    f.Add(1, 2)
    f.Fuzz(func(t *testing.T, a, b int) { add(a, b) })
}
```

**16. How do you run tests?**
`go test ./...` runs all tests in all packages. `go test .` runs tests in the current package.

**17. How do you run a specific test?**
```go
go test -run TestAdd ./...
```

**18. How do you run tests with coverage?**
```go
go test -cover ./...
```

**19. How do you generate an HTML coverage report?**
```bash
go test -coverprofile=cover.out ./...
go tool cover -html=cover.out
```

**20. How do you test HTTP handlers?**
Use `httptest.NewRecorder()` as the `ResponseWriter` and `httptest.NewRequest()` to create a fake request.

**21. What is `httptest`?**
The `net/http/httptest` package provides utilities for testing HTTP handlers without starting a real server.

**22. How do you mock dependencies?**
Define an interface for the dependency, create a fake struct that implements it in the test file, and inject it.

**23. Should you use interfaces for testing?**
Yes, when you need to substitute dependencies. But don't define interfaces just for testing — define them where they add value for the design too.

**24. What is dependency injection?**
Passing dependencies (DB, logger, HTTP client) into a function/struct instead of creating them internally. Makes code testable and configurable.

**25. How do you test goroutines?**
Use `sync.WaitGroup` or channels to wait for goroutines to complete before asserting. The race detector (`-race`) catches data races in tests.

**26. How do you test timeouts?**
Use `context.WithTimeout` and check that the function respects cancellation. Use short timeouts in tests.

**27. How do you avoid flaky tests?**
Don't rely on timing. Use channels/WaitGroups to wait for events. Avoid global state. Run with `-race`.

**28. What is race testing?**
Running tests with the race detector: `go test -race`. Detects concurrent memory access issues.

**29. How do you test private functions?**
Test them through the exported API they support. If you must test a private function directly, the test file in the same package has access.

**30. Should you test private functions directly?**
Generally no. Test behavior through public APIs. If a private function is complex enough to warrant direct testing, consider making it a separate exported function or package.

---

## 20. Packages, Modules, and Project Structure

**1. What is a package?**
A directory of Go source files. All files in a directory share the same package name and can access each other's unexported identifiers.

**2. What is a module?**
A collection of related packages with a defined module path and versioning, declared in `go.mod`.

**3. What is `go.mod`?**
The module definition file. Contains the module path, Go version, and direct dependencies:
```
module github.com/user/myapp
go 1.21
require github.com/some/lib v1.2.3
```

**4. What is `go.sum`?**
A file containing cryptographic hashes of module contents. Ensures reproducible builds and detects tampering.

**5. What is semantic import versioning?**
Major versions >= 2 must be included in the module path: `github.com/user/lib/v2`. This allows multiple major versions to coexist.

**6. How do you create a new Go module?**
```bash
go mod init github.com/user/myapp
```

**7. How do you add dependencies?**
```bash
go get github.com/some/package@v1.2.3
```
Or just import and run `go mod tidy`.

**8. How do you update dependencies?**
```bash
go get -u ./...         # update all
go get package@latest   # update specific
```

**9. What is `go mod tidy`?**
Adds missing and removes unused dependencies from `go.mod` and `go.sum`. Run before committing.

**10. What is vendoring?**
Copying all dependencies into a `vendor/` directory in your repo. Use `go mod vendor`. The build uses local copies instead of downloading.

**11. What is the `internal` package?**
A special directory name. Code inside `internal/` can only be imported by code in the parent tree of the `internal` directory. Prevents external packages from importing your private code.

**12. What is the `cmd` directory?**
A convention for placing the `main` packages (executables) of a project. Each subdirectory is a separate binary: `cmd/server`, `cmd/worker`.

**13. What is the `pkg` directory?**
A historical convention for shared library code. It's largely unnecessary and often discouraged now — just use meaningful package names at the root.

**14. What is the difference between package name and folder name?**
The folder name is used in import paths. The package name (in `package xxx`) is used in code. They should match, but the canonical exception is `main`.

**15. How do you avoid circular imports?**
Design packages with clear layers. Lower-level packages should not import higher-level ones. Refactor shared code into a separate package.

**16. Does Go allow circular dependencies?**
No. The compiler rejects circular imports. This enforces clean, acyclic dependency graphs.

**17. How do you structure a REST API project?**
Common layout:
```
cmd/api/main.go
internal/handler/
internal/service/
internal/repository/
internal/domain/
```

**18. Should every project use Clean Architecture?**
No. For small projects it's overkill. Use the simplest structure that keeps concerns separated. Add layers as complexity grows.

**19. What is idiomatic Go project structure?**
Keep it flat for small projects. Use `internal/` for private code. Group by feature/domain rather than by layer (handler/service/repo). Avoid deep nesting.

**20. How do you organize domain, service, repository, and handler layers?**
- `domain/` or `model/`: data types, interfaces
- `repository/`: database access
- `service/` or `usecase/`: business logic
- `handler/` or `transport/`: HTTP, gRPC handlers
Each layer depends on the one below via interfaces, not concrete types.
---

## 21. Performance and Profiling

**1. How do you profile Go applications?**
Use `pprof`. Expose the debug endpoints or write profiles to files, then analyze with `go tool pprof`.

**2. What is `pprof`?**
A profiling tool built into Go. It can profile CPU, memory, goroutines, mutex contention, and blocking.

**3. How do you profile CPU usage?**
```go
import "runtime/pprof"
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
// or use net/http/pprof and: go tool pprof http://localhost:6060/debug/pprof/profile
```

**4. How do you profile memory usage?**
```go
pprof.WriteHeapProfile(f)
// or: go tool pprof http://localhost:6060/debug/pprof/heap
```

**5. How do you profile goroutines?**
```bash
go tool pprof http://localhost:6060/debug/pprof/goroutine
```
Shows all running goroutines and their stack traces.

**6. How do you profile mutex contention?**
Enable mutex profiling: `runtime.SetMutexProfileFraction(1)`, then:
```bash
go tool pprof http://localhost:6060/debug/pprof/mutex
```

**7. What is a benchmark?**
A `BenchmarkXxx(b *testing.B)` function that runs your code `b.N` times to measure its performance. Run with `go test -bench=.`.

**8. How do you reduce allocations?**
Pre-allocate slices with known capacity, use `sync.Pool` for reusable buffers, avoid interface boxing of value types, use `strings.Builder` for string building.

**9. How do you optimize string concatenation?**
Don't use `+` in a loop — it allocates a new string each time. Use `strings.Builder`:
```go
var b strings.Builder
for _, s := range parts { b.WriteString(s) }
result := b.String()
```

**10. What is `strings.Builder`?**
A buffer for efficiently building strings. Write to it with `WriteString`/`WriteByte`, get the result with `String()`. Minimizes allocations.

**11. What is `bytes.Buffer`?**
Like `strings.Builder` but for `[]byte`. Used when you need byte-level output.

**12. When should you use `strings.Builder`?**
When building strings from many pieces, especially in loops. It avoids the O(n²) cost of repeated `+` concatenation.

**13. How do you avoid unnecessary conversions?**
Avoid converting `string` ↔ `[]byte` unnecessarily — each conversion allocates. Use `[]byte` throughout if you need to modify, or keep as `string` if not.

**14. How do you optimize JSON encoding?**
Use `json.NewEncoder` to stream directly to a writer (avoids intermediate buffer). For hot paths, use faster libraries like `jsoniter` or `sonic`.

**15. How do you optimize database access?**
Use connection pooling, prepared statements, select only needed columns, avoid N+1 queries, use batch inserts, index properly.

**16. How do you identify slow code?**
CPU profiling shows which functions use the most time. Benchmarks measure specific functions. Tracing (`go tool trace`) shows goroutine activity.

**17. What is Big O notation?**
A way to describe algorithm complexity — how time or space grows with input size. O(1) is constant, O(n) is linear, O(n²) is quadratic.

**18. How do you avoid premature optimization?**
Write correct, readable code first. Profile to find actual bottlenecks. Only optimize what the profiler shows is slow.

**19. How do you tune garbage collection?**
Adjust `GOGC` (default 100). Use `GOMEMLIMIT` (Go 1.19+) to set a memory ceiling. Reduce allocations to lower GC frequency.

**20. What is backpressure?**
When downstream slowness propagates upstream to slow down producers. In Go, buffered channels naturally provide backpressure — producers block when the buffer is full.

---

## 22. JSON and Encoding

**1. How do you encode JSON in Go?**
```go
data, err := json.Marshal(v)
```

**2. How do you decode JSON in Go?**
```go
var v MyStruct
err := json.Unmarshal(data, &v)
```

**3. What is `json.Marshal`?**
Converts a Go value to a JSON byte slice.

**4. What is `json.Unmarshal`?**
Parses a JSON byte slice into a Go value.

**5. What is `json.NewEncoder`?**
Creates a streaming JSON encoder that writes directly to an `io.Writer`:
```go
json.NewEncoder(w).Encode(v)
```

**6. What is `json.NewDecoder`?**
Creates a streaming JSON decoder that reads from an `io.Reader`:
```go
json.NewDecoder(r.Body).Decode(&v)
```

**7. What is the difference between `Marshal` and `Encoder`?**
`Marshal` works with an in-memory byte slice. `Encoder` writes directly to a stream, which is more efficient for HTTP responses or file output.

**8. What are JSON struct tags?**
Annotations on struct fields that control JSON encoding:
```go
Name string `json:"name"`
```

**9. What does `omitempty` do?**
Omits the field from JSON output if it is the zero value (`0`, `""`, `false`, `nil`):
```go
Name string `json:"name,omitempty"`
```

**10. Why does `omitempty` sometimes behave unexpectedly?**
It checks for the zero value, not "empty" in a semantic sense. A struct is never omitted even if all its fields are zero. An `int` with value `0` will be omitted even if that's meaningful.

**11. How do you handle unknown JSON fields?**
By default, unknown fields are silently ignored. Use a decoder and `DisallowUnknownFields()` to reject them.

**12. What does `DisallowUnknownFields` do?**
```go
dec := json.NewDecoder(r.Body)
dec.DisallowUnknownFields()
```
Returns an error if the JSON contains a field not in the struct. Useful for strict APIs.

**13. How do you decode dynamic JSON?**
Use `map[string]interface{}` or `json.RawMessage` to defer decoding:
```go
var result map[string]interface{}
json.Unmarshal(data, &result)
```

**14. How do you handle nullable JSON fields?**
Use a pointer: `*string` is nil when JSON is null; `string` would be empty string. Or use `json.RawMessage`.

**15. What is the difference between missing field and zero value?**
`json.Unmarshal` can't distinguish — both leave the field as its zero value. Use `*string` (pointer) to tell: nil = missing/null, non-nil = present.

**16. How do you customize JSON marshaling?**
Implement `json.Marshaler` and/or `json.Unmarshaler` interfaces on your type.

**17. What are `MarshalJSON` and `UnmarshalJSON`?**
Methods you implement to fully control how your type is encoded/decoded:
```go
func (t MyTime) MarshalJSON() ([]byte, error) { ... }
func (t *MyTime) UnmarshalJSON(data []byte) error { ... }
```

**18. How do you avoid exposing sensitive fields in JSON?**
Tag them with `json:"-"` to always exclude:
```go
Password string `json:"-"`
```

**19. How do you stream large JSON responses?**
Use `json.NewEncoder` with `io.Pipe` or write an array element by element, flushing periodically:
```go
enc := json.NewEncoder(w)
for _, item := range items { enc.Encode(item) }
```

**20. How do you handle time in JSON?**
`time.Time` marshals to RFC 3339 format by default. For custom formats, implement `MarshalJSON`/`UnmarshalJSON`.

---

## 23. Time

**1. How do you work with dates in Go?**
Use the `time` package. `time.Time` represents a moment, `time.Duration` represents a span.

**2. What is `time.Time`?**
A struct representing a specific instant in time with nanosecond precision and timezone info.

**3. What is `time.Duration`?**
An `int64` representing nanoseconds. Constants like `time.Second`, `time.Minute`, `time.Hour` make it readable.

**4. How do you parse time?**
```go
t, err := time.Parse("2006-01-02", "2024-01-15")
```

**5. Why does Go use `2006-01-02 15:04:05` as layout?**
It's Go's "reference time" — a fixed moment (Mon Jan 2 15:04:05 MST 2006). Each number is unique so you can't confuse month and day. The mnemonic: 1-2-3-4-5-6-7 (month-day-hour-minute-second-year-timezone).

**6. How do you format time?**
```go
t.Format("2006-01-02 15:04:05")
```
Use the same reference time as a template.

**7. How do you compare times?**
```go
t1.Before(t2)
t1.After(t2)
t1.Equal(t2)
```
Don't use `==` directly — it also compares monotonic clock and timezone.

**8. How do you add duration to time?**
```go
later := t.Add(24 * time.Hour)
```

**9. What is the difference between `time.Add` and `time.AddDate`?**
`Add` adds a `Duration` (precise, handles DST consistently). `AddDate` adds calendar years/months/days (handles month-length differences).

**10. How do you handle time zones?**
```go
loc, _ := time.LoadLocation("America/New_York")
t = t.In(loc)
```
Use UTC for storage; convert to local timezone for display.

**11. What is UTC?**
Coordinated Universal Time — the world's time standard. No daylight saving, no timezone offset. Always store and process times in UTC.

**12. How do you create a ticker?**
```go
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()
for t := range ticker.C {
    fmt.Println(t)
}
```

**13. What is the difference between `time.Timer` and `time.Ticker`?**
`Timer` fires once after a duration. `Ticker` fires repeatedly at a fixed interval.

**14. Why should you stop tickers?**
Not stopping a ticker leaks a goroutine in the runtime that keeps firing. Always call `ticker.Stop()`.

**15. How do you implement timeout?**
```go
select {
case result := <-work:
    return result
case <-time.After(5 * time.Second):
    return errors.New("timeout")
}
```
Or use `context.WithTimeout`.

---

## 24. Senior Go Questions

**1. What does idiomatic Go mean?**
Code that follows Go's conventions: small interfaces, explicit errors, simple structure, no over-engineering, clear naming, and proper use of concurrency primitives.

**2. How do you design a Go service?**
Define domain types and interfaces first. Build thin layers: transport (HTTP/gRPC) → service (business logic) → repository (data). Depend on interfaces, not implementations.

**3. How do you decide between channels and mutexes?**
Use channels for passing data and ownership between goroutines. Use mutexes for protecting shared state accessed by multiple goroutines simultaneously.

**4. How do you design a worker pool?**
Fixed number of goroutines reading from a jobs channel. Main goroutine sends work; workers send results to a results channel. Use context for cancellation.

**5. How do you prevent goroutine leaks?**
Always provide a way out: context cancellation, done channels. Use `defer cancel()`. Test with `goleak` library. Profile goroutine count in production.

**6. How do you handle graceful shutdown?**
Listen for OS signals (SIGTERM, SIGINT). Stop accepting new work. Wait for in-flight work to complete using a WaitGroup. Cancel the root context to signal all goroutines.

**7. How do you handle request cancellation?**
Pass `context.Context` to all layers. Check `ctx.Done()` at blocking points. Use context-aware versions of all I/O operations.

**8. How do you design a retry mechanism?**
Retry only idempotent operations. Use exponential backoff with jitter. Set a max retry count and total timeout. Check if the error is retryable before retrying.

**9. How do you design rate limiting?**
Use `golang.org/x/time/rate` token bucket: `limiter.Wait(ctx)` blocks until a token is available. For per-user rate limiting, maintain a map of limiters protected by a mutex.

**10. How do you design middleware?**
Write functions that take an `http.Handler` and return an `http.Handler`. Chain them together. Each middleware does one thing: logging, auth, recovery, etc.

**11. How do you handle observability?**
Three pillars: logs (structured, via `slog` or `zap`), metrics (Prometheus counters/histograms), and traces (OpenTelemetry). Instrument at the transport and service layers.

**12. How do you add logging?**
Use structured logging (`slog` standard library, or `zap`/`zerolog`). Include request ID, user ID, and duration. Log errors at the appropriate level.

**13. How do you add metrics?**
Use Prometheus client library. Expose `/metrics` endpoint. Track request count, error rate, and latency (histograms) per endpoint.

**14. How do you add tracing?**
Use OpenTelemetry SDK. Create spans around significant operations. Propagate trace context via HTTP headers or gRPC metadata.

**15. How do you handle configuration?**
Read from environment variables (12-factor app). Parse on startup into a typed config struct. Validate required fields. Use libraries like `envconfig` or `viper`.

**16. How do you handle secrets?**
Never hardcode or commit secrets. Use environment variables, Vault, AWS Secrets Manager, or Kubernetes secrets. Rotate secrets regularly.

**17. How do you design health checks?**
`/healthz` — liveness: is the process alive and not deadlocked? `readyz` — readiness: is the service ready to accept traffic (DB connected, caches warm)?

**18. How do you design readiness and liveness endpoints?**
Liveness: always return 200 unless the process is stuck. Readiness: check all dependencies (DB, cache, upstream services) before returning 200.

**19. How do you handle database migrations?**
Use a migration tool like `golang-migrate` or `goose`. Store migrations in version-controlled SQL files. Run migrations on startup or as a separate step in CI/CD.

**20. How do you handle backward-compatible API changes?**
Add new fields (never remove or rename). Use `omitempty` for new optional fields. Version breaking changes as `/v2`. Document deprecated fields.

**21. How do you version APIs?**
Path versioning (`/v1/`, `/v2/`) is most common. Header versioning is cleaner but harder to test. Semantic versioning for libraries.

**22. How do you design a clean package boundary?**
A package should have one clear responsibility. Its API should be minimal. It should only depend on packages below it in the dependency tree. Avoid circular imports by design.

**23. How do you avoid over-engineering in Go?**
Start with the simplest thing that works. Add abstractions when you feel real pain from duplication or coupling, not in anticipation of future needs.

**24. How do you avoid Java-style OOP in Go?**
Avoid creating classes with getter/setter methods. Prefer plain functions over methods. Use small interfaces. Don't create `Manager`, `Factory`, `Handler` structs just for the pattern.

**25. How do you keep Go code simple?**
Prefer short, focused functions. Return early on errors. Avoid deep nesting. Use flat package structures. Don't abstract until you have to.

**26. What are common mistakes by developers from PHP/JavaScript/Python?**
Using exceptions instead of error returns, global state, ignoring goroutine lifecycle, not closing resources, building string SQL instead of parameterized queries, not understanding interface satisfaction.

**27. How do you review Go code?**
Check error handling (are all errors checked?), resource cleanup (defer Close), goroutine safety (race conditions), interface design (too large?), and naming conventions.

**28. How do you handle high-throughput HTTP services?**
Reuse `http.Client` and DB connections, minimize allocations in hot paths, use `sync.Pool` for buffers, profile and optimize bottlenecks, use proper timeouts, consider connection limits.

**29. How do you debug production issues?**
Enable `pprof` endpoints (with auth). Use distributed tracing. Check structured logs. Use `go tool trace` for goroutine analysis. Capture heap dumps.

**30. How do you profile a production service safely?**
Expose `pprof` on a separate, internal-only port. CPU profiling adds ~5% overhead — acceptable for short periods. Heap profiling is cheaper. Always require auth on debug endpoints.

---

## 25. Common Coding Tasks

**1. Reverse a string safely with Unicode.**
```go
func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

**2. Check if a string is a palindrome.**
```go
func isPalindrome(s string) bool {
    r := []rune(strings.ToLower(s))
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        if r[i] != r[j] { return false }
    }
    return true
}
```

**3. Count word frequency.**
```go
func wordFreq(s string) map[string]int {
    freq := map[string]int{}
    for _, w := range strings.Fields(s) { freq[w]++ }
    return freq
}
```

**4. Remove duplicates from a slice.**
```go
func unique(s []int) []int {
    seen := map[int]bool{}
    result := s[:0]
    for _, v := range s {
        if !seen[v] { seen[v] = true; result = append(result, v) }
    }
    return result
}
```

**5. Implement safe concurrent counter.**
```go
type Counter struct {
    mu    sync.Mutex
    value int
}
func (c *Counter) Inc() { c.mu.Lock(); defer c.mu.Unlock(); c.value++ }
func (c *Counter) Get() int { c.mu.Lock(); defer c.mu.Unlock(); return c.value }
```

**6. Implement timeout with context.**
```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()
select {
case result := <-doWork(ctx):
    return result, nil
case <-ctx.Done():
    return nil, ctx.Err()
}
```

**7. Implement worker pool.**
```go
func workerPool(ctx context.Context, jobs <-chan Job, numWorkers int) <-chan Result {
    results := make(chan Result)
    var wg sync.WaitGroup
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for job := range jobs {
                select {
                case results <- process(job):
                case <-ctx.Done(): return
                }
            }
        }()
    }
    go func() { wg.Wait(); close(results) }()
    return results
}
```

**8. Implement rate limiter (token bucket).**
```go
import "golang.org/x/time/rate"
limiter := rate.NewLimiter(rate.Every(time.Second), 10) // 10 per second
if err := limiter.Wait(ctx); err != nil { return err }
// proceed
```

**9. Write table-driven tests.**
```go
func TestAdd(t *testing.T) {
    cases := []struct{ a, b, want int }{
        {1, 2, 3}, {0, 0, 0}, {-1, 1, 0},
    }
    for _, tc := range cases {
        t.Run(fmt.Sprintf("%d+%d", tc.a, tc.b), func(t *testing.T) {
            if got := add(tc.a, tc.b); got != tc.want {
                t.Errorf("got %d, want %d", got, tc.want)
            }
        })
    }
}
```

---

## 26. Tricky Go Questions

**1. What is wrong with this code?**
```go
for _, v := range values {
    go func() { fmt.Println(v) }()
}
```
**Problem:** All goroutines capture the same variable `v`. By the time they run, `v` likely holds the last value.
**Fix:**
```go
for _, v := range values {
    v := v // shadow with a new variable
    go func() { fmt.Println(v) }()
}
```

**2. What is the output?**
```go
defer fmt.Println("1")
defer fmt.Println("2")
defer fmt.Println("3")
```
Output: `3`, `2`, `1` — defers execute in LIFO order.

**3. What is the output?**
```go
s := []int{1, 2, 3}
modify(s)
fmt.Println(s)

func modify(s []int) { s[0] = 100 }
```
Output: `[100 2 3]` — The slice header is copied, but the underlying array is shared. Modifying elements via the copy affects the original.

**4. What is the output?**
```go
var s []int
fmt.Println(s == nil) // true
fmt.Println(len(s))   // 0
```
Output: `true`, `0` — A nil slice has length 0, and `nil` comparison works correctly.

**5. What happens here?**
```go
var m map[string]int
m["a"] = 1
```
**Panic:** `assignment to entry in nil map`. Must initialize with `make` first.

**6. What happens here?**
```go
ch := make(chan int)
ch <- 1
fmt.Println(<-ch)
```
**Deadlock.** Sending to an unbuffered channel blocks until a receiver is ready. Since the send and receive are on the same goroutine, it deadlocks.

**7. What happens here?**
```go
var ch chan int
ch <- 1
```
**Blocks forever.** Sending to a nil channel blocks forever — goroutine leak.

**8. What happens here?**
```go
ch := make(chan int)
close(ch)
ch <- 1
```
**Panic:** `send on closed channel`.

**9. What is the output?**
```go
ch := make(chan int)
close(ch)
v, ok := <-ch
fmt.Println(v, ok)
```
Output: `0 false` — receiving from a closed empty channel returns the zero value with `ok = false`.

**10. What is the output?**
```go
var i interface{} = nil
fmt.Println(i == nil) // true
```
Output: `true` — both type and value are nil.

**11. What is the output?**
```go
var p *int = nil
var i interface{} = p
fmt.Println(i == nil) // false!
```
Output: `false` — the interface `i` has type `*int` (non-nil type) even though the pointer value is nil. An interface is only nil when both type AND value are nil.

**12. Why is this code dangerous?**
```go
mu.Lock()
doSomething()
mu.Unlock()
```
If `doSomething()` panics, `Unlock()` is never called — the mutex is permanently locked, causing a deadlock.

**13. Why is this better?**
```go
mu.Lock()
defer mu.Unlock()
doSomething()
```
`defer` ensures the mutex is always unlocked, even if `doSomething()` panics.

**14. What is wrong with this?**
```go
rows, _ := db.Query("SELECT * FROM users")
```
The error is ignored. If `Query` fails, `rows` could be nil and later code will panic. Always check errors.

**15. What is missing here?**
```go
resp, err := http.Get(url)
if err != nil { return err }
```
`resp.Body` is never closed. This leaks the connection. Always:
```go
defer resp.Body.Close()
```

---

## 27. Good Questions to Ask the Interviewer

1. How do you use Go in your production systems?
2. What Go version are you currently using?
3. Do you rely mostly on the standard library, or do you use frameworks?
4. How do you structure Go services — monolith, microservices, or something in between?
5. How do you handle observability in your Go services (logs, metrics, traces)?
6. Do you use goroutines heavily? How do you prevent goroutine leaks?
7. How do you handle graceful shutdown across services?
8. What database libraries do you use — `database/sql`, `sqlx`, `pgx`, or an ORM?
9. How do you test your Go services? Do you use integration tests or mostly unit tests?
10. Do you containerize Go services with Docker and orchestrate with Kubernetes?
11. What are the biggest performance challenges in your Go codebase?
12. How do you do code review for Go — what are your main focus areas?
13. What does idiomatic, good Go code look like on your team?
14. How do you manage database migrations in production?

---

## Most Important Topics to Prepare First

For interviews, focus on these in priority order:

1. **Slices** — internal structure, append, capacity, aliasing
2. **Maps** — nil map, concurrent access, key types
3. **Pointers** — escape analysis, nil pointer, value vs pointer receivers
4. **Interfaces** — implicit implementation, nil interface trap, empty interface
5. **Errors** — error interface, wrapping, `errors.Is`/`errors.As`, sentinel errors
6. **Goroutines** — GMP model, goroutine leaks, WaitGroup
7. **Channels** — buffered vs unbuffered, nil/closed channel behavior, select
8. **Context** — cancellation, timeout, propagation, always defer cancel
9. **Mutex / WaitGroup** — race conditions, RWMutex, deadlocks
10. **HTTP server/client** — graceful shutdown, timeouts, middleware
11. **Testing** — table-driven tests, httptest, race detector
12. **Worker pools** — fan-out/fan-in, semaphores, pipelines
13. **Graceful shutdown** — signal handling, context cancellation
14. **database/sql** — connection pooling, transactions, prepared statements
15. **Profiling basics** — pprof, escape analysis, allocation reduction
