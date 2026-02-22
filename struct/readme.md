# Go Structs - Understanding Custom Types

A practical guide to understanding structs in Go based on hands-on learning.

## Table of Contents

- [What is a Struct?](#what-is-a-struct)
- [Type Definition vs Instance](#type-definition-vs-instance)
- [Memory Organization](#memory-organization)
- [Creating Instances](#creating-instances)
- [Accessing Properties](#accessing-properties)
- [Complete Example](#complete-example)

---

## What is a Struct?

A **struct** is a custom type that groups together related data. It's a way to create your own data types with multiple fields.

```go
type User struct {
    Name string
    Age  int
}
```

This defines a new type called `User` with two fields: `Name` and `Age`.

---

## Type Definition vs Instance

### Type Definition (Read-Only Blueprint)

```go
type User struct {
    Name string
    Age  int
}
```

- This is just a **definition** or **blueprint**
- It's **read-only** - it describes what a User looks like
- No memory is allocated for data yet
- Stored in the **code segment** of memory
- Think of it like a template or a class definition

### Instance (Actual Object)

```go
user1 := User{
    Name: "Arman",
    Age:  30,
}
```

- This is an **instance** or **object**
- Actual memory is allocated to store the data
- Contains real values ("Arman", 30)
- Each instance is independent
- Stored in the **stack** or **heap** (depending on usage)

**Key Point:** The type definition is like a blueprint for a house. An instance is an actual house built from that blueprint.

---

## Memory Organization

Go programs organize memory into segments:

```
┌─────────────────────────────────────┐
│        CODE SEGMENT                 │
│  ┌─────────────────────────────┐   │
│  │ type User struct {          │   │
│  │     Name string             │   │
│  │     Age  int                │   │
│  │ }                           │   │
│  │                             │   │
│  │ func main() {...}           │   │
│  └─────────────────────────────┘   │
└─────────────────────────────────────┘

┌─────────────────────────────────────┐
│        STACK/HEAP                   │
│  ┌─────────────────────────────┐   │
│  │ user1: User {               │   │
│  │     Name: "Arman"           │   │
│  │     Age:  30                │   │
│  │ }                           │   │
│  │                             │   │
│  │ user2: User {               │   │
│  │     Name: "Nusrat"          │   │
│  │     Age:  28                │   │
│  │ }                           │   │
│  └─────────────────────────────┘   │
└─────────────────────────────────────┘
```

### Code Segment
- Stores type definitions
- Stores function definitions
- Read-only
- Shared by all instances

### Stack/Heap
- Stores actual instance data
- Each instance has its own memory
- Can be modified
- Independent from each other

---

## Creating Instances

### Method 1: Declare then Assign

```go
var user1 User          // Declare a variable of type User
user1 = User{           // Assign values to it
    Name: "Arman",
    Age:  30,
}
```

**Steps:**
1. `var user1 User` - Creates a variable, zero-initialized
2. `user1 = User{...}` - Assigns values to the fields

### Method 2: Declare and Initialize Together

```go
user2 := User{          // Short declaration with initialization
    Name: "Nusrat",
    Age:  28,
}
```

**This is shorter and more common** - combines declaration and initialization.

### Important Notes

✅ **Each instance is independent:**
```go
user1 := User{Name: "Arman", Age: 30}
user2 := User{Name: "Nusrat", Age: 28}
```
- `user1` and `user2` are **separate** objects in memory
- Changing `user1` does NOT affect `user2`
- They don't replace each other

✅ **Terminology:**
- **Instance** = An actual object created from the type
- **Object** = Same as instance (can be used interchangeably in Go)
- **Instantiate** = The act of creating an instance

---

## Accessing Properties

### Member Variables / Properties

```go
type User struct {
    Name string    // Member variable / Property
    Age  int       // Member variable / Property
}
```

The fields inside a struct are called:
- **Member variables**
- **Properties** 
- **Fields**

All three terms mean the same thing.

### Dot Notation (.)

To access or modify a property, use the **dot operator (.)**:

```go
user1 := User{
    Name: "Arman",
    Age:  30,
}

// Reading properties
fmt.Println("Name:", user1.Name)  // Output: Name: Arman
fmt.Println("Age:", user1.Age)    // Output: Age: 30

// Modifying properties
user1.Name = "Arman Khan"
user1.Age = 31

fmt.Println("Updated Name:", user1.Name)  // Output: Updated Name: Arman Khan
fmt.Println("Updated Age:", user1.Age)    // Output: Updated Age: 31
```

**Syntax:**
```
variableName.PropertyName
```

---

## Complete Example

```go
package main

import "fmt"

// Type definition - stored in CODE SEGMENT
type User struct {
    Name string  // Property/Member variable
    Age  int     // Property/Member variable
}

func main() {
    // Method 1: Declare then assign
    var user1 User
    user1 = User{
        Name: "Arman",
        Age:  30,
    }
    
    // Accessing properties using dot notation
    fmt.Println("Name=", user1.Name)  // Output: Name= Arman
    fmt.Println("Age=", user1.Age)    // Output: Age= 30

    // Method 2: Declare and initialize together
    user2 := User{  // Creating an instance/object
        Name: "Nusrat",
        Age:  28,
    }
    
    // Accessing properties
    fmt.Println("Name", user2.Name)   // Output: Name Nusrat
    fmt.Println("Age", user2.Age)     // Output: Age 28
    
    // Instances are independent
    user1.Age = 31
    fmt.Println("user1 Age:", user1.Age)  // Output: user1 Age: 31
    fmt.Println("user2 Age:", user2.Age)  // Output: user2 Age: 28 (unchanged)
}
```

### What Happens in Memory:

```
CODE SEGMENT:
┌──────────────────────────────┐
│ type User struct {           │
│     Name string              │
│     Age  int                 │
│ }                            │
│                              │
│ func main() {...}            │
└──────────────────────────────┘

STACK:
┌──────────────────────────────┐
│ user1 (Instance 1)           │
│   Name: "Arman"              │
│   Age:  30                   │
├──────────────────────────────┤
│ user2 (Instance 2)           │
│   Name: "Nusrat"             │
│   Age:  28                   │
└──────────────────────────────┘
```

---

## Key Concepts Summary

| Concept | Description |
|---------|-------------|
| **Type Definition** | Blueprint/template stored in code segment (read-only) |
| **Instance** | Actual object created from the type definition |
| **Object** | Same as instance |
| **Instantiate** | Creating an instance |
| **Property/Member Variable/Field** | Variables inside a struct |
| **Dot Notation (.)** | Used to access properties: `user.Name` |
| **Independence** | Each instance has its own memory; changing one doesn't affect others |

---

## Terminology Equivalents

Different terms that mean the same thing:

```go
type User struct {
    Name string  // ← This is called:
    Age  int     //   - Member variable
                 //   - Property
                 //   - Field
}

user1 := User{...}  // ← This is called:
                    //   - Instance
                    //   - Object
                    //   - Creating/Instantiating
```

---

## Practice Exercise

Try this to reinforce your understanding:

```go
package main

import "fmt"

type Product struct {
    Name  string
    Price float64
    Stock int
}

func main() {
    // Create two product instances
    product1 := Product{
        Name:  "Laptop",
        Price: 1200.50,
        Stock: 5,
    }
    
    product2 := Product{
        Name:  "Mouse",
        Price: 25.99,
        Stock: 50,
    }
    
    // Access properties
    fmt.Println("Product 1:", product1.Name, "-", product1.Price)
    fmt.Println("Product 2:", product2.Name, "-", product2.Price)
    
    // Modify properties
    product1.Stock = product1.Stock - 1  // Sold one laptop
    fmt.Println("Laptop stock after sale:", product1.Stock)  // Output: 4
    
    // product2 is unaffected
    fmt.Println("Mouse stock:", product2.Stock)  // Output: 50 (unchanged)
}
```

---

**Remember:**
- Type definition = Blueprint (code segment)
- Instance = Actual object (stack/heap)
- Use dot (`.`) to access properties
- Each instance is independent

Happy coding! 🚀