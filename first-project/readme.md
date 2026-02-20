# Functional Programming in Go 🚀

A comprehensive guide to understanding functional programming concepts in Go, covering first-order functions, higher-order functions, callbacks, and first-class citizens.

## Table of Contents

- [Introduction to Functional Programming](#introduction-to-functional-programming)
- [Functional Programming Languages](#functional-programming-languages)
- [Core Concepts in Go](#core-concepts-in-go)
  - [1. First-Order Functions](#1-first-order-functions)
  - [2. Higher-Order Functions](#2-higher-order-functions)
  - [3. Callback Functions](#3-callback-functions)
  - [4. First-Class Citizens](#4-first-class-citizens)
- [Practical Examples](#practical-examples)
- [Best Practices](#best-practices)
- [Common Use Cases](#common-use-cases)

---

## Introduction to Functional Programming

**Functional Programming (FP)** is a programming paradigm that treats computation as the evaluation of mathematical functions. It emphasizes writing code that is predictable, testable, and easier to reason about.

### Key Principles

| Principle | Description |
|-----------|-------------|
| **Pure Functions** | Same input always produces the same output, no side effects |
| **Immutability** | Data structures don't change after creation |
| **First-Class Functions** | Functions can be assigned to variables, passed as arguments, and returned |
| **Function Composition** | Building complex operations by combining simpler functions |
| **Declarative Style** | Focus on *what* to compute, not *how* to compute it |

### Benefits

- ✅ Easier to test (pure functions)
- ✅ Easier to reason about (no hidden state)
- ✅ Better for concurrent programming
- ✅ More reusable code
- ✅ Reduced bugs from side effects

---

## Functional Programming Languages

### Pure Functional Languages
Languages designed specifically around functional programming:
- **Haskell** - Strong, static typing with lazy evaluation
- **Elm** - For building web UIs
- **PureScript** - Strongly-typed functional language that compiles to JavaScript

### Multi-Paradigm Languages
Languages that support functional programming alongside other paradigms:
- **JavaScript/TypeScript** - First-class functions, closures, map/filter/reduce
- **Python** - Lambda functions, map/filter, list comprehensions
- **Scala** - Runs on JVM, combines OOP and FP
- **Rust** - Ownership system with functional features
- **Go** - Limited FP support but covers essential concepts
- **Swift** - Apple's language with strong functional support
- **F#** - Functional-first language on .NET
- **Clojure** - Lisp dialect on JVM
- **Erlang/Elixir** - Built for concurrent, distributed systems

---

## Core Concepts in Go

### 1. First-Order Functions

First-order functions are regular functions that **do not** take functions as parameters or return functions. They're the building blocks of your program.

#### i. Standard Function (Named Function)

A traditional function with a name, defined at package or block level.

**Syntax:**
```go
func functionName(parameters) returnType {
    // function body
}
```

**Example:**
```go
package main

import "fmt"

// Standard named function
func greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

// Function with multiple parameters and return value
func add(a, b int) int {
    return a + b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

func main() {
    // Calling standard functions
    message := greet("Alice")
    fmt.Println(message) // Output: Hello, Alice!
    
    sum := add(10, 20)
    fmt.Println("Sum:", sum) // Output: Sum: 30
    
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result) // Output: Result: 5
    }
}
```

**When to use:**
- Reusable logic that will be called multiple times
- Public API functions (exported from packages)
- When you need a descriptive name for clarity

---

#### ii. Anonymous Function

A function without a name, defined inline. Also called a **function literal**.

**Syntax:**
```go
func(parameters) returnType {
    // function body
}
```

**Example:**
```go
package main

import "fmt"

func main() {
    // Anonymous function assigned to a variable
    multiply := func(x, y int) int {
        return x * y
    }
    
    result := multiply(4, 5)
    fmt.Println("Multiplication:", result) // Output: Multiplication: 20
    
    // Anonymous function used directly without assignment
    func(msg string) {
        fmt.Println("Direct call:", msg)
    }("Hello from anonymous function!") // Output: Direct call: Hello from anonymous function!
    
    // Anonymous function with closure (captures outer variable)
    counter := 0
    increment := func() int {
        counter++
        return counter
    }
    
    fmt.Println(increment()) // Output: 1
    fmt.Println(increment()) // Output: 2
    fmt.Println(increment()) // Output: 3
}
```

**When to use:**
- Short, one-time operations
- When passing functions as arguments
- Implementing closures
- Quick utility functions

**Closures Example:**
```go
func makeAdder(base int) func(int) int {
    // Anonymous function that "closes over" the base variable
    return func(x int) int {
        return base + x // base is captured from outer scope
    }
}

func main() {
    add5 := makeAdder(5)
    add10 := makeAdder(10)
    
    fmt.Println(add5(3))  // Output: 8  (5 + 3)
    fmt.Println(add10(3)) // Output: 13 (10 + 3)
}
```

---

#### iii. IIFE (Immediately Invoked Function Expression)

An anonymous function that is **defined and executed immediately**. Borrowed from JavaScript, this pattern is useful for initialization and encapsulation.

**Syntax:**
```go
func(parameters) returnType {
    // function body
}(arguments) // <- Note the immediate invocation
```

**Example:**
```go
package main

import "fmt"

func main() {
    // Simple IIFE
    result := func(x, y int) int {
        return x * y
    }(6, 7) // Immediately invoked with arguments
    
    fmt.Println("Result:", result) // Output: Result: 42
    
    // IIFE for initialization
    config := func() map[string]string {
        settings := make(map[string]string)
        settings["host"] = "localhost"
        settings["port"] = "8080"
        settings["protocol"] = "https"
        return settings
    }() // Execute immediately and assign result
    
    fmt.Println("Config:", config)
    // Output: Config: map[host:localhost port:8080 protocol:https]
    
    // IIFE with complex initialization logic
    validatedAge := func(age int) int {
        if age < 0 {
            fmt.Println("Invalid age, defaulting to 0")
            return 0
        }
        if age > 120 {
            fmt.Println("Age too high, defaulting to 120")
            return 120
        }
        return age
    }(150) // Pass 150, but it will be clamped to 120
    
    fmt.Println("Validated Age:", validatedAge) // Output: Validated Age: 120
}
```

**Advanced IIFE Example - Database Connection:**
```go
// IIFE for initializing a database connection with error handling
db := func() *sql.DB {
    conn, err := sql.Open("postgres", "connection_string")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    
    // Test the connection
    if err := conn.Ping(); err != nil {
        log.Fatal("Database unreachable:", err)
    }
    
    fmt.Println("Database connected successfully")
    return conn
}()

// db is now ready to use
defer db.Close()
```

**When to use:**
- One-time initialization with complex logic
- Encapsulating setup code
- Creating isolated scopes
- Validating or transforming values during assignment

---

#### iv. Function Expression

Assigning a function (named or anonymous) to a variable, making it a reusable function value.

**Syntax:**
```go
variableName := func(parameters) returnType {
    // function body
}
```

**Example:**
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // Basic function expression
    divide := func(a, b float64) float64 {
        if b == 0 {
            fmt.Println("Warning: division by zero")
            return 0
        }
        return a / b
    }
    
    // Now 'divide' can be reused multiple times
    fmt.Println(divide(10, 2))  // Output: 5
    fmt.Println(divide(15, 3))  // Output: 5
    fmt.Println(divide(7, 0))   // Output: Warning: division by zero, then 0
    
    // Function expression with complex logic
    validateEmail := func(email string) bool {
        return strings.Contains(email, "@") && 
               strings.Contains(email, ".") &&
               len(email) > 5
    }
    
    fmt.Println(validateEmail("test@example.com")) // Output: true
    fmt.Println(validateEmail("invalid"))          // Output: false
    
    // Reassigning function expressions
    operation := func(x, y int) int {
        return x + y
    }
    
    fmt.Println("Add:", operation(5, 3)) // Output: Add: 8
    
    // Reassign to different function
    operation = func(x, y int) int {
        return x * y
    }
    
    fmt.Println("Multiply:", operation(5, 3)) // Output: Multiply: 15
}
```

**Function Expression with Type Declaration:**
```go
package main

import "fmt"

// Define a function type
type MathOperation func(int, int) int

func main() {
    // Function expressions using the type
    var operation MathOperation
    
    operation = func(a, b int) int {
        return a + b
    }
    fmt.Println("Addition:", operation(10, 5)) // Output: Addition: 15
    
    operation = func(a, b int) int {
        return a - b
    }
    fmt.Println("Subtraction:", operation(10, 5)) // Output: Subtraction: 5
    
    // Array of function expressions
    operations := []MathOperation{
        func(a, b int) int { return a + b },
        func(a, b int) int { return a - b },
        func(a, b int) int { return a * b },
        func(a, b int) int { return a / b },
    }
    
    operationNames := []string{"Add", "Subtract", "Multiply", "Divide"}
    
    for i, op := range operations {
        fmt.Printf("%s: %d\n", operationNames[i], op(20, 4))
    }
    // Output:
    // Add: 24
    // Subtract: 16
    // Multiply: 80
    // Divide: 5
}
```

**When to use:**
- Creating flexible, swappable behaviors
- Strategy pattern implementation
- Dynamic function selection
- Building function registries or maps

---

### 2. Higher-Order Functions

A **higher-order function** is a function that does one or both of:
1. **Accepts** one or more functions as parameters
2. **Returns** a function as its result

Higher-order functions treat functions as **first-class citizens**, enabling powerful abstraction patterns.

#### Types of Higher-Order Functions

##### A. Functions that Accept Functions

**Example: Generic Operation Applier**
```go
package main

import "fmt"

// Higher-order function that accepts a function
func applyOperation(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func main() {
    // Define various operations
    add := func(x, y int) int { return x + y }
    multiply := func(x, y int) int { return x * y }
    subtract := func(x, y int) int { return x - y }
    
    // Use the same higher-order function with different operations
    fmt.Println("Add:", applyOperation(10, 5, add))           // Output: Add: 15
    fmt.Println("Multiply:", applyOperation(10, 5, multiply)) // Output: Multiply: 50
    fmt.Println("Subtract:", applyOperation(10, 5, subtract)) // Output: Subtract: 5
    
    // Pass anonymous function directly
    fmt.Println("Power:", applyOperation(2, 3, func(x, y int) int {
        result := 1
        for i := 0; i < y; i++ {
            result *= x
        }
        return result
    })) // Output: Power: 8
}
```

**Example: Map Function**
```go
// Map transforms a slice using a provided function
func mapInts(slice []int, transform func(int) int) []int {
    result := make([]int, len(slice))
    for i, val := range slice {
        result[i] = transform(val)
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    
    // Double each number
    doubled := mapInts(numbers, func(n int) int {
        return n * 2
    })
    fmt.Println("Doubled:", doubled) // Output: Doubled: [2 4 6 8 10]
    
    // Square each number
    squared := mapInts(numbers, func(n int) int {
        return n * n
    })
    fmt.Println("Squared:", squared) // Output: Squared: [1 4 9 16 25]
}
```

**Example: Filter Function**
```go
// Filter returns elements that satisfy the predicate
func filter(slice []int, predicate func(int) bool) []int {
    result := []int{}
    for _, val := range slice {
        if predicate(val) {
            result = append(result, val)
        }
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Get even numbers
    evens := filter(numbers, func(n int) bool {
        return n%2 == 0
    })
    fmt.Println("Evens:", evens) // Output: Evens: [2 4 6 8 10]
    
    // Get numbers greater than 5
    greaterThanFive := filter(numbers, func(n int) bool {
        return n > 5
    })
    fmt.Println("Greater than 5:", greaterThanFive) // Output: Greater than 5: [6 7 8 9 10]
}
```

##### B. Functions that Return Functions

**Example: Function Factory**
```go
package main

import "fmt"

// Higher-order function that returns a function
func createMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

// Returns a function with closure over prefix
func createGreeter(prefix string) func(string) string {
    return func(name string) string {
        return fmt.Sprintf("%s, %s!", prefix, name)
    }
}

func main() {
    // Create specific multiplier functions
    double := createMultiplier(2)
    triple := createMultiplier(3)
    tenTimes := createMultiplier(10)
    
    fmt.Println(double(5))    // Output: 10
    fmt.Println(triple(5))    // Output: 15
    fmt.Println(tenTimes(5))  // Output: 50
    
    // Create specific greeter functions
    casualGreet := createGreeter("Hey")
    formalGreet := createGreeter("Good morning")
    
    fmt.Println(casualGreet("Alice"))  // Output: Hey, Alice!
    fmt.Println(formalGreet("Bob"))    // Output: Good morning, Bob!
}
```

**Example: Counter with Closure**
```go
// Returns a function that maintains state
func createCounter(start int) func() int {
    count := start
    return func() int {
        count++
        return count
    }
}

func main() {
    counter1 := createCounter(0)
    counter2 := createCounter(100)
    
    fmt.Println(counter1()) // Output: 1
    fmt.Println(counter1()) // Output: 2
    fmt.Println(counter1()) // Output: 3
    
    fmt.Println(counter2()) // Output: 101
    fmt.Println(counter2()) // Output: 102
}
```

##### C. Functions that Do Both

**Example: Middleware Pattern**
```go
package main

import (
    "fmt"
    "time"
)

// Function type for handlers
type Handler func(string) string

// Middleware: accepts a handler and returns a wrapped handler
func loggingMiddleware(handler Handler) Handler {
    return func(input string) string {
        fmt.Printf("[LOG] Processing: %s\n", input)
        result := handler(input)
        fmt.Printf("[LOG] Result: %s\n", result)
        return result
    }
}

func timingMiddleware(handler Handler) Handler {
    return func(input string) string {
        start := time.Now()
        result := handler(input)
        duration := time.Since(start)
        fmt.Printf("[TIMING] Took: %v\n", duration)
        return result
    }
}

func main() {
    // Base handler
    baseHandler := func(input string) string {
        return "Processed: " + input
    }
    
    // Wrap with middleware
    handler := loggingMiddleware(timingMiddleware(baseHandler))
    
    handler("test data")
    // Output:
    // [LOG] Processing: test data
    // [TIMING] Took: ...
    // [LOG] Result: Processed: test data
}
```

**Example: Function Composition**
```go
// Compose two functions
func compose(f func(int) int, g func(int) int) func(int) int {
    return func(x int) int {
        return f(g(x))
    }
}

func main() {
    addOne := func(x int) int { return x + 1 }
    double := func(x int) int { return x * 2 }
    
    // Compose: first double, then add one
    doubleAndAddOne := compose(addOne, double)
    
    fmt.Println(doubleAndAddOne(5)) // Output: 11  (5*2 + 1)
}
```

---

### 3. Callback Functions

A **callback function** is a function that is passed as an argument to another function (a higher-order function) and is "called back" at an appropriate time.

#### Why Use Callbacks?

- ✅ Asynchronous operations
- ✅ Event handling
- ✅ Custom behavior injection
- ✅ Decoupling logic
- ✅ Flexible APIs

#### Basic Callback Example

```go
package main

import "fmt"

// processData accepts a callback for each element
func processData(data []int, callback func(int)) {
    for _, value := range data {
        callback(value) // Call the callback function
    }
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    
    // Callback 1: Print each number
    processData(numbers, func(n int) {
        fmt.Printf("Number: %d\n", n)
    })
    
    // Callback 2: Print squares
    processData(numbers, func(n int) {
        fmt.Printf("%d squared is %d\n", n, n*n)
    })
}
```

#### Async-Style Callbacks

```go
package main

import (
    "fmt"
    "time"
)

// Simulates fetching data with success/error callbacks
func fetchData(url string, onSuccess func(string), onError func(error)) {
    // Simulate network delay
    time.Sleep(1 * time.Second)
    
    if url == "" {
        onError(fmt.Errorf("URL cannot be empty"))
        return
    }
    
    // Simulate successful data fetch
    data := fmt.Sprintf("Data from %s", url)
    onSuccess(data)
}

func main() {
    fmt.Println("Fetching data...")
    
    // Define callbacks
    successCallback := func(data string) {
        fmt.Println("✓ Success:", data)
    }
    
    errorCallback := func(err error) {
        fmt.Println("✗ Error:", err)
    }
    
    // Call with callbacks
    fetchData("https://api.example.com", successCallback, errorCallback)
    // Output after 1 second: ✓ Success: Data from https://api.example.com
    
    fetchData("", successCallback, errorCallback)
    // Output: ✗ Error: URL cannot be empty
}
```

#### Predicate Callbacks (Filtering)

```go
// filter uses a predicate callback to determine which elements to keep
func filter(numbers []int, predicate func(int) bool) []int {
    result := []int{}
    for _, num := range numbers {
        if predicate(num) { // Callback determines inclusion
            result = append(result, num)
        }
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Different predicates (callbacks)
    isEven := func(n int) bool { return n%2 == 0 }
    isOdd := func(n int) bool { return n%2 != 0 }
    greaterThanFive := func(n int) bool { return n > 5 }
    
    fmt.Println("Evens:", filter(numbers, isEven))
    // Output: Evens: [2 4 6 8 10]
    
    fmt.Println("Odds:", filter(numbers, isOdd))
    // Output: Odds: [1 3 5 7 9]
    
    fmt.Println("Greater than 5:", filter(numbers, greaterThanFive))
    // Output: Greater than 5: [6 7 8 9 10]
}
```

#### Comparator Callbacks (Sorting)

```go
package main

import (
    "fmt"
    "sort"
)

type Person struct {
    Name string
    Age  int
}

// sortBy sorts a slice using a custom comparator callback
func sortPeople(people []Person, comparator func(Person, Person) bool) {
    sort.Slice(people, func(i, j int) bool {
        return comparator(people[i], people[j])
    })
}

func main() {
    people := []Person{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
        {"Diana", 28},
    }
    
    // Sort by age (ascending)
    sortPeople(people, func(a, b Person) bool {
        return a.Age < b.Age
    })
    fmt.Println("By Age:", people)
    // Output: By Age: [{Bob 25} {Diana 28} {Alice 30} {Charlie 35}]
    
    // Sort by name (alphabetically)
    sortPeople(people, func(a, b Person) bool {
        return a.Name < b.Name
    })
    fmt.Println("By Name:", people)
    // Output: By Name: [{Alice 30} {Bob 25} {Charlie 35} {Diana 28}]
}
```

#### Event Handler Callbacks

```go
package main

import "fmt"

type EventHandler func(string)

type Button struct {
    label    string
    onClick  EventHandler
    onHover  EventHandler
}

func (b *Button) Click() {
    if b.onClick != nil {
        b.onClick(b.label)
    }
}

func (b *Button) Hover() {
    if b.onHover != nil {
        b.onHover(b.label)
    }
}

func main() {
    submitButton := &Button{
        label: "Submit",
        onClick: func(label string) {
            fmt.Printf("Button '%s' was clicked!\n", label)
        },
        onHover: func(label string) {
            fmt.Printf("Hovering over '%s' button\n", label)
        },
    }
    
    submitButton.Hover()  // Output: Hovering over 'Submit' button
    submitButton.Click()  // Output: Button 'Submit' was clicked!
}
```

#### Real-World Example: HTTP Middleware

```go
package main

import (
    "fmt"
    "time"
)

type Request struct {
    Method string
    Path   string
}

type Response struct {
    StatusCode int
    Body       string
}

// Handler is a callback type
type Handler func(Request) Response

// Middleware wraps a handler with additional functionality
func authMiddleware(next Handler) Handler {
    return func(req Request) Response {
        fmt.Println("🔐 Checking authentication...")
        // Authentication logic here
        return next(req)
    }
}

func loggingMiddleware(next Handler) Handler {
    return func(req Request) Response {
        start := time.Now()
        fmt.Printf("📝 %s %s\n", req.Method, req.Path)
        
        response := next(req)
        
        duration := time.Since(start)
        fmt.Printf("✓ Completed in %v\n", duration)
        return response
    }
}

func main() {
    // Base handler
    baseHandler := func(req Request) Response {
        return Response{
            StatusCode: 200,
            Body:       "Hello, World!",
        }
    }
    
    // Wrap with middleware (callbacks)
    handler := loggingMiddleware(authMiddleware(baseHandler))
    
    req := Request{Method: "GET", Path: "/api/users"}
    response := handler(req)
    
    fmt.Printf("Response: %d - %s\n", response.StatusCode, response.Body)
}
```

---

### 4. First-Class Citizens

In Go, functions are **first-class citizens**, meaning they can be:

1. ✅ Assigned to variables
2. ✅ Passed as arguments to other functions
3. ✅ Returned from functions
4. ✅ Stored in data structures (slices, maps, structs)

This is what enables higher-order functions and functional programming patterns.

#### 1. Functions Assigned to Variables

```go
package main

import "fmt"

func main() {
    // Assign functions to variables
    add := func(a, b int) int { return a + b }
    subtract := func(a, b int) int { return a - b }
    multiply := func(a, b int) int { return a * b }
    divide := func(a, b int) int {
        if b == 0 {
            return 0
        }
        return a / b
    }
    
    // Use them like any other variable
    fmt.Println("Add:", add(10, 5))        // Output: Add: 15
    fmt.Println("Subtract:", subtract(10, 5)) // Output: Subtract: 5
    fmt.Println("Multiply:", multiply(10, 5)) // Output: Multiply: 50
    fmt.Println("Divide:", divide(10, 5))     // Output: Divide: 2
    
    // Reassign variables
    operation := add
    fmt.Println("Operation result:", operation(7, 3)) // Output: Operation result: 10
    
    operation = multiply
    fmt.Println("Operation result:", operation(7, 3)) // Output: Operation result: 21
}
```

#### 2. Functions Passed as Arguments

```go
// Execute a function with given arguments
func execute(a, b int, fn func(int, int) int) int {
    return fn(a, b)
}

func main() {
    result1 := execute(10, 5, func(x, y int) int { return x + y })
    fmt.Println("Sum:", result1) // Output: Sum: 15
    
    result2 := execute(10, 5, func(x, y int) int { return x * y })
    fmt.Println("Product:", result2) // Output: Product: 50
}
```

#### 3. Functions Returned from Functions

```go
func mathOperationFactory(operation string) func(int, int) int {
    switch operation {
    case "add":
        return func(a, b int) int { return a + b }
    case "subtract":
        return func(a, b int) int { return a - b }
    case "multiply":
        return func(a, b int) int { return a * b }
    case "divide":
        return func(a, b int) int {
            if b == 0 {
                return 0
            }
            return a / b
        }
    default:
        return func(a, b int) int { return 0 }
    }
}

func main() {
    addFunc := mathOperationFactory("add")
    multiplyFunc := mathOperationFactory("multiply")
    
    fmt.Println(addFunc(8, 2))      // Output: 10
    fmt.Println(multiplyFunc(8, 2)) // Output: 16
}
```

#### 4. Functions Stored in Data Structures

##### A. Functions in Maps

```go
package main

import "fmt"

// Define function type
type MathOperation func(int, int) int

func main() {
    // Store functions in a map
    operations := map[string]MathOperation{
        "add": func(a, b int) int { return a + b },
        "sub": func(a, b int) int { return a - b },
        "mul": func(a, b int) int { return a * b },
        "div": func(a, b int) int {
            if b == 0 {
                return 0
            }
            return a / b
        },
    }
    
    // Use functions from map
    fmt.Println("Add:", operations["add"](10, 5))  // Output: Add: 15
    fmt.Println("Mul:", operations["mul"](10, 5))  // Output: Mul: 50
    
    // Dynamic operation selection
    operationName := "sub"
    result := operations[operationName](20, 8)
    fmt.Println("Result:", result) // Output: Result: 12
}
```

##### B. Functions in Slices

```go
func main() {
    // Array of transformation functions
    transformers := []func(int) int{
        func(x int) int { return x * 2 },      // Double
        func(x int) int { return x * x },      // Square
        func(x int) int { return x + 10 },     // Add 10
        func(x int) int { return x - 5 },      // Subtract 5
    }
    
    value := 5
    fmt.Printf("Starting value: %d\n", value)
    
    for i, transform := range transformers {
        value = transform(value)
        fmt.Printf("After transform %d: %d\n", i+1, value)
    }
    // Output:
    // Starting value: 5
    // After transform 1: 10   (5 * 2)
    // After transform 2: 100  (10 * 10)
    // After transform 3: 110  (100 + 10)
    // After transform 4: 105  (110 - 5)
}
```

##### C. Functions in Structs

```go
package main

import "fmt"

type Calculator struct {
    Add      func(int, int) int
    Subtract func(int, int) int
    Multiply func(int, int) int
    History  []string
}

func NewCalculator() *Calculator {
    calc := &Calculator{
        Add:      func(a, b int) int { return a + b },
        Subtract: func(a, b int) int { return a - b },
        Multiply: func(a, b int) int { return a * b },
        History:  []string{},
    }
    return calc
}

func (c *Calculator) Compute(a, b int, operation string) int {
    var result int
    var opFunc func(int, int) int
    
    switch operation {
    case "add":
        opFunc = c.Add
    case "subtract":
        opFunc = c.Subtract
    case "multiply":
        opFunc = c.Multiply
    }
    
    result = opFunc(a, b)
    c.History = append(c.History, fmt.Sprintf("%d %s %d = %d", a, operation, b, result))
    return result
}

func main() {
    calc := NewCalculator()
    
    calc.Compute(10, 5, "add")
    calc.Compute(10, 5, "multiply")
    calc.Compute(10, 5, "subtract")
    
    fmt.Println("Calculation History:")
    for _, entry := range calc.History {
        fmt.Println(entry)
    }
    // Output:
    // Calculation History:
    // 10 add 5 = 15
    // 10 multiply 5 = 50
    // 10 subtract 5 = 5
}
```

#### Complete Example: Command Pattern with First-Class Functions

```go
package main

import "fmt"

// Command is a function type
type Command func()

// CommandRegistry stores commands as first-class citizens
type CommandRegistry struct {
    commands map[string]Command
}

func NewCommandRegistry() *CommandRegistry {
    return &CommandRegistry{
        commands: make(map[string]Command),
    }
}

func (cr *CommandRegistry) Register(name string, cmd Command) {
    cr.commands[name] = cmd
}

func (cr *CommandRegistry) Execute(name string) {
    if cmd, exists := cr.commands[name]; exists {
        cmd() // Execute the stored function
    } else {
        fmt.Printf("Command '%s' not found\n", name)
    }
}

func main() {
    registry := NewCommandRegistry()
    
    // Register commands (functions as first-class citizens)
    registry.Register("greet", func() {
        fmt.Println("Hello, World!")
    })
    
    registry.Register("time", func() {
        fmt.Println("Current time: 2024-01-15 10:30:00")
    })
    
    registry.Register("status", func() {
        fmt.Println("System Status: OK")
    })
    
    // Execute commands by name
    registry.Execute("greet")   // Output: Hello, World!
    registry.Execute("time")    // Output: Current time: 2024-01-15 10:30:00
    registry.Execute("status")  // Output: System Status: OK
    registry.Execute("invalid") // Output: Command 'invalid' not found
}
```

---

## Practical Examples

### Example 1: Data Processing Pipeline

```go
package main

import (
    "fmt"
    "strings"
)

type StringTransformer func(string) string

// Chain multiple transformers
func chain(transformers ...StringTransformer) StringTransformer {
    return func(input string) string {
        result := input
        for _, transformer := range transformers {
            result = transformer(result)
        }
        return result
    }
}

func main() {
    // Define transformers
    trim := func(s string) string {
        return strings.TrimSpace(s)
    }
    
    lowercase := func(s string) string {
        return strings.ToLower(s)
    }
    
    removeSpaces := func(s string) string {
        return strings.ReplaceAll(s, " ", "")
    }
    
    addPrefix := func(s string) string {
        return "processed_" + s
    }
    
    // Create processing pipeline
    pipeline := chain(trim, lowercase, removeSpaces, addPrefix)
    
    input := "  Hello World  "
    output := pipeline(input)
    
    fmt.Printf("Input:  '%s'\n", input)
    fmt.Printf("Output: '%s'\n", output)
    // Output: processed_helloworld
}
```

### Example 2: Validation Framework

```go
package main

import (
    "fmt"
    "regexp"
    "strings"
)

type ValidationRule func(string) (bool, string)

// Validator holds multiple validation rules
type Validator struct {
    rules []ValidationRule
}

func NewValidator() *Validator {
    return &Validator{rules: []ValidationRule{}}
}

func (v *Validator) AddRule(rule ValidationRule) *Validator {
    v.rules = append(v.rules, rule)
    return v
}

func (v *Validator) Validate(input string) (bool, []string) {
    var errors []string
    
    for _, rule := range v.rules {
        if valid, errMsg := rule(input); !valid {
            errors = append(errors, errMsg)
        }
    }
    
    return len(errors) == 0, errors
}

// Validation rule factories
func MinLength(min int) ValidationRule {
    return func(s string) (bool, string) {
        if len(s) < min {
            return false, fmt.Sprintf("must be at least %d characters", min)
        }
        return true, ""
    }
}

func MaxLength(max int) ValidationRule {
    return func(s string) (bool, string) {
        if len(s) > max {
            return false, fmt.Sprintf("must be at most %d characters", max)
        }
        return true, ""
    }
}

func MustContain(char string) ValidationRule {
    return func(s string) (bool, string) {
        if !strings.Contains(s, char) {
            return false, fmt.Sprintf("must contain '%s'", char)
        }
        return true, ""
    }
}

func MatchPattern(pattern string) ValidationRule {
    re := regexp.MustCompile(pattern)
    return func(s string) (bool, string) {
        if !re.MatchString(s) {
            return false, fmt.Sprintf("must match pattern: %s", pattern)
        }
        return true, ""
    }
}

func main() {
    // Build email validator
    emailValidator := NewValidator().
        AddRule(MinLength(5)).
        AddRule(MustContain("@")).
        AddRule(MustContain(".")).
        AddRule(MatchPattern(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`))
    
    // Build password validator
    passwordValidator := NewValidator().
        AddRule(MinLength(8)).
        AddRule(MaxLength(32)).
        AddRule(MatchPattern(`[A-Z]`)).  // Must have uppercase
        AddRule(MatchPattern(`[a-z]`)).  // Must have lowercase
        AddRule(MatchPattern(`[0-9]`))   // Must have digit
    
    // Test email
    testEmails := []string{"user@example.com", "invalid", "test@", "a@b.c"}
    
    fmt.Println("Email Validation:")
    for _, email := range testEmails {
        valid, errors := emailValidator.Validate(email)
        if valid {
            fmt.Printf("✓ '%s' is valid\n", email)
        } else {
            fmt.Printf("✗ '%s' is invalid:\n", email)
            for _, err := range errors {
                fmt.Printf("  - %s\n", err)
            }
        }
    }
    
    // Test passwords
    fmt.Println("\nPassword Validation:")
    testPasswords := []string{"MyP@ssw0rd", "weak", "NoDigits", "alllowercase123"}
    
    for _, pwd := range testPasswords {
        valid, errors := passwordValidator.Validate(pwd)
        if valid {
            fmt.Printf("✓ Password is valid\n")
        } else {
            fmt.Printf("✗ Password is invalid:\n")
            for _, err := range errors {
                fmt.Printf("  - %s\n", err)
            }
        }
    }
}
```

### Example 3: Retry Mechanism with Callbacks

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

type RetryableFunc func() error

// Retry executes a function with retry logic
func Retry(fn RetryableFunc, maxAttempts int, delay time.Duration, onRetry func(int, error)) error {
    var err error
    
    for attempt := 1; attempt <= maxAttempts; attempt++ {
        err = fn()
        
        if err == nil {
            return nil // Success
        }
        
        if attempt < maxAttempts {
            if onRetry != nil {
                onRetry(attempt, err) // Callback on retry
            }
            time.Sleep(delay)
        }
    }
    
    return fmt.Errorf("failed after %d attempts: %w", maxAttempts, err)
}

func main() {
    rand.Seed(time.Now().UnixNano())
    
    // Simulate unreliable operation
    unreliableOperation := func() error {
        if rand.Float32() < 0.7 { // 70% chance of failure
            return fmt.Errorf("operation failed")
        }
        return nil
    }
    
    // Retry with callback
    err := Retry(
        unreliableOperation,
        5,                    // max 5 attempts
        500*time.Millisecond, // wait 500ms between retries
        func(attempt int, err error) {
            fmt.Printf("Attempt %d failed: %v. Retrying...\n", attempt, err)
        },
    )
    
    if err != nil {
        fmt.Println("Final error:", err)
    } else {
        fmt.Println("✓ Operation succeeded!")
    }
}
```

---

## Best Practices

### 1. ✅ Keep Functions Pure When Possible

```go
// ✗ BAD: Function with side effects
var counter int
func incrementAndReturn() int {
    counter++ // Modifies external state
    return counter
}

// ✓ GOOD: Pure function
func increment(value int) int {
    return value + 1
}
```

### 2. ✅ Use Descriptive Function Types

```go
// ✗ BAD: Generic function signature
type Func func(interface{}) interface{}

// ✓ GOOD: Descriptive types
type UserValidator func(User) error
type PriceCalculator func(Product, int) float64
type ErrorHandler func(error)
```

### 3. ✅ Prefer Small, Focused Functions

```go
// ✗ BAD: Function doing too much
func processUser(user User) error {
    // Validate
    // Transform
    // Save to database
    // Send email
    // Update cache
    // ... too many responsibilities
}

// ✓ GOOD: Separate concerns
func validateUser(user User) error { /* ... */ }
func transformUser(user User) User { /* ... */ }
func saveUser(user User) error { /* ... */ }
```

### 4. ✅ Use Closures for Configuration

```go
type ServerConfig struct {
    Port    int
    Timeout time.Duration
}

func createServer(config ServerConfig) func() {
    return func() {
        fmt.Printf("Server running on port %d with timeout %v\n", 
            config.Port, config.Timeout)
        // Server logic uses config from closure
    }
}
```

### 5. ✅ Document Higher-Order Functions Clearly

```go
// Map applies a transformation function to each element in the slice
// and returns a new slice with the transformed values.
//
// Example:
//   doubled := Map([]int{1,2,3}, func(x int) int { return x * 2 })
//   // Result: [2, 4, 6]
func Map[T any, R any](slice []T, transform func(T) R) []R {
    result := make([]R, len(slice))
    for i, val := range slice {
        result[i] = transform(val)
    }
    return result
}
```

---

## Common Use Cases

### 1. 📊 Data Transformation Pipelines

```go
type Pipeline func([]int) []int

func FilterEven(data []int) []int {
    result := []int{}
    for _, v := range data {
        if v%2 == 0 {
            result = append(result, v)
        }
    }
    return result
}

func Double(data []int) []int {
    result := make([]int, len(data))
    for i, v := range data {
        result[i] = v * 2
    }
    return result
}

func Sum(data []int) int {
    total := 0
    for _, v := range data {
        total += v
    }
    return total
}

// Compose pipelines
numbers := []int{1, 2, 3, 4, 5, 6}
even := FilterEven(numbers)
doubled := Double(even)
sum := Sum(doubled)
fmt.Println(sum) // Output: 24
```

### 2. 🔌 Middleware/Plugin Systems

```go
type Middleware func(HandlerFunc) HandlerFunc
type HandlerFunc func(Context)

func Logger() Middleware {
    return func(next HandlerFunc) HandlerFunc {
        return func(ctx Context) {
            fmt.Println("Logging request...")
            next(ctx)
        }
    }
}

func Auth() Middleware {
    return func(next HandlerFunc) HandlerFunc {
        return func(ctx Context) {
            fmt.Println("Checking auth...")
            next(ctx)
        }
    }
}
```

### 3. 🎯 Strategy Pattern

```go
type PaymentStrategy func(amount float64) error

func CreditCardPayment(amount float64) error {
    fmt.Printf("Processing credit card payment: $%.2f\n", amount)
    return nil
}

func PayPalPayment(amount float64) error {
    fmt.Printf("Processing PayPal payment: $%.2f\n", amount)
    return nil
}

func CryptoPayment(amount float64) error {
    fmt.Printf("Processing crypto payment: $%.2f\n", amount)
    return nil
}

// Select strategy at runtime
func processPayment(amount float64, strategy PaymentStrategy) error {
    return strategy(amount)
}
```

### 4. 🧪 Testing with Mocks

```go
type UserRepository interface {
    GetUser(id int) (*User, error)
}

// Production implementation
func RealUserRepo() UserRepository { /* ... */ }

// Test implementation (mock)
func MockUserRepo(users map[int]*User) UserRepository {
    return func(id int) (*User, error) {
        if user, ok := users[id]; ok {
            return user, nil
        }
        return nil, errors.New("user not found")
    }
}
```

### 5. ⏱️ Debouncing/Throttling

```go
func Debounce(interval time.Duration, fn func()) func() {
    var timer *time.Timer
    
    return func() {
        if timer != nil {
            timer.Stop()
        }
        timer = time.AfterFunc(interval, fn)
    }
}

// Usage
saveData := Debounce(300*time.Millisecond, func() {
    fmt.Println("Saving data...")
})

// Multiple rapid calls only execute once after 300ms
saveData()
saveData()
saveData() // Only this one eventually executes
```

---

## Summary

### Key Takeaways

| Concept | Description | Use When |
|---------|-------------|----------|
| **First-Order Functions** | Regular functions (named, anonymous, IIFE) | Basic operations, utilities |
| **Higher-Order Functions** | Accept/return functions | Building abstractions, reusable logic |
| **Callbacks** | Functions passed as arguments | Event handling, async operations |
| **First-Class Citizens** | Functions as values | Flexible, dynamic behavior |

### The Power of Functional Programming in Go

While Go is not a pure functional language, it provides enough functional features to:

- ✅ Write cleaner, more modular code
- ✅ Reduce code duplication
- ✅ Create flexible, extensible systems
- ✅ Implement powerful design patterns
- ✅ Build composable, testable functions

### Next Steps

1. **Practice**: Implement map, filter, and reduce functions
2. **Experiment**: Build middleware systems or validation frameworks
3. **Refactor**: Convert imperative code to functional style
4. **Explore**: Study functional patterns in popular Go libraries
5. **Combine**: Mix functional and object-oriented approaches

---

## Additional Resources

- [Go Official Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example - Functions](https://gobyexample.com/functions)
- [Functional Options Pattern](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)

---

**Happy Coding! 🚀**

*Remember: Functions are first-class citizens in Go. Treat them well, and they'll make your code elegant and powerful.*